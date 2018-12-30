package ddz

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	"trunk/cellnet"
	"trunk/cellnet/game"
	"trunk/cellnet/myrpc"
	pbddz "trunk/cellnet/pb/ddz"
	"trunk/cellnet/util"
)

func init() {
	game.RegeditLogicProcessFunction("ddz",ddz_handle)
}

var ddz_handle = func(ev cellnet.Event){

	player := ev.Session().Player().(*DDZ_Player)

	switch msg := ev.Message().(type) {
	case *cellnet.SessionConnectError:
		myrpc.Rpcqueue <- "SessionConnectError"
	case *cellnet.SessionClosed:
		myrpc.Rpcqueue <- fmt.Sprintf("close player=%s",player.Att.UserID)

	case *cellnet.SessionConnected:
		myrpc.Rpcqueue <- "connect success"
		login(ev,util.GenIntValue())
		//player.IncReqTimes()
	case *pbddz.RoomInfoPush:
		myrpc.Rpcqueue <- "RoomInfoPush"

	case *pbddz.LoginResp:
		myrpc.Rpcqueue <- "LoginResp"
		go heartbeat(ev)

		code := ev.Message().(*pbddz.LoginResp).Code
		if code != pbddz.LoginCode_login_success{
			myrpc.Rpcqueue <- fmt.Sprintf("error 登陆失败 服务器返回的错误码 %d",code)
			return
		}
		role := ev.Message().(*pbddz.LoginResp).GetRole()
		if nil == role{
			myrpc.Rpcqueue <- "error role=========nil"
			return
		}
		player.Att.UserID = role.Userid
		player.Att.Chips = role.Chips
		player.Att.Nickname = role.Nickname
		player.Att.Level = role.Level
		player.Att.Avatar = role.Avatar
		ev.Session().Send(&pbddz.GetChipsReq{})
		//player.IncReqTimes()
		myrpc.Rpcqueue <- fmt.Sprintf("LoginResp playerid=%s",player.Att.UserID)

	case *pbddz.ChipsPush:

		Room1 := 1
		Room4 := 1
		player.Att.Chips = ev.Message().(*pbddz.ChipsPush).GetChips()
		player.Room = util.GenerateRangeNum(Room1,Room4)

		//开始游戏请求
		ev.Session().Send(&pbddz.StartMatchReq{
			Type:uint32(player.Room),
		})
		myrpc.Rpcqueue <- fmt.Sprintf("%s 请求进入游戏",player.Att.UserID)

	case *pbddz.StartMatchResp:
		result := ev.Message().(*pbddz.StartMatchResp).GetCode()
		if result == pbddz.LobbyCode_already_in_queue{
			myrpc.Rpcqueue <- "正在匹配中，请稍后。。。"
		}else if result == pbddz.LobbyCode_lobby_success{
			myrpc.Rpcqueue <- fmt.Sprintf("%s 开始新一轮游戏 当前金币: %d 进入 %d 房",
				player.Att.UserID,player.Att.Chips,player.Room)
		}else{
			myrpc.Rpcqueue <- "error 进入游戏失败"
		}
	case *pbddz.PushRoomInfo:
		player.Game.Point = msg.SelfPos
		player.Game.Cardlist = msg.Seats.Cards
		myrpc.Rpcqueue <- fmt.Sprintf("玩家%s 座位号%d 牌%v",player.Att.UserID,player.Game.Point,player.Game.Cardlist)

	case *pbddz.HeartbeatResp:

	case *pbddz.RobBankerOper:
		if msg.CurrPos == player.Game.Point{

			time.Sleep(time.Second*3)

			ev.Session().Send(&pbddz.RobBankerResp{
				CurrPos:player.Game.Point,
				SureRob:true,//是否选地主
			})
		}else{
			//没轮到自己选地主 选地主按钮变灰
		}
	case *pbddz.DipaiPush:
		if msg.CurrPos == player.Game.Point{
			//3张底牌是自己的,存起来
			player.Game.Cardlist = append(player.Game.Cardlist,msg.Cards[0],msg.Cards[1],msg.Cards[2])
		}else{
			//UI展示3张底牌
		}
	case *pbddz.StagePush:
		if ev.Message().(*pbddz.StagePush).CurrPos == player.Game.Point {
			//轮到自己 可以达到超时时间之前出牌 sleep一会
			l := player.PlayCard()
			ev.Session().Send(&pbddz.OperationReq{
				Pos:l.Point,
				Operation:pbddz.OperationType(l.OperType),
				Data:&pbddz.CardData{
					Cardtype:uint32(l.Cardtype),
					Cardvalue:l.CardValue,
					Extra:l.Extra,
				},
			})
		}else{
			//没轮到自己 转圈
		}
	case *pbddz.PushPosOperation:
		//UI展示服务器推送过来的上家出的牌 并保存
		if msg.Pos == player.Game.Point{
			preoper := &Previouser{
				OperType:uint32(msg.Operation),
				Point:msg.Pos,
				Cardtype:CARD_TYPE(msg.Data.Cardtype),
				CardValue:msg.Data.Cardvalue,
				Extra:msg.Data.Extra,
			}
			player.Game.Pre = preoper
		}


/*	case *ddz.GamePlayerInfoPush:
		list := ev.Message().(*jdzjh.GamePlayerInfoPush).GetPlayerInfos()
		for _,v := range list{
			if v.IsYours == 1{
				player.Pos = v.Pos
				//myrpc.Rpcqueue <- fmt.Sprintf("这局有 %d 人参与 %s pos==%d",len(list),player.Att.Account,v.Pos)
			}
		}*/

	default:
		fmt.Println("未知消息",msg)
		myrpc.Rpcqueue <- fmt.Sprintf("error 未知消息 %v",msg)
	}

}

func login(ev cellnet.Event,account int64){

	ev.Session().Send(&pbddz.LoginAccountReq{
		Userid: strconv.FormatInt(account,10),
		Token:   "123",
	})

}


func heartbeat(ev interface{}) {

	ll := ev.(cellnet.Event)
	session := ll.Session()
	//player := session.Player().(*DDZ_Player)

	for {
		conn, ok := session.Raw().(*websocket.Conn)
		if !ok || conn == nil {
			myrpc.Rpcqueue <- "退出心跳GO程"
			break
		}else{
			//player.IncReqTimes()
			session.Send(&pbddz.HeartbeatReq{Id: 3000,})
			time.Sleep(time.Second * 1)
		}
	}
}
