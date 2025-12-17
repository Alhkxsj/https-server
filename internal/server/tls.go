package server

import (
	"crypto/tls"
	"fmt"

	tlspolicy "github.com/Alhkxsj/hserve/internal/tls"
)

func LoadTLSConfig(certPath, keyPath string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("加载 TLS 证书失败: %w", err)
	}
	return tlspolicy.DefaultConfig(cert), nil
}
