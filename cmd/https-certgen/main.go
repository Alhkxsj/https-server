package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/Alhkxsj/https-server/pkg/certgen"
)

var (
	force   = flag.Bool("force", false, "强制重新生成")
	install = flag.Bool("install", false, "安装证书到系统")
	version = flag.Bool("version", false, "显示版本信息")
	help    = flag.Bool("help", false, "显示帮助信息")
)

func main() {
	flag.Parse()
	
	if *version {
		fmt.Println("HTTPS证书生成工具")
		fmt.Println("版本: 1.1.0")
		fmt.Println("作者: 快手阿泠好困想睡觉")
		fmt.Println("描述: 用于生成HTTPS服务器和CA证书的工具")
		os.Exit(0)
	}
	
	if *help {
		fmt.Println("HTTPS证书生成工具 - 用于生成HTTPS服务器和CA证书")
		fmt.Println("")
		fmt.Println("用法:")
		fmt.Println("  https-certgen [选项]")
		fmt.Println("")
		fmt.Println("选项:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("示例:")
		fmt.Println("  https-certgen --install          # 生成并安装证书")
		fmt.Println("  https-certgen --force            # 强制重新生成证书")
		os.Exit(0)
	}
	
	fmt.Println("🔐 HTTPS证书生成工具")
	fmt.Println(strings.Repeat("=", 50))
	
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/data/data/com.termux/files/home"
	}
	
	caCertPath := filepath.Join(home, "https-ca.crt")
	
	var serverCertPath, serverKeyPath string
	if IsInTermux() {
		prefix := os.Getenv("PREFIX")
		if prefix != "" {
			serverCertPath = prefix + "/etc/https-server/cert.pem"
			serverKeyPath = prefix + "/etc/https-server/key.pem"
		} else {
			serverCertPath = "/data/data/com.termux/files/usr/etc/https-server/cert.pem"
			serverKeyPath = "/data/data/com.termux/files/usr/etc/https-server/key.pem"
		}
	} else {
		serverCertPath = "/etc/https-server/cert.pem"
		serverKeyPath = "/etc/https-server/key.pem"
	}
	
	if !*force {
		if _, err := os.Stat(serverCertPath); err == nil {
			fmt.Println("✅ 系统证书已存在")
			fmt.Printf("📄 CA证书: %s\n", caCertPath)
			fmt.Println()
			certgen.ShowInstructions(caCertPath)
			return
		}
	}
	
	ip := certgen.GetLocalIP()
	fmt.Printf("📡 检测到本机IP: %s\n", ip)
	
	fmt.Println("\n📝 生成CA根证书...")
	caCert, caKey, err := certgen.GenerateCACert()
	if err != nil {
		log.Fatal("生成CA证书失败:", err)
	}
	
	fmt.Println("📝 生成服务器证书...")
	serverCert, serverKey, err := certgen.GenerateServerCert(caCert, caKey, ip)
	if err != nil {
		log.Fatal("生成服务器证书失败:", err)
	}
	
	fmt.Printf("💾 保存CA证书到: %s\n", caCertPath)
	certgen.SaveCertFile(caCertPath, caCert, 0644)
	
	if *install {
		fmt.Println("📦 安装证书到系统...")
		
		dir := filepath.Dir(serverCertPath)
		os.MkdirAll(dir, 0755)
		
		certgen.SaveCertFile(serverCertPath, serverCert, 0644)
		certgen.SaveKeyFile(serverKeyPath, serverKey, 0644)
		
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
	certgen.ShowInstructions(caCertPath)
}

func IsInTermux() bool {
	prefix := os.Getenv("PREFIX")
	if prefix != "" && len(prefix) > 4 && prefix[len(prefix)-4:] == "/usr" {
		return true
	}
	_, err := os.Stat("/data/data/com.termux/files/usr/bin/termux-setup-storage")
	return err == nil
}