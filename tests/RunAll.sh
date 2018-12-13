#!/usr/bin/env bash
export CURRDIR=`pwd`
cd ../../../../..
export GOPATH=`pwd`
cd ${CURRDIR}

set -e

go test -v .

# 编译例子
mkdir -p examplebin
go build -p 4 -v -o ./examplebin/echo github.com/cellnet/examples/echo
go build -p 4 -v -o ./examplebin/echo github.com/cellnet/examples/chat/client
go build -p 4 -v -o ./examplebin/echo github.com/cellnet/examples/chat/server
go build -p 4 -v -o ./examplebin/echo github.com/cellnet/examples/fileserver
go build -p 4 -v -o ./examplebin/echo github.com/cellnet/examples/websocket

function Cleanup()
{
    echo "cleanup"
    rm -rf examplebin
}

trap Cleanup EXIT