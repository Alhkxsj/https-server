package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

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

func IsInTermux() bool {
	prefix := os.Getenv("PREFIX")
	if prefix != "" && len(prefix) > 4 && prefix[len(prefix)-4:] == "/usr" {
		return true
	}
	_, err := os.Stat("/data/data/com.termux/files/usr/bin/termux-setup-storage")
	return err == nil
}