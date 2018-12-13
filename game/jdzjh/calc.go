package jdzjh

import (
	"fmt"
	"sync"
	"time"
	"trunk/cellnet/game"
	"trunk/cellnet/myrpc"
)

var gametimes int32			//成功进入游戏次数
var gameFailtimes int32		//进入游戏失败次数
var discardtimes int32		//弃牌次数
var cingltimes int32			//跟注次数
var raisetimes int32			//加注次数

var room1times int32        //进入room1次数
var room2times int32        //进入room2次数
var room3times int32        //进入room3次数
var room4times int32        //进入room4次数
var roomOperList [5]RoomOperTimes   //所有玩家在每个房间 弃牌 跟注 加注次数统计 第一个0节点不用 只用 1 2 3 4

var calcGuard sync.Mutex

func PrintPlayerOperator(begin time.Time,end time.Time){

	runtimes := end.Sub(begin)
	myrpc.Rpcqueue <- fmt.Sprint("运行时长:",runtimes)
	myrpc.Rpcqueue <- fmt.Sprint("启动机器人:",game.Robots)
	myrpc.Rpcqueue <- fmt.Sprint("在线干活机器人:",game.RealWorkRobots)
	myrpc.Rpcqueue <- fmt.Sprint("成功进入游戏次数:",gametimes)
	myrpc.Rpcqueue <- fmt.Sprint("进入游戏失败次数:",gameFailtimes)
	myrpc.Rpcqueue <- fmt.Sprint("弃牌次数:",discardtimes)
	myrpc.Rpcqueue <- fmt.Sprint("跟注次数:",cingltimes)
	myrpc.Rpcqueue <- fmt.Sprint("加注次数:",raisetimes)
	myrpc.Rpcqueue <- fmt.Sprint("在房间1游戏次数:",room1times)
	myrpc.Rpcqueue <- fmt.Sprint("在房间2游戏次数:",room2times)
	myrpc.Rpcqueue <- fmt.Sprint("在房间3游戏次数:",room3times)
	myrpc.Rpcqueue <- fmt.Sprint("在房间4游戏次数:",room4times)

	for i := 1; i <= 4; i++{
		myrpc.Rpcqueue <- fmt.Sprintf("房间 %d 弃牌次数 %d 跟注次数 %d 加注次数 %d",i,
			roomOperList[i].Discardtimes,roomOperList[i].Cingltimes,roomOperList[i].Raisetimes)
	}
}

//统计所有机器人做的各种操作次数
func IncreAllPlayerData(p *Zjh_Player){

	calcGuard.Lock()
	defer calcGuard.Unlock()

	gametimes += p.Gametimes
	gameFailtimes += p.GameFailtimes
	discardtimes += p.Discardtimes
	cingltimes += p.Cingltimes
	raisetimes += p.Raisetimes

	room1times += p.Room1times
	room2times += p.Room2times
	room3times += p.Room3times
	room4times += p.Room4times

	for i := 1; i <= 4; i++{
		roomOperList[i].Cingltimes 	 += p.RoomOperList[i].Cingltimes
		roomOperList[i].Discardtimes += p.RoomOperList[i].Discardtimes
		roomOperList[i].Raisetimes   += p.RoomOperList[i].Raisetimes
	}

}

