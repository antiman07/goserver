package tbnn

import (
	"fmt"
	"sync"
	"time"
	"trunk/cellnet/myrpc"
)

var gametimes int32			//成功进入游戏次数
var gameFailtimes int32		//进入游戏失败次数
var gamePasstimes int32		//过牌次数
var discardtimes int32		//弃牌次数
var cingltimes int32			//跟注次数
var raisetimes int32			//加注次数
var Showhandtimes int32	   //梭哈次数

var calcGuard sync.Mutex

func PrintPlayerOperator(begin time.Time,end time.Time){

	runtimes := end.Sub(begin)

	myrpc.Rpcqueue <- fmt.Sprint("运行时长:",runtimes)

	myrpc.Rpcqueue <- fmt.Sprint("成功进入游戏次数:",gametimes)

	myrpc.Rpcqueue <- fmt.Sprint("进入游戏失败次数:",gameFailtimes)

	myrpc.Rpcqueue <- fmt.Sprint("过牌次数:",gamePasstimes)

	myrpc.Rpcqueue <- fmt.Sprint("弃牌次数:",discardtimes)

	myrpc.Rpcqueue <- fmt.Sprint("跟注次数:",cingltimes)

	myrpc.Rpcqueue <- fmt.Sprint("加注次数:",raisetimes)

	myrpc.Rpcqueue <- fmt.Sprint("梭哈次数:",Showhandtimes)
}

//统计所有机器人做的各种操作次数

func IncreAllPlayerData(p *Tbnn_Player){

	calcGuard.Lock()
	defer calcGuard.Unlock()

	gametimes += p.Gametimes
	gameFailtimes += p.GameFailtimes
	gamePasstimes += p.Passtimes
	discardtimes += p.Discardtimes
	cingltimes += p.Cingltimes
	raisetimes += p.Raisetimes
	raisetimes += p.Raisetimes
	Showhandtimes += p.Showhandtimes
}

