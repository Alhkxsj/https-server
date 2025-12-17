#!/bin/bash
set -e

echo "🔨 构建 hserve 多架构版本"

# 创建 dist 目录
mkdir -p dist

# 获取版本号
VERSION=$(grep "VERSION :=" Makefile | cut -d' ' -f3)
if [ -z "$VERSION" ]; then
    VERSION="1.2.3"
fi

echo "📦 版本号: $VERSION"

# 根据当前架构决定构建哪些架构
CURRENT_ARCH=$(uname -m)
echo "📍 当前系统架构: $CURRENT_ARCH"

# 根据当前平台支持的交叉编译能力选择架构
if [ "$CURRENT_ARCH" = "aarch64" ]; then
    # 在 aarch64 上可以尝试构建多个架构，但可能有些会失败
    ARCHS=("aarch64" "arm" "i686" "x86_64")
elif [ "$CURRENT_ARCH" = "x86_64" ]; then
    # 在 x86_64 上可以尝试构建多个架构
    ARCHS=("x86_64" "i686" "aarch64" "arm")
else
    # 其他架构只构建当前架构
    ARCHS=("$CURRENT_ARCH")
fi

for arch in "${ARCHS[@]}"; do
    echo "🔄 构建 $arch 架构..."
    
    case $arch in
        "aarch64")
            GOARCH=arm64
            ;;
        "arm")
            GOARCH=arm
            ;;
        "i686")
            GOARCH=386
            ;;
        "x86_64")
            GOARCH=amd64
            ;;
        *)
            # 如果架构不在预设列表中，使用架构名作为 GOARCH
            GOARCH=$arch
            ;;
    esac
    
    # 尝试构建
    if GOOS=android GOARCH=$GOARCH go build -o dist/hserve-$arch-$VERSION ./cmd/hserve; then
        echo "✅ $arch 架构构建完成"
    else
        echo "⚠️  $arch 架构构建失败，跳过..."
    fi
done

echo "🎉 架构构建完成！"
echo "📁 输出文件位于 dist/ 目录中："
ls -la dist/