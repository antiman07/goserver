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

		//go heartbeat(ev)

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
		myrpc.Rpcqueue <- fmt.Sprintf("登陆成功 playerid=%s,金币=%s",player.Att.UserID,player.Att.Chips)

	case *pbddz.ChipsPush:

		Room1 := 1
		Room4 := 1
		player.Att.Chips = ev.Message().(*pbddz.ChipsPush).GetChips()
		player.Room = util.GenerateRangeNum(Room1,Room4)

		//开始游戏请求
		ev.Session().Send(&pbddz.StartMatchReq{
			Type:uint32(player.Room),
		})
		myrpc.Rpcqueue <- fmt.Sprintf("%s 请求进入房间%d",player.Att.UserID,player.Room)

	case *pbddz.StartMatchResp:
		result := ev.Message().(*pbddz.StartMatchResp).GetCode()
		if result == pbddz.LobbyCode_already_in_queue{
			myrpc.Rpcqueue <- fmt.Sprintf("%s 正在匹配中,请稍后",player.Att.UserID)
		}else if result == pbddz.LobbyCode_lobby_success{
			myrpc.Rpcqueue <- fmt.Sprintf("成功进入游戏 %s  当前金币 %s 进入 %d 房",
				player.Att.UserID,player.Att.Chips,player.Room)
		}else{
			myrpc.Rpcqueue <- "error 进入游戏失败"
		}

	case *pbddz.PushRoomInfo:
		player.Game.Point = msg.SelfPos
		player.Game.Cardlist = msg.Seats.Cards
		myrpc.Rpcqueue <- fmt.Sprintf("玩家%s 座位号%d 牌%v",player.Att.UserID,player.Game.Point,player.Game.Cardlist)

		//初始化 保存前一个玩家出牌内容的结构体
		player.Game.Pre = &Previouser{
			Cardtype:0,//赋0 标示还没出过牌
		}

	case *pbddz.HeartbeatResp:

	case *pbddz.RobBankerOper:

		if msg.CurrPos == player.Game.Point{

			//小逻辑 ID是10000的玩家当选地主
/*			choose_dizhu := false
			if player.Att.UserID == "10001" {
				choose_dizhu = true
				myrpc.Rpcqueue <- "10001 当选地主"
			}*/

			time.Sleep(time.Second*1)
			ev.Session().Send(&pbddz.RobBankerResp{
				CurrPos:player.Game.Point,
				//SureRob:choose_dizhu,//是否选地主
				SureRob:true,//是否选地主
			})
		}else{
			//没轮到自己选地主 选地主按钮变灰
		}
	case *pbddz.DipaiPush:
		if msg.CurrPos == player.Game.Point{

			player.Game.IsMain = true //地主

			//3张底牌是自己的,存起来
			player.Game.Cardlist = append(player.Game.Cardlist,msg.Cards...)
			myrpc.Rpcqueue <- fmt.Sprintf("玩家%s是地主,底牌是:%v,完整手牌是:%+v",player.Att.UserID,msg.Cards,player.Game.Cardlist)

			//牌已下发 客户端拆牌
			player.Game.Parse.GetBeautifulCard(player.Game.Cardlist)

		}else{
			//UI展示3张底牌
			myrpc.Rpcqueue <- fmt.Sprintf("玩家%s是农民,底牌是:%v,完整手牌是:%+v",player.Att.UserID,msg.Cards,player.Game.Cardlist)

			//牌已下发 客户端拆牌
			player.Game.Parse.GetBeautifulCard(player.Game.Cardlist)
		}
	case *pbddz.StagePush:
		if ev.Message().(*pbddz.StagePush).CurrPos == player.Game.Point {
			//轮到自己 可以达到超时时间之前出牌 sleep一会
			l := player.PlayCard()

			if l == nil{
				myrpc.Rpcqueue <- "error 手里没牌出了?"
				return
			}

			isover := player.check_no_card()
			showcard(l,player)
			//过牌的话 OperationReq协议中有3个字段没值 分别是下面两种情况
			if l.OperType == uint32(pbddz.OperationType_pass){
				ev.Session().Send(&pbddz.OperationReq{
					Pos:l.Point,
					Operation:pbddz.OperationType_pass,
				})
			}else{
				ev.Session().Send(&pbddz.OperationReq{
					Pos:l.Point,
					Operation:pbddz.OperationType(l.OperType),
					Data:&pbddz.CardData{
						Cardtype:uint32(l.Cardtype),
						Cardvalue:l.CardValue,
						Extra:l.Extra,
					},
					Nocards:isover,
				})
			}
		}else{
			//没轮到自己 转圈
		}

	case *pbddz.PushPosOperation:
		//UI展示服务器推送过来的上家出的牌 并保存
		//上一家出的是要不起牌(过牌)的话,不能保存到Game.Pre结构体中,该结构体仅保存上一家出过的非过牌信息,以便于轮到自己出牌时 要出大于上家或上上家的牌。
		if msg.Operation != pbddz.OperationType_pass{
			preoper := &Previouser{
				OperType:uint32(msg.Operation),
				Point:msg.Pos,
				Cardtype:CARD_TYPE(msg.Data.Cardtype),
				CardValue:msg.Data.Cardvalue,
				Extra:msg.Data.Extra,
				Nocards:msg.Nocards,
			}
			player.Game.Pre = preoper
		}


	case *pbddz.GameoverPush:
		if player.Game.IsMain{
			if msg.DizhuWin {
				//自己是地主 并且赢了 UI动画
				fmt.Println("我是地主 我赢了")
			}else{
				//自己是地主 并且输了 UI动画
				fmt.Println("我是地主 我输了")
			}
		}else{
			if msg.DizhuWin {
				//自己是农民 并且输了 UI动画
				fmt.Println("我是农民 我输了")
			}else{
				//自己是农民 并且赢了 UI动画
				fmt.Println("我是农民 我赢了")
			}
		}

	default:
		fmt.Println("未知消息",msg)
		myrpc.Rpcqueue <- fmt.Sprintf("error 未知消息 %v",msg)
	}

}

func showcard(info *Previouser,player *DDZ_Player){

	switch info.OperType{

	case uint32(pbddz.OperationType_pass):
		fmt.Printf("玩家%s pos=%d 要不起\n",player.Att.UserID,info.Point)

	case uint32(pbddz.OperationType_doit):

		if info.Cardtype == SINGLE_TYPE{
			fmt.Printf("玩家%s pos=%d 出的单牌:%d\n",player.Att.UserID,info.Point,info.CardValue)

		}else if info.Cardtype == DOUBLE_TYPE{
			fmt.Printf("玩家%s pos=%d 出的对牌:%d%d\n",player.Att.UserID,info.Point,info.CardValue,info.CardValue)

		}else if info.Cardtype == THREE_ONE_TYPE{
			fmt.Printf("玩家%s pos=%d 出的三带一牌:%d%d%d%d\n",player.Att.UserID,info.Point,info.CardValue,info.CardValue,info.CardValue,info.Extra)

		}else if info.Cardtype == THREE_TYPE{
			fmt.Printf("玩家%s pos=%d 出的三牌:%d%d%d\n",player.Att.UserID,info.Point,info.CardValue,info.CardValue,info.CardValue)

		}else if info.Cardtype == SINGLE_LINK_TYPE{
			begin_card := info.CardValue
			var ll []uint32
			var i uint32
			ll = append(ll,begin_card)
			for i = 1; i < info.Extra; i++{
				ll = append(ll,begin_card + i)
			}
			fmt.Printf("玩家%s pos=%d 出的单顺牌:%+v\n",player.Att.UserID,info.Point,ll)

		}else if info.Cardtype == DOULEL_INK_TYPE{
			begin_card := info.CardValue
			var ll []uint32
			var i uint32
			ll = append(ll,begin_card)
			for i = 1; i < info.Extra; i++{
				ll = append(ll,begin_card + i)
				ll = append(ll,begin_card + i)
			}
			fmt.Printf("玩家%s pos=%d 出的双顺牌:%+v\n",player.Att.UserID,info.Point,ll)

		}else if info.Cardtype == ZHADAN_TYPE{
			fmt.Printf("玩家%s pos=%d 出的炸弹牌:%d%d%d%d\n",player.Att.UserID,info.Point,info.CardValue,info.CardValue,info.CardValue,info.CardValue)

		}else if info.Cardtype == ROCKET_TYPE{
			fmt.Printf("玩家%s pos=%d 出的火箭牌:大鬼小鬼\n",player.Att.UserID,info.Point)
		}
	}

}

func login(ev cellnet.Event,account int64){

	myrpc.Rpcqueue <- fmt.Sprintf("账号 %d 请求登陆",account)

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
