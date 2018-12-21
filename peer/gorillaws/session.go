package gorillaws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"trunk/cellnet"
	"trunk/cellnet/game"
	"trunk/cellnet/game/ddz"
	"trunk/cellnet/game/jdzjh"
	"trunk/cellnet/game/suoha"
	"trunk/cellnet/game/tbnn"
	"trunk/cellnet/myrpc"
	"trunk/cellnet/peer"
	"trunk/cellnet/util"
)

// Socket会话
type WsSession struct {
	peer.CoreContextSet
	peer.CoreSessionIdentify
	*peer.CoreProcBundle

	pInterface cellnet.Peer

	conn *websocket.Conn

	// 退出同步器
	exitSync sync.WaitGroup

	// 发送队列
	sendQueue *util.Pipe

	cleanupGuard sync.Mutex

	endNotify func()

	//新增玩家数据结构
	//player *role.Player
	player game.Player //具体的player 是指game目录下具体游戏的player实体(Zjh_Player or Suoha_Player)
}


func (self *WsSession) Player() interface{}{
	return self.player
}

func (self *WsSession) Peer() cellnet.Peer {
	return self.pInterface
}

// 取原始连接
func (self *WsSession) Raw() interface{} {
	if self.conn == nil {
		return nil
	}

	return self.conn
}

func (self *WsSession) Close() {
	self.sendQueue.Add(nil)
}

// 发送封包
func (self *WsSession) Send(msg interface{}) {
	self.sendQueue.Add(msg)
}

// 接收循环
func (self *WsSession) recvLoop() {

	//退出sendLoop会导致self.conn置NULL 所以在for循环外面也捕获一下该GO程退出的日志信息
	isWirte := false
	for self.conn != nil {

		msg, err := self.ReadMessage(self)

		if err != nil {
			isWirte = true

			myrpc.Rpcqueue <- fmt.Sprintf("退出接收数据GO程 error1=%v",err.Error())

			if !util.IsEOFOrNetReadError(err) {
				myrpc.Rpcqueue <- fmt.Sprintf("recvLoop closed error1=%v",err.Error())
			}

			self.PostEvent(&cellnet.RecvMsgEvent{self, &cellnet.SessionClosed{}})
			break
		}

		self.PostEvent(&cellnet.RecvMsgEvent{self, msg})
	}

	if isWirte == false {
		self.PostEvent(&cellnet.RecvMsgEvent{self, &cellnet.SessionClosed{}}) //加这行代码仅仅是便于统计"close player=" 的准确性
		myrpc.Rpcqueue <- fmt.Sprintf("退出接收数据GO程")
	}


	self.cleanup()
}

// 发送循环
func (self *WsSession) sendLoop() {

	var writeList []interface{}

	for {
		writeList = writeList[0:0]
		exit := self.sendQueue.Pick(&writeList)

		// 遍历要发送的数据
		for _, msg := range writeList {

			// TODO SendMsgEvent并不是很有意义
			err := self.SendMessage(&cellnet.SendMsgEvent{self, msg})
			if err != nil{
				myrpc.Rpcqueue <- fmt.Sprintf("退出发送数据GO程 error=%v",err.Error())
				self.cleanup()
				return
			}
		}

		if exit {
			myrpc.Rpcqueue <- fmt.Sprintf("退出发送数据GO程")
			break
		}
	}

	self.cleanup()
}

// 清理资源
func (self *WsSession) cleanup() {

	self.cleanupGuard.Lock()

	defer self.cleanupGuard.Unlock()

	// 关闭连接
	if self.conn != nil {
		self.conn.Close()
		self.conn = nil
	}
	self.Close() //add by wangtao 给队列喂药 退出sendLoop协程
	// 通知完成
	self.exitSync.Done()
}

//统计该用户游戏行为次数
func (self* WsSession) statis_player_data(){

	switch game.GameName{
	case "suoha":
		p := self.player.(*suoha.Suoha_Player)
		suoha.IncreAllPlayerData(p)
	case "jdzjh":
		p := self.player.(*jdzjh.Zjh_Player)
		jdzjh.IncreAllPlayerData(p)
	case "tbnn":
		p := self.player.(*tbnn.Tbnn_Player)
		tbnn.IncreAllPlayerData(p)
	default:
		myrpc.Rpcqueue <- "error 统计失败"
	}
}

// 启动会话的各种资源
func (self *WsSession) Start() {

	// 将会话添加到管理器
	self.Peer().(peer.SessionManager).Add(self)

	// 需要接收和发送线程同时完成时才算真正的完成
	self.exitSync.Add(2)

	switch game.GameName {
	case "suoha":
		self.player = new(suoha.Suoha_Player)
	case "jdzjh":
		self.player = new(jdzjh.Zjh_Player)
	case "tbnn":
		self.player = new(tbnn.Tbnn_Player)
	case "ddz":
		self.player = new(ddz.DDZ_Player)
	default:
		panic("not game name")
	}

	go func() {

		// 等待2个任务结束
		self.exitSync.Wait()

		self.statis_player_data()

		// 将会话从管理器移除
		self.Peer().(peer.SessionManager).Remove(self)

		if self.endNotify != nil {
			self.endNotify()
		}

	}()

	//客户端主动关闭链接情况下 会先停止sendLoop,然后停止recvLoop
	//服务端主动关闭链接情况下 会先导致recvLoop先退出,进而引发sendLoop退出.判断依据是日志中查找“send failed”
	// 启动并发接收goroutine
	go self.recvLoop()

	// 启动并发发送goroutine
	go self.sendLoop()
}

func newSession(conn *websocket.Conn, p cellnet.Peer, endNotify func()) *WsSession {
	self := &WsSession{
		conn:       conn,
		endNotify:  endNotify,
		sendQueue:  util.NewPipe(),
		pInterface: p,
		CoreProcBundle: p.(interface {
			GetBundle() *peer.CoreProcBundle
		}).GetBundle(),
	}


	return self
}
