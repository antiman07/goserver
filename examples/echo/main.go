package main

import (
	"flag"
	"fmt"

    /*有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包.
    golang对没有使用的导入包会编译报错，但是有时我们只想调用该包的init函数，不使用包导出的变量或者方法，这时就采用下划线导入方案。
    */
	_ "github.com/cellnet/peer/tcp" // 注册TCP Peer
	_ "github.com/cellnet/proc/tcp" // 注册TCP Processor
)

const peerAddress = "127.0.0.1:17701"

var clientmode = flag.Int("clientmode", 0, "0: for async recv, 1: for async rpc, 2: for sync rpc")

func main() {

	flag.Parse()

	server()

	switch *clientmode {
	case 0:
		fmt.Println("client mode: async callback")
		clientAsyncCallback()
	case 1:
		fmt.Println("client mode: async rpc")
		clientAsyncRPC()
	case 2:
		fmt.Println("client mode: sync rpc")
		clientSyncRPC()
	}

}
