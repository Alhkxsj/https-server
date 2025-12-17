package server

import (
	"crypto/tls"
	"fmt"

	"github.com/Alhkxsj/hserve/internal/i18n"
	tlspolicy "github.com/Alhkxsj/hserve/internal/tls"
)

func LoadTLSConfig(certPath, keyPath string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf(i18n.T(i18n.GetLanguage(), "tls_config_failed"), err)
	}
	return tlspolicy.DefaultConfig(cert), nil
}
