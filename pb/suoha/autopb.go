// Generated by trunk/cellnet/protoc-gen-msg
// DO NOT EDIT!
// Source: msg.proto
// Source: msg_base.proto
// Source: msg_common.proto
// Source: msg_broadcast.proto
// Source: msg_gm.proto
// Source: msg_lobby.proto
// Source: msg_role_common.proto
// Source: msg_login.proto
// Source: msg_role.proto
// Source: msg_room_6303.proto

package suoha

import (
	"trunk/cellnet"
	"reflect"
	_ "trunk/cellnet/codec/gogopb"
	"trunk/cellnet/codec"
)

func NotWriteInit() {

	// msg.proto
	// msg_base.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*HeartbeatReq)(nil)).Elem(),
		ID:    13510,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*HeartbeatResp)(nil)).Elem(),
		ID:    10040,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*TimeReq)(nil)).Elem(),
		ID:    10245,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*TimeResp)(nil)).Elem(),
		ID:    33367,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*PopupMsgResp)(nil)).Elem(),
		ID:    4355,
	})
	// msg_common.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*Undefined)(nil)).Elem(),
		ID:    15776,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*DmsgContent)(nil)).Elem(),
		ID:    61428,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*Dmsg)(nil)).Elem(),
		ID:    54425,
	})
	// msg_broadcast.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*BcastworldResp)(nil)).Elem(),
		ID:    26109,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*BcastMsgResp)(nil)).Elem(),
		ID:    33724,
	})
	// msg_gm.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*GmCardInfo)(nil)).Elem(),
		ID:    2760,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*GmCardReq)(nil)).Elem(),
		ID:    5732,
	})
	// msg_lobby.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*RoomInfo)(nil)).Elem(),
		ID:    12183,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*RoomInfoPush)(nil)).Elem(),
		ID:    41559,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*StartMatchReq)(nil)).Elem(),
		ID:    11409,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*StartMatchResp)(nil)).Elem(),
		ID:    6243,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*CancelMatchReq)(nil)).Elem(),
		ID:    15369,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*CancelMatchResp)(nil)).Elem(),
		ID:    5851,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LeaveRoomReq)(nil)).Elem(),
		ID:    48320,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LeaveRoomResp)(nil)).Elem(),
		ID:    44658,
	})
	// msg_role_common.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*RoleInfo)(nil)).Elem(),
		ID:    64556,
	})
	// msg_login.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LoginReq)(nil)).Elem(),
		ID:    39439,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LoginResp)(nil)).Elem(),
		ID:    13729,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LoginAccountReq)(nil)).Elem(),
		ID:    55068,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*KickResp)(nil)).Elem(),
		ID:    27818,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*UpdateTokenResp)(nil)).Elem(),
		ID:    12620,
	})
	// msg_role.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*ChipsPush)(nil)).Elem(),
		ID:    59429,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SetAvatarReq)(nil)).Elem(),
		ID:    40801,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SetAvatarResp)(nil)).Elem(),
		ID:    58675,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*GetChipsReq)(nil)).Elem(),
		ID:    39757,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*UpdateGtReq)(nil)).Elem(),
		ID:    48468,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*GetWalletReq)(nil)).Elem(),
		ID:    57215,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*GetWalletResp)(nil)).Elem(),
		ID:    10513,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SetWalletReq)(nil)).Elem(),
		ID:    34315,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SetWalletResp)(nil)).Elem(),
		ID:    41245,
	})
	// msg_room_6303.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*PushRoomInfo)(nil)).Elem(),
		ID:    62935,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SeatInfo)(nil)).Elem(),
		ID:    8103,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*CardInfo)(nil)).Elem(),
		ID:    53172,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*BaseChipsPush)(nil)).Elem(),
		ID:    61408,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*DealCardsPush)(nil)).Elem(),
		ID:    37713,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*FinalDealResp)(nil)).Elem(),
		ID:    29800,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SpreadInfo)(nil)).Elem(),
		ID:    39289,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SpreadCardsPush)(nil)).Elem(),
		ID:    58234,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*CheckCardsPush)(nil)).Elem(),
		ID:    29113,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*StagePush)(nil)).Elem(),
		ID:    13954,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*OperationReq)(nil)).Elem(),
		ID:    2407,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LookResp)(nil)).Elem(),
		ID:    3869,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*PushPosOperation)(nil)).Elem(),
		ID:    5969,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*ChipsInfo)(nil)).Elem(),
		ID:    59409,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*SettleChipsPush)(nil)).Elem(),
		ID:    46358,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*LeaveRoomPush)(nil)).Elem(),
		ID:    31928,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*ReloginInfoResp)(nil)).Elem(),
		ID:    63940,
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*RoomCodeResp)(nil)).Elem(),
		ID:    25792,
	})
}
