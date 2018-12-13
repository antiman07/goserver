package main

import (
	"fmt"
	"github.com/cellnet"
	"github.com/cellnet/peer"
	"github.com/cellnet/proc"
)

func clientAsyncCallback() {

	// 等待服务器返回数据
	done := make(chan struct{})

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("tcp.Connector", "clientAsyncCallback", peerAddress, queue) // 实例化对象见connector.go

	proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected: // 已经连接上
			fmt.Println("clientAsyncCallback connected")
			ev.Session().Send(&TestEchoACK{
				Msg:   "hello",
				Value: 1234,
			})
		case *TestEchoACK: //收到服务器发送的消息

			fmt.Printf("clientAsyncCallback recv %+v\n", msg)

			// 完成操作
			done <- struct{}{}

		case *cellnet.SessionClosed:
			fmt.Println("clientAsyncCallback closed")
		}
	})

	p.Start()

	queue.StartLoop()

	// 等待客户端收到消息
	<-done
}
