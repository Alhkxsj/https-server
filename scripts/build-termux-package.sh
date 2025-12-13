#!/bin/bash
# 为Termux构建和打包HTTPS服务器

set -e

echo "🔧 为Termux构建和打包HTTPS服务器..."

# 检查是否在Termux环境中
if [ -z "$PREFIX" ]; then
    echo "❌ 未检测到Termux环境"
    echo "此脚本只能在Termux中运行"
    exit 1
fi

echo "✅ 检测到Termux环境: $PREFIX"

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "❌ Go未安装"
    echo "请先安装Go: pkg install golang"
    exit 1
fi

# 检查dpkg-deb是否安装
if ! command -v dpkg-deb &> /dev/null; then
    echo "❌ dpkg-deb未安装"
    echo "请先安装dpkg: pkg install dpkg"
    exit 1
fi

# 创建构建目录
mkdir -p build/bin

# 编译程序
echo "📦 编译程序..."
cd src
go build -o ../build/bin/https-server https-server.go
go build -o ../build/bin/https-certgen https-certgen.go
cd ..

# 设置执行权限
chmod +x build/bin/*

echo "✅ 构建完成！"

# 构建Deb包
echo "📦 构建Deb包..."
make deb

echo "✅ Deb包构建完成！"
echo "包文件位于: build/https-server_1.0.0_aarch64.deb"

# 询问是否安装
read -p "是否要安装Deb包? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "📦 安装Deb包..."
    dpkg -i build/https-server_1.0.0_aarch64.deb
    echo "✅ Deb包安装完成！"
fi

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