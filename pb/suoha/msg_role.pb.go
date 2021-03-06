// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg_role.proto

package suoha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RoleCode int32

const (
	RoleCode_role_reserver RoleCode = 0
	RoleCode_success       RoleCode = 1
	RoleCode_fail          RoleCode = 2
	RoleCode_frequent      RoleCode = 3
	RoleCode_frequenty     RoleCode = 4
	RoleCode_not_enough    RoleCode = 5
)

var RoleCode_name = map[int32]string{
	0: "role_reserver",
	1: "success",
	2: "fail",
	3: "frequent",
	4: "frequenty",
	5: "not_enough",
}
var RoleCode_value = map[string]int32{
	"role_reserver": 0,
	"success":       1,
	"fail":          2,
	"frequent":      3,
	"frequenty":     4,
	"not_enough":    5,
}

func (x RoleCode) Enum() *RoleCode {
	p := new(RoleCode)
	*p = x
	return p
}
func (x RoleCode) String() string {
	return proto.EnumName(RoleCode_name, int32(x))
}
func (x *RoleCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RoleCode_value, data, "RoleCode")
	if err != nil {
		return err
	}
	*x = RoleCode(value)
	return nil
}
func (RoleCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{0}
}

// 推送金额
type ChipsPush struct {
	Chips uint64 `protobuf:"varint,1,req,name=chips" json:"chips"`
}

func (m *ChipsPush) Reset()         { *m = ChipsPush{} }
func (m *ChipsPush) String() string { return proto.CompactTextString(m) }
func (*ChipsPush) ProtoMessage()    {}
func (*ChipsPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{0}
}
func (m *ChipsPush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChipsPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChipsPush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ChipsPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChipsPush.Merge(dst, src)
}
func (m *ChipsPush) XXX_Size() int {
	return m.Size()
}
func (m *ChipsPush) XXX_DiscardUnknown() {
	xxx_messageInfo_ChipsPush.DiscardUnknown(m)
}

var xxx_messageInfo_ChipsPush proto.InternalMessageInfo

func (m *ChipsPush) GetChips() uint64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

// 设置头像请求
type SetAvatarReq struct {
	Avatar uint32 `protobuf:"varint,1,req,name=avatar" json:"avatar"`
}

func (m *SetAvatarReq) Reset()         { *m = SetAvatarReq{} }
func (m *SetAvatarReq) String() string { return proto.CompactTextString(m) }
func (*SetAvatarReq) ProtoMessage()    {}
func (*SetAvatarReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{1}
}
func (m *SetAvatarReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetAvatarReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetAvatarReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetAvatarReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAvatarReq.Merge(dst, src)
}
func (m *SetAvatarReq) XXX_Size() int {
	return m.Size()
}
func (m *SetAvatarReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAvatarReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetAvatarReq proto.InternalMessageInfo

func (m *SetAvatarReq) GetAvatar() uint32 {
	if m != nil {
		return m.Avatar
	}
	return 0
}

// 设置头像返回
type SetAvatarResp struct {
	Avatar uint32 `protobuf:"varint,1,req,name=avatar" json:"avatar"`
}

func (m *SetAvatarResp) Reset()         { *m = SetAvatarResp{} }
func (m *SetAvatarResp) String() string { return proto.CompactTextString(m) }
func (*SetAvatarResp) ProtoMessage()    {}
func (*SetAvatarResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{2}
}
func (m *SetAvatarResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetAvatarResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetAvatarResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetAvatarResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAvatarResp.Merge(dst, src)
}
func (m *SetAvatarResp) XXX_Size() int {
	return m.Size()
}
func (m *SetAvatarResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAvatarResp.DiscardUnknown(m)
}

var xxx_messageInfo_SetAvatarResp proto.InternalMessageInfo

func (m *SetAvatarResp) GetAvatar() uint32 {
	if m != nil {
		return m.Avatar
	}
	return 0
}

// 获取金额请求
type GetChipsReq struct {
}

func (m *GetChipsReq) Reset()         { *m = GetChipsReq{} }
func (m *GetChipsReq) String() string { return proto.CompactTextString(m) }
func (*GetChipsReq) ProtoMessage()    {}
func (*GetChipsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{3}
}
func (m *GetChipsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChipsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChipsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetChipsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChipsReq.Merge(dst, src)
}
func (m *GetChipsReq) XXX_Size() int {
	return m.Size()
}
func (m *GetChipsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChipsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetChipsReq proto.InternalMessageInfo

// 更新终端类型请求
type UpdateGtReq struct {
	Gt GameTerminalType `protobuf:"varint,1,req,name=gt,enum=suoha.GameTerminalType" json:"gt"`
}

func (m *UpdateGtReq) Reset()         { *m = UpdateGtReq{} }
func (m *UpdateGtReq) String() string { return proto.CompactTextString(m) }
func (*UpdateGtReq) ProtoMessage()    {}
func (*UpdateGtReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{4}
}
func (m *UpdateGtReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateGtReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateGtReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UpdateGtReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGtReq.Merge(dst, src)
}
func (m *UpdateGtReq) XXX_Size() int {
	return m.Size()
}
func (m *UpdateGtReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGtReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGtReq proto.InternalMessageInfo

func (m *UpdateGtReq) GetGt() GameTerminalType {
	if m != nil {
		return m.Gt
	}
	return GameTerminalType_game_terminal_type_reserve
}

// 中心钱包额度请求
type GetWalletReq struct {
}

func (m *GetWalletReq) Reset()         { *m = GetWalletReq{} }
func (m *GetWalletReq) String() string { return proto.CompactTextString(m) }
func (*GetWalletReq) ProtoMessage()    {}
func (*GetWalletReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{5}
}
func (m *GetWalletReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetWalletReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetWalletReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetWalletReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletReq.Merge(dst, src)
}
func (m *GetWalletReq) XXX_Size() int {
	return m.Size()
}
func (m *GetWalletReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletReq proto.InternalMessageInfo

// 中心钱包额度返回
type GetWalletResp struct {
	Code       RoleCode `protobuf:"varint,1,req,name=code,enum=suoha.RoleCode" json:"code"`
	Chips      int64    `protobuf:"varint,2,opt,name=chips" json:"chips"`
	AutoWallet int64    `protobuf:"varint,3,opt,name=auto_wallet,json=autoWallet" json:"auto_wallet"`
}

func (m *GetWalletResp) Reset()         { *m = GetWalletResp{} }
func (m *GetWalletResp) String() string { return proto.CompactTextString(m) }
func (*GetWalletResp) ProtoMessage()    {}
func (*GetWalletResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{6}
}
func (m *GetWalletResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetWalletResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetWalletResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetWalletResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletResp.Merge(dst, src)
}
func (m *GetWalletResp) XXX_Size() int {
	return m.Size()
}
func (m *GetWalletResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletResp proto.InternalMessageInfo

func (m *GetWalletResp) GetCode() RoleCode {
	if m != nil {
		return m.Code
	}
	return RoleCode_role_reserver
}

func (m *GetWalletResp) GetChips() int64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func (m *GetWalletResp) GetAutoWallet() int64 {
	if m != nil {
		return m.AutoWallet
	}
	return 0
}

// 钱包转入请求
type SetWalletReq struct {
	Chips  uint64 `protobuf:"varint,1,req,name=chips" json:"chips"`
	IsAuto bool   `protobuf:"varint,2,req,name=is_auto,json=isAuto" json:"is_auto"`
}

func (m *SetWalletReq) Reset()         { *m = SetWalletReq{} }
func (m *SetWalletReq) String() string { return proto.CompactTextString(m) }
func (*SetWalletReq) ProtoMessage()    {}
func (*SetWalletReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{7}
}
func (m *SetWalletReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetWalletReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetWalletReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetWalletReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetWalletReq.Merge(dst, src)
}
func (m *SetWalletReq) XXX_Size() int {
	return m.Size()
}
func (m *SetWalletReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetWalletReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetWalletReq proto.InternalMessageInfo

func (m *SetWalletReq) GetChips() uint64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func (m *SetWalletReq) GetIsAuto() bool {
	if m != nil {
		return m.IsAuto
	}
	return false
}

// 钱包转入返回
type SetWalletResp struct {
	Code  RoleCode `protobuf:"varint,1,req,name=code,enum=suoha.RoleCode" json:"code"`
	Chips int64    `protobuf:"varint,2,opt,name=chips" json:"chips"`
}

func (m *SetWalletResp) Reset()         { *m = SetWalletResp{} }
func (m *SetWalletResp) String() string { return proto.CompactTextString(m) }
func (*SetWalletResp) ProtoMessage()    {}
func (*SetWalletResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_1ceffa6ba2e3a8f6, []int{8}
}
func (m *SetWalletResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetWalletResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetWalletResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetWalletResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetWalletResp.Merge(dst, src)
}
func (m *SetWalletResp) XXX_Size() int {
	return m.Size()
}
func (m *SetWalletResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SetWalletResp.DiscardUnknown(m)
}

var xxx_messageInfo_SetWalletResp proto.InternalMessageInfo

func (m *SetWalletResp) GetCode() RoleCode {
	if m != nil {
		return m.Code
	}
	return RoleCode_role_reserver
}

func (m *SetWalletResp) GetChips() int64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func init() {
	proto.RegisterType((*ChipsPush)(nil), "suoha.ChipsPush")
	proto.RegisterType((*SetAvatarReq)(nil), "suoha.SetAvatarReq")
	proto.RegisterType((*SetAvatarResp)(nil), "suoha.SetAvatarResp")
	proto.RegisterType((*GetChipsReq)(nil), "suoha.GetChipsReq")
	proto.RegisterType((*UpdateGtReq)(nil), "suoha.UpdateGtReq")
	proto.RegisterType((*GetWalletReq)(nil), "suoha.GetWalletReq")
	proto.RegisterType((*GetWalletResp)(nil), "suoha.GetWalletResp")
	proto.RegisterType((*SetWalletReq)(nil), "suoha.SetWalletReq")
	proto.RegisterType((*SetWalletResp)(nil), "suoha.SetWalletResp")
	proto.RegisterEnum("suoha.RoleCode", RoleCode_name, RoleCode_value)
}
func (m *ChipsPush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChipsPush) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Chips))
	return i, nil
}

func (m *SetAvatarReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetAvatarReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Avatar))
	return i, nil
}

func (m *SetAvatarResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetAvatarResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Avatar))
	return i, nil
}

func (m *GetChipsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChipsReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UpdateGtReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateGtReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Gt))
	return i, nil
}

func (m *GetWalletReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetWalletReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetWalletResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetWalletResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Code))
	dAtA[i] = 0x10
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Chips))
	dAtA[i] = 0x18
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.AutoWallet))
	return i, nil
}

func (m *SetWalletReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetWalletReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Chips))
	dAtA[i] = 0x10
	i++
	if m.IsAuto {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *SetWalletResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetWalletResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Code))
	dAtA[i] = 0x10
	i++
	i = encodeVarintMsgRole(dAtA, i, uint64(m.Chips))
	return i, nil
}

func encodeVarintMsgRole(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ChipsPush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgRole(uint64(m.Chips))
	return n
}

func (m *SetAvatarReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgRole(uint64(m.Avatar))
	return n
}

func (m *SetAvatarResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgRole(uint64(m.Avatar))
	return n
}

func (m *GetChipsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UpdateGtReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgRole(uint64(m.Gt))
	return n
}

func (m *GetWalletReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetWalletResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgRole(uint64(m.Code))
	n += 1 + sovMsgRole(uint64(m.Chips))
	n += 1 + sovMsgRole(uint64(m.AutoWallet))
	return n
}

func (m *SetWalletReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgRole(uint64(m.Chips))
	n += 2
	return n
}

func (m *SetWalletResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgRole(uint64(m.Code))
	n += 1 + sovMsgRole(uint64(m.Chips))
	return n
}

func sovMsgRole(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMsgRole(x uint64) (n int) {
	return sovMsgRole(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChipsPush) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChipsPush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChipsPush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chips", wireType)
			}
			m.Chips = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chips |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("chips")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetAvatarReq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetAvatarReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetAvatarReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			m.Avatar = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Avatar |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("avatar")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetAvatarResp) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetAvatarResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetAvatarResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			m.Avatar = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Avatar |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("avatar")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetChipsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetChipsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChipsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdateGtReq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateGtReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateGtReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gt", wireType)
			}
			m.Gt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gt |= (GameTerminalType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("gt")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetWalletReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetWalletReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetWalletReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetWalletResp) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetWalletResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetWalletResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (RoleCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chips", wireType)
			}
			m.Chips = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chips |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoWallet", wireType)
			}
			m.AutoWallet = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AutoWallet |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("code")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetWalletReq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetWalletReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetWalletReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chips", wireType)
			}
			m.Chips = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chips |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsAuto", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsAuto = bool(v != 0)
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("chips")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("is_auto")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetWalletResp) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetWalletResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetWalletResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (RoleCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chips", wireType)
			}
			m.Chips = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chips |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgRole(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgRole
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("code")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsgRole(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgRole
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgRole
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMsgRole
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMsgRole
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMsgRole(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMsgRole = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgRole   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("msg_role.proto", fileDescriptor_msg_role_1ceffa6ba2e3a8f6) }

var fileDescriptor_msg_role_1ceffa6ba2e3a8f6 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xc1, 0xca, 0xd3, 0x40,
	0x10, 0xc7, 0x93, 0x34, 0xf5, 0x6b, 0x27, 0x4d, 0x88, 0x0b, 0x42, 0xfc, 0xd0, 0x58, 0x16, 0xc4,
	0xf2, 0xa1, 0x15, 0x7c, 0x00, 0xe1, 0xab, 0x87, 0xe2, 0x4d, 0x2a, 0xd2, 0x63, 0x58, 0xd3, 0x69,
	0x12, 0x48, 0xb2, 0xe9, 0xee, 0xa6, 0xd2, 0x8b, 0xcf, 0xe0, 0x63, 0xf5, 0xd8, 0xa3, 0x27, 0x91,
	0xf6, 0x45, 0x24, 0xdb, 0xa4, 0xd6, 0x8b, 0x5e, 0x3c, 0xfe, 0x67, 0xfe, 0xf3, 0x9f, 0x99, 0x1f,
	0x78, 0x85, 0x4c, 0x22, 0xc1, 0x73, 0x9c, 0x56, 0x82, 0x2b, 0x4e, 0xfa, 0xb2, 0xe6, 0x29, 0xbb,
	0x7d, 0xd4, 0x95, 0xa3, 0x98, 0x17, 0x05, 0x2f, 0xcf, 0x5d, 0xfa, 0x02, 0x86, 0xef, 0xd2, 0xac,
	0x92, 0x1f, 0x6a, 0x99, 0x92, 0x5b, 0xe8, 0xc7, 0x8d, 0x08, 0xcc, 0xb1, 0x35, 0xb1, 0x67, 0xf6,
	0xfe, 0xc7, 0x33, 0x63, 0x71, 0x2e, 0xd1, 0x97, 0x30, 0xfa, 0x88, 0xea, 0x7e, 0xcb, 0x14, 0x13,
	0x0b, 0xdc, 0x90, 0x27, 0xf0, 0x80, 0x69, 0xa1, 0xcd, 0x6e, 0x6b, 0x6e, 0x6b, 0xf4, 0x15, 0xb8,
	0x57, 0x6e, 0x59, 0xfd, 0xc3, 0xee, 0x82, 0x33, 0x47, 0xa5, 0x0f, 0x59, 0xe0, 0x86, 0xbe, 0x05,
	0xe7, 0x53, 0xb5, 0x62, 0x0a, 0xe7, 0xaa, 0x59, 0xf5, 0x1a, 0xac, 0x44, 0xe9, 0x39, 0xef, 0xcd,
	0xe3, 0xa9, 0x7e, 0x67, 0x9a, 0xb0, 0x02, 0x23, 0x85, 0xa2, 0xc8, 0x4a, 0x96, 0x47, 0x6a, 0x57,
	0x61, 0x1b, 0x69, 0x25, 0x8a, 0x7a, 0x30, 0x9a, 0xa3, 0x5a, 0xb2, 0x3c, 0xc7, 0x26, 0x80, 0x7e,
	0x05, 0xf7, 0x4a, 0xcb, 0x8a, 0xdc, 0x81, 0x1d, 0xf3, 0x15, 0xb6, 0x99, 0x7e, 0x9b, 0xd9, 0xd2,
	0x59, 0x75, 0x51, 0xda, 0xf3, 0x1b, 0x8a, 0x35, 0x36, 0x27, 0xbd, 0x3f, 0xa0, 0x90, 0xe7, 0xe0,
	0xb0, 0x5a, 0xf1, 0xe8, 0x8b, 0x8e, 0x0e, 0x7a, 0x57, 0x0e, 0x68, 0x1a, 0xe7, 0x95, 0xf4, 0xbd,
	0x66, 0x77, 0xb9, 0xe7, 0x6f, 0x9c, 0xc9, 0x53, 0xb8, 0xc9, 0x64, 0xd4, 0x0c, 0x07, 0xd6, 0xd8,
	0x9a, 0x0c, 0x3a, 0x52, 0x99, 0xbc, 0xaf, 0x15, 0xa7, 0x4b, 0x0d, 0xf6, 0xff, 0xbf, 0x72, 0xf7,
	0x19, 0x86, 0x97, 0x21, 0xf2, 0x10, 0x5c, 0x2d, 0x04, 0x4a, 0x14, 0x5b, 0x14, 0xbe, 0x41, 0x1c,
	0xb8, 0x91, 0x75, 0x1c, 0xa3, 0x94, 0xbe, 0x49, 0x06, 0x60, 0xaf, 0x59, 0x96, 0xfb, 0x16, 0x19,
	0xc1, 0x60, 0x2d, 0x70, 0x53, 0x63, 0xa9, 0xfc, 0x1e, 0x71, 0x61, 0xd8, 0xa9, 0x9d, 0x6f, 0x13,
	0x0f, 0xa0, 0xe4, 0x2a, 0xc2, 0x92, 0xd7, 0x49, 0xea, 0xf7, 0x67, 0xc1, 0xfe, 0x18, 0x9a, 0x87,
	0x63, 0x68, 0xfe, 0x3c, 0x86, 0xe6, 0xb7, 0x53, 0x68, 0x1c, 0x4e, 0xa1, 0xf1, 0xfd, 0x14, 0x1a,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x0d, 0x3f, 0xd7, 0xb5, 0x02, 0x00, 0x00,
}
