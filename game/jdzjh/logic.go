package jdzjh

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	"trunk/cellnet"
	"trunk/cellnet/game"
	"trunk/cellnet/pb/jdzjh"
	"trunk/cellnet/myrpc"
	"trunk/cellnet/util"
)

func init() {
	game.RegeditLogicProcessFunction("jdzjh",zjh_handle)
}

var zjh_handle = func(ev cellnet.Event){

	player := ev.Session().Player().(*Zjh_Player)

	switch msg := ev.Message().(type) {
	case *cellnet.SessionConnectError:
		myrpc.Rpcqueue <- "SessionConnectError"
	case *cellnet.SessionClosed:
		myrpc.Rpcqueue <- fmt.Sprintf("close player=%s",player.Att.Account)

	case *cellnet.SessionConnected:
		myrpc.Rpcqueue <- "connect success"
		login(ev,util.GenIntValue())
		player.IncReqTimes()

	case *jdzjh.UpdateTokenResp:
	case *jdzjh.HeartbeatResp:
	case *jdzjh.TimeResp:
	case *jdzjh.PlayerReconnect:
		myrpc.Rpcqueue <- "PlayerReconnect"
	case *jdzjh.PlayerMsg:
		myrpc.Rpcqueue <- "PlayerMsg"
	case *jdzjh.LoginResp:

		go heartbeat(ev)

		code := ev.Message().(*jdzjh.LoginResp).Code
		if code != 1{
			myrpc.Rpcqueue <- fmt.Sprintf("error 登陆失败 服务器返回的错误码 %d",code)
			return
		}
		role := ev.Message().(*jdzjh.LoginResp).GetRole()
		if nil == role{
			myrpc.Rpcqueue <- "error role=========nil"
			return
		}
		player.Att.Account = role.Account
		player.Att.Agentid = role.AgentId
		player.Att.Ownerid = role.OwnerId
		player.Att.Chips = role.Chips
		player.Att.Nickname = role.Nickname
		player.Att.Level = role.Level
		player.Att.Avatar = role.Avatar
		player.Att.Isnew = role.IsNew
		player.Att.Ismainwallet = role.IsMainWallet
		player.Att.Mainwallet = role.MainWallet
		player.Att.Autowallet = role.AutoWallet
		ev.Session().Send(&jdzjh.GetChipsReq{})
		player.IncReqTimes()
		myrpc.Rpcqueue <- fmt.Sprintf("LoginResp playerid=%s",player.Att.Account)

	case *jdzjh.GamePlayerInfoPush:
		list := ev.Message().(*jdzjh.GamePlayerInfoPush).GetPlayerInfos()
		for _,v := range list{
			if v.IsYours == 1{
				player.Pos = v.Pos
				//myrpc.Rpcqueue <- fmt.Sprintf("这局有 %d 人参与 %s pos==%d",len(list),player.Att.Account,v.Pos)
			}
		}
	case *jdzjh.RoomInfoPush:
		//myrpc.Rpcqueue <- "RoomInfoPush"

	case *jdzjh.ChipsPush:
		player.Att.Chips = ev.Message().(*jdzjh.ChipsPush).GetChips()
		player.Room = util.GenerateRangeNum(Room1,Room4)



		//开始游戏请求
		ev.Session().Send(&jdzjh.StartMatchReq{
			Type:uint32(player.Room),
		})
		myrpc.Rpcqueue <- fmt.Sprintf("%s 请求进入游戏",player.Att.Account)
		player.IncReqTimes()

	case *jdzjh.StartMatchResp:
		result := ev.Message().(*jdzjh.StartMatchResp).GetCode()
		if result != jdzjh.LobbyCode_lobby_success{
			player.CalcTimes(InsertGameFailed)
			myrpc.Rpcqueue <- fmt.Sprintf("error 玩家 %s 金币 %d 进入%d房间失败,错误码是%d",player.Att.Account,
				player.Att.Chips,player.Room,result)
		}else{
			player.CalcTimes(InsertGameSuccess)
			player.CalcRooms(uint32(player.Room))
			myrpc.Rpcqueue <- fmt.Sprintf("%s 开始新一轮游戏 当前金币: %d 进入 %d 房",
				player.Att.Account,player.Att.Chips,player.Room)
		}


	case *jdzjh.UserOperationResp:
		if ev.Message().(*jdzjh.UserOperationResp).Result != 1{
			myrpc.Rpcqueue <- fmt.Sprintf("error %s UserOperationResp %v",player.Att.Account,msg)
		}

	case *jdzjh.BcastMsgResp:
	case *jdzjh.GamePush:
		tmptype := ev.Message().(*jdzjh.GamePush).Type
		pos := ev.Message().(*jdzjh.GamePush).KeyId
		if pos == uint64(player.Pos) && tmptype == 4{
			time.Sleep(time.Duration(util.GenerateRangeNum(4,6))*time.Second)
			//time.Sleep(time.Duration(game.Suspend)*time.Second)
			//关于 PreType =[7,8] 移步查看proto GamePush协议
			if player.PrePos != 0/* && player.PreType == 7*/{
				//跟进上家7 跟注操作
				//myrpc.Rpcqueue <- fmt.Sprintf("%s 跟注操作1",player.Att.Account)
				player.CalcTimes(CinglOperator)
				player.RoomOperList[player.Room].Cingltimes += 1
				ev.Session().Send(&jdzjh.UserOperationReq{
					//OperationKey:   player.PreType - 4,
					OperationKey:3,
					//OperationValue: 0,//跟注时 这个值无意义 可以随便写
				})
				player.IncReqTimes()

			}else if player.PrePos != 0 && player.PreType == 8{
				//跟进上家8 加注操作
				//myrpc.Rpcqueue <- fmt.Sprintf("%s 加注操作1",player.Att.Account)
				player.CalcTimes(RaisetimesOperator)
				player.RoomOperList[player.Room].Raisetimes += 1
				ev.Session().Send(&jdzjh.UserOperationReq{
					//OperationKey:   player.PreType - 4,
					//OperationValue: uint32(player.PreValueArry[1]),//加注时 这个值是倍数
					OperationKey:3,
					//OperationValue:0,
				})
				player.IncReqTimes()

			} else if player.PrePos == 0{
				//这是新开局自己当庄,自己先出牌  2 弃牌，3 跟注， 4 加注
				which_oper := util.GenerateRangeNum(DiscardOperator,RaisetimesOperator)
				player.CalcTimes(uint32(which_oper))
				player.IncReqTimes()
				if which_oper == RaisetimesOperator{
					//myrpc.Rpcqueue <- fmt.Sprintf("%s 加注操作2",player.Att.Account)
					player.RoomOperList[player.Room].Raisetimes += 1
					ev.Session().Send(&jdzjh.UserOperationReq{
						OperationKey:uint32(which_oper),
						OperationValue:2, //加注默认2倍
					})
				}else{
					if which_oper == DiscardOperator{
						//myrpc.Rpcqueue <- fmt.Sprintf("%s 弃牌操作",player.Att.Account)
						player.RoomOperList[player.Room].Discardtimes += 1
					}else{
						//myrpc.Rpcqueue <- fmt.Sprintf("%s 跟注操作2",player.Att.Account)
						player.RoomOperList[player.Room].Cingltimes += 1
					}
					ev.Session().Send(&jdzjh.UserOperationReq{
						OperationKey:uint32(which_oper),
						//OperationValue:NoMeaning, //2 弃牌，3 跟注 的时候这个值没意义
					})
				}
			}

		}else{
			//保存前一个玩家执行的 跟注 加注操作
			//if  tmptype == 7 || tmptype == 8 {
				player.PreType = tmptype
				player.PrePos = pos
				player.PreValueArry = ev.Message().(*jdzjh.GamePush).ValueList
			//}
		}

	case *jdzjh.GameSettlementPush:
		myrpc.Rpcqueue <- fmt.Sprintf("%s 开始结算  ",player.Att.Account)
		//清零状态
		player.Pos = 0
		player.PreType = 0
		player.PrePos = 0
		player.PreValueArry = player.PreValueArry[:0]
		ev.Session().Send(&jdzjh.GetChipsReq{})
		player.IncReqTimes()

	default:
		fmt.Println("未知消息",msg)
		myrpc.Rpcqueue <- fmt.Sprintf("error 未知消息 %v",msg)
	}

}

func login(ev cellnet.Event,account int64){

	ev.Session().Send(&jdzjh.LoginAccountReq{
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
	player := session.Player().(*Zjh_Player)

	for {
		conn, ok := session.Raw().(*websocket.Conn)
		if !ok || conn == nil {
			myrpc.Rpcqueue <- "退出心跳GO程"
			break
		}else{
			player.IncReqTimes()
			session.Send(&jdzjh.HeartbeatReq{Id: 2000,})
			time.Sleep(time.Second * 1)
		}
	}

}