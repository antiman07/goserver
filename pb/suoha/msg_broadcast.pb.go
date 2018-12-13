// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg_broadcast.proto

package suoha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// 全局广播信息
type BcastworldResp struct {
	Content string `protobuf:"bytes,1,req,name=content" json:"content"`
}

func (m *BcastworldResp) Reset()         { *m = BcastworldResp{} }
func (m *BcastworldResp) String() string { return proto.CompactTextString(m) }
func (*BcastworldResp) ProtoMessage()    {}
func (*BcastworldResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_broadcast_0572a5c704e92824, []int{0}
}
func (m *BcastworldResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BcastworldResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BcastworldResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BcastworldResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BcastworldResp.Merge(dst, src)
}
func (m *BcastworldResp) XXX_Size() int {
	return m.Size()
}
func (m *BcastworldResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BcastworldResp.DiscardUnknown(m)
}

var xxx_messageInfo_BcastworldResp proto.InternalMessageInfo

func (m *BcastworldResp) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// 广播消息 多语言
type BcastMsgResp struct {
	Msg *Dmsg `protobuf:"bytes,1,req,name=msg" json:"msg,omitempty"`
}

func (m *BcastMsgResp) Reset()         { *m = BcastMsgResp{} }
func (m *BcastMsgResp) String() string { return proto.CompactTextString(m) }
func (*BcastMsgResp) ProtoMessage()    {}
func (*BcastMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_broadcast_0572a5c704e92824, []int{1}
}
func (m *BcastMsgResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BcastMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BcastMsgResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BcastMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BcastMsgResp.Merge(dst, src)
}
func (m *BcastMsgResp) XXX_Size() int {
	return m.Size()
}
func (m *BcastMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BcastMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_BcastMsgResp proto.InternalMessageInfo

func (m *BcastMsgResp) GetMsg() *Dmsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*BcastworldResp)(nil), "suoha.BcastworldResp")
	proto.RegisterType((*BcastMsgResp)(nil), "suoha.BcastMsgResp")
}
func (m *BcastworldResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BcastworldResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMsgBroadcast(dAtA, i, uint64(len(m.Content)))
	i += copy(dAtA[i:], m.Content)
	return i, nil
}

func (m *BcastMsgResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BcastMsgResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Msg == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("msg")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMsgBroadcast(dAtA, i, uint64(m.Msg.Size()))
		n1, err := m.Msg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintMsgBroadcast(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BcastworldResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	n += 1 + l + sovMsgBroadcast(uint64(l))
	return n
}

func (m *BcastMsgResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovMsgBroadcast(uint64(l))
	}
	return n
}

func sovMsgBroadcast(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMsgBroadcast(x uint64) (n int) {
	return sovMsgBroadcast(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BcastworldResp) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgBroadcast
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
			return fmt.Errorf("proto: BcastworldResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BcastworldResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBroadcast
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
				return ErrInvalidLengthMsgBroadcast
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgBroadcast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgBroadcast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("content")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BcastMsgResp) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgBroadcast
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
			return fmt.Errorf("proto: BcastMsgResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BcastMsgResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBroadcast
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
				return ErrInvalidLengthMsgBroadcast
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &Dmsg{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgBroadcast(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgBroadcast
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("msg")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsgBroadcast(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgBroadcast
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
					return 0, ErrIntOverflowMsgBroadcast
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
					return 0, ErrIntOverflowMsgBroadcast
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
				return 0, ErrInvalidLengthMsgBroadcast
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMsgBroadcast
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
				next, err := skipMsgBroadcast(dAtA[start:])
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
	ErrInvalidLengthMsgBroadcast = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgBroadcast   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("msg_broadcast.proto", fileDescriptor_msg_broadcast_0572a5c704e92824) }

var fileDescriptor_msg_broadcast_0572a5c704e92824 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2d, 0x4e, 0x8f,
	0x4f, 0x2a, 0xca, 0x4f, 0x4c, 0x49, 0x4e, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2d, 0x2e, 0xcd, 0xcf, 0x48, 0x94, 0x12, 0x00, 0xc9, 0x25, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7,
	0x41, 0x24, 0x94, 0x0c, 0xb8, 0xf8, 0x9c, 0x40, 0xea, 0xca, 0xf3, 0x8b, 0x72, 0x52, 0x82, 0x52,
	0x8b, 0x0b, 0x84, 0xe4, 0xb8, 0xd8, 0x93, 0xf3, 0xf3, 0x4a, 0x52, 0xf3, 0x4a, 0x24, 0x18, 0x15,
	0x98, 0x34, 0x38, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0x09, 0x2a, 0xe9, 0x72, 0xf1,
	0x80, 0x75, 0xf8, 0x16, 0xa7, 0x83, 0xd5, 0xcb, 0x72, 0x31, 0xe7, 0x16, 0xa7, 0x83, 0xd5, 0x72,
	0x1b, 0x71, 0xeb, 0x81, 0x2d, 0xd2, 0x73, 0xc9, 0x2d, 0x4e, 0x0f, 0x02, 0x89, 0x3b, 0x49, 0x9c,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x03, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xad,
	0x98, 0x96, 0xa9, 0x00, 0x00, 0x00,
}
