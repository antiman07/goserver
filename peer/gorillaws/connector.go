package gorillaws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"sync"
	"time"
	"trunk/cellnet"
	"trunk/cellnet/myrpc"
	"trunk/cellnet/peer"
)

type websocketConnector struct{
	peer.SessionManager

	peer.CorePeerProperty
	peer.CoreContextSet
	peer.CoreRunningTag
	peer.CoreProcBundle
	peer.CoreTCPSocketOption

	defaultSes *WsSession

	tryConnTimes int // 尝试连接次数

	sesEndSignal sync.WaitGroup

	reconDur time.Duration
}

func (self *websocketConnector) Start() cellnet.Peer {

	self.WaitStopFinished()

	if self.IsRunning() {
		return self
	}

	go self.connect(self.Address())

	return self
}

func (self *websocketConnector) Session() cellnet.Session {
	return self.defaultSes
}

func (self *websocketConnector) SetSessionManager(raw interface{}) {
	self.SessionManager = raw.(peer.SessionManager)
}

func (self *websocketConnector) Stop() {
	if !self.IsRunning() {
		return
	}

	if self.IsStopping() {
		return
	}

	self.StartStopping()

	//往发送GO程队列投送nil,待当前队列所有包发送完毕遍历到nil时,退出发送GO程sendLoop,调用cleanup关闭链接导致接收GO程recvLoop也退出。
	self.defaultSes.Close()

	// 等待线程结束
	self.WaitStopFinished()

}

func (self *websocketConnector) ReconnectDuration() time.Duration {

	return self.reconDur
}

func (self *websocketConnector) SetReconnectDuration(v time.Duration) {
	self.reconDur = v
}

const reportConnectFailedLimitTimes = 3
//var addr = flag.String("addr", "localhost:8080", "http service address")

// 连接器，传入连接地址和发送封包次数
func (self *websocketConnector) connect(address string) {

	self.SetRunning(true)

	for {
		self.tryConnTimes++

		// 尝试用Socket连接地址
		//u := url.URL{Scheme: "ws", Host: address, Path: "/echo"}
		u := url.URL{Scheme: "ws", Host: address, Path: "/"}

		conn,_,err := websocket.DefaultDialer.Dial(u.String(),nil)

		self.defaultSes.conn = conn

		// 发生错误时退出
		if err != nil {

			if self.tryConnTimes <= reportConnectFailedLimitTimes {
				myrpc.Rpcqueue <- fmt.Sprintf("tcp.connect error(%s) %v",self.NameOrAddress(),err.Error())
			}

			if self.tryConnTimes == reportConnectFailedLimitTimes {
				myrpc.Rpcqueue <- fmt.Sprintf("(%s) error continue reconnecting, but mute log",self.NameOrAddress())
			}

			// 没重连就退出
			if self.ReconnectDuration() == 0 {
				myrpc.Rpcqueue <- fmt.Sprintf("tcp.connect error(%s)@%d address: %s",self.Name(), self.defaultSes.ID(), self.Address())
				self.PostEvent(&cellnet.RecvMsgEvent{self.defaultSes, &cellnet.SessionConnectError{}})
				break
			}

			// 有重连就等待
			time.Sleep(self.ReconnectDuration())

			// 继续连接
			continue
		}

		self.sesEndSignal.Add(1)

		//self.ApplySocketOption(conn)

		self.defaultSes.Start()

		self.tryConnTimes = 0

		self.PostEvent(&cellnet.RecvMsgEvent{self.defaultSes, &cellnet.SessionConnected{}})

		myrpc.Rpcqueue <- "-----建链成功----"

		self.sesEndSignal.Wait()

		myrpc.Rpcqueue <- "-----释放链接----"

		self.defaultSes.conn = nil

		// 没重连就退出/主动退出
		if self.IsStopping() || self.ReconnectDuration() == 0 {
			break
		}

		// 有重连就等待
		time.Sleep(self.ReconnectDuration())

		// 继续连接
		continue

	}

	self.SetRunning(false)

	self.EndStopping()
}

func (self *websocketConnector) IsReady() bool {

	return self.SessionCount() != 0
}

func (self *websocketConnector) TypeName() string {
	return "websocket.Connector"
}

var GsesManager *peer.CoreSessionManager
func init() {
	globalSessionManager := new(peer.CoreSessionManager)
	GsesManager = globalSessionManager
	peer.RegisterPeerCreator(func() cellnet.Peer {
		self := &websocketConnector{
			//SessionManager: new(peer.CoreSessionManager),
			SessionManager: globalSessionManager,
		}

		self.defaultSes = newSession(nil, self, func() {
			self.sesEndSignal.Done()
		})

		self.CoreTCPSocketOption.Init()

		//设置重连间隔时间 5s
		//self.SetReconnectDuration(time.Second * 5)

		return self
	})
}