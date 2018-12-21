package ddz


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
	ProtocolMap["UpdateTokenResp"]		= 	((1 & 0x007F) << 9) | (4 & 0x01FF)

	ProtocolMap["ChipsPush"]			= 	((21 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["GetChipsReq"]			= 	((21 & 0x007F) << 9) | (3 & 0x01FF)

	ProtocolMap["BcastMsgResp"]			= 	((23 & 0x007F) << 9) | (2 & 0x01FF)

	ProtocolMap["RoomInfoPush"]			= 	((22 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["StartMatchReq"]		= 	((22 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["StartMatchResp"]		= 	((22 & 0x007F) << 9) | (2 & 0x01FF)

	ProtocolMap["GamePlayerInfoPush"]	= 	((24 & 0x007F) << 9) | (7 & 0x01FF)
	ProtocolMap["GamePush"]				= 	((24 & 0x007F) << 9) | (8 & 0x01FF)
	ProtocolMap["UserOperationReq"]		= 	((24 & 0x007F) << 9) | (9 & 0x01FF)
	ProtocolMap["UserOperationResp"]	= 	((24 & 0x007F) << 9) | (10 & 0x01FF)

	ProtocolMap["GameSettlementPush"]	= 	((24 & 0x007F) << 9) | (12 & 0x01FF)
	ProtocolMap["PlayerReconnect"]		= 	((24 & 0x007F) << 9) | (15 & 0x01FF)
	ProtocolMap["PlayerMsg"]			= 	((24 & 0x007F) << 9) | (16 & 0x01FF)


}

func init(){

	fill_protocol_cmd()
}