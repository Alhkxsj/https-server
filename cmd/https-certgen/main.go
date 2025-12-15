package main

import (
	"flag"
	"fmt"
	"os"

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
	force := flag.Bool("force", false, "强制重新生成证书")
	version := flag.Bool("version", false, "显示版本信息")
	flag.Parse()

	if *version {
		fmt.Println("HTTPS 证书生成工具 v1.1.0")
		return
	}

	fmt.Println("🔐 HTTPS 证书生成工具")
	fmt.Println("================================")

	if err := certgen.Generate(*force); err != nil {
		fatal("证书生成失败", err)
	}

	fmt.Println("================================")
}