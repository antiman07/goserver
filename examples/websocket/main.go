package main

import (
	"flag"
	"fmt"
	"trunk/cellnet"
	"trunk/cellnet/ai"
	"trunk/cellnet/db"
	"trunk/cellnet/game/ddz"
	"trunk/cellnet/myrpc"
	"trunk/cellnet/peer"
	_ "trunk/cellnet/peer/gorillaws"
	"trunk/cellnet/proc"
	_ "trunk/cellnet/proc/gorillaws"

	"trunk/cellnet/game"
	pb_ddz "trunk/cellnet/pb/ddz"
)


func server_parse_cmd_params(){

	tmp_which_game := flag.String("game","ddz","choose game to test")
	tmp_ipaddr := flag.String("addr","http://192.168.10.10:6302/echo","choose game server ip port like as 192.168.200.54:10303")
	tmp_rpclogaddr := flag.String("rl","192.168.10.10:1234","log server ip port input like as 10.200.124.122:1234")

	flag.Parse()

	game.GameName = *tmp_which_game
	game.GameAddr = *tmp_ipaddr
	game.LogserverAddr 	= *tmp_rpclogaddr
}

func main() {

	pb_ddz.NotWriteInit()
	server_parse_cmd_params()

	go myrpc.Rpclog(game.LogserverAddr,1)

	match := ai.Build()
	go match.StartLoop()
	// 创建一个事件处理队列，整个服务器只有这一个队列处理事件，服务器属于单线程服务器
	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("gorillaws.Acceptor", "server", game.GameAddr, queue)

	proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {

		player := ev.Session().Player().(*ddz.DDZ_Player)

		switch msg := ev.Message().(type) {

		case *cellnet.SessionAccepted:
			myrpc.Rpcqueue <- "SessionAccepted"
		case *cellnet.SessionClosed:
			myrpc.Rpcqueue <- fmt.Sprintf("close player=%s",player.Att.UserID)

		case *pb_ddz.LoginAccountReq:
			myrpc.Rpcqueue <- "LoginAccountReq"
			rst,err := db.Query(msg.Userid)
			if err != nil{
				ev.Session().Send(&pb_ddz.LoginResp{
				Code:pb_ddz.LoginCode_login_fail,
			})
			}else{
				for _,v := range rst{
					ev.Session().Send(&pb_ddz.LoginResp{
						Code:pb_ddz.LoginCode_login_success,
						Role:&pb_ddz.RoleInfo{
							Userid:v["uid"],
							Chips:v["chips"],
							Nickname:v["nickname"],
							Level:v["level"],
							Avatar:v["avatar"],
						},
					})
					player.Att.Chips = v["chips"]
					player.Att.UserID = v["uid"]
					player.Att.Nickname = v["nickname"]
				}
			}

		case *pb_ddz.HeartbeatReq:
			ev.Session().Send(&pb_ddz.HeartbeatResp{})

		case *pb_ddz.GetChipsReq:
			myrpc.Rpcqueue <- "GetChipsReq"
			ev.Session().Send(&pb_ddz.ChipsPush{
				Chips:player.Att.Chips,
			})

		case *pb_ddz.StartMatchReq:
			myrpc.Rpcqueue <- "StartMatchReq"
			player.Game.Status = ddz.Type_matching

			ev.Session().Send(&pb_ddz.StartMatchResp{
				Code:pb_ddz.LobbyCode_already_in_queue, //通知客户端展示状态 正在匹配中。。。
				Type:msg.Type,
			})

			match.Add(ev.Session())

		case *pb_ddz.RobBankerResp:
			if player.Game.NowIsLead == true{
				//NowIsLead 是 true时
				player.Game.IsMain = msg.SureRob
				player.ExitSync.Done() //Done 调用2次会崩溃(该协议发来的速度晚于定时器回调,会导致xitSync.Done()调用2次 会崩溃的)
			}
		case *pb_ddz.OperationReq:
			if player.Game.IsDoit == true{
				//客户端在超时之前出牌了 记录下出的牌 存放到该玩家的结构体内
				//player.Game.Pre = ev.Message().(*ddz.Previouser)
				if msg.Operation == pb_ddz.OperationType_pass{
					player.Game.IsPassCard = true
				}else{
					player.Game.IsPassCard = false
					player.Game.Pre =&ddz.Previouser{
						Point:msg.Pos,
						OperType:uint32(msg.Operation),
						Cardtype:ddz.CARD_TYPE(msg.Data.Cardtype),
						CardValue:msg.Data.Cardvalue,
						Extra:msg.Data.Extra,
						Nocards:msg.Nocards,
					}
				}

				player.ExitSync.Done()

			}

		}
	})

	// 开始侦听
	p.Start()

	// 事件队列开始循环
	queue.StartLoop()

	// 阻塞等待事件队列结束退出( 在另外的goroutine调用queue.StopLoop() )
	queue.Wait()

}
