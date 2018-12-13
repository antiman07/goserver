package tbnn

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	"trunk/cellnet"
	"trunk/cellnet/game"
	"trunk/cellnet/myrpc"
	"trunk/cellnet/pb/tbnn"
	"trunk/cellnet/util"
)

func init(){
	game.RegeditLogicProcessFunction("tbnn",tbnn_handle)
}

var tbnn_handle = func(ev cellnet.Event){

	player := ev.Session().Player().(*Tbnn_Player)

	switch msg := ev.Message().(type) {
	case *cellnet.SessionConnectError:
		myrpc.Rpcqueue <- "SessionConnectError"
	case *cellnet.SessionClosed:
		myrpc.Rpcqueue <- fmt.Sprintf("close player=%s",player.Att.Account)
	case *cellnet.SessionConnected:
		myrpc.Rpcqueue <- "connect success"
		login(ev,util.GenIntValue())

	case *tbnn.LoginResp:
		go heartbeat(ev)
		player.Att.Account = ev.Message().(*tbnn.LoginResp).GetRole().Account
		player.Att.Agentid = ev.Message().(*tbnn.LoginResp).GetRole().AgentId
		player.Att.Ownerid = ev.Message().(*tbnn.LoginResp).GetRole().OwnerId
		player.Att.Chips = ev.Message().(*tbnn.LoginResp).GetRole().Chips
		player.Att.Nickname = ev.Message().(*tbnn.LoginResp).GetRole().Nickname
		player.Att.Level = ev.Message().(*tbnn.LoginResp).GetRole().Level
		player.Att.Avatar = ev.Message().(*tbnn.LoginResp).GetRole().Avatar
		player.Att.Isnew = ev.Message().(*tbnn.LoginResp).GetRole().IsNew
		player.Att.Ismainwallet = ev.Message().(*tbnn.LoginResp).GetRole().IsMainWallet
		player.Att.Mainwallet = ev.Message().(*tbnn.LoginResp).GetRole().MainWallet
		player.Att.Autowallet = ev.Message().(*tbnn.LoginResp).GetRole().AutoWallet

	case *tbnn.RoomInfoPush:
		ev.Session().Send(&tbnn.StartMatchReq{
			Type:uint32(game.RoomID),
		})

	case *tbnn.StartMatchResp:

		//服务器下发的所有玩家进入房间信息列表 这里保存起来
	case *tbnn.EnterRoomPush:
		list := ev.Message().(*tbnn.EnterRoomPush).GetRole()
		for i := 0; i < len(list); i++{
			node := list[i]
			var s *tbnn.RoomRoleInfo = &tbnn.RoomRoleInfo{
				Avatar:node.Avatar,
				Nickname:node.Nickname,
				Pos: node.Pos,
				Chips:node.Chips,
				Profit:node.Profit,
				CardType:node.CardType,
				BetMulti:node.BetMulti,
			}

			player.EnterRoomPushList = append(player.EnterRoomPushList,s)
		}


		//服务器下发的下注倍数列表	这里保存起来
	case *tbnn.MultiListPush:
		list := ev.Message().(*tbnn.MultiListPush).GetMultiList()
		for i := 0; i < len(list); i++{
			node := list[i]
			player.MutilListPushList = append(player.MutilListPushList, node)
		}

		//发送下注倍数请求 发送时延(0-5秒之间) 下注倍数
		time.Sleep(time.Duration(game.BetmultiTime) * time.Second)
		ev.Session().Send(&tbnn.BetMultiReq{
			Multi:uint32(game.Betmulti),
		})

		//下注结果
	case *tbnn.BetMultiPush:
		node := ev.Message().(*tbnn.BetMultiPush)
		player.BetMutilPushList = append(player.BetMutilPushList, node)

		//发牌协议
	case *tbnn.DealCardResp:
		list := ev.Message().(*tbnn.DealCardResp).GetCards()
		for i := 0; i < len(list); i++{
			card := list[i]
			player.DealCardRespInfo = append(player.DealCardRespInfo, card)
		}

		//发送翻牌请求 发送时延(0-5秒之间)
		time.Sleep(time.Duration(game.SpreadcardTime) * time.Second)
		ev.Session().Send(&tbnn.SpreadCardReq{})


		//翻牌协议
	case *tbnn.SpreadCardPush:
		list := ev.Message().(*tbnn.SpreadCardPush).GetInfo()
		for i := 0; i < len(list); i++{
			node := list[i]
			player.SpreadCardPushList = append(player.SpreadCardPushList, node)
		}


		//结算收益
	case *tbnn.SettleChipPush:
		list := ev.Message().(*tbnn.SettleChipPush).GetInfo()
		for i := 0; i < len(list); i++{
			node := list[i]
			player.SettleChipPushList = append(player.SettleChipPushList, node)
		}

		//请求余额 暂且在这里请求
		ev.Session().Send(&tbnn.GetChipsReq{
		})

		time.Sleep(time.Duration(5) * time.Second)
		//log_client.Debugf("-----------5s后继续下一局游戏--------------")
		//清除之前服务器下发的数据
		player.EnterRoomPushList = nil
		player.MutilListPushList = nil
		player.BetMutilPushList = nil
		player.DealCardRespInfo = nil
		player.SpreadCardPushList = nil
		player.SettleChipPushList = nil

		ev.Session().Send(&tbnn.StartMatchReq{
			Type:uint32(game.RoomID),
		})


	case *tbnn.ChipsPush:
		player.Att.Chips = ev.Message().(*tbnn.ChipsPush).GetChips()

	case *tbnn.HeartbeatResp:
	case *tbnn.TimeResp:
	case *tbnn.SyncRoomStagePush:
	case *tbnn.BcastMsgResp:


	default:
		fmt.Println("未知消息",msg)
	}
}


func login(ev cellnet.Event,account int64){

	ev.Session().Send(&tbnn.LoginAccountReq{
		Account: strconv.FormatInt(account,10),
		Token:   "123",
		Lang:    "lang",
		Gt:      1,
		GameId:  100,
	})
}


func heartbeat(ev interface{}) {

	ll := ev.(cellnet.Event)
	session := ll.Session()

	for {
		conn, ok := session.Raw().(*websocket.Conn)
		if !ok || conn == nil {
			myrpc.Rpcqueue <- "退出心跳GO程"
			break
		}else{
			session.Send(&tbnn.HeartbeatReq{Id: 2000,})
			time.Sleep(time.Second * 1)
		}
	}

}
