package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alhkxsj/https-server/internal/server"
	"github.com/Alhkxsj/https-server/pkg/certgen"
)

func fatal(msg string, err error) {
	fmt.Println("❌ 错误:", msg)
	if err != nil {
		fmt.Println("   详情:", err.Error())
	}
	os.Exit(1)
}

func main() {
	port := flag.Int("port", 8443, "监听端口（默认 8443）")
	dir := flag.String("dir", ".", "共享目录（默认当前目录）")
	quiet := flag.Bool("quiet", false, "安静模式（不输出访问日志）")
	version := flag.Bool("version", false, "显示版本信息")
	flag.Parse()

	if *version {
		fmt.Println("HTTPS 文件服务器 v1.1.0")
		return
	}

	root, err := filepath.Abs(*dir)
	if err != nil {
		fatal("获取目录路径失败", err)
	}

	certPath, keyPath := certgen.GetCertPaths()
	if !certgen.CheckCertificateExists(certPath) {
		fmt.Println("⚠️  未检测到服务器证书")
		fmt.Println("请先运行：https-certgen")
		os.Exit(1)
	}

	opts := server.Options{
		Addr:     fmt.Sprintf(":%d", *port),
		Root:     root,
		Quiet:    *quiet,
		CertPath: certPath,
		KeyPath:  keyPath,
	}

	if err := server.Run(opts); err != nil {
		fatal("启动 HTTPS 服务器失败", err)
	}
}