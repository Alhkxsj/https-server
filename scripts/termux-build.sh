#!/bin/bash

set -e

if [ -z "$PREFIX" ]; then
    echo "❌ 未检测到Termux环境"
    echo "此脚本只能在Termux中运行"
    exit 1
fi

echo "🔧 构建HTTPS服务器 (Termux包)"

if ! command -v go &> /dev/null; then
    echo "❌ Go未安装"
    echo "请先安装Go: pkg install golang"
    exit 1
fi

if ! command -v git &> /dev/null; then
    echo "⚠️  Git未安装，将跳过版本检查"
fi

PKG_NAME="https-server"
PKG_VERSION="1.0.0"
BUILD_DIR="$HOME/.cache/${PKG_NAME}-build"
SRC_DIR="$BUILD_DIR/src"

mkdir -p "$BUILD_DIR" "$SRC_DIR"

cp src/https-server.go "$SRC_DIR/"
cp src/https-certgen.go "$SRC_DIR/"

cd "$SRC_DIR"

echo "📦 编译程序..."
go build -o https-server https-server.go
go build -o https-certgen https-certgen.go

echo "🚚 安装到Termux..."
install -Dm755 https-server "$PREFIX/bin/https-server"
install -Dm755 https-certgen "$PREFIX/bin/https-certgen"

mkdir -p "$PREFIX/etc/https-server"

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