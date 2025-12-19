#!/bin/bash
# hserve 使用示例

# 生成证书
hserve cert

# 启动服务器，默认端口 8443，共享当前目录
hserve

# 指定端口和目录
hserve serve -port=9443 -dir=/path/to/files

# 安静模式运行（不输出访问日志）
hserve serve -quiet -port=8080