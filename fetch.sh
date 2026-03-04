#!/data/data/com.termux/files/usr/bin/bash
# 切换到工作目录
cd /data/data/com.termux/files/home/git/AVmerger
git checkout .
git pull
# 在这个目录编译二进制文件
go build -o /data/data/com.termux/files/usr/bin/merge main.go
