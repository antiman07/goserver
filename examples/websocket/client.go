package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
	"trunk/cellnet"
	_ "trunk/cellnet/codec/gogopb"
	_ "trunk/cellnet/codec/json"
	"trunk/cellnet/game"
	"trunk/cellnet/game/jdzjh"
	"trunk/cellnet/game/suoha"
	"trunk/cellnet/game/tbnn"
	"trunk/cellnet/myrpc"
	pb_jdzjh "trunk/cellnet/pb/jdzjh"
	pb_suoha "trunk/cellnet/pb/suoha"
	pb_tbnn "trunk/cellnet/pb/tbnn"
	pb_ddz "trunk/cellnet/pb/ddz"
	"trunk/cellnet/peer"
	"trunk/cellnet/peer/gorillaws"
	_ "trunk/cellnet/peer/gorillaws"
	"trunk/cellnet/proc"
	_ "trunk/cellnet/proc/gorillaws"
	"trunk/cellnet/util"
)

func ReadBackProcessExitMsg(){
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)
	signal.Notify(sigs,syscall.SIGTERM,syscall.SIGKILL)
	go func(){
		sig := <- sigs
		myrpc.Rpcqueue <- "后台进程接收到退出信号,开始释放资源,准备停止运行..."
		myrpc.Rpcqueue <- fmt.Sprint(sig)
		done <- true
	}()

	myrpc.Rpcqueue <- "监听kill信号中..."
	<- done

	game.RealWorkRobots = gorillaws.GsesManager.Count() //退出之前统计在线玩家数量
	game.Goover = "exit" //设置状态 开始退出进程工作 具体见robot函数
}

func ReadConsole(){
	for{
		text,err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil{
			break
		}
		//输入"exit" 开始退出进程工作 具体见robot函数
		text = strings.TrimSpace(text)
		if text == "exit"{
			game.RealWorkRobots = gorillaws.GsesManager.Count() //退出之前统计在线玩家数量
			game.Goover = text
		}
	}
}

func robot(sig *sync.WaitGroup){
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("websocket.Connector", "client", game.GameAddr, queue)

	if it, ok := game.RegGameLogicFunction[game.GameName]; ok {
		proc.BindProcessorHandler(p,"gorillaws.ltv",it)
	}else{
		panic("game logic process function not found:" + game.GameName)
	}

	p.Start()
	queue.StartLoop()

	for{
		if game.Goover == "exit" {
			p.Stop() //这个函数会阻塞等待 即是 websocketConnector Stop函数
			//先给session的发送GO程队列喂药(即 self.sendQueue.Add(nil)),发送GO程吃到药退出之前调用self.cleanup(),这样会立即导致
			//session的接收GO程退出,收发GO程都退出后 会导致session.go中 self.endNotify()调用,进而导致connector.go中的
			// self.sesEndSignal.Wait()不再阻塞,继续执行代码到 self.EndStopping() 导致 websocketConnector中 Stop()
			// 函数中的 self.WaitStopFinished()执行。到此p.Stop()这行代码执行完毕。

			queue.StopLoop() //执行到这一行,给队列喂药 设置退出标示,喂药之前队列内已存在的事件会处理完毕,
			//同一时刻 session的接收GO程会接收网络事件继续投递到该事件循环GO程队列中,但这些网络事件是不会被处理的 (见queue.go)
			queue.Wait()//这个函数会阻塞等待,执行到喂药标示后 退出事件循环GO程

			//调用上一行的 queue.Wait()返回后,事件循环GO程已经停止工作,虽然此刻session的接收GO程还会接收网络事件
			// 继续投递到该事件循环GO程队列中,但不会得到执行了。


			sig.Done()

			break
		}else{
			time.Sleep(time.Second*3)
		}

	}

	fmt.Println("player quit")
}


func parse_cmd_params()(){
	//通过命令行方式改变robots roomid 等变量的值。命令行输入方式 websocketclient.exe -b 5 -r 2 -s 3 -t 4 -p 4
	//使用flag来操作命令行参数，支持的格式如下： -b=1   --b=1  -b 1  --b 1

	//第一个参数表示参数名称，第二个参数表示默认值，第三个参数表示使用说明和描述
	tmp_runmode := flag.Int("rm", 1, "进程启动模式 1:前台运行模式 2:后台运行模式")
	tmp_robots := flag.Int("b", 3, "choose robot numbers")
	tmp_accid := flag.Int("i",10000,"批量开启机器人时 第一个账号ID")
	tmp_roomid := flag.Int("r", 1, "choose roomid")
	tmp_betmulti := flag.Int("s", 1, "choose betmulti")
	tmp_betmulti_time := flag.Int("t", 5, "choose betmulti_time")
	tmp_spreadcard_time := flag.Int("p", 2, "choose spreadcard_time")
	tmp_which_game := flag.String("game","ddz","choose game to test")

	//tmp_ipaddr := flag.String("addr","10.200.124.133:11302","choose game server ip port like as 10.200.124.133:6302")
	tmp_ipaddr := flag.String("addr","192.168.10.10:6302","choose game server ip port like as 192.168.200.54:10303")

	//tmp_rpclogaddr := flag.String("rl","10.200.124.122:1234","log server ip port input like as 10.200.124.122:1234")
	tmp_rpclogaddr := flag.String("rl","192.168.10.10:1234","log server ip port input like as 10.200.124.122:1234")

	tmp_sleep := flag.Int("sleep",10,"每个机器人启动间隔时长")
	tmp_suspend := flag.Int("st", 5, "轮到自己出牌时暂停的时长 单位是秒")
	flag.Parse()

	game.Runmode        = *tmp_runmode
	game.Robots 		= *tmp_robots
	game.BeginAccid 	= *tmp_accid
	game.RoomID 		= *tmp_roomid
	game.Betmulti 		= *tmp_betmulti
	game.BetmultiTime 	= *tmp_betmulti_time
	game.SpreadcardTime = *tmp_spreadcard_time
	game.GameName 		= *tmp_which_game
	game.GameAddr 		= *tmp_ipaddr
	game.LogserverAddr 	= *tmp_rpclogaddr
	game.SpaceTime      = *tmp_sleep
	game.Suspend        = *tmp_suspend
	fmt.Printf("游戏名称=%s,服务器IP=%s,日志服IP=%s,机器人数量=%d,运行模式=%d,并发启动机器人间隔时长=%d\n",
				game.GameName,game.GameAddr,game.LogserverAddr,game.Robots,game.Runmode,game.SpaceTime)
}


func main() {

	begin := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	parse_cmd_params()

	switch game.GameName {
	case "jdzjh":
		pb_jdzjh.NotWriteInit()
		jdzjh.Init_game_protocol()
	case "suoha":
		pb_suoha.NotWriteInit()
		suoha.Init_game_protocol()
	case "tbnn":
		tbnn.Init_game_protocol()
		pb_tbnn.NotWriteInit()
	case "ddz":
		fmt.Println("斗地主前后端自己定义协议 不需要转换了,但这里不要让系统自动调用Init函数,因为有多个游戏,容易引起混乱")
		pb_ddz.NotWriteInit()
		//Init_game_protocol  这款游戏前后端自己写 不需要调用协议转换这个函数了
	default:
		panic("choose game first")
	}


	util.SetFrontRobotID(int64(game.BeginAccid))

	//服务端代码 type 填写1
	go myrpc.Rpclog(game.LogserverAddr,2)

	var sig sync.WaitGroup

	for i := 0; i< game.Robots; i++{
		sig.Add(1)
		go robot(&sig) //坑,一定要传引用 见https://liudanking.com/golang/golang-waitgroup-usage/
		time.Sleep(time.Duration(game.SpaceTime)*time.Millisecond)
	}

	switch game.Runmode {
	case 1:					//前台进程启动模式,监控屏幕输入退出字符串
		go ReadConsole()
	case 2:					//后台进程启动模式,监控KILL信号 不能用kill -9,不然进程接收不到信号
		go ReadBackProcessExitMsg()
	default:
		go ReadBackProcessExitMsg()
	}

	myrpc.Rpcqueue <- "主GO程等待退出信号中..."

	sig.Wait()

	switch game.GameName{
	case "suoha":
		suoha.PrintPlayerOperator(begin,time.Now())
	case "jdzjh":
		jdzjh.PrintPlayerOperator(begin,time.Now())
	}

	myrpc.Rpcqueue <- "game over"
	//给予足够的时间等待GO程们都退出(正常退出情况流程中,发送数据GO程早于接收数据GO程退出)
	time.Sleep(time.Second*15)

}