package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Alhkxsj/hserve/pkg/certgen"
)

func fatal(msg string, err error) {
	fmt.Fprintln(os.Stderr, "❌ 错误:", msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "   详情:", err.Error())
	}
	os.Exit(1)
}

func main() {
	fmt.Println("💡 注意: hserve-certgen 功能已合并到主程序中")
	fmt.Println("💡 请使用 'hserve certgen' 命令来生成证书")
	fmt.Println()

	flag.Usage = func() {
		fmt.Println("🔐 HTTPS 证书生成工具 - 为您的安全访问保驾护航")
		fmt.Println("(此工具已合并到主程序中，请使用 hserve certgen 命令)")
		fmt.Println()
		fmt.Println("📖 使用方法:")
		fmt.Printf("  hserve certgen [选项]\n")
		fmt.Println()
		fmt.Println("✨ 可用选项:")
		fmt.Println("  -force")
		fmt.Println("      强制重新生成证书")
		fmt.Println("  -version")
		fmt.Println("      显示版本信息")
		fmt.Println("  -help")
		fmt.Println("      显示此帮助信息")
		fmt.Println()
		fmt.Println("💡 小贴士: 生成的证书用于 hserve 工具的 HTTPS 连接哦~")
		fmt.Println("🌟 愿代码如诗，生活如歌 ~")
	}

	force := flag.Bool("force", false, "强制重新生成证书")
	version := flag.Bool("version", false, "显示版本信息")
	help := flag.Bool("help", false, "显示此帮助信息")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *version {
		fmt.Println("🔐 hserve 证书生成工具 v1.3.0")
		fmt.Println("👤 作者: 快手阿泠 (Alexa Haley)")
		fmt.Println("🏠 项目地址: https://github.com/Alhkxsj/hserve")
		fmt.Println("✨ 愿代码如诗，生活如歌 ~")
		return
	}

	fmt.Println("🔐 HTTPS 证书生成工具 - 为您的安全访问保驾护航")
	fmt.Println("🌟 正在为您生成安全证书，请稍候...")

	if err := certgen.Generate(*force); err != nil {
		fatal("证书生成失败", err)
	}

	fmt.Println("================================")
}
