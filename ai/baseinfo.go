package ai

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"trunk/cellnet/game/ddz"
	pbddz "trunk/cellnet/pb/ddz"
	"trunk/cellnet/peer/gorillaws"
)

var Paidata = [5][13]int {
	{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d},//方块(Diamond)
	{0x11,0x12,0x13,0x14,0x15,0x16,0x17,0x18,0x19,0x1a,0x1b,0x1c,0x1d},//梅花(Club)
	{0x21,0x22,0x23,0x24,0x25,0x26,0x27,0x28,0x29,0x2a,0x2b,0x2c,0x2d},//红桃(Heart)
	{0x31,0x32,0x33,0x34,0x35,0x36,0x37,0x38,0x39,0x3a,0x3b,0x3c,0x3d},//黑桃(Spade)
	{0x4E,0x4F}, //王(Joker)
}
//说明 21 22 分别标示 A 2,没用1 2 或者14 15 标示是为了拆牌算法需要,避免组单顺牌时把 A 2 这样的大牌组进去
//大小鬼 用100 80标示,没用紧接着12 13后面的数字也是为了拆牌算法的考虑,避免组单顺时把大小鬼这样的大牌组进去
var paidata = []uint32{
	3,4,5,6,7,8,9,10,11,12,13,21,22,
	3,4,5,6,7,8,9,10,11,12,13,21,22,
	3,4,5,6,7,8,9,10,11,12,13,21,22,
	3,4,5,6,7,8,9,10,11,12,13,21,22,
	100,80,
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

func sendEnterSuccess(list []*gorillaws.WsSession){
	for _,ses := range list{
		ses.Send(&pbddz.StartMatchResp{
			Code:pbddz.LobbyCode_lobby_success,
		})
	}
}

func send17card(list []*gorillaws.WsSession) []uint32{

	pailist,dipailist := fapai()

	var selfpos uint32
	for it,ses := range list{
		selfpos = uint32(it + 1)//玩家座位号从 1 开始往后排
		ses.Player().(*ddz.DDZ_Player).Game.Point = selfpos
		ses.Player().(*ddz.DDZ_Player).Game.Cardlist = pailist[it]
		ses.Player().(*ddz.DDZ_Player).Game.Status = ddz.Type_matched

		//这个操作是为了利用int类型API做排序 begin
		var tmplist []int
		for _,v := range pailist[it]{
			tmplist = append(tmplist,int(v))
		}
		sort.Ints(tmplist)
		fmt.Printf("玩家%s 17张牌是:%+v\n",ses.Player().(*ddz.DDZ_Player).Att.UserID,tmplist)
		//end

		ses.Send(&pbddz.PushRoomInfo{
			SelfPos:selfpos,
			Seats:&pbddz.SeatInfo{
				Pos:selfpos,
				Cards:pailist[it],
			},
		})
	}

	fmt.Printf("底牌是%+v\n",dipailist)

	return dipailist
}

//产生地主 根据进入房间的次序 挨个询问玩家
func find_landowner(list []*gorillaws.WsSession) (bool,*ddz.DDZ_Player){

	var curoper_playid int = 1
	var curoper_playerSession *gorillaws.WsSession = nil
	for {
		//没轮到自己选地主的另外两个空闲玩家 服务器仅仅下发消息展示UI倒计时
		//轮到自己选地主的玩家 服务器通过定时器超时机制等待玩家是否选定地主
		//先给没轮到自己的两个空闲玩家发消息 注意CurrPos 要填写当前正在选地主的玩家point

		switch curoper_playid{
		case 1:
			curoper_playerSession = list[0]
		case 2:
			curoper_playerSession = list[1]
		case 3:
			curoper_playerSession = list[2]

		}

		curoper_player := curoper_playerSession.Player().(*ddz.DDZ_Player)

		//默认设置true,超时调用函数中会从新设置flase, 挡住超时后收到客户端发来的协议,防止再次调用ExitSync.Done 导致崩溃
		curoper_player.Game.NowIsLead = true//用于控制ExitSync.Done仅被调用一次

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
			return true,curoper_player
		}else{
			curoper_playid += 1
			//轮了一圈 没人选地主 [尴尬] 从新发牌
			if curoper_playid == 3{
				return false,nil
			}
		}
	}

	return false,nil
}


func EnterRoom(s1 interface{},s2 interface{},s3 interface{}) error{

	var list []*gorillaws.WsSession
	ses1 := s1.(*gorillaws.WsSession)
	ses2 := s2.(*gorillaws.WsSession)
	ses3 := s3.(*gorillaws.WsSession)
	list = append(list, ses1,ses2,ses3)

	sendEnterSuccess(list)
	dipai := send17card(list)
	findmain,dizhu := find_landowner(list)

	if findmain{
		//服务器保存地主玩家的3张底牌 地主有20张 农民有17张
		dizhu.Game.Cardlist = append(dizhu.Game.Cardlist,dipai...)

		//3张底牌发给客户端 地主玩家
		for _,ses := range list{
			ses.Send(&pbddz.DipaiPush{
				CurrPos:dizhu.Game.Point,//群发3人,这是地主的Point
				Cards:dipai,
			})
		}
	}else{
		//重新选地主
		EnterRoom(s1,s2,s3)
	}

	//确定3个人的出牌顺序 从地主开始,后面2人顺延 begin
	var sortlist []*gorillaws.WsSession

	var dizhu_index int
	for index,ses := range list{
		if ses.Player().(*ddz.DDZ_Player) == dizhu{
			dizhu_index =  index
			break
		}
	}
	switch dizhu_index{
	case 0:
		sortlist = append(sortlist,list[0],list[1],list[2])
	case 1:
		sortlist = append(sortlist,list[1],list[2],list[0])
	case 2:
		sortlist = append(sortlist,list[2],list[0],list[1])
	}
	//end

	//拆出一手好牌的逻辑
	{
		player1 := sortlist[0].Player().(*ddz.DDZ_Player)
		player2 := sortlist[1].Player().(*ddz.DDZ_Player)
		player3 := sortlist[2].Player().(*ddz.DDZ_Player)

		player1.Game.Parse.GetBeautifulCard(player1.Game.Cardlist)
		player2.Game.Parse.GetBeautifulCard(player2.Game.Cardlist)
		player3.Game.Parse.GetBeautifulCard(player3.Game.Cardlist)
	}


	//开始出牌流程
	var curoper_playid int = 0
	var curoper_playerSession *gorillaws.WsSession = nil

	for{
		switch curoper_playid{
		case 0:
			curoper_playerSession = sortlist[0]
		case 1:
			curoper_playerSession = sortlist[1]
		case 2:
			curoper_playerSession = sortlist[2]
		}

		curoper_player := curoper_playerSession.Player().(*ddz.DDZ_Player)

		//设置允许出牌的flag 默认是true 超时后设置flase,晚于超时后再发来出牌不被允许。
		curoper_player.Game.IsDoit = true//用于控制ExitSync.Done仅被调用一次

		for _,ses := range sortlist{
			ses.Send(&pbddz.StagePush{
				CurrPos:curoper_player.Game.Point,
				RoundTime:12,
			})
		}

		//启动超时处理
		f := func(){
			curoper_player.Game.IsDoit = false
			curoper_player.ExitSync.Done()
			fmt.Println("该玩家没出牌,默认成过牌操作")
		}
		var t *time.Timer
		t = time.AfterFunc(time.Second*12,f)
		curoper_player.ExitSync.Add(1)
		curoper_player.ExitSync.Wait()
		t.Stop()


		//判断客户端是否手里没牌了,没了就结束进入结算环节。

		if curoper_player.Game.IsPassCard != true{
			if curoper_player.Game.Pre.Nocards{
				for _,ses := range sortlist{
					ses.Send(&pbddz.GameoverPush{
						DizhuWin:curoper_player.Game.IsMain,
					})
				}
				//结算收益

				return nil
			}
			//判断玩家出的牌是否合法 (不合法要通知到出牌的玩家)begin 服务端也需移除自己存储的玩家出过的牌
			//todo...
			//end
		}


		//判断玩家是否已经出牌了 出牌了就通知另外2个玩家 该玩家出牌内容 然后继续下一个人出牌
		var otherTwoPlayer []*gorillaws.WsSession
		switch curoper_playid{
		case 0:
			otherTwoPlayer = append(otherTwoPlayer,sortlist[1],sortlist[2])
		case 1:
			otherTwoPlayer = append(otherTwoPlayer,sortlist[0],sortlist[2])
		case 2:
			otherTwoPlayer = append(otherTwoPlayer,sortlist[0],sortlist[1])
		}

		var info *pbddz.PushPosOperation
		if curoper_player.Game.IsPassCard {
			info = &pbddz.PushPosOperation{
				//Pos:curoper_player.Game.Pre.Point,
				Pos:curoper_player.Game.Point, //用Pre.Point 可能会崩溃(原因:Pre是nil)  这个字段意义不大
				Operation:pbddz.OperationType(pbddz.OperationType_pass),
			}
		}else{
			info = &pbddz.PushPosOperation{
				Pos:curoper_player.Game.Pre.Point,
				Operation:pbddz.OperationType(curoper_player.Game.Pre.OperType),
				Data:&pbddz.CardData{
					Cardtype:uint32(curoper_player.Game.Pre.Cardtype),
					Cardvalue:curoper_player.Game.Pre.CardValue,
					Extra:curoper_player.Game.Pre.Extra,
				},
				Nocards:curoper_player.Game.Pre.Nocards,
			}
		}

		for _,ses := range otherTwoPlayer{
			ses.Send(info)
		}

		curoper_playid += 1 //轮到下一个玩家,到最后一个玩家后 赋0进入下一轮
		if curoper_playid == 3{
			curoper_playid = 0
		}
	}

	return nil
}


