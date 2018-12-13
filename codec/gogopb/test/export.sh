#!/usr/bin/env bash

# 设置GOPATH
#CURR=`pwd`
#cd ../../../../../../..
#export GOPATH=`pwd`
#cd ${CURR}


# 插件及protoc存放路径
BIN_PATH=${GOPATH}bin

# 错误时退出
set -e

#在${GOPATH}bin 目录下安装 protoc-gen-gogofaster.exe
go install -v github.com/gogo/protobuf/protoc-gen-gogofaster

#在${GOPATH}bin 目录下安装 protoc-gen-msg.exe
go install -v github.com/cellnet/protoc-gen-msg

# windows下时，添加后缀名
if [ `go env GOHOSTOS` == "windows" ];then
	EXESUFFIX=.exe
fi

#protoc_win32.exe 是在https://github.com/protocolbuffers/protobuf/releases 下载的 protoc-3.6.1-win32.zip
# 生成协议
${BIN_PATH}/protoc_win32.exe --plugin=protoc-gen-gogofaster=${BIN_PATH}/protoc-gen-gogofaster${EXESUFFIX} --gogofaster_out=. --proto_path="." *.proto

# 生成cellnet 消息注册文件
${BIN_PATH}/protoc_win32.exe --plugin=protoc-gen-msg=${BIN_PATH}/protoc-gen-msg${EXESUFFIX} --msg_out=msgid.go:. --proto_path="." *.proto
#注册消息文件的生成依赖pbmeta(https://github.com/davyxu/pbmeta),而pbmeta依赖golexer(https://github.com/davyxu/golexer)

#注意：好像这个脚本只适合生成 syntax = "proto3";版本的proto文件
