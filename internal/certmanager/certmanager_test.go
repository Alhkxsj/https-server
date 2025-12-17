package certmanager

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsInTermux(t *testing.T) {
	// 保存原始环境变量
	originalPrefix := os.Getenv("PREFIX")
	originalTermuxVersion := os.Getenv("TERMUX_VERSION")

	// 清理环境变量
	os.Unsetenv("PREFIX")
	os.Unsetenv("TERMUX_VERSION")

	// 测试非Termux环境
	if IsInTermux() {
		t.Error("Expected IsInTermux() to return false when not in Termux")
	}

	// 设置Termux环境
	os.Setenv("PREFIX", "/data/data/com.termux/files/usr")
	os.Setenv("TERMUX_VERSION", "0.118.0")

	if !IsInTermux() {
		t.Error("Expected IsInTermux() to return true when in Termux")
	}

	// 恢复原始环境变量
	os.Setenv("PREFIX", originalPrefix)
	os.Setenv("TERMUX_VERSION", originalTermuxVersion)
}

func TestGetCertPaths(t *testing.T) {
	// 这个测试会验证证书路径生成逻辑
	certPath, keyPath := GetCertPaths()

	// 检查路径是否包含正确的文件名
	if filepath.Ext(certPath) != ".crt" && filepath.Ext(certPath) != ".pem" {
		t.Errorf("Certificate path does not have expected extension: %s", certPath)
	}
	if filepath.Ext(keyPath) != ".key" && filepath.Ext(keyPath) != ".pem" {
		t.Errorf("Key path does not have expected extension: %s", keyPath)
	}
}

func TestCheckCertificateExists(t *testing.T) {
	// 测试不存在的文件
	if CheckCertificateExists("/nonexistent/path/to/cert") {
		t.Error("Expected non-existent file to not exist")
	}

	// 测试当前目录下不存在的文件
	if CheckCertificateExists("nonexistent.cert") {
		t.Error("Expected non-existent file to not exist")
	}
}
