.PHONY: all build clean install deb multiarch install-deb

APP_NAME := hserve
VERSION  := 1.2.3

PREFIX ?= /data/data/com.termux/files/usr
BIN_DIR := build/bin
DIST_DIR := dist
PKG_DIR := build/pkg

all: build

build:
	@echo "🔧 构建程序..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/hserve ./cmd/hserve
	@echo "✅ 构建完成"

install: build
	@echo "📦 安装到 Termux..."
	install -Dm755 $(BIN_DIR)/hserve $(PREFIX)/bin/hserve
	mkdir -p $(PREFIX)/etc/hserve
	@echo "✅ 安装完成"

deb:
	@echo "📦 构建当前架构的 deb 包..."
	./scripts/build-deb.sh

deb-all:
	@echo "📦 构建所有架构的 deb 包..."
	./scripts/build-deb-multiarch.sh

multiarch:
	@echo "📦 构建多架构版本..."
	./scripts/build-multiarch.sh

install-deb: deb
	@echo "📦 安装 deb 包 (aarch64)..."
	dpkg -i $(DIST_DIR)/$(APP_NAME)_$(VERSION)_aarch64.deb

install-deb-all: deb
	@echo "📦 安装所有架构的 deb 包..."
	@for arch in aarch64 arm i686 x86_64; do \
		if [ -f $(DIST_DIR)/$(APP_NAME)_$(VERSION)_$arch.deb ]; then \
			dpkg -i $(DIST_DIR)/$(APP_NAME)_$(VERSION)_$arch.deb; \
		fi \
	done

install-deb-arch:
	@echo "📦 安装指定架构的 deb 包..."
	@if [ -z "$(ARCH)" ]; then \
		echo "请指定架构: make install-deb-arch ARCH=aarch64"; \
		exit 1; \
	fi
	dpkg -i $(DIST_DIR)/$(APP_NAME)_$(VERSION)_$(ARCH).deb

clean:
	rm -rf build dist

fmt:
	@echo "🎨 格式化代码..."
	go fmt ./...

vet:
	@echo "🔍 检查代码..."
	go vet ./...

test:
	@echo "🧪 运行测试..."
	go test ./...