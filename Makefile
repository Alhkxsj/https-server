# 为Termux优化的Makefile

.PHONY: all build clean install deb termux-install

# 检测是否在Termux环境中
ifeq ($(PREFIX),)
  # 非Termux环境
  TARGET_ARCH = arm64
  PACKAGE_NAME = https-server
  VERSION = 1.0.0
  SYSTEM_CERT_DIR = /etc/https-server
else
  # Termux环境
  TARGET_ARCH = aarch64
  PACKAGE_NAME = https-server
  VERSION = 1.0.0
  SYSTEM_CERT_DIR = $(PREFIX)/etc/https-server
endif

# 源代码目录
CMD_DIR = cmd
BUILD_DIR = build
DEB_DIR = $(BUILD_DIR)/deb
BIN_DIR = $(BUILD_DIR)/bin

all: build

# 编译程序
build: $(BIN_DIR)/https-server $(BIN_DIR)/https-certgen

$(BIN_DIR)/https-server: $(CMD_DIR)/https-server/main.go
	@mkdir -p $(BIN_DIR)
	go build -o $@ ./cmd/https-server

$(BIN_DIR)/https-certgen: $(CMD_DIR)/https-certgen/main.go
	@mkdir -p $(BIN_DIR)
	go build -o $@ ./cmd/https-certgen

# 清理
clean:
	rm -rf $(BUILD_DIR)

# 安装到系统（Termux）
install: build
	install -Dm755 $(BIN_DIR)/https-server $(PREFIX)/bin/https-server
	install -Dm755 $(BIN_DIR)/https-certgen $(PREFIX)/bin/https-certgen
	install -Dm755 debian/postinst $(PREFIX)/share/https-server/postinst
	# 创建证书目录
	mkdir -p $(PREFIX)/etc/https-server
	echo "✅ 安装完成!"

# 专为Termux的安装
termux-install: build
	@echo "🔧 安装到Termux..."
	install -Dm755 $(BIN_DIR)/https-server $(PREFIX)/bin/https-server
	install -Dm755 $(BIN_DIR)/https-certgen $(PREFIX)/bin/https-certgen
	# 创建证书目录
	mkdir -p $(PREFIX)/etc/https-server
	@echo "✅ Termux安装完成!"
	@echo ""
	@echo "📋 使用:"
	@echo "  1. 生成证书: https-certgen --install"
	@echo "  2. 安装CA证书到安卓"
	@echo "  3. 启动: https-server"

# 构建Deb包
deb: build
	@echo "📦 构建Deb包..."
	
	# 创建Deb包目录结构 - 符合Termux标准
	@mkdir -p $(DEB_DIR)/DEBIAN
	@chmod 755 $(DEB_DIR)/DEBIAN
	@mkdir -p $(DEB_DIR)/data/data/com.termux/files/usr/bin
	@mkdir -p $(DEB_DIR)/data/data/com.termux/files/usr/etc
	@mkdir -p $(DEB_DIR)/data/data/com.termux/files/usr/share/doc/https-server
	@mkdir -p $(DEB_DIR)/data/data/com.termux/files/usr/share/licenses/https-server
	
	# 复制控制文件
	cp debian/control $(DEB_DIR)/DEBIAN/
	cp debian/postinst $(DEB_DIR)/DEBIAN/
	cp debian/prerm $(DEB_DIR)/DEBIAN/
	chmod 755 $(DEB_DIR)/DEBIAN/postinst $(DEB_DIR)/DEBIAN/prerm
	
	# 复制二进制文件到Termux路径
	cp $(BIN_DIR)/https-server $(DEB_DIR)/data/data/com.termux/files/usr/bin/
	cp $(BIN_DIR)/https-certgen $(DEB_DIR)/data/data/com.termux/files/usr/bin/
	chmod 755 $(DEB_DIR)/data/data/com.termux/files/usr/bin/https-server $(DEB_DIR)/data/data/com.termux/files/usr/bin/https-certgen
	
	# 复制文档
	echo "HTTPS服务器 v$(VERSION)" > $(DEB_DIR)/data/data/com.termux/files/usr/share/doc/https-server/README
	echo "使用: https-server [选项]" >> $(DEB_DIR)/data/data/com.termux/files/usr/share/doc/https-server/README
	echo "MIT License" > $(DEB_DIR)/data/data/com.termux/files/usr/share/licenses/https-server/LICENSE
	
	# 构建包
	dpkg-deb --build $(DEB_DIR) $(BUILD_DIR)/$(PACKAGE_NAME)_$(VERSION)_$(TARGET_ARCH).deb
	
	@echo "✅ Deb包构建完成: $(BUILD_DIR)/$(PACKAGE_NAME)_$(VERSION)_$(TARGET_ARCH).deb"

# 快速安装到Termux
quick-install:
	@echo "⚡ 快速安装到Termux..."
	
	# 编译
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/https-server $(SRC_DIR)/https-server.go
	go build -o $(BIN_DIR)/https-certgen $(SRC_DIR)/https-certgen.go
	
	# 安装到Termux的bin目录
	install -Dm755 $(BIN_DIR)/https-server $(HOME)/../usr/bin/https-server
	install -Dm755 $(BIN_DIR)/https-certgen $(HOME)/../usr/bin/https-certgen
	
	# 创建配置目录
	mkdir -p $(HOME)/../usr/etc/https-server
	mkdir -p /etc/https-server  # 也创建标准路径以兼容
	
	@echo "✅ 安装完成!"
	@echo ""
	@echo "📋 使用:"
	@echo "  1. 生成证书: https-certgen --install"
	@echo "  2. 安装CA证书到安卓"
	@echo "  3. 启动: https-server"