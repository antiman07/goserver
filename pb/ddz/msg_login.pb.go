// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg_login.proto

package ddz

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

type LoginModeCode int32

const (
	LoginModeCode_login_mode_code_reserve1 LoginModeCode = 0
	LoginModeCode_normal                   LoginModeCode = 1
	LoginModeCode_reconnect                LoginModeCode = 2
)

var LoginModeCode_name = map[int32]string{
	0: "login_mode_code_reserve1",
	1: "normal",
	2: "reconnect",
}
var LoginModeCode_value = map[string]int32{
	"login_mode_code_reserve1": 0,
	"normal":                   1,
	"reconnect":                2,
}

func (x LoginModeCode) Enum() *LoginModeCode {
	p := new(LoginModeCode)
	*p = x
	return p
}
func (x LoginModeCode) String() string {
	return proto.EnumName(LoginModeCode_name, int32(x))
}
func (x *LoginModeCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LoginModeCode_value, data, "LoginModeCode")
	if err != nil {
		return err
	}
	*x = LoginModeCode(value)
	return nil
}
func (LoginModeCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_msg_login_637925b29ed15f70, []int{0}
}

type LoginCode int32

const (
	LoginCode_login_code_reserve LoginCode = 0
	LoginCode_login_success      LoginCode = 1
	LoginCode_login_fail         LoginCode = 2
	LoginCode_err_account        LoginCode = 3
	LoginCode_token_timeout      LoginCode = 4
	LoginCode_token_error        LoginCode = 5
	LoginCode_other_login        LoginCode = 6
	LoginCode_force_kick         LoginCode = 7
)

var LoginCode_name = map[int32]string{
	0: "login_code_reserve",
	1: "login_success",
	2: "login_fail",
	3: "err_account",
	4: "token_timeout",
	5: "token_error",
	6: "other_login",
	7: "force_kick",
}
var LoginCode_value = map[string]int32{
	"login_code_reserve": 0,
	"login_success":      1,
	"login_fail":         2,
	"err_account":        3,
	"token_timeout":      4,
	"token_error":        5,
	"other_login":        6,
	"force_kick":         7,
}

func (x LoginCode) Enum() *LoginCode {
	p := new(LoginCode)
	*p = x
	return p
}
func (x LoginCode) String() string {
	return proto.EnumName(LoginCode_name, int32(x))
}
func (x *LoginCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LoginCode_value, data, "LoginCode")
	if err != nil {
		return err
	}
	*x = LoginCode(value)
	return nil
}
func (LoginCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_msg_login_637925b29ed15f70, []int{1}
}

type LoginReq struct {
	Code  LoginModeCode    `protobuf:"varint,1,req,name=code,enum=ddz.LoginModeCode" json:"code"`
	Token string           `protobuf:"bytes,2,req,name=token" json:"token"`
	Lang  string           `protobuf:"bytes,3,req,name=lang" json:"lang"`
	Gt    GameTerminalType `protobuf:"varint,4,req,name=gt,enum=ddz.GameTerminalType" json:"gt"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_login_637925b29ed15f70, []int{0}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(dst, src)
}
func (m *LoginReq) XXX_Size() int {
	return m.Size()
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetCode() LoginModeCode {
	if m != nil {
		return m.Code
	}
	return LoginModeCode_login_mode_code_reserve1
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReq) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *LoginReq) GetGt() GameTerminalType {
	if m != nil {
		return m.Gt
	}
	return GameTerminalType_game_terminal_type_reserve
}

type LoginResp struct {
	Code LoginCode `protobuf:"varint,1,req,name=code,enum=ddz.LoginCode" json:"code"`
	Role *RoleInfo `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_login_637925b29ed15f70, []int{1}
}
func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(dst, src)
}
func (m *LoginResp) XXX_Size() int {
	return m.Size()
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetCode() LoginCode {
	if m != nil {
		return m.Code
	}
	return LoginCode_login_code_reserve
}

func (m *LoginResp) GetRole() *RoleInfo {
	if m != nil {
		return m.Role
	}
	return nil
}

type LoginAccountReq struct {
	Userid string `protobuf:"bytes,1,req,name=userid" json:"userid"`
	Token  string `protobuf:"bytes,2,req,name=token" json:"token"`
}

func (m *LoginAccountReq) Reset()         { *m = LoginAccountReq{} }
func (m *LoginAccountReq) String() string { return proto.CompactTextString(m) }
func (*LoginAccountReq) ProtoMessage()    {}
func (*LoginAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_login_637925b29ed15f70, []int{2}
}
func (m *LoginAccountReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginAccountReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LoginAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAccountReq.Merge(dst, src)
}
func (m *LoginAccountReq) XXX_Size() int {
	return m.Size()
}
func (m *LoginAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAccountReq proto.InternalMessageInfo

func (m *LoginAccountReq) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *LoginAccountReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type KickResp struct {
	Code LoginCode `protobuf:"varint,1,req,name=code,enum=ddz.LoginCode" json:"code"`
}

func (m *KickResp) Reset()         { *m = KickResp{} }
func (m *KickResp) String() string { return proto.CompactTextString(m) }
func (*KickResp) ProtoMessage()    {}
func (*KickResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_login_637925b29ed15f70, []int{3}
}
func (m *KickResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KickResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KickResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *KickResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickResp.Merge(dst, src)
}
func (m *KickResp) XXX_Size() int {
	return m.Size()
}
func (m *KickResp) XXX_DiscardUnknown() {
	xxx_messageInfo_KickResp.DiscardUnknown(m)
}

var xxx_messageInfo_KickResp proto.InternalMessageInfo

func (m *KickResp) GetCode() LoginCode {
	if m != nil {
		return m.Code
	}
	return LoginCode_login_code_reserve
}

type UpdateTokenResp struct {
	ReconnectToken string `protobuf:"bytes,1,opt,name=reconnect_token,json=reconnectToken" json:"reconnect_token"`
}

func (m *UpdateTokenResp) Reset()         { *m = UpdateTokenResp{} }
func (m *UpdateTokenResp) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenResp) ProtoMessage()    {}
func (*UpdateTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_login_637925b29ed15f70, []int{4}
}
func (m *UpdateTokenResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateTokenResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UpdateTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenResp.Merge(dst, src)
}
func (m *UpdateTokenResp) XXX_Size() int {
	return m.Size()
}
func (m *UpdateTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenResp proto.InternalMessageInfo

func (m *UpdateTokenResp) GetReconnectToken() string {
	if m != nil {
		return m.ReconnectToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "ddz.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "ddz.LoginResp")
	proto.RegisterType((*LoginAccountReq)(nil), "ddz.LoginAccountReq")
	proto.RegisterType((*KickResp)(nil), "ddz.KickResp")
	proto.RegisterType((*UpdateTokenResp)(nil), "ddz.UpdateTokenResp")
	proto.RegisterEnum("ddz.LoginModeCode", LoginModeCode_name, LoginModeCode_value)
	proto.RegisterEnum("ddz.LoginCode", LoginCode_name, LoginCode_value)
}
func (m *LoginReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(m.Code))
	dAtA[i] = 0x12
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(len(m.Token)))
	i += copy(dAtA[i:], m.Token)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(len(m.Lang)))
	i += copy(dAtA[i:], m.Lang)
	dAtA[i] = 0x20
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(m.Gt))
	return i, nil
}

func (m *LoginResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(m.Code))
	if m.Role != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMsgLogin(dAtA, i, uint64(m.Role.Size()))
		n1, err := m.Role.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *LoginAccountReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginAccountReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(len(m.Userid)))
	i += copy(dAtA[i:], m.Userid)
	dAtA[i] = 0x12
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(len(m.Token)))
	i += copy(dAtA[i:], m.Token)
	return i, nil
}

func (m *KickResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KickResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(m.Code))
	return i, nil
}

func (m *UpdateTokenResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateTokenResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMsgLogin(dAtA, i, uint64(len(m.ReconnectToken)))
	i += copy(dAtA[i:], m.ReconnectToken)
	return i, nil
}

func encodeVarintMsgLogin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoginReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgLogin(uint64(m.Code))
	l = len(m.Token)
	n += 1 + l + sovMsgLogin(uint64(l))
	l = len(m.Lang)
	n += 1 + l + sovMsgLogin(uint64(l))
	n += 1 + sovMsgLogin(uint64(m.Gt))
	return n
}

func (m *LoginResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgLogin(uint64(m.Code))
	if m.Role != nil {
		l = m.Role.Size()
		n += 1 + l + sovMsgLogin(uint64(l))
	}
	return n
}

func (m *LoginAccountReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Userid)
	n += 1 + l + sovMsgLogin(uint64(l))
	l = len(m.Token)
	n += 1 + l + sovMsgLogin(uint64(l))
	return n
}

func (m *KickResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovMsgLogin(uint64(m.Code))
	return n
}

func (m *UpdateTokenResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ReconnectToken)
	n += 1 + l + sovMsgLogin(uint64(l))
	return n
}

func sovMsgLogin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMsgLogin(x uint64) (n int) {
	return sovMsgLogin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginReq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgLogin
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
			return fmt.Errorf("proto: LoginReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (LoginModeCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lang = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gt", wireType)
			}
			m.Gt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
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
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgLogin
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
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("token")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("lang")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("gt")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoginResp) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgLogin
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
			return fmt.Errorf("proto: LoginResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (LoginCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgLogin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Role == nil {
				m.Role = &RoleInfo{}
			}
			if err := m.Role.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgLogin
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
func (m *LoginAccountReq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgLogin
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
			return fmt.Errorf("proto: LoginAccountReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginAccountReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("userid")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("token")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *KickResp) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgLogin
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
			return fmt.Errorf("proto: KickResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KickResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (LoginCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgLogin
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
func (m *UpdateTokenResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgLogin
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
			return fmt.Errorf("proto: UpdateTokenResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateTokenResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReconnectToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReconnectToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgLogin
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
func skipMsgLogin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgLogin
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
					return 0, ErrIntOverflowMsgLogin
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
					return 0, ErrIntOverflowMsgLogin
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
				return 0, ErrInvalidLengthMsgLogin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMsgLogin
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
				next, err := skipMsgLogin(dAtA[start:])
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
	ErrInvalidLengthMsgLogin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgLogin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("msg_login.proto", fileDescriptor_msg_login_637925b29ed15f70) }

var fileDescriptor_msg_login_637925b29ed15f70 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xf6, 0xba, 0x6e, 0x68, 0xa6, 0x4a, 0xbc, 0xac, 0xf8, 0xb1, 0xaa, 0xca, 0x04, 0x9f, 0x42,
	0xa5, 0x46, 0xa2, 0x12, 0x77, 0xe8, 0x0d, 0xca, 0x29, 0x82, 0x03, 0xa7, 0x95, 0xb5, 0x9e, 0x18,
	0x2b, 0xf6, 0x4e, 0x58, 0x6f, 0x90, 0xe8, 0x53, 0xf0, 0x02, 0x88, 0xd7, 0xe9, 0xb1, 0x47, 0x4e,
	0x08, 0x25, 0x2f, 0x82, 0x76, 0x9d, 0xa4, 0xa5, 0x48, 0x48, 0x1c, 0xfd, 0x7d, 0xf3, 0xfd, 0xcc,
	0x78, 0x21, 0x6e, 0xda, 0x52, 0xd6, 0x54, 0x56, 0x7a, 0xb2, 0x30, 0x64, 0x49, 0xec, 0x15, 0xc5,
	0xe5, 0xd1, 0x43, 0x87, 0x1a, 0xaa, 0x51, 0x2a, 0x6a, 0x1a, 0xda, 0x70, 0xd9, 0x77, 0x06, 0x07,
	0x6f, 0xdd, 0xec, 0x14, 0x3f, 0x89, 0x09, 0x44, 0x8a, 0x0a, 0x4c, 0xd8, 0x28, 0x1c, 0x0f, 0xcf,
	0x1e, 0x4c, 0x8a, 0xe2, 0x72, 0xe2, 0x8d, 0x64, 0x43, 0x85, 0x13, 0x16, 0x78, 0x1e, 0x5d, 0xfd,
	0x7c, 0x12, 0x4c, 0xfd, 0x9c, 0x38, 0x82, 0x7d, 0x4b, 0x73, 0xd4, 0x49, 0x38, 0x0a, 0xc7, 0xfd,
	0x0d, 0xd5, 0x41, 0x22, 0x81, 0xa8, 0xce, 0x75, 0x99, 0xec, 0xdd, 0xa2, 0x3c, 0x22, 0x4e, 0x21,
	0x2c, 0x6d, 0x12, 0xf9, 0x8c, 0xc7, 0x3e, 0xa3, 0xcc, 0x1b, 0x94, 0x16, 0x4d, 0x53, 0xe9, 0xbc,
	0x96, 0xf6, 0xcb, 0x62, 0x1b, 0x13, 0x96, 0x36, 0xfb, 0x00, 0xfd, 0x4d, 0xc1, 0x76, 0x21, 0x9e,
	0xfd, 0xd1, 0x30, 0xbe, 0xd5, 0xf0, 0xaf, 0x72, 0x4f, 0x21, 0x72, 0xeb, 0x26, 0xe1, 0x88, 0x8d,
	0x0f, 0xcf, 0x06, 0x7e, 0x74, 0x4a, 0x35, 0xbe, 0xd6, 0x33, 0x9a, 0x7a, 0x2a, 0xbb, 0x80, 0xd8,
	0x5b, 0xbf, 0x52, 0x8a, 0x96, 0xda, 0xba, 0x13, 0x1c, 0x43, 0x6f, 0xd9, 0xa2, 0xa9, 0x0a, 0x1f,
	0xb1, 0x2d, 0xbe, 0xc1, 0xfe, 0xb5, 0x70, 0xf6, 0x02, 0x0e, 0x2e, 0x2a, 0x35, 0xff, 0xcf, 0x9a,
	0xd9, 0x4b, 0x88, 0xdf, 0x2f, 0x8a, 0xdc, 0xe2, 0x3b, 0xe7, 0xe2, 0xd5, 0xa7, 0x10, 0x1b, 0x54,
	0xa4, 0x35, 0x2a, 0x2b, 0xbb, 0x3c, 0x36, 0x62, 0xbb, 0xbc, 0xe1, 0x8e, 0xf4, 0x92, 0x93, 0x37,
	0x10, 0xdf, 0xf9, 0x49, 0xe2, 0x18, 0x92, 0x3b, 0x90, 0x34, 0xd8, 0xa2, 0xf9, 0x8c, 0xcf, 0x79,
	0x20, 0x00, 0x7a, 0x9a, 0x4c, 0x93, 0xd7, 0x9c, 0x89, 0x01, 0xf4, 0x77, 0x76, 0x3c, 0x3c, 0xf9,
	0xc6, 0x00, 0x6e, 0x8a, 0x8a, 0x47, 0x20, 0x6e, 0xbe, 0xb6, 0x16, 0x3c, 0x10, 0xf7, 0x61, 0xd0,
	0xe1, 0xed, 0x52, 0x29, 0x6c, 0x5b, 0xce, 0xc4, 0x70, 0x2b, 0x9c, 0xe5, 0x55, 0xcd, 0x43, 0x11,
	0xc3, 0x21, 0x1a, 0x23, 0xf3, 0xee, 0xb4, 0x7c, 0xcf, 0x69, 0xfc, 0x2e, 0xd2, 0x56, 0x0d, 0xd2,
	0xd2, 0xf2, 0xc8, 0xcd, 0x74, 0x10, 0x1a, 0x43, 0x86, 0xef, 0x3b, 0x80, 0xec, 0x47, 0x34, 0xdd,
	0xf3, 0xe5, 0x3d, 0xe7, 0x3a, 0x23, 0xa3, 0x50, 0xce, 0x2b, 0x35, 0xe7, 0xf7, 0xce, 0x93, 0xab,
	0x55, 0xca, 0xae, 0x57, 0x29, 0xfb, 0xb5, 0x4a, 0xd9, 0xd7, 0x75, 0x1a, 0x5c, 0xaf, 0xd3, 0xe0,
	0xc7, 0x3a, 0x0d, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x5d, 0x2b, 0x41, 0xf6, 0x02, 0x00,
	0x00,
}