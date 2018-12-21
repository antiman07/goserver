package ddz

import (
	"sync"
	"trunk/cellnet/pb/ddz"
)

type Attr struct{
	UserID string
	Nickname string
	Chips string
	Level string
	Avatar string
	Sex 	string
}

const (
	Type_matching int = 0
	Type_matched  int = 1
)
type Gamer struct{
	IsMain bool //true:地主
	Point uint32
	Cardlist []uint32
	Status int

	NowIsLead bool  //当前是否活跃状态 用途:防止该玩家的ExitSync.Done()调用2次导致崩溃(玩家返回结果时间晚于超时处理后)
}

var DDZ_RoomInfoList []*ddz.RoomInfo

type DDZ_Player struct{
	Att Attr
	Game Gamer

	ExitSync sync.WaitGroup

	Room int
}

func (self* DDZ_Player) PlayerName() string{
	return "ddz"
}



