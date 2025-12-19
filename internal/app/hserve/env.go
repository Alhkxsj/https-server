package server

import (
	"fmt"
	"net"
	"os"
)

// 检测端口是否可用
func checkPort(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("端口 %s 无法监听，可能已被占用", addr)
	}
	_ = ln.Close()
	return nil
}

// 运行前环境自检
func PreflightCheck(addr, certPath, keyPath string) error {
	if _, err := os.Stat(certPath); err != nil {
		return fmt.Errorf("未找到证书文件：%s\n请先运行 hserve cert 生成证书", certPath)
	}

	if _, err := os.Stat(keyPath); err != nil {
		return fmt.Errorf("未找到私钥文件：%s\n请先运行 hserve cert 生成证书", keyPath)
	}

	if err := checkPort(addr); err != nil {
		return err
	}

	return nil
}
