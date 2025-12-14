# HTTPS服务器

一个简单易用的HTTPS文件服务器，支持自动证书生成，适用于本地开发和文件共享。

## 特性

- 简单易用，零配置
- 自动生成HTTPS证书
- 支持安卓设备安装CA证书
- 文件系统访问日志
- 安全头设置

## 安装
记住，记住，记住，每一次重新安装软件包，都得重新生成ca证书，并到设置里面删除以前的旧证书，重新安装
### 快速安装 to Termux

```bash
# 克隆项目
git clone https://github.com/Alhkxsj/https-server.git
cd https-server

# 构建并安装
chmod +x scripts/build.sh
./scripts/build.sh
make quick-install
```

### 构建Deb包安装

```bash
# 构建Deb包
make deb

# 安装Deb包（Termux中不需要sudo）
dpkg -i build/https-server_1.0.0_aarch64.deb
```

### 手动构建

```bash
# 确保已安装Go 1.19+
go version

# 构建项目
make build

# 或者直接使用Go命令
mkdir -p build/bin
go build -o build/bin/https-server ./cmd/https-server
go build -o build/bin/https-certgen ./cmd/https-certgen

# 安装到Termux路径（不需要sudo）
install -Dm755 build/bin/https-server $PREFIX/bin/https-server
install -Dm755 build/bin/https-certgen $PREFIX/bin/https-certgen
```

## 使用方法

### 1. 生成证书

```bash
# 生成并安装证书到系统目录
https-certgen --install
```

### 2. 安装CA证书到安卓

- 找到CA证书文件：`~/https-ca.crt`
- 将证书复制到手机存储
- 设置 → 安全 → 加密与凭据 → 安装证书 → CA证书
- 选择证书文件

### 3. 启动服务器

```bash
# 在当前目录启动服务器（默认行为）
https-server

# 指定端口
https-server -port=8888

# 指定服务目录（默认为当前目录）
https-server -dir=/path/to/directory
https-server -dir=/sdcard/Download
https-server -dir=/home/user/mywebsite

# 安静模式（不显示访问日志）
https-server -quiet

# 查看版本信息
https-server --version

# 查看帮助信息
https-server --help
```

## 项目结构

```
https-server/
├── go.mod                     # Go模块文件
├── Makefile                   # 构建脚本
├── LICENSE                    # 许可证文件
├── README.md                  # 说明文档
├── cmd/                       # 命令行工具
│   ├── https-server/          # 主服务器命令
│   │   └── main.go
│   └── https-certgen/         # 证书生成工具命令
│       └── main.go
├── internal/                  # 内部包
│   └── server/                # 服务器核心逻辑
│       └── server.go
├── pkg/                       # 可重用包
│   └── certgen/               # 证书生成工具
│       └── certgen.go
├── build/                     # 构建输出目录
├── debian/                    # Debian包配置文件
│   ├── control                # 包信息
│   ├── postinst               # 安装后脚本
│   └── prerm                  # 卸载前脚本
└── scripts/                   # 构建和部署脚本
    ├── build.sh               # 构建脚本
    ├── build-termux-package.sh # Termux包构建脚本
    └── termux-build.sh        # Termux构建脚本
```

## 文件系统布局

安装后的文件结构 (Termux)：

```
# 程序文件
$PREFIX/bin/https-server          # 主程序
$PREFIX/bin/https-certgen         # 证书工具

# 系统证书文件
$PREFIX/etc/https-server/cert.pem # 服务器证书
$PREFIX/etc/https-server/key.pem  # 服务器私钥

# 用户证书文件
~/https-ca.crt                    # CA证书（手动安装到安卓）
```

## 安全特性

- TLS 1.2+ 最低版本要求
- 安全头设置 (X-Content-Type-Options, X-Frame-Options)
- 访问日志记录
- 证书验证

## 开发

```bash
# 构建
make build

# 清理
make clean
```

## 许可证

MIT
