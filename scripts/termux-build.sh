#!/bin/bash
# Termux包构建脚本

set -e  # 遇到错误时退出

# 检查是否在Termux环境中
if [ -z "$PREFIX" ]; then
    echo "❌ 未检测到Termux环境"
    echo "此脚本只能在Termux中运行"
    exit 1
fi

echo "🔧 构建HTTPS服务器 (Termux包)"

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "❌ Go未安装"
    echo "请先安装Go: pkg install golang"
    exit 1
fi

# 检查git是否安装
if ! command -v git &> /dev/null; then
    echo "⚠️  Git未安装，将跳过版本检查"
fi

# 设置变量
PKG_NAME="https-server"
PKG_VERSION="1.0.0"
BUILD_DIR="$HOME/.cache/${PKG_NAME}-build"
SRC_DIR="$BUILD_DIR/src"

# 创建构建目录
mkdir -p "$BUILD_DIR" "$SRC_DIR"

# 复制源代码
cp src/https-server.go "$SRC_DIR/"
cp src/https-certgen.go "$SRC_DIR/"

# 进入构建目录
cd "$SRC_DIR"

# 构建程序
echo "📦 编译程序..."
go build -o https-server https-server.go
go build -o https-certgen https-certgen.go

# 安装到Termux
echo "🚚 安装到Termux..."
install -Dm755 https-server "$PREFIX/bin/https-server"
install -Dm755 https-certgen "$PREFIX/bin/https-certgen"

# 创建证书目录
mkdir -p "$PREFIX/etc/https-server"

# 创建用户数据目录
mkdir -p "$HOME/.local/share/https-server"

echo "✅ 构建和安装完成！"
echo ""
echo "📋 使用说明："
echo "  1. 生成证书: https-certgen --install"
echo "  2. 安装CA证书到安卓系统"
echo "  3. 启动服务器: https-server"
echo ""
echo "📁 安装位置："
echo "  - 可执行文件: $PREFIX/bin/"
echo "  - 配置目录: $PREFIX/etc/https-server/"
echo "  - 用户证书: $HOME/https-ca.crt"