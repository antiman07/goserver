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
	ProtocolMap["SetAvatarResp"]		= 	((21 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["GetChipsReq"]			= 	((21 & 0x007F) << 9) | (3 & 0x01FF)



	ProtocolMap["SetAvatarReq"]		    = 	((21 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["UpdateGtReq"]		    = 	((21 & 0x007F) << 9) | (4 & 0x01FF)
	ProtocolMap["SetWalletReq"]			= 	((21 & 0x007F) << 9) | (7 & 0x01FF)
	ProtocolMap["GetWalletResp"]		= 	((21 & 0x007F) << 9) | (6 & 0x01FF)
	ProtocolMap["BcastworldResp"]		= 	((23 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["GmCardReq"]		    = 	((24 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["UpdateTokenResp"]		= 	((1 & 0x007F) << 9)  | (4 & 0x01FF)
	ProtocolMap["KickResp"]			    = 	((1 & 0x007F) << 9)  | (3 & 0x01FF)

	ProtocolMap["LookResp"]		        = 	((64 & 0x007F) << 9) | (8 & 0x01FF)

	ProtocolMap["CancelMatchReq"]		= 	((22 & 0x007F) << 9) | (3 & 0x01FF)
	ProtocolMap["CancelMatchResp"]		= 	((22 & 0x007F) << 9) | (4 & 0x01FF)
	ProtocolMap["LeaveRoomReq"]		    = 	((22 & 0x007F) << 9) | (5 & 0x01FF)
	ProtocolMap["LeaveRoomResp"]		= 	((22 & 0x007F) << 9) | (6 & 0x01FF)
	ProtocolMap["LeaveRoomPush"]		= 	((64 & 0x007F) << 9) | (13 & 0x01FF)
	ProtocolMap["GetWalletReq"]			= 	((21 & 0x007F) << 9) | (5 & 0x01FF)
	ProtocolMap["SetWalletResp"]		= 	((21 & 0x007F) << 9) | (8 & 0x01FF)
	ProtocolMap["PopupMsgResp"]			= 	((0 & 0x007F) << 9)  | (4 & 0x01FF)


	ProtocolMap["RoomInfoPush"]			= 	((22 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["StartMatchReq"]		= 	((22 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["StartMatchResp"]		= 	((22 & 0x007F) << 9) | (2 & 0x01FF)

	ProtocolMap["BcastMsgResp"]			= 	((23 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["StopNoticePush"]		= 	((23 & 0x007F) << 9) | (3 & 0x01FF)

	ProtocolMap["PushRoomInfo"]			= 	((64 & 0x007F) << 9) | (0 & 0x01FF)
	ProtocolMap["StagePush"]			= 	((64 & 0x007F) << 9) | (1 & 0x01FF)
	ProtocolMap["OperationReq"]			= 	((64 & 0x007F) << 9) | (2 & 0x01FF)
	ProtocolMap["BaseChipsPush"]		= 	((64 & 0x007F) << 9) | (3 & 0x01FF)
	ProtocolMap["DealCardsPush"]		= 	((64 & 0x007F) << 9) | (4 & 0x01FF)
	ProtocolMap["FinalDealResp"]		= 	((64 & 0x007F) << 9) | (5 & 0x01FF)
	ProtocolMap["SpreadCardsPush"]		= 	((64 & 0x007F) << 9) | (6 & 0x01FF)
	ProtocolMap["CheckCardsPush"]		= 	((64 & 0x007F) << 9) | (7 & 0x01FF)
	ProtocolMap["PushPosOperation"]		= 	((64 & 0x007F) << 9) | (9 & 0x01FF)
	ProtocolMap["ReloginInfoResp"]		= 	((64 & 0x007F) << 9) | (10 & 0x01FF)
	ProtocolMap["RoomCodeResp"]			= 	((64 & 0x007F) << 9) | (14 & 0x01FF)
	ProtocolMap["SettleChipsPush"]		= 	((64 & 0x007F) << 9) | (11 & 0x01FF)
}

func init(){

	fill_protocol_cmd()
}