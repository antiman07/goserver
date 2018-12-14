package jdzjh

import (
	"sync"
	"trunk/cellnet/pb/jdzjh"
)

var Zjh_RoomInfoList []*jdzjh.RoomInfo


const(
	InsertGameSuccess 		= 0 //成功进入游戏操作
	InsertGameFailed 		= 1	//进入游戏失败操作
	DiscardOperator 		= 2	//弃牌操作
	CinglOperator			= 3	//跟注操作
	RaisetimesOperator 		= 4	//加注操作
)
const(
	NoMeaning = 0
)
const(
	Room1 = 1
	Room2 = 2
	Room3 = 3
	Room4 = 4
)


type RoomOperTimes struct{
	Discardtimes int32		//该玩家弃牌次数
	Cingltimes int32		//该玩家跟注次数
	Raisetimes int32		//该玩家加注次数
}

type Zjh_Player struct{

	Att Atter
	Pos uint32 //玩家的座位点
	Room int //当前进入的房间号
	//Previous jdzjh.PushPosOperation//最新出牌玩家的数据
	//前一个出牌玩家信息 begin
	PreType     uint32  //弃牌 跟注 加注
	PrePos		uint64  //玩家座位号
	PreValueArry[] uint64
	//前一个出牌玩家信息 end

	//统计玩家行为数据
	Gametimes int32					//该玩家成功进入游戏次数
	GameFailtimes int32				//该玩家进入游戏失败次数
	Discardtimes int32				//该玩家弃牌次数
	Cingltimes int32				//该玩家跟注次数
	Raisetimes int32				//该玩家加注次数
	Room1times int32        		//进入room1次数
	Room2times int32        		//进入room2次数
	Room3times int32        		//进入room3次数
	Room4times int32        		//进入room4次数
	RoomOperList [5]RoomOperTimes   //该玩家在每个房间 弃牌 跟注 加注次数统计 第一个0节点不用 只用 1 2 3 4
	AllRequestTimes int64           //玩家请求次数
}

type Atter struct{
	Account string
	Agentid int64
	Ownerid int64
	Chips uint64
	Nickname string
	Level uint32
	Avatar uint32
	Sex int32
	Isnew bool
	Ismainwallet bool
	Mainwallet uint64
	Autowallet bool
}


func (self* Zjh_Player) PlayerName() string{
	return "jdzjh"
}

var ccGuard sync.Mutex

//有2个GO程访问这个函数  心跳GO程 和 事件处理GO程 所以加锁
func (self* Zjh_Player) IncReqTimes()  {

	ccGuard.Lock()
	defer ccGuard.Unlock()

	self.AllRequestTimes += 1
}

func (self* Zjh_Player) CalcTimes(type1 uint32)  {
	switch type1{
	case 0:	self.Gametimes += 1
	case 1:	self.GameFailtimes += 1
	case 2: self.Discardtimes += 1
	case 3:	self.Cingltimes += 1
	case 4:	self.Raisetimes += 1
	}
}

func (self* Zjh_Player) CalcRooms(which uint32){
	switch which {
	case 1:self.Room1times += 1
	case 2:self.Room2times += 1
	case 3:self.Room3times += 1
	case 4:self.Room4times += 1
	}
}