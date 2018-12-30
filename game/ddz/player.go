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

type Previouser struct{
	OperType uint32       //操作类型  过牌 出牌
	Point uint32 			//出牌人坐标
	Cardtype CARD_TYPE 		//出牌类型
	CardValue uint32 	//出牌内容
	Extra uint32 //配合出牌内容组成有意义的信息 比如 3带1
}

type Gamer struct{
	IsMain bool //true:地主
	Point uint32
	Cardlist []uint32
	Status int
	NowIsLead bool  // 选地主阶段的flag 当前是否活跃状态 用途:防止该玩家的ExitSync.Done()调用2次导致崩溃(玩家返回结果时间晚于超时处理后)

	IsDoit bool //出牌阶段的flag 用途:防止该玩家的ExitSync.Done()调用2次导致崩溃(玩家返回结果时间晚于超时处理后)

	Parse Parser //拆牌处理

	Pre *Previouser  //上一家出的牌
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

//首次出牌策略 先出单张 然后对子 第一次出牌 绝对有这两种牌 就这就够了
func (self* DDZ_Player) easy_play() *Previouser {

	if len(self.Game.Parse.SingleCards) != 0{
		preoper := &Previouser{
			OperType:uint32(ddz.OperationType_doit),
			Point:self.Game.Point,
			Cardtype:SINGLE_TYPE,
			CardValue:uint32(self.Game.Parse.SingleCards[0]),
			Extra:0,
		}
		//保存到上家出牌内容结构体中
		self.Game.Pre = preoper

		//移除出过的牌
		self.Game.Parse.SingleCards = self.Game.Parse.SingleCards[1:]

		return preoper

	}else if len(self.Game.Parse.DoubleCards) != 0 {
		preoper := &Previouser{
			OperType:uint32(ddz.OperationType_doit),
			Point:self.Game.Point,
			Cardtype:DOUBLE_TYPE,
			CardValue:uint32(self.Game.Parse.DoubleCards[0]),
			Extra:0,
		}

		//保存到上家出牌内容结构体中
		self.Game.Pre = preoper

		//移除出过的牌
		self.Game.Parse.DoubleCards = self.Game.Parse.DoubleCards[1:]

		return preoper
	}

	return nil
}

//根据上家出牌内容 分析出合适的牌

func (self* DDZ_Player) judge()  *Previouser{

	switch self.Game.Pre.Cardtype{

	case ROCKET_TYPE://火箭无人能敌 要不起返回过牌
		preoper := &Previouser{
			OperType:uint32(ddz.OperationType_pass),
			Point:self.Game.Point,
		}
		//保存到上家出牌内容结构体中
		self.Game.Pre = preoper

		//移除出过的牌 nothing

		return preoper

	case ZHADAN_TYPE://分析自己手中是否有大于炸弹的牌
		bigercard := -1;index := -1
		for i,v := range self.Game.Parse.Bombes{
			if v > int(self.Game.Pre.CardValue) {
				bigercard = v
				index = i
				break
			}
		}

		if index == -1{
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_pass),
				Point:self.Game.Point,
			}
			self.Game.Pre = preoper
			return preoper
		}else{
			//出过的牌移除之
			self.Game.Parse.Bombes = append(self.Game.Parse.Bombes[0:index],self.Game.Parse.Bombes[index+1:]...)

			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_doit),
				Point:self.Game.Point,
				Cardtype:ZHADAN_TYPE,
				CardValue:uint32(bigercard),
				Extra:0,
			}
			self.Game.Pre = preoper
			return preoper
		}


	case DOULEL_INK_TYPE://这个类型下 CardValue 标示双链开头值 Extra标示双链长度
		bigcard := -1
		bigindex := -1
		biglen := -1
		for k,v := range self.Game.Parse.DoubleLink{
			if v.BeginValue > int(self.Game.Pre.CardValue) && v.Len == int(self.Game.Pre.Extra){
				bigindex = k
				bigcard = v.BeginValue
				biglen = v.Len
				break
			}
		}
		if bigindex == -1{
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_pass),
				Point:self.Game.Point,
			}
			self.Game.Pre = preoper
			return preoper
		}else{
			//出过的牌移除之
			self.Game.Parse.DoubleLink = append(self.Game.Parse.DoubleLink[0:bigindex],self.Game.Parse.DoubleLink[bigindex+1:]...)
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_doit),
				Point:self.Game.Point,
				Cardtype:DOULEL_INK_TYPE,
				CardValue:uint32(bigcard),
				Extra:uint32(biglen),
			}
			self.Game.Pre = preoper
			return preoper
		}
	case SINGLE_LINK_TYPE:
		bigcard := -1
		bigindex := -1
		biglen := -1
		for k,v := range self.Game.Parse.SingleLink{
			if v.BeginValue > int(self.Game.Pre.CardValue) && v.Len == int(self.Game.Pre.Extra){
				bigindex = k
				bigcard = v.BeginValue
				biglen = v.Len
				break
			}
		}
		if bigindex == -1{
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_pass),
				Point:self.Game.Point,
			}
			self.Game.Pre = preoper
			return preoper
		}else{
			//出过的牌移除之
			self.Game.Parse.SingleLink = append(self.Game.Parse.SingleLink[0:bigindex],self.Game.Parse.SingleLink[bigindex+1:]...)
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_doit),
				Point:self.Game.Point,
				Cardtype:SINGLE_LINK_TYPE,
				CardValue:uint32(bigcard),
				Extra:uint32(biglen),
			}
			self.Game.Pre = preoper
			return preoper
		}

	case THREE_ONE_TYPE:
		bigcard := -1
		singlecard :=-1
		bigindex := -1
		for k,v := range self.Game.Parse.ThreeOneCards{
			if v.ThreeCard > int(self.Game.Pre.CardValue){
				bigindex = k
				bigcard = v.ThreeCard
				singlecard = v.SingleCard
				break
			}
		}
		if bigindex == -1{
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_pass),
				Point:self.Game.Point,
			}
			self.Game.Pre = preoper
			return preoper
		}else{
			//出过的牌移除之
			self.Game.Parse.ThreeOneCards = append(self.Game.Parse.ThreeOneCards[0:bigindex],self.Game.Parse.ThreeOneCards[bigindex+1:]...)
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_doit),
				Point:self.Game.Point,
				Cardtype:THREE_ONE_TYPE,
				CardValue:uint32(bigcard),
				Extra:uint32(singlecard),
			}
			self.Game.Pre = preoper
			return preoper
		}

	case THREE_TYPE:
		bigcard := -1
		bigindex := -1
		for k,v := range self.Game.Parse.ThreeCards{
			if v.ThreeCard > int(self.Game.Pre.CardValue){
				bigindex = k
				bigcard = v.ThreeCard
				break
			}
		}
		if bigindex == -1{
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_pass),
				Point:self.Game.Point,
			}
			self.Game.Pre = preoper
			return preoper
		}else{
			//出过的牌移除之
			self.Game.Parse.ThreeCards = append(self.Game.Parse.ThreeCards[0:bigindex],self.Game.Parse.ThreeCards[bigindex+1:]...)
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_doit),
				Point:self.Game.Point,
				Cardtype:THREE_TYPE,
				CardValue:uint32(bigcard),
				Extra:0,
			}
			self.Game.Pre = preoper
			return preoper
		}
	case DOUBLE_TYPE:
		bigcard :=  -1
		bigindex := -1
		for k,v := range self.Game.Parse.DoubleCards{
			if v > int(self.Game.Pre.CardValue){
				bigindex = k
				bigcard = v
				break
			}
		}
		if bigindex == -1{
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_pass),
				Point:self.Game.Point,
			}
			self.Game.Pre = preoper
			return preoper
		}else{
			//出过的牌移除之
			self.Game.Parse.DoubleCards = append(self.Game.Parse.DoubleCards[0:bigindex],self.Game.Parse.DoubleCards[bigindex+1:]...)
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_doit),
				Point:self.Game.Point,
				Cardtype:DOUBLE_TYPE,
				CardValue:uint32(bigcard),
				Extra:0,
			}
			self.Game.Pre = preoper
			return preoper
		}
	case SINGLE_TYPE:
		bigcard :=  -1
		bigindex := -1
		for k,v := range self.Game.Parse.SingleCards{
			if v > int(self.Game.Pre.CardValue){
				bigindex = k
				bigcard = v
				break
			}
		}
		if bigindex == -1{
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_pass),
				Point:self.Game.Point,
			}
			self.Game.Pre = preoper
			return preoper
		}else{
			//出过的牌移除之
			self.Game.Parse.SingleCards = append(self.Game.Parse.SingleCards[0:bigindex],self.Game.Parse.SingleCards[bigindex+1:]...)
			preoper := &Previouser{
				OperType:uint32(ddz.OperationType_doit),
				Point:self.Game.Point,
				Cardtype:SINGLE_TYPE,
				CardValue:uint32(bigcard),
				Extra:0,
			}
			self.Game.Pre = preoper
			return preoper
		}
	}
	return nil
}

func (self* DDZ_Player) PlayCard() *Previouser {
	//分析是否开局自己先出第一手牌 Cardtype为0标示这局刚开始 3个玩家都还没出过牌
	if self.Game.Pre.Cardtype == 0{
		return self.easy_play()
	}else{
		return self.judge()
	}
	//分析是否自己出牌 另外2人要不起 又轮到自己出牌
	return nil
}

