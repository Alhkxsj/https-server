package certgen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func GetLocalIP() string {
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

func GenerateCACert() ([]byte, *rsa.PrivateKey, error) {
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
		NotAfter:              time.Now().AddDate(100, 0, 0),
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

func GenerateServerCert(caCertDER []byte, caKey *rsa.PrivateKey, ip string) ([]byte, *rsa.PrivateKey, error) {
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
		NotAfter:    time.Now().AddDate(100, 0, 0),
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

func SaveCertFile(path string, certDER []byte, mode os.FileMode) {
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

func SaveKeyFile(path string, key *rsa.PrivateKey, mode os.FileMode) {
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

func ShowInstructions(caCertPath string) {
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