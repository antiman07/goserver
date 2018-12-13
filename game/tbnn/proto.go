package tbnn


var ProtocolMap map[string]uint16

func fill_protocol_cmd(){
	ProtocolMap = make(map[string]uint16)
	ProtocolMap["HeartbeatReq"] 		= 	((0 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["HeartbeatResp"]		= 	((0 & 0x007F) << 9) | (1 & 0x01FF)

	ProtocolMap["LoginReq"]				=  	((1 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["LoginAccountReq"]		=  	((1 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["LoginResp"]			= 	((1 & 0x007F) << 9) | (1 & 0x01FF)

	ProtocolMap["RoomInfoPush"]			= 	((22 & 0x007F) << 9) | (0 & 0x01FF) //登陆完成后 不用请求 服务器主动下发

	ProtocolMap["StartMatchReq"]		= 	((22 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["StartMatchResp"]		= 	((22 & 0x007F) << 9) | (2 & 0x01FF)

	ProtocolMap["TimeReq"]				=  	((0 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["TimeResp"]				= 	((0 & 0x007F) << 9) | (3 & 0x01FF)


	//以下协议还没调通 begin
	ProtocolMap["EnterRoomPush"]		= 	((69 & 0x007F) << 9) | (0 & 0x01FF)

	ProtocolMap["SyncRoomStagePush"]	= 	((69 & 0x007F) << 9) | (9 & 0x01FF)


	ProtocolMap["MultiListPush"]		= 	((69 & 0x007F) << 9) | (1 & 0x01FF)

	ProtocolMap["BetMultiReq"]			= 	((69 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["BetMultiPush"]			= 	((69 & 0x007F) << 9) | (3 & 0x01FF)

	ProtocolMap["DealCardResp"]			= 	((69 & 0x007F) << 9) | (4 & 0x01FF)

	ProtocolMap["SpreadCardReq"]		= 	((69 & 0x007F) << 9) | (5 & 0x01FF)
	ProtocolMap["SpreadCardPush"]		= 	((69 & 0x007F) << 9) | (6 & 0x01FF)

	ProtocolMap["SettleChipPush"]		= 	((69 & 0x007F) << 9) | (7 & 0x01FF)

	ProtocolMap["ChipsPush"]			= 	((21 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["GetChipsReq"]			= 	((21 & 0x007F) << 9) | (3 & 0x01FF)


	ProtocolMap["BcastMsgResp"]			= 	((23 & 0x007F) << 9) | (2 & 0x01FF)



}

func init(){

	fill_protocol_cmd()
}