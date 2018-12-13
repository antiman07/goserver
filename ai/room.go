package ai

import (
	"trunk/cellnet/util"
)

//var RoomMager RoomManager

type RoomInfo struct{
	ID int
	Name string //初级场 中级场 高级场 至尊场
	Dizhu float32 //底注 0.1,1,2,3
	MinCoin int //准入金币 10，50,100
	MaxDeskNum int //最多有多少桌
	DeskQueue util.Queue
}

//4个房间不放到列表中统一管理 便于开4个GO程分别处理,免去加锁操作
var Room1,Room2,Room3,Room4 *RoomInfo

/*type RoomManager struct{
	RoomList []*RoomInfo
}*/

func get_free_desk_from_room1() (interface{}) {

	return Room1.DeskQueue.Dequeue()
}

func put_free_desk_to_room1(data interface{}){

	Room1.DeskQueue.Enqueue(data)
}


func load_room_info(){
	room1 := &RoomInfo{
		ID:100,
		Name:"初级场",
		Dizhu:0.1,
		MinCoin:10,
		MaxDeskNum:1000,
	}

	Room1 = room1

	for index := 0; index<room1.MaxDeskNum; index++ {
		desk := &Desk{
			IsGameing:false,
			ID:index,
			CurPlay:"",
			PeopleList:[]*People{
				&People{},
				&People{},
				&People{}},
		}
		room1.DeskQueue.Enqueue(desk)
	}
}

func init()  {
	load_room_info()
}
