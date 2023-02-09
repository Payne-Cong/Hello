#!/bin/bash


# host_ip=$(cat /etc/resolv.conf |grep "nameserver" |cut -f 2 -d " ")

# 利用WSL 安装的 Linux, 代理通过主机局域网共享设置
# 代理设置为主机的ip:端口

echo "Host ip: $(cat /etc/resolv.conf | grep nameserver | awk '{ print $2 }')"

echo "WSL client ip: $(hostname -I | awk '{print $1}')"