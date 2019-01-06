package suoha

import (
	"sync"
	"trunk/cellnet/pb/suoha"
)

var Suoha_RoomInfoList []*suoha.RoomInfo

type RoomOperTimes struct{
	Passtimes int32         //该玩家过牌次数
	Discardtimes int32		//该玩家弃牌次数
	Cingltimes int32		//该玩家跟注次数
	Raisetimes int32		//该玩家加注次数
}

type Suoha_Player struct{
	Room int //当前进入的房间号
	Att Attr
	Pos uint32 //玩家的座位点
	Previous suoha.PushPosOperation//最新出牌玩家的数据

	//统计玩家行为数据
	Gametimes int32			//该玩家成功进入游戏次数
	GameFailtimes int32		//该玩家进入游戏失败次数
	Passtimes     int32     //该玩家过牌次数
	Discardtimes int32		//该玩家弃牌次数
	Cingltimes int32		//该玩家跟注次数
	Raisetimes int32		//该玩家加注次数
	Showhandtimes int32		//该玩家梭哈次数

	Room1times int32        		//进入room1次数
	Room2times int32        		//进入room2次数
	Room3times int32        		//进入room3次数
	Room4times int32        		//进入room4次数
	RoomOperList [5]RoomOperTimes   //该玩家在每个房间 弃牌 跟注 加注次数统计 第一个0节点不用 只用 1 2 3 4
	AllRequestTimes int64           //玩家请求次数
}

func (self* Suoha_Player) PlayerName() string{
	return "suoha"
}

var ccGuard sync.Mutex

//有2个GO程访问这个函数  心跳GO程 和 事件处理GO程 所以加锁
func (self* Suoha_Player) IncReqTimes()  {

	ccGuard.Lock()
	defer ccGuard.Unlock()

	self.AllRequestTimes += 1
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


func (self* Suoha_Player) CalcTimes(type1 suoha.OperationType)  {
	switch type1{
	case suoha.OperationType_pass: self.Passtimes += 1
	case suoha.OperationType_discard:self.Discardtimes += 1
	case suoha.OperationType_cingl:self.Cingltimes += 1
	case suoha.OperationType_raise:self.Raisetimes += 1
	case suoha.OperationType_showhand:self.Showhandtimes += 1
	case 6:		self.Gametimes += 1
	case 7:		self.GameFailtimes += 1
	}
}

func (self* Suoha_Player) CalcRooms(which uint32){
	switch which {
	case 1:self.Room1times += 1
	case 2:self.Room2times += 1
	case 3:self.Room3times += 1
	case 4:self.Room4times += 1
	}
}

//从数据库中加载上线玩家数据 begin
/*type PlayerAttr struct{
	UserId 			int32
	AcctName 		string
}

func (self* PlayerAttr) LoadData(name interface{}) interface{}{
	rows,err := db.GDB.Query("select * from t_user where acct_name= ?",name)
	defer rows.Close()
	if err != nil{
		panic(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&self.AcctName,&self.UserId)
		if err != nil {
			panic(err.Error())
		}
	}
	return self
}

func (self* PlayerAttr) SaveData() interface{}{
	_,err := db.GDB.Exec("replace into t_user(user_id,acct_name) values(?,?)",self.UserId,self.AcctName)
	if err != nil{
		panic(err.Error())
	}
	return self
}

type PlayerMail struct {
	userId 			int32
	mailId			int32
}

func (self* PlayerMail) LoadData(name interface{}) interface{}{
	return self
}
func (self* PlayerMail) SaveData()interface{} {
	return nil
}

type PlayerTask struct {
	userId 			int32
	taskId			int32
}

func (self* PlayerTask) LoadData(name interface{}) interface{}{
	return self
}
func (self* PlayerTask) SaveData()interface{} {
	return nil
}

type DBPlayer struct{
	Attr  *PlayerAttr
	Mail  *PlayerMail
	Task  *PlayerTask
}

func (self* DBPlayer) CreatePlayer(name interface{}){
	attr := new(PlayerAttr)
	attr.LoadData(name)

	mail := new(PlayerMail)
	mail.LoadData(name)

	task := new(PlayerTask)
	task.LoadData(name)
}*/
//从数据库中加载上线玩家数据 end