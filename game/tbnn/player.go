package tbnn

import (
	"trunk/cellnet/pb/tbnn"
)

var Tbnn_RoomInfoList []*tbnn.RoomInfo

type Tbnn_Player struct{

	Att Attr

	EnterRoomPushList []*tbnn.RoomRoleInfo

	MutilListPushList []uint32

	BetMutilPushList []*tbnn.BetMultiPush

	DealCardRespInfo []*tbnn.PokerInfo

	SpreadCardPushList []*tbnn.SpreadInfo

	SettleChipPushList []*tbnn.SettleInfo

	//统计玩家行为数据
	Gametimes int32			//该玩家成功进入游戏次数
	GameFailtimes int32		//该玩家进入游戏失败次数
	Passtimes     int32     //该玩家过牌次数
	Discardtimes int32		//该玩家弃牌次数
	Cingltimes int32		//该玩家跟注次数
	Raisetimes int32		//该玩家加注次数
	Showhandtimes int32		//该玩家梭哈次数
}

type Attr struct{
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

func (self* Tbnn_Player) PlayerName() string{
	return "tongbiniuniu"
}
/*
func (self* Tbnn_Player) CalcTimes(type1 tbnn.OperationType)  {
	switch type1{
	case suoha.OperationType_pass: self.Passtimes += 1
	case suoha.OperationType_discard:self.Discardtimes += 1
	case suoha.OperationType_cingl:self.Cingltimes += 1
	case suoha.OperationType_raise:self.Raisetimes += 1
	case suoha.OperationType_showhand:self.Showhandtimes += 1
	case 6:		self.Gametimes += 1
	case 7:		self.GameFailtimes += 1
	}
}*/