// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg_role.proto

package tbnn

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
	RoleCode_role_code_reserve RoleCode = 0
	RoleCode_role_success      RoleCode = 1
	RoleCode_role_fail         RoleCode = 2
	RoleCode_frequent          RoleCode = 3
	RoleCode_frequenty         RoleCode = 4
	RoleCode_not_enough        RoleCode = 5
)

var RoleCode_name = map[int32]string{
	0: "role_code_reserve",
	1: "role_success",
	2: "role_fail",
	3: "frequent",
	4: "frequenty",
	5: "not_enough",
}
var RoleCode_value = map[string]int32{
	"role_code_reserve": 0,
	"role_success":      1,
	"role_fail":         2,
	"frequent":          3,
	"frequenty":         4,
	"not_enough":        5,
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
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{0}
}

// 推送金额
type ChipsPush struct {
	Chips uint64 `protobuf:"varint,1,req,name=chips" json:"chips"`
}

func (m *ChipsPush) Reset()         { *m = ChipsPush{} }
func (m *ChipsPush) String() string { return proto.CompactTextString(m) }
func (*ChipsPush) ProtoMessage()    {}
func (*ChipsPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{0}
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
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{1}
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
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{2}
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
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{3}
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
type UpdateGetReq struct {
	Gt GameTerminalType `protobuf:"varint,1,req,name=gt,enum=tbnn.GameTerminalType" json:"gt"`
}

func (m *UpdateGetReq) Reset()         { *m = UpdateGetReq{} }
func (m *UpdateGetReq) String() string { return proto.CompactTextString(m) }
func (*UpdateGetReq) ProtoMessage()    {}
func (*UpdateGetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{4}
}
func (m *UpdateGetReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateGetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateGetReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UpdateGetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGetReq.Merge(dst, src)
}
func (m *UpdateGetReq) XXX_Size() int {
	return m.Size()
}
func (m *UpdateGetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGetReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGetReq proto.InternalMessageInfo

func (m *UpdateGetReq) GetGt() GameTerminalType {
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
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{5}
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
	Code       RoleCode `protobuf:"varint,1,req,name=code,enum=tbnn.RoleCode" json:"code"`
	Chips      int64    `protobuf:"varint,2,opt,name=chips" json:"chips"`
	AutoWallet int64    `protobuf:"varint,3,opt,name=auto_wallet,json=autoWallet" json:"auto_wallet"`
}

func (m *GetWalletResp) Reset()         { *m = GetWalletResp{} }
func (m *GetWalletResp) String() string { return proto.CompactTextString(m) }
func (*GetWalletResp) ProtoMessage()    {}
func (*GetWalletResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{6}
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
	return RoleCode_role_code_reserve
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
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{7}
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
	Code  RoleCode `protobuf:"varint,1,req,name=code,enum=tbnn.RoleCode" json:"code"`
	Chips int64    `protobuf:"varint,2,opt,name=chips" json:"chips"`
}

func (m *SetWalletResp) Reset()         { *m = SetWalletResp{} }
func (m *SetWalletResp) String() string { return proto.CompactTextString(m) }
func (*SetWalletResp) ProtoMessage()    {}
func (*SetWalletResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_role_673eca56b68ebad3, []int{8}
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
	return RoleCode_role_code_reserve
}

func (m *SetWalletResp) GetChips() int64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func init() {
	proto.RegisterType((*ChipsPush)(nil), "tbnn.ChipsPush")
	proto.RegisterType((*SetAvatarReq)(nil), "tbnn.SetAvatarReq")
	proto.RegisterType((*SetAvatarResp)(nil), "tbnn.SetAvatarResp")
	proto.RegisterType((*GetChipsReq)(nil), "tbnn.GetChipsReq")
	proto.RegisterType((*UpdateGetReq)(nil), "tbnn.UpdateGetReq")
	proto.RegisterType((*GetWalletReq)(nil), "tbnn.GetWalletReq")
	proto.RegisterType((*GetWalletResp)(nil), "tbnn.GetWalletResp")
	proto.RegisterType((*SetWalletReq)(nil), "tbnn.SetWalletReq")
	proto.RegisterType((*SetWalletResp)(nil), "tbnn.SetWalletResp")
	proto.RegisterEnum("tbnn.RoleCode", RoleCode_name, RoleCode_value)
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

func (m *UpdateGetReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateGetReq) MarshalTo(dAtA []byte) (int, error) {
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

func (m *UpdateGetReq) Size() (n int) {
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
func (m *UpdateGetReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpdateGetReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateGetReq: illegal tag %d (wire type %d)", fieldNum, wire)
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

func init() { proto.RegisterFile("msg_role.proto", fileDescriptor_msg_role_673eca56b68ebad3) }

var fileDescriptor_msg_role_673eca56b68ebad3 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0x1d, 0x17, 0xda, 0xa9, 0x6d, 0xcc, 0x4a, 0x95, 0xac, 0x0a, 0x4c, 0xb4, 0x12, 0x22,
	0x20, 0xf0, 0x81, 0x0f, 0x40, 0x6a, 0x39, 0x44, 0xdc, 0x50, 0x10, 0x70, 0x5c, 0x2d, 0xce, 0xd4,
	0xb1, 0x64, 0x7b, 0x1d, 0xef, 0xb8, 0xa8, 0x07, 0xfe, 0x81, 0xcf, 0xea, 0x31, 0x47, 0x4e, 0x08,
	0x25, 0x3f, 0x82, 0xbc, 0xb6, 0x43, 0xb8, 0xc0, 0xa5, 0xc7, 0xf7, 0xe6, 0xbd, 0x37, 0x33, 0x0f,
	0x82, 0x52, 0x67, 0xa2, 0x51, 0x05, 0x26, 0x75, 0xa3, 0x48, 0x31, 0x97, 0xbe, 0x54, 0xd5, 0xf9,
	0xd9, 0xc8, 0x8a, 0x54, 0x95, 0xa5, 0xaa, 0xfa, 0x21, 0x7f, 0x06, 0x27, 0x6f, 0x57, 0x79, 0xad,
	0xdf, 0xb7, 0x7a, 0xc5, 0xce, 0xe1, 0x28, 0xed, 0x40, 0x64, 0x4f, 0x9d, 0x99, 0x7b, 0xe9, 0xde,
	0xfe, 0x7c, 0x62, 0x2d, 0x7a, 0x8a, 0xbf, 0x04, 0xef, 0x03, 0xd2, 0xc5, 0xb5, 0x24, 0xd9, 0x2c,
	0x70, 0xcd, 0x1e, 0xc1, 0x3d, 0x69, 0x80, 0x11, 0xfb, 0x83, 0x78, 0xe0, 0xf8, 0x2b, 0xf0, 0x0f,
	0xd4, 0xba, 0xfe, 0x8f, 0xdc, 0x87, 0xd3, 0x39, 0x92, 0x39, 0x64, 0x81, 0x6b, 0xfe, 0x06, 0xbc,
	0x8f, 0xf5, 0x52, 0x12, 0xce, 0x91, 0xba, 0x5d, 0x09, 0x38, 0x19, 0x19, 0x63, 0xf0, 0x3a, 0x4a,
	0xba, 0x77, 0x92, 0x4c, 0x96, 0x28, 0x08, 0x9b, 0x32, 0xaf, 0x64, 0x21, 0xe8, 0xa6, 0xc6, 0x21,
	0xd2, 0xc9, 0x88, 0x07, 0xe0, 0xcd, 0x91, 0x3e, 0xcb, 0xa2, 0x30, 0x7e, 0xfe, 0x0d, 0xfc, 0x03,
	0xac, 0x6b, 0xf6, 0x1c, 0xdc, 0x54, 0x2d, 0x71, 0x88, 0x7c, 0xd0, 0x47, 0x0e, 0xe5, 0x2c, 0xc7,
	0x24, 0x23, 0xf9, 0xd3, 0x89, 0x33, 0xb5, 0x67, 0x93, 0xbf, 0x3a, 0x61, 0x4f, 0xe1, 0x54, 0xb6,
	0xa4, 0xc4, 0x57, 0x93, 0x1c, 0x4d, 0x0e, 0x14, 0xd0, 0x0d, 0xfa, 0x8d, 0xfc, 0x9d, 0xa9, 0x6e,
	0x7f, 0xce, 0xbf, 0x6a, 0x66, 0x8f, 0xe1, 0x7e, 0xae, 0x45, 0x67, 0x8e, 0x9c, 0xa9, 0x33, 0x3b,
	0x1e, 0x8b, 0xca, 0xf5, 0x45, 0x4b, 0x8a, 0x7f, 0x32, 0xbd, 0xde, 0xf9, 0x27, 0x2f, 0x6a, 0x38,
	0xd9, 0x9b, 0xd8, 0x19, 0x3c, 0xdc, 0x03, 0xd1, 0xa0, 0xc6, 0xe6, 0x1a, 0x43, 0x8b, 0x85, 0xe0,
	0x19, 0x5a, 0xb7, 0x69, 0x8a, 0x5a, 0x87, 0x36, 0xf3, 0x07, 0xd7, 0x95, 0xcc, 0x8b, 0xd0, 0x61,
	0x1e, 0x1c, 0x5f, 0x35, 0xb8, 0x6e, 0xb1, 0xa2, 0x70, 0xd2, 0x0d, 0x47, 0x74, 0x13, 0xba, 0x2c,
	0x00, 0xa8, 0x14, 0x09, 0xac, 0x54, 0x9b, 0xad, 0xc2, 0xa3, 0xcb, 0xe8, 0x76, 0x1b, 0xdb, 0x9b,
	0x6d, 0x6c, 0xff, 0xda, 0xc6, 0xf6, 0xf7, 0x5d, 0x6c, 0x6d, 0x76, 0xb1, 0xf5, 0x63, 0x17, 0x5b,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x02, 0x1a, 0xea, 0x7a, 0xc0, 0x02, 0x00, 0x00,
}