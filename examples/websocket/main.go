package main

import (
	"trunk/cellnet"
	"trunk/cellnet/peer"
	_ "trunk/cellnet/peer/gorillaws"
	"trunk/cellnet/proc"
	_ "trunk/cellnet/proc/gorillaws"
	"trunk/cellnet/util"
)


func main() {
	// 创建一个事件处理队列，整个服务器只有这一个队列处理事件，服务器属于单线程服务器
	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("gorillaws.Acceptor", "server", "http://192.168.10.10:6302/echo", queue)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {

		switch msg := ev.Message().(type) {

		case *cellnet.SessionAccepted:
			util.Slog.Debugf("server accepted client conn [%+v]",ev,msg)
		case *cellnet.SessionClosed:
			util.Slog.Debugf("session closed [%+v]",ev)
		}
	})

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()

}
