package ai

import (
	"fmt"
	"math/rand"
	"time"
	"trunk/cellnet/game/ddz"
	pbddz "trunk/cellnet/pb/ddz"
	"trunk/cellnet/myrpc"
	"trunk/cellnet/peer/gorillaws"
)

var Paidata = [5][13]int {
	{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d},//方块(Diamond)
	{0x11,0x12,0x13,0x14,0x15,0x16,0x17,0x18,0x19,0x1a,0x1b,0x1c,0x1d},//梅花(Club)
	{0x21,0x22,0x23,0x24,0x25,0x26,0x27,0x28,0x29,0x2a,0x2b,0x2c,0x2d},//红桃(Heart)
	{0x31,0x32,0x33,0x34,0x35,0x36,0x37,0x38,0x39,0x3a,0x3b,0x3c,0x3d},//黑桃(Spade)
	{0x4E,0x4F}, //王(Joker)
}

var paidata = []uint32{
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	1,2,3,4,5,6,7,8,9,10,11,12,13,
	14,15,
}

func fapai() (map[int][]uint32,[]uint32) {
	flag := [54]int{}
	rand.Seed(time.Now().Unix())

	carmap := make(map[int][]uint32) //发牌列表(每人17张) 17 * 3
	var rest []uint32				//3张余牌
	for i := 0; i < 3; i++ {
		node := make([]uint32, 0)
		for index := 0; index < 17; index++ {
			m := rand.Intn(54)
			if flag[ m ] == 1{
				index--
				continue
			}

			flag[m] = 1
			node = append(node, paidata[m])
		}
		//sort.Ints(rst)
		carmap[i] = node
	}
	for k,_ := range flag{
		if flag[k] == 0{
			//fmt.Printf("底牌是 %d\n",paidata[k])
			rest = append(rest,paidata[k])
		}
	}
	//fmt.Println(flag)
	return carmap,rest
}

func firstcard(s1 interface{},s2 interface{},s3 interface{}){

	var list []*gorillaws.WsSession
	ses1 := s1.(*gorillaws.WsSession)
	ses2 := s2.(*gorillaws.WsSession)
	ses3 := s3.(*gorillaws.WsSession)

	player1 := ses1.Player().(*ddz.DDZ_Player)
	player2 := ses2.Player().(*ddz.DDZ_Player)
	player3 := ses3.Player().(*ddz.DDZ_Player)

	fmt.Printf("firstcard %s %s %s",player1.Att.UserID,player2.Att.UserID,player3.Att.UserID)
	list = append(list, ses1,ses2,ses3)

	pailist,restlist := fapai()
	fmt.Println(pailist,restlist)

	var selfpos uint32
	for it,ses := range list{
		myrpc.Rpcqueue <- "PushRoomInfo"
		selfpos = uint32(it + 1)//玩家座位号从 1 开始往后排
		ses.Player().(*ddz.DDZ_Player).Game.Point = selfpos
		ses.Player().(*ddz.DDZ_Player).Game.Cardlist = pailist[it]
		ses.Player().(*ddz.DDZ_Player).Game.Status = ddz.Type_matched

		ses.Send(&pbddz.PushRoomInfo{
			SelfPos:selfpos,
			Seats:&pbddz.SeatInfo{
				Pos:selfpos,
				Cards:pailist[it],
			},
		})
	}
}

//产生地主 根据进入房间的次序 挨个询问玩家
func find_landowner(s1 interface{},s2 interface{},s3 interface{}) bool{

	var list []*gorillaws.WsSession
	ses1 := s1.(*gorillaws.WsSession)
	ses2 := s2.(*gorillaws.WsSession)
	ses3 := s3.(*gorillaws.WsSession)
	list = append(list, ses1,ses2,ses3)
	var curoper_playid int = 1
	var curoper_playerSession *gorillaws.WsSession = nil
	for {
		//没轮到自己选地主的另外两个空闲玩家 服务器仅仅下发消息展示UI倒计时
		//轮到自己选地主的玩家 服务器通过定时器超时机制等待玩家是否选定地主
		//先给没轮到自己的两个空闲玩家发消息 注意CurrPos 要填写当前正在选地主的玩家point

		switch curoper_playid{
		case 1:
			curoper_playerSession = ses1
		case 2:
			curoper_playerSession = ses2
		case 3:
			curoper_playerSession = ses3

		}

		curoper_player := curoper_playerSession.Player().(*ddz.DDZ_Player)

		curoper_player.Game.NowIsLead = true

		for _,ses := range list{
			ses.Send(&pbddz.RobBankerOper{
				CurrPos:curoper_player.Game.Point,
				RoundTime:12,
			})
		}

		//启动超时处理
		f := func(){
			curoper_player.Game.NowIsLead = false
			curoper_player.ExitSync.Done()
			fmt.Println("callback not called")
		}
		var t *time.Timer
		t = time.AfterFunc(time.Second*12,f)
		curoper_player.ExitSync.Add(1)
		curoper_player.ExitSync.Wait()
		t.Stop()
		//判断是否已经有玩家当选地主了
		if curoper_player.Game.IsMain{
			return true
		}else{
			curoper_playid += 1
			//轮了一圈 没人选地主 [尴尬] 从新发牌
			if curoper_playid == 3{
				return false
			}
		}
	}
	fmt.Println("not come hereeeee")
	return true
}


func EnterRoom(s1 interface{},s2 interface{},s3 interface{}) error{

	firstcard(s1,s2,s3)
	findmain := find_landowner(s1,s2,s3)
	if findmain{
		//开始出牌流程
		fmt.Println("not come here2")
	}else{
		//重新选地主
		fmt.Println("come here")
		EnterRoom(s1,s2,s3)
	}

	for{
		fmt.Println("not come here1")
		time.Sleep(time.Second*600)

	}

/*	session0 := s1.(cellnet.Session)
	session1 := s2.(cellnet.Session)
	session2 := s3.(cellnet.Session)
	switch whichroom {
	case 1:
		rst := get_free_desk_from_room1()
		if rst != nil{
			desk := rst.(*Desk)
			desk.IsGameing = true
			desk.ID = 1000
			desk.CurPlay = ""
			pailist := fapai().(map[int][]int)
			pson0 := desk.PeopleList[0]
			pson0.MySelfName = session0.GetPlayer().(role.Player).Att.Account
			pson0.MyHandPai = pailist[0]
			pson0.Ses = session0
			pson0.Sesslist = append(pson0.Sesslist,session1)
			pson0.Sesslist = append(pson0.Sesslist,session2)
			session0.GetPlayer().(role.Player).Desk = desk
			//session0.Send()

			pson1 := desk.PeopleList[1]
			pson1.MySelfName = session1.GetPlayer().(role.Player).Att.Account
			pson1.MyHandPai = pailist[1]
			pson1.Ses = session1
			pson1.Sesslist = append(pson0.Sesslist,session0)
			pson1.Sesslist = append(pson0.Sesslist,session2)
			session1.GetPlayer().(role.Player).Desk = desk

			pson2 := desk.PeopleList[2]
			pson2.MySelfName = session2.GetPlayer().(role.Player).Att.Account
			pson2.MyHandPai = pailist[2]
			pson2.Ses = session2
			pson2.Sesslist = append(pson0.Sesslist,session0)
			pson2.Sesslist = append(pson0.Sesslist,session1)
			session2.GetPlayer().(role.Player).Desk = desk

			fmt.Println(desk.ID,desk.IsGameing,desk.PeopleList)
			for _,node := range desk.PeopleList{
				fmt.Println(node.MySelfName,node.MyHandPai)
			}
		}else{
			return errors.New("房间没空余牌桌")
		}

	default:
		return errors.New("选择房间出错")
	}*/

	return nil
}


