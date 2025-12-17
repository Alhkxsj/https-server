package certmanager

import "os"

// CheckCertificateExists 检查证书是否存在
func CheckCertificateExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsInTermux 检测是否在 Termux 环境中
func IsInTermux() bool {
	return os.Getenv("PREFIX") != "" && os.Getenv("TERMUX_VERSION") != ""
}
