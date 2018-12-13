package suoha


var ProtocolMap map[string]uint16

//暂时手工添加协议映射关系
func fill_protocol_cmd(){
	ProtocolMap = make(map[string]uint16)
	ProtocolMap["HeartbeatReq"] 		= 	((0 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["HeartbeatResp"]		= 	((0 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["TimeReq"]				=  	((0 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["TimeResp"]				= 	((0 & 0x007F) << 9) | (3 & 0x01FF)

	ProtocolMap["LoginReq"]				=  	((1 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["LoginAccountReq"]		=  	((1 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["LoginResp"]			= 	((1 & 0x007F) << 9) | (1 & 0x01FF)

	ProtocolMap["ChipsPush"]			= 	((21 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["GetChipsReq"]			= 	((21 & 0x007F) << 9) | (3 & 0x01FF)

	ProtocolMap["BcastMsgResp"]			= 	((23 & 0x007F) << 9) | (2 & 0x01FF)

	ProtocolMap["RoomInfoPush"]			= 	((22 & 0x007F) << 9) | (0 & 0x01FF) //登陆完成后 不用请求 服务器主动下发
	ProtocolMap["StartMatchReq"]		= 	((22 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["StartMatchResp"]		= 	((22 & 0x007F) << 9) | (2 & 0x01FF)

	ProtocolMap["PushRoomInfo"]			= 	((64 & 0x007F) << 9) | (0 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["StagePush"]			= 	((64 & 0x007F) << 9) | (1 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["OperationReq"]			= 	((64 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["BaseChipsPush"]		= 	((64 & 0x007F) << 9) | (3 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["DealCardsPush"]		= 	((64 & 0x007F) << 9) | (4 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["FinalDealResp"]		= 	((64 & 0x007F) << 9) | (5 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["SpreadCardsPush"]		= 	((64 & 0x007F) << 9) | (6 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["CheckCardsPush"]		= 	((64 & 0x007F) << 9) | (7 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["PushPosOperation"]		= 	((64 & 0x007F) << 9) | (9 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["ReloginInfoResp"]		= 	((64 & 0x007F) << 9) | (10 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["RoomCodeResp"]			= 	((64 & 0x007F) << 9) | (13 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
	ProtocolMap["SettleChipsPush"]		= 	((64 & 0x007F) << 9) | (11 & 0x01FF) //登陆完成后 不用请求 服务器主动下发(suoha)
}

func init(){

	fill_protocol_cmd()
}