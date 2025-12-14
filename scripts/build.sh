#!/bin/bash

set -e

echo "🔧 为Termux构建HTTPS服务器..."

if ! command -v go &> /dev/null; then
    echo "❌ Go未安装"
    echo "请先安装Go: pkg install golang"
    exit 1
fi

if [ -n "$PREFIX" ] && [ -d "$PREFIX" ]; then
    echo "✅ 检测到Termux环境: $PREFIX"
else
    echo "⚠️  未检测到Termux环境"
    exit 1
fi

rm -rf build
mkdir -p build/bin

echo "📦 编译程序..."
cd src
go build -o ../build/bin/https-server https-server.go
go build -o ../build/bin/https-certgen https-certgen.go
cd ..

chmod +x build/bin/*

echo "✅ 构建完成!"

echo "📦 安装到Termux..."
install -Dm755 build/bin/https-server $PREFIX/bin/https-server
install -Dm755 build/bin/https-certgen $PREFIX/bin/https-certgen

mkdir -p $PREFIX/etc/https-server
mkdir -p $HOME/.local/share/https-server

echo "✅ 安装完成!"
echo ""
echo "📋 使用方法:"
echo "  https-certgen --install    # 生成并安装证书"
echo "  https-server               # 启动服务器"
echo ""
echo "  证书位置:"
echo "  - 服务器证书: $PREFIX/etc/https-server/"
echo "  - CA证书: $HOME/https-ca.crt (用于安装到安卓)"