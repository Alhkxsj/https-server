// internal/server/server.go
package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

// RunServer 启动HTTPS服务器
func RunServer(addr string, handler http.Handler, tlsConfig *tls.Config, quiet bool) {
	server := &http.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: tlsConfig,
	}
	
	if !quiet {
		log.Printf("服务启动: https://localhost%s", addr)
	}
	
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal("服务器错误:", err)
	}
}

// IsInTermux 检查是否在Termux环境中
func IsInTermux() bool {
	// 检查是否在Termux环境中
	prefix := os.Getenv("PREFIX")
	if prefix != "" && len(prefix) > 4 && prefix[len(prefix)-4:] == "/usr" {
		return true
	}
	// 检查Termux特有的目录
	_, err := os.Stat("/data/data/com.termux/files/usr/bin/termux-setup-storage")
	return err == nil
}