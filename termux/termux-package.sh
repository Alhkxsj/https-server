# Termux包构建配置
# 文件名: build.sh (需要重命名为 termux-build.sh 或放入 .termux-build/ 目录)

TERMUX_PKG_HOMEPAGE=https://github.com/yourname/https-server
TERMUX_PKG_DESCRIPTION="一个简单易用的HTTPS文件服务器，支持自动证书生成"
TERMUX_PKG_LICENSE="MIT"
TERMUX_PKG_MAINTAINER="Your Name <your.email@example.com>"
TERMUX_PKG_VERSION=1.0.0
TERMUX_PKG_SRCURL=https://github.com/yourname/https-server/archive/v${TERMUX_PKG_VERSION}.tar.gz
TERMUX_PKG_SHA256=SKIP_CHECKSUM
TERMUX_PKG_BUILD_IN_SRC=true

termux_step_make() {
	# 编译Go程序
	go build -o https-server src/https-server.go
	go build -o https-certgen src/https-certgen.go
}

termux_step_make_install() {
	# 安装二进制文件
	install -Dm755 https-server $TERMUX_PREFIX/bin/https-server
	install -Dm755 https-certgen $TERMUX_PREFIX/bin/https-certgen

	# 创建证书目录
	mkdir -p $TERMUX_PREFIX/etc/https-server
}