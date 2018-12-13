package suoha

import (
	"trunk/cellnet/pb/suoha"
)

var Suoha_RoomInfoList []*suoha.RoomInfo

type Suoha_Player struct{
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
}

func (self* Suoha_Player) PlayerName() string{
	return "suoha"
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