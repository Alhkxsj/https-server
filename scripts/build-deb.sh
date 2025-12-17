#!/bin/bash
set -e

echo "📦 构建 hserve deb 包"

APP_NAME="hserve"
VERSION="1.2.3"
DIST_DIR="dist"
PKG_DIR="build/pkg"

# 只构建当前架构
CURRENT_ARCH=$(uname -m)
echo "📍 当前系统架构: $CURRENT_ARCH"
ARCHS=("$CURRENT_ARCH")

for arch in "${ARCHS[@]}"; do
    echo "🔄 构建 $arch 架构 deb 包..."
    
    # 根据架构设置 GOARCH
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
    
    # 创建临时目录
    TMP_PKG_DIR="${PKG_DIR}_${arch}"
    rm -rf $TMP_PKG_DIR
    mkdir -p $TMP_PKG_DIR/DEBIAN
    mkdir -p $TMP_PKG_DIR/data/data/com.termux/files/usr/bin
    mkdir -p $TMP_PKG_DIR/data/data/com.termux/files/usr/etc/hserve

    # 尝试构建对应架构的二进制文件
    if GOOS=android GOARCH=$GOARCH go build -o $TMP_PKG_DIR/data/data/com.termux/files/usr/bin/hserve ./cmd/hserve; then
        echo "✅ $arch 架构二进制文件构建成功"
        
        # 复制控制文件
        cat > $TMP_PKG_DIR/DEBIAN/control << EOF
Package: hserve
Version: $VERSION
Architecture: $arch
Maintainer: Alhkxsj <fan343908@@gmail.com>
Homepage: https://github.com/Alhkxsj/hserve
Depends: openssl, ca-certificates
Description: Simple and easy-to-use HTTPS file server for Termux
 A zero-configuration HTTPS file server with built-in certificate generation tool.
 Supports quick sharing of local files in Termux environment and running pure frontend web pages.
 Achieves HTTPS secure access through self-signed CA.
EOF

        # 复制 postinst 脚本 (已更新，英文输出，自动清理旧证书)
        cp packaging/termux/postinst $TMP_PKG_DIR/DEBIAN/postinst
        sed -i 's|#!/bin/bash|#!/data/data/com.termux/files/usr/bin/sh|' $TMP_PKG_DIR/DEBIAN/postinst

        # 复制 prerm 脚本 (已更新，移除了 emoji)
        cp packaging/termux/prerm $TMP_PKG_DIR/DEBIAN/prerm
        sed -i 's|#!/bin/bash|#!/data/data/com.termux/files/usr/bin/sh|' $TMP_PKG_DIR/DEBIAN/prerm

        # 设置权限
        chmod 755 $TMP_PKG_DIR/DEBIAN
        chmod 755 $TMP_PKG_DIR/DEBIAN/postinst
        chmod 755 $TMP_PKG_DIR/DEBIAN/prerm

        # 构建 deb 包
        dpkg-deb --build $TMP_PKG_DIR $DIST_DIR/$APP_NAME"_"$VERSION"_"$arch.deb

        echo "✅ $arch 架构 deb 包构建完成: $DIST_DIR/$APP_NAME"_"$VERSION"_"$arch.deb"
    else
        echo "⚠️  $arch 架构构建失败，跳过..."
        # 清理失败的临时目录
        rm -rf $TMP_PKG_DIR
    fi
done

echo "🎉 deb 包构建完成！"