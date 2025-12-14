package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	port    = flag.Int("port", 8443, "HTTPS端口")
	dir     = flag.String("dir", ".", "服务目录")
	quiet   = flag.Bool("quiet", false, "安静模式")
	version = flag.Bool("version", false, "显示版本信息")
	help    = flag.Bool("help", false, "显示帮助信息")
)

func main() {
	flag.Parse()
	
	if *version {
		fmt.Println("HTTPS服务器")
		fmt.Println("版本: 1.1.0")
		fmt.Println("作者: 快手阿泠好困想睡觉")
		fmt.Println("描述: 一个简单易用的HTTPS文件服务器，支持自动证书生成")
		os.Exit(0)
	}
	
	if *help {
		fmt.Println("HTTPS服务器 - 简单易用的HTTPS文件服务器")
		fmt.Println("")
		fmt.Println("用法:")
		fmt.Println("  https-server [选项]")
		fmt.Println("")
		fmt.Println("选项:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("示例:")
		fmt.Println("  https-server                    # 在当前目录启动服务器")
		fmt.Println("  https-server -port=8080         # 指定端口启动")
		fmt.Println("  https-server -dir=/sdcard       # 指定服务目录")
		fmt.Println("  https-server -quiet             # 安静模式启动")
		os.Exit(0)
	}
	
	if *dir != "." {
		if err := os.Chdir(*dir); err != nil {
			log.Fatal("无法切换到目录:", err)
		}
	}
	
	cwd, _ := os.Getwd()
	absPath, _ := filepath.Abs(cwd)
	
	if !*quiet {
		fmt.Println("🚀 HTTPS服务器启动")
		fmt.Printf("📁 目录: %s\n", absPath)
		fmt.Printf("🔐 端口: %d\n", *port)
		fmt.Println("🛑 按Ctrl+C停止")
		fmt.Println()
	}
	
	var certPath, keyPath string
	if isInTermux() {
		certPath = os.Getenv("PREFIX") + "/etc/https-server/cert.pem"
		keyPath = os.Getenv("PREFIX") + "/etc/https-server/key.pem"
	} else {
		certPath = "/etc/https-server/cert.pem"
		keyPath = "/etc/https-server/key.pem"
	}
	
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		log.Println("⚠️  警告: 未找到系统证书")
		log.Println("请运行 'https-certgen' 生成证书")
		log.Fatal("或检查证书是否已安装")
	}
	
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal("加载证书失败:", err)
	}
	
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
	
	fs := http.FileServer(http.Dir("."))
	
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !*quiet && r.URL.Path != "/favicon.ico" {
			fmt.Printf("[%s] %s %s\n", 
				time.Now().Format("15:04:05"), 
				r.Method, 
				r.URL.Path)
		}
		
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		
		fs.ServeHTTP(w, r)
	})
	
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", *port),
		Handler:   handler,
		TLSConfig: tlsConfig,
	}
	
	if !*quiet {
		log.Printf("服务启动: https://localhost:%d", *port)
	}
	
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal("服务器错误:", err)
	}
}

func isInTermux() bool {
	prefix := os.Getenv("PREFIX")
	if prefix != "" && len(prefix) > 4 && prefix[len(prefix)-4:] == "/usr" {
		return true
	}
	_, err := os.Stat("/data/data/com.termux/files/usr/bin/termux-setup-storage")
	return err == nil
}