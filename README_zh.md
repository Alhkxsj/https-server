# hserve

一个快速搭建本地HTTPS服务器的工具。

## 特性

- [x] 单一可执行文件，多子命令设计
- [x] 智能证书管理（自动生成、安装）
- [x] 访问路径白名单控制
- [x] 支持外部TLS证书
- [x] 安装时交互式语言选择
- [x] 动态中英文界面切换
- [x] 多架构交叉编译支持
- [x] Termux环境优化

## 使用说明

> [警告] 注意：每次安装新版本后，建议重新生成证书并安装到系统中。

详细使用说明请参见 [使用说明文档](./docs/usage_zh.md)。

## 子命令详解

hserve 支持以下子命令：

### gen-cert - 生成证书
```bash
hserve gen-cert           # 生成证书
hserve gen-cert --force   # 强制重新生成证书
```

### serve - 启动服务器
```bash
hserve serve                                # 启动服务器（默认端口8443，当前目录）
hserve serve --port 9443 --dir /sdcard     # 指定端口和目录
hserve serve --quiet                       # 安静模式（不输出访问日志）
hserve serve --allow /sdcard --allow /home # 设置访问白名单
hserve serve --auto-gen                    # 自动为首次运行生成证书
hserve serve --tls-cert-file cert.pem --tls-key-file key.pem # 使用外部证书
```

### language - 切换语言
```bash
hserve language en    # 切换为英文
hserve language zh    # 切换为中文
```

### install-ca - 安装CA证书到Termux信任库
```bash
hserve install-ca    # 将CA证书安装到Termux信任库
```

### export-ca - 导出CA证书
```bash
hserve export-ca     # 导出CA证书到下载目录，便于手动安装到安卓系统
```

## 构建与安装

### 环境要求
- Go 1.21 或更高版本
- Termux 环境（用于安装和使用）

### 构建命令

```bash
# 构建二进制文件
make build

# 构建多架构版本
make multiarch

# 构建 deb 包
make deb

# 格式化代码
make fmt

# 检查代码
make vet

# 运行测试
make test

# 清理构建文件
make clean
```

### 安装方式

**方式一：直接安装二进制文件**
```bash
make install
```

**方式二：安装 deb 包**
```bash
# 构建并安装
make install-deb

# 或手动安装
dpkg -i dist/*.deb
```
## 许可证

本项目采用 [LICENSE](./LICENSE) 许可证。

## 使用说明
注意注意。每一次。安装新版本之后，都要重新生成证书，并从系统里删除以前的旧证书。重新安装新证书。
详细使用说明请参见 [使用说明文档](./docs/usage_zh.md)。

## 安全模型

关于安全模型的信息请参见 [安全模型文档](./docs/security-model_zh.md)。

## Android CA安装

在Android设备上安装CA证书的详细步骤请参见 [Android CA安装文档](./docs/android-ca-install_zh.md)。

## 构建与安装

### 构建

```bash
make build
```

### 构建deb包

```bash
make deb
```

### 安装

```bash
make install
```

安装deb包：

```bash
dpkg -i dist/*.deb
```

## 许可证

本项目采用 [LICENSE](./LICENSE) 许可证。