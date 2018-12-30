package suoha

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	"trunk/cellnet"
	"trunk/cellnet/game"
	"trunk/cellnet/pb/suoha"
	"trunk/cellnet/myrpc"
	"trunk/cellnet/util"
)

func init(){
	game.RegeditLogicProcessFunction("suoha",suoha_handle)
}

var suoha_handle = func(ev cellnet.Event){
	player := ev.Session().Player().(*Suoha_Player)

	switch msg := ev.Message().(type) {
	case *cellnet.SessionConnectError:
		myrpc.Rpcqueue <- "SessionConnectError"
	case *cellnet.SessionClosed:
		myrpc.Rpcqueue <- fmt.Sprintf("close player=%s",player.Att.Account)

	case *cellnet.SessionConnected:
		myrpc.Rpcqueue <- "connect success"
		suoha_login(ev,util.GenIntValue())

	case *suoha.HeartbeatResp:
		//myrpc.Rpcqueue <- fmt.Sprintf("HeartbeatResp %s",player.Att.Account)

	case *suoha.LoginResp:

		go suoha_heartbeat(ev)


		code := ev.Message().(*suoha.LoginResp).Code
		if code != 1{
			myrpc.Rpcqueue <- fmt.Sprintf("error 登陆失败 服务器返回的错误码 %d",code)
			return
		}
		role := ev.Message().(*suoha.LoginResp).GetRole()
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

		myrpc.Rpcqueue <- fmt.Sprintf("LoginResp playerid=%s",player.Att.Account)

		ev.Session().Send(&suoha.GetChipsReq{})

	case *suoha.RoomInfoPush:
		list := ev.Message().(*suoha.RoomInfoPush).GetRoomList()
		for _,v := range list{
			var s *suoha.RoomInfo = &suoha.RoomInfo{
				RoomType:v.RoomType,
				Ante:v.Ante,
				Lowest:v.Lowest,
				MemberNum:v.MemberNum,
				IsOpen:v.IsOpen,
				Limit:v.Limit,
			}
			Suoha_RoomInfoList = append(Suoha_RoomInfoList, s)
		}

	case *suoha.ChipsPush:
		player.Att.Chips = ev.Message().(*suoha.ChipsPush).GetChips()
		myrpc.Rpcqueue <- fmt.Sprintf("玩家 %s 当前金币数量 %d",player.Att.Account,player.Att.Chips)

		//随机出房间
		player.Room = util.GenerateRangeNum(1,4)
		var into_room_id uint32
		switch player.Room{
		case 1:
			into_room_id = 1
		case 2:
			into_room_id = 5
		case 3:
			into_room_id = 10
		case 4:
			into_room_id = 30
		}

		//开始游戏请求
		ev.Session().Send(&suoha.StartMatchReq{
			Type:into_room_id,
		})

	case *suoha.StartMatchResp:
		result := ev.Message().(*suoha.StartMatchResp).GetCode()
		if result != 1{
			player.CalcTimes(7)
			myrpc.Rpcqueue <- fmt.Sprintf("error 玩家 %s 金币 %d 进入游戏失败,错误码是%d",player.Att.Account,
				player.Att.Chips,result)
		}else{
			player.CalcTimes(6)
			player.CalcRooms(uint32(player.Room))
			myrpc.Rpcqueue <- fmt.Sprintf("%s进入房间",player.Att.Account)
		}

	case *suoha.LeaveRoomPush:
		myrpc.Rpcqueue <- fmt.Sprintf("%s离开房间",player.Att.Account)

	case *suoha.PushRoomInfo:
		player.Pos = ev.Message().(*suoha.PushRoomInfo).GetSelfPos()
		//myrpc.Rpcqueue <- fmt.Sprintf("玩家 %s 自己的位置是%d",player.Att.Account,player.Pos)

	case *suoha.StagePush:
		if ev.Message().(*suoha.StagePush).CurrPos == player.Pos{
			//有界面的话 此刻展示UI出牌倒计时效果
			randtime := util.GenerateRangeNum(1,game.BetmultiTime)
			time.Sleep(time.Second * time.Duration(randtime))
			//myrpc.Rpcqueue <- fmt.Sprintf("玩家 %s 随机暂停%d秒后出牌",player.Att.Account,randtime)
			//轮到自己出牌

			//去掉梭哈 不然短时间玩家金币输光 不能持续进入游戏
			//which_operator := util.GenerateRangeNum(int(suoha.OperationType_pass),int(suoha.OperationType_showhand))
			which_operator := util.GenerateRangeNum(int(suoha.OperationType_pass),int(suoha.OperationType_raise))

			if player.Previous.CurrPos != 0{
				if which_operator == int(suoha.OperationType_cingl){
					// begin 根据前一个玩家出牌 跟随之
					switch suoha.OperationType(player.Previous.Operation){
					case suoha.OperationType_pass:
						player.CalcTimes(suoha.OperationType_pass)
						player.RoomOperList[player.Room].Passtimes += 1
					case suoha.OperationType_discard:
						player.CalcTimes(suoha.OperationType_discard)
						player.RoomOperList[player.Room].Discardtimes += 1
					case suoha.OperationType_cingl:
						player.CalcTimes(suoha.OperationType_cingl)
						player.RoomOperList[player.Room].Cingltimes += 1
					case suoha.OperationType_raise:
						player.CalcTimes(suoha.OperationType_raise)
						player.RoomOperList[player.Room].Raisetimes += 1
					}
					//end

					//如果随机到跟注操作 就跟进上一家
					ev.Session().Send(&suoha.OperationReq{Operation:player.Previous.Operation,
						Value:player.Previous.ReqValue})

				}else if which_operator == int(suoha.OperationType_raise){
					player.CalcTimes(suoha.OperationType_raise)
					player.RoomOperList[player.Room].Raisetimes += 1
					//如果随机到加注操作 就默认2倍
					ev.Session().Send(&suoha.OperationReq{Operation:suoha.OperationType_raise,
						Value:2,
					})

				}else{
					//其他操作 Value值无意义
					// begin
					player.CalcTimes(suoha.OperationType(which_operator))

					switch suoha.OperationType(which_operator){
					case suoha.OperationType_pass:
						player.RoomOperList[player.Room].Passtimes += 1
					case suoha.OperationType_discard:
						player.RoomOperList[player.Room].Discardtimes += 1
					case suoha.OperationType_cingl:
						player.RoomOperList[player.Room].Cingltimes += 1
					case suoha.OperationType_raise:
						player.RoomOperList[player.Room].Raisetimes += 1
					}
					//end

					ev.Session().Send(&suoha.OperationReq{
						Operation:suoha.OperationType(which_operator),
						Value:0,
					})
				}

			}else{
				//新开局游戏 并且第一轮自己先出牌(在过牌,弃牌,跟注,加注,梭哈中随机选择一种操作方式)
				player.CalcTimes(suoha.OperationType(which_operator))

				if which_operator == int(suoha.OperationType_raise) {
					player.RoomOperList[player.Room].Raisetimes += 1
					//是加注 就默认2倍
					ev.Session().Send(&suoha.OperationReq{Operation:suoha.OperationType_raise,
						Value:2,
					})
				}else{
					// begin
					switch suoha.OperationType(which_operator){
					case suoha.OperationType_pass:
						player.RoomOperList[player.Room].Passtimes += 1
					case suoha.OperationType_discard:
						player.RoomOperList[player.Room].Discardtimes += 1
					case suoha.OperationType_cingl:
						player.RoomOperList[player.Room].Cingltimes += 1
					}
					//end

					//其他操作 Value值无意义
					ev.Session().Send(&suoha.OperationReq{
						Operation:suoha.OperationType(which_operator),
						Value:0,
					})
				}
			}
		}

	case *suoha.PushPosOperation:
		if player.Pos != ev.Message().(*suoha.PushPosOperation).CurrPos{
			//保存上一个玩家的出牌信息
			info := ev.Message().(*suoha.PushPosOperation)
			player.Previous.CurrPos = info.GetCurrPos()
			player.Previous.Operation = info.GetOperation()
			player.Previous.OperationValue = info.GetOperationValue()
			player.Previous.ReqValue = info.GetReqValue()

			//myrpc.Rpcqueue <- fmt.Sprintf("玩家 %s 保存上一个玩家出牌操作类型%d 附带操作的值%d 请求值%d",
			//player.Att.Account,info.GetOperation(), info.GetOperationValue(),info.GetReqValue())
		}

	case *suoha.SettleChipsPush:
		//清空上一局缓存数据
		player.Previous.CurrPos = 0
		player.Previous.Operation = 0
		player.Previous.OperationValue = 0
		player.Previous.ReqValue = 0
		player.Att.Chips = 0

		ev.Session().Send(&suoha.GetChipsReq{})
		//myrpc.Rpcqueue <- fmt.Sprintf("玩家%s 准备开始下一局游戏",player.Att.Account)

	case *suoha.BcastMsgResp:
	case *suoha.TimeResp:
	case *suoha.FinalDealResp:
	case *suoha.SpreadCardsPush:
	case *suoha.CheckCardsPush:
	case *suoha.BaseChipsPush:
	case *suoha.DealCardsPush:

		//重连会发这些包过来 begin
	case *suoha.ReloginInfoResp:
	case *suoha.RoomCodeResp:
		//重连会发这些包过来 end

	default:
		fmt.Printf("未知消息%+v",ev.Message(),msg)
		myrpc.Rpcqueue <- fmt.Sprintf("error 未知消息 %v",ev.Message())
	}
}

func suoha_login(ev cellnet.Event,account int64){

	ev.Session().Send(&suoha.LoginAccountReq{
		Account: strconv.FormatInt(account,10),
		Token:   "123",
		Lang:    "lang",
		Gt:      1,
		GameId:  100,
	})
}


func suoha_heartbeat(ev interface{}) {

	ll := ev.(cellnet.Event)
	session := ll.Session()
	for {
		conn, ok := session.Raw().(*websocket.Conn)
		if !ok || conn == nil {
			myrpc.Rpcqueue <- "退出心跳GO程"
			break
		}else{
			session.Send(&suoha.HeartbeatReq{Id: 2000,})
			time.Sleep(time.Second * 1)
		}
	}

}
