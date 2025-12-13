// src/https-certgen.go - 证书生成工具
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	force = flag.Bool("force", false, "强制重新生成")
	install = flag.Bool("install", false, "安装证书到系统")
)

func main() {
	flag.Parse()
	
	fmt.Println("🔐 HTTPS证书生成工具")
	fmt.Println(strings.Repeat("=", 50))
	
	// 获取用户主目录
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/data/data/com.termux/files/home"
	}
	
	// 证书保存路径
	caCertPath := filepath.Join(home, "https-ca.crt")
	
	// 在Termux中使用特定的证书路径
	var serverCertPath, serverKeyPath string
	if isTermux() {
		prefix := os.Getenv("PREFIX")
		if prefix != "" {
			serverCertPath = prefix + "/etc/https-server/cert.pem"
			serverKeyPath = prefix + "/etc/https-server/key.pem"
		} else {
			// fallback
			serverCertPath = "/data/data/com.termux/files/usr/etc/https-server/cert.pem"
			serverKeyPath = "/data/data/com.termux/files/usr/etc/https-server/key.pem"
		}
	} else {
		serverCertPath = "/etc/https-server/cert.pem"
		serverKeyPath = "/etc/https-server/key.pem"
	}
	
	// 检查是否已存在
	if !*force {
		if _, err := os.Stat(serverCertPath); err == nil {
			fmt.Println("✅ 系统证书已存在")
			fmt.Printf("📄 CA证书: %s\n", caCertPath)
			fmt.Println()
			showInstructions(caCertPath)
			return
		}
	}
	
	// 获取本机IP
	ip := getLocalIP()
	fmt.Printf("📡 检测到本机IP: %s\n", ip)
	
	// 生成CA证书
	fmt.Println("\n📝 生成CA根证书...")
	caCert, caKey, err := generateCACert()
	if err != nil {
		log.Fatal("生成CA证书失败:", err)
	}
	
	// 生成服务器证书
	fmt.Println("📝 生成服务器证书...")
	serverCert, serverKey, err := generateServerCert(caCert, caKey, ip)
	if err != nil {
		log.Fatal("生成服务器证书失败:", err)
	}
	
	// 保存CA证书到用户目录（用于安装到安卓）
	fmt.Printf("💾 保存CA证书到: %s\n", caCertPath)
	saveCertFile(caCertPath, caCert, 0644)
	
	// 保存服务器证书到系统目录
	if *install {
		fmt.Println("📦 安装证书到系统...")
		
		// 创建系统目录 - 在Termux中使用不同的路径
		dir := filepath.Dir(serverCertPath)
		os.MkdirAll(dir, 0755)
		
		// 保存服务器证书和密钥
		saveCertFile(serverCertPath, serverCert, 0644)
		saveKeyFile(serverKeyPath, serverKey, 0644)
		
		fmt.Printf("✅ 证书安装完成:\n")
		fmt.Printf("   📄 服务器证书: %s\n", serverCertPath)
		fmt.Printf("   🔑 服务器密钥: %s\n", serverKeyPath)
	} else {
		fmt.Println("\n📋 手动安装证书:")
		fmt.Printf("   mkdir -p %s\n", filepath.Dir(serverCertPath))
		fmt.Printf("   cp %s %s\n", filepath.Join(home, "server.crt"), serverCertPath)
		fmt.Printf("   cp %s %s\n", filepath.Join(home, "server.key"), serverKeyPath)
	}
	
	fmt.Println("\n" + strings.Repeat("=", 50))
	showInstructions(caCertPath)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func generateCACert() ([]byte, *rsa.PrivateKey, error) {
	caKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:      []string{"CN"},
			Organization: []string{"Local HTTPS CA"},
			CommonName:   "Local HTTPS Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	
	caCertDER, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}
	
	return caCertDER, caKey, nil
}

func generateServerCert(caCertDER []byte, caKey *rsa.PrivateKey, ip string) ([]byte, *rsa.PrivateKey, error) {
	caCert, err := x509.ParseCertificate(caCertDER)
	if err != nil {
		return nil, nil, err
	}
	
	serverKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:      []string{"CN"},
			Organization: []string{"Local HTTPS Server"},
			CommonName:   "localhost",
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:    []string{"localhost", ip},
		IPAddresses: []net.IP{
			net.IPv4(127, 0, 0, 1),
			net.IPv6loopback,
			net.ParseIP(ip),
		},
	}
	
	serverCertDER, err := x509.CreateCertificate(rand.Reader, &template, caCert, &serverKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}
	
	return serverCertDER, serverKey, nil
}

func saveCertFile(path string, certDER []byte, mode os.FileMode) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		log.Fatal("创建证书文件失败:", err)
	}
	defer file.Close()
	
	pem.Encode(file, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})
}

func saveKeyFile(path string, key *rsa.PrivateKey, mode os.FileMode) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		log.Fatal("创建私钥文件失败:", err)
	}
	defer file.Close()
	
	pem.Encode(file, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
}

func isTermux() bool {
	// 检查是否在Termux环境中
	_, err := os.Stat("/data/data/com.termux/files/usr/bin/termux-setup-storage")
	return err == nil
}

func showInstructions(caCertPath string) {
	fmt.Println("\n📱 安卓证书安装步骤:")
	fmt.Println("  1. 找到CA证书文件:", caCertPath)
	fmt.Println("  2. 将证书复制到手机存储")
	fmt.Println("  3. 设置 → 安全 → 加密与凭据")
	fmt.Println("  4. 安装证书 → CA证书")
	fmt.Println("  5. 选择证书文件，命名为 'Local HTTPS'")
	fmt.Println()
	fmt.Println("🚀 启动服务器:")
	fmt.Println("  cd /path/to/website")
	fmt.Println("  https-server")
	fmt.Println()
	fmt.Println("📖 更多信息:")
	fmt.Println("  https-server --help")
}