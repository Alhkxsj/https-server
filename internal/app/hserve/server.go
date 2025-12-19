package server

import (
	"fmt"
	"net/http"
)

type Options struct {
	Addr     string
	Root     string
	Quiet    bool
	CertPath string
	KeyPath  string
}

func Run(opt Options) error {
	if err := PreflightCheck(opt.Addr, opt.CertPath, opt.KeyPath); err != nil {
		return err
	}

	tlsConfig, err := LoadTLSConfig(opt.CertPath, opt.KeyPath)
	if err != nil {
		return err
	}

	handler := NewHandler(opt.Root, opt.Quiet)

	srv := &http.Server{
		Addr:      opt.Addr,
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	if !opt.Quiet {
		fmt.Printf("🚀 hserve 已启动\n")
		fmt.Printf("📁 共享目录: %s\n", opt.Root)
		fmt.Printf("🌐 访问地址: https://localhost%s\n", opt.Addr)
		fmt.Printf("🔐 监听地址: %s\n", opt.Addr)
		fmt.Println("💡 提示: 在浏览器中打开访问地址即可浏览文件")
		fmt.Print("🛑 按 Ctrl+C 停止\n\n")
	}

	return srv.ListenAndServeTLS("", "")
}
