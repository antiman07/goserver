package main

import (
	"github.com/cellnet"
	"github.com/cellnet/examples/chat/proto"
	"github.com/cellnet/peer"
	"github.com/cellnet/proc"
	"github.com/golog"
	"strconv"
	"time"

	/*有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包.
golang对没有使用的导入包会编译报错，但是有时我们只想调用该包的init函数，不使用包导出的变量或者方法，这时就采用下划线导入方案。
*/
	_ "github.com/cellnet/peer/tcp"
	_ "github.com/cellnet/proc/tcp"
)

var log = golog.New("client")
var pFile = golog.Create_MyWrite_File("./clientlog.txt")
var logex = golog.New_ex("myclient",pFile)

func ReadConsole(index int,callback func(string)) {

	/*for {

		// 从标准输入读取字符串，以\n为分割
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			break
		}

		// 去掉读入内容的空白符
		text = strings.TrimSpace(text)

		callback(text)

	}*/

	for {
		sendinfo := strconv.Itoa(index)
		sendinfo += " send data is: " + time.Now().String()
		callback(sendinfo)
		log.Infof(sendinfo)
		time.Sleep(60 * time.Second)
	}

}

func main() {
	ch := make(chan int)

	for i := 0; i< 100; i++ {
		logex.Infof("i=======%d",i)
		// 创建一个事件处理队列，整个客户端只有这一个队列处理事件，客户端属于单线程模型
		queue := cellnet.NewEventQueue()

		// 创建一个tcp的连接器，名称为client，连接地址为127.0.0.1:8801，将事件投递到queue队列,单线程的处理（收发封包过程是多线程）
		p := peer.NewGenericPeer("tcp.Connector", "client", "127.0.0.1:18801", queue)

		// 设定封包收发处理的模式为tcp的ltv(Length-Type-Value), Length为封包大小，Type为消息ID，Value为消息内容
		// 并使用switch处理收到的消息
		proc.BindProcessorHandler(p, "tcp.ltv", func(ev cellnet.Event) {
			switch msg := ev.Message().(type) {
			case *cellnet.SessionConnected:
				log.Debugln("client connected")
			case *cellnet.SessionClosed:
				log.Debugln("client error")
			case *proto.ChatACK:
				log.Infof("sid%d say: %s", msg.Id, msg.Content)
			}
		})

		// 开始发起到服务器的连接
		p.Start()

		// 事件队列开始循环
		queue.StartLoop()

		// 阻塞的从命令行获取聊天输入
		/*1:ReadConsole中的参数属于闭包函数
		2:p变量是tcpConnector类实例
		3:tcpConnector.Session()返回 已经实现cellnet.Session 接口类型的 defaultSes *tcpSession,最终调用tcpSession的Send方法
		相似的用法见tcp\session.go newSession 函数
		*/
		go ReadConsole(i,func(str string) {

			//临时转换一个接口并调用其方法
			p.(interface {
				Session() cellnet.Session
			}).Session().Send(&proto.ChatREQ{
				Content: str,
			})

		})
	}

	<- ch
	pFile.CloseFile()
}
