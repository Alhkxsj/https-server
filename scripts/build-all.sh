#!/bin/bash
# 统一构建脚本，自动选择构建方式

echo "🚀 hserve 统一构建脚本"

if [ "$1" == "deb" ]; then
    echo "📦 构建 deb 包..."
    make deb
elif [ "$1" == "multiarch" ]; then
    echo "📦 构建多架构版本..."
    make multiarch
elif [ "$1" == "all" ]; then
    echo "📦 构建所有版本..."
    make build
    make multiarch
    make deb
else
    echo "🔧 构建 hserve..."
    make build
fi

echo "✅ 构建完成！"