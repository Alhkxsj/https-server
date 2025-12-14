TERMUX_PKG_HOMEPAGE=https://github.com/yourname/https-server
TERMUX_PKG_DESCRIPTION="一个简单易用的HTTPS文件服务器，支持自动证书生成"
TERMUX_PKG_LICENSE="MIT"
TERMUX_PKG_MAINTAINER="Your Name <your.email@example.com>"
TERMUX_PKG_VERSION=1.1.0
TERMUX_PKG_SRCURL=https://github.com/yourname/https-server/archive/v${TERMUX_PKG_VERSION}.tar.gz
TERMUX_PKG_SHA256=SKIP_CHECKSUM
TERMUX_PKG_BUILD_IN_SRC=true

termux_step_make() {
	go build -o https-server src/https-server.go
	go build -o https-certgen src/https-certgen.go
}

termux_step_make_install() {
	install -Dm755 https-server $TERMUX_PREFIX/bin/https-server
	install -Dm755 https-certgen $TERMUX_PREFIX/bin/https-certgen

	mkdir -p $TERMUX_PREFIX/etc/https-server
}