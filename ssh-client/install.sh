#!/bin/bash

GOPATH=$(go env GOPATH)

# 一些工具函数
function isMacos() {
  if [ "$(uname)" == "Darwin" ]; then
    return 1
  else
    return 0
  fi
}

function isLinux() {
  if [ "$(uname)" == "Linux" ]; then
    return 1
  else
    return 0
  fi
}

isLinux
linux_platform=$?
isMacos
macos_platform=$?

usr_bin_path=/usr/local/bin

# 创建必要的文件夹
mkdir -p /usr/local/etc/minssh

echo "======================== download ==========================="
go mod tidy

echo "======================== compile and install minssh ==========================="
go install ./cmd/minssh
cp "$GOPATH"/bin/minssh "$usr_bin_path"/minssh # 拷贝到 /usr/local/bin
echo "minssh install to $GOPATH/bin/minssh and $usr_bin_path/minssh"
echo ""

echo "minssh installed successfully"



