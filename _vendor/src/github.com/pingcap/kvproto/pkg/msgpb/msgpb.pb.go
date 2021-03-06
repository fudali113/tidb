// Code generated by protoc-gen-gogo.
// source: msgpb.proto
// DO NOT EDIT!

/*
	Package msgpb is a generated protocol buffer package.

	It is generated from these files:
		msgpb.proto

	It has these top-level messages:
		Message
*/
package msgpb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	raft_cmdpb "github.com/pingcap/kvproto/pkg/raft_cmdpb"

	raft_serverpb "github.com/pingcap/kvproto/pkg/raft_serverpb"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	coprocessor "github.com/pingcap/kvproto/pkg/coprocessor"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_None    MessageType = 0
	MessageType_Cmd     MessageType = 1
	MessageType_CmdResp MessageType = 2
	MessageType_Raft    MessageType = 3
	MessageType_KvReq   MessageType = 4
	MessageType_KvResp  MessageType = 5
	MessageType_CopReq  MessageType = 6
	MessageType_CopResp MessageType = 7
)

var MessageType_name = map[int32]string{
	0: "None",
	1: "Cmd",
	2: "CmdResp",
	3: "Raft",
	4: "KvReq",
	5: "KvResp",
	6: "CopReq",
	7: "CopResp",
}
var MessageType_value = map[string]int32{
	"None":    0,
	"Cmd":     1,
	"CmdResp": 2,
	"Raft":    3,
	"KvReq":   4,
	"KvResp":  5,
	"CopReq":  6,
	"CopResp": 7,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMsgpb, []int{0} }

// Message holds all messages communicating with TiKV.
type Message struct {
	MsgType          MessageType                 `protobuf:"varint,1,opt,name=msg_type,json=msgType,enum=msgpb.MessageType" json:"msg_type"`
	CmdReq           *raft_cmdpb.RaftCmdRequest  `protobuf:"bytes,2,opt,name=cmd_req,json=cmdReq" json:"cmd_req,omitempty"`
	CmdResp          *raft_cmdpb.RaftCmdResponse `protobuf:"bytes,3,opt,name=cmd_resp,json=cmdResp" json:"cmd_resp,omitempty"`
	Raft             *raft_serverpb.RaftMessage  `protobuf:"bytes,4,opt,name=raft" json:"raft,omitempty"`
	KvReq            *kvrpcpb.Request            `protobuf:"bytes,5,opt,name=kv_req,json=kvReq" json:"kv_req,omitempty"`
	KvResp           *kvrpcpb.Response           `protobuf:"bytes,6,opt,name=kv_resp,json=kvResp" json:"kv_resp,omitempty"`
	CopReq           *coprocessor.Request        `protobuf:"bytes,7,opt,name=cop_req,json=copReq" json:"cop_req,omitempty"`
	CopResp          *coprocessor.Response       `protobuf:"bytes,8,opt,name=cop_resp,json=copResp" json:"cop_resp,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorMsgpb, []int{0} }

func (m *Message) GetMsgType() MessageType {
	if m != nil {
		return m.MsgType
	}
	return MessageType_None
}

func (m *Message) GetCmdReq() *raft_cmdpb.RaftCmdRequest {
	if m != nil {
		return m.CmdReq
	}
	return nil
}

func (m *Message) GetCmdResp() *raft_cmdpb.RaftCmdResponse {
	if m != nil {
		return m.CmdResp
	}
	return nil
}

func (m *Message) GetRaft() *raft_serverpb.RaftMessage {
	if m != nil {
		return m.Raft
	}
	return nil
}

func (m *Message) GetKvReq() *kvrpcpb.Request {
	if m != nil {
		return m.KvReq
	}
	return nil
}

func (m *Message) GetKvResp() *kvrpcpb.Response {
	if m != nil {
		return m.KvResp
	}
	return nil
}

func (m *Message) GetCopReq() *coprocessor.Request {
	if m != nil {
		return m.CopReq
	}
	return nil
}

func (m *Message) GetCopResp() *coprocessor.Response {
	if m != nil {
		return m.CopResp
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "msgpb.Message")
	proto.RegisterEnum("msgpb.MessageType", MessageType_name, MessageType_value)
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintMsgpb(dAtA, i, uint64(m.MsgType))
	if m.CmdReq != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMsgpb(dAtA, i, uint64(m.CmdReq.Size()))
		n1, err := m.CmdReq.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.CmdResp != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMsgpb(dAtA, i, uint64(m.CmdResp.Size()))
		n2, err := m.CmdResp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Raft != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMsgpb(dAtA, i, uint64(m.Raft.Size()))
		n3, err := m.Raft.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.KvReq != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMsgpb(dAtA, i, uint64(m.KvReq.Size()))
		n4, err := m.KvReq.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.KvResp != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMsgpb(dAtA, i, uint64(m.KvResp.Size()))
		n5, err := m.KvResp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.CopReq != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMsgpb(dAtA, i, uint64(m.CopReq.Size()))
		n6, err := m.CopReq.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.CopResp != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMsgpb(dAtA, i, uint64(m.CopResp.Size()))
		n7, err := m.CopResp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Msgpb(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Msgpb(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMsgpb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Message) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMsgpb(uint64(m.MsgType))
	if m.CmdReq != nil {
		l = m.CmdReq.Size()
		n += 1 + l + sovMsgpb(uint64(l))
	}
	if m.CmdResp != nil {
		l = m.CmdResp.Size()
		n += 1 + l + sovMsgpb(uint64(l))
	}
	if m.Raft != nil {
		l = m.Raft.Size()
		n += 1 + l + sovMsgpb(uint64(l))
	}
	if m.KvReq != nil {
		l = m.KvReq.Size()
		n += 1 + l + sovMsgpb(uint64(l))
	}
	if m.KvResp != nil {
		l = m.KvResp.Size()
		n += 1 + l + sovMsgpb(uint64(l))
	}
	if m.CopReq != nil {
		l = m.CopReq.Size()
		n += 1 + l + sovMsgpb(uint64(l))
	}
	if m.CopResp != nil {
		l = m.CopResp.Size()
		n += 1 + l + sovMsgpb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMsgpb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMsgpb(x uint64) (n int) {
	return sovMsgpb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgpb
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			m.MsgType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgType |= (MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CmdReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
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
				return ErrInvalidLengthMsgpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CmdReq == nil {
				m.CmdReq = &raft_cmdpb.RaftCmdRequest{}
			}
			if err := m.CmdReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CmdResp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
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
				return ErrInvalidLengthMsgpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CmdResp == nil {
				m.CmdResp = &raft_cmdpb.RaftCmdResponse{}
			}
			if err := m.CmdResp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raft", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
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
				return ErrInvalidLengthMsgpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Raft == nil {
				m.Raft = &raft_serverpb.RaftMessage{}
			}
			if err := m.Raft.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KvReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
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
				return ErrInvalidLengthMsgpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.KvReq == nil {
				m.KvReq = &kvrpcpb.Request{}
			}
			if err := m.KvReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KvResp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
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
				return ErrInvalidLengthMsgpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.KvResp == nil {
				m.KvResp = &kvrpcpb.Response{}
			}
			if err := m.KvResp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CopReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
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
				return ErrInvalidLengthMsgpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CopReq == nil {
				m.CopReq = &coprocessor.Request{}
			}
			if err := m.CopReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CopResp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgpb
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
				return ErrInvalidLengthMsgpb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CopResp == nil {
				m.CopResp = &coprocessor.Response{}
			}
			if err := m.CopResp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsgpb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgpb
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
					return 0, ErrIntOverflowMsgpb
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
					return 0, ErrIntOverflowMsgpb
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
				return 0, ErrInvalidLengthMsgpb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMsgpb
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
				next, err := skipMsgpb(dAtA[start:])
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
	ErrInvalidLengthMsgpb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgpb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("msgpb.proto", fileDescriptorMsgpb) }

var fileDescriptorMsgpb = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe3, 0xc6, 0xf6, 0xba, 0x63, 0x81, 0xb6, 0x4b, 0x91, 0xac, 0x20, 0x85, 0x8a, 0x0b,
	0x55, 0x25, 0x0c, 0x6a, 0x25, 0x1e, 0xa0, 0xb9, 0xb5, 0x82, 0xc3, 0x8a, 0x7b, 0x14, 0xd6, 0x5b,
	0x1f, 0x22, 0x77, 0xa7, 0x3b, 0xc6, 0x52, 0xdf, 0x84, 0xa7, 0xe0, 0x39, 0x72, 0xe4, 0x09, 0x10,
	0x0a, 0x2f, 0x82, 0x76, 0xec, 0x18, 0x0b, 0xf5, 0x36, 0x33, 0xfb, 0xfd, 0xe3, 0x7f, 0x7e, 0x43,
	0xde, 0x50, 0x8d, 0x5f, 0x4b, 0xf4, 0xae, 0x75, 0x2a, 0xe1, 0x66, 0x21, 0xfd, 0xe6, 0xae, 0x5d,
	0x9b, 0xa6, 0x3a, 0x3c, 0x2c, 0x5e, 0xf0, 0x84, 0xac, 0xef, 0xac, 0x1f, 0x87, 0xcf, 0xb6, 0x9d,
	0x47, 0x33, 0xb6, 0x27, 0xc6, 0xa1, 0x77, 0xc6, 0x12, 0x39, 0x3f, 0x8c, 0x4e, 0x6b, 0x57, 0x3b,
	0x2e, 0xdf, 0x87, 0xaa, 0x9f, 0xbe, 0xf9, 0x31, 0x07, 0xf1, 0xc9, 0x12, 0x6d, 0x6a, 0xab, 0xae,
	0x20, 0x6b, 0xa8, 0x5e, 0xb7, 0x8f, 0x68, 0x8b, 0xe8, 0x2c, 0x3a, 0x7f, 0x7e, 0xa9, 0xca, 0xde,
	0xd1, 0x40, 0x7c, 0x79, 0x44, 0x7b, 0x1d, 0xef, 0x7e, 0xbd, 0x9e, 0x69, 0xd1, 0x50, 0x1d, 0x5a,
	0x75, 0x05, 0xc2, 0x34, 0xd5, 0xda, 0xdb, 0x87, 0xe2, 0xe8, 0x2c, 0x3a, 0xcf, 0x2f, 0x17, 0xe5,
	0xc4, 0xb1, 0xde, 0xdc, 0xb5, 0xab, 0xa6, 0xd2, 0xf6, 0xe1, 0x9b, 0xa5, 0x56, 0xa7, 0x86, 0x6b,
	0xf5, 0x11, 0xb2, 0x5e, 0x44, 0x58, 0xcc, 0x59, 0xf5, 0xea, 0x49, 0x15, 0xa1, 0xbb, 0x27, 0xab,
	0x85, 0xe9, 0x1b, 0x55, 0x42, 0x1c, 0xb0, 0x22, 0x9e, 0x7e, 0x69, 0x4c, 0x22, 0xc8, 0x06, 0xa7,
	0x9a, 0x39, 0xf5, 0x16, 0xd2, 0x6d, 0xc7, 0xde, 0x12, 0x56, 0xc8, 0xf2, 0x10, 0xd3, 0xc1, 0x51,
	0xb2, 0xed, 0x82, 0xa1, 0x0b, 0x10, 0x0c, 0x12, 0x16, 0x29, 0x93, 0x27, 0x13, 0x72, 0x70, 0x91,
	0x06, 0x94, 0x50, 0xbd, 0x03, 0x61, 0x1c, 0xf2, 0x56, 0xc1, 0xec, 0x69, 0x39, 0x4d, 0xfb, 0xdf,
	0xad, 0x0e, 0xc3, 0xea, 0x0f, 0x90, 0xf5, 0x38, 0x61, 0x91, 0x31, 0xff, 0xf2, 0x3f, 0x7e, 0xbc,
	0x32, 0x08, 0x08, 0x6f, 0xe2, 0xec, 0x58, 0xc2, 0x4d, 0x9c, 0x81, 0xcc, 0x75, 0x8a, 0x9c, 0xad,
	0x16, 0xd8, 0xc7, 0x75, 0x51, 0x43, 0x3e, 0xf9, 0x1b, 0x2a, 0x83, 0xf8, 0xb3, 0xbb, 0xb7, 0x72,
	0xa6, 0x04, 0xcc, 0x57, 0x4d, 0x25, 0x23, 0x95, 0x83, 0x18, 0xc2, 0x93, 0x47, 0xe1, 0x3d, 0xc4,
	0x22, 0xe7, 0xea, 0x18, 0x92, 0xdb, 0x70, 0xab, 0x8c, 0x15, 0x40, 0x7a, 0xcb, 0xb7, 0xc8, 0x24,
	0xd4, 0x2b, 0x36, 0x2a, 0x53, 0x56, 0xf6, 0x1e, 0xa4, 0xb8, 0x96, 0xbb, 0xfd, 0x32, 0xfa, 0xb9,
	0x5f, 0x46, 0xbf, 0xf7, 0xcb, 0xe8, 0xfb, 0x9f, 0xe5, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x9e, 0xa4, 0x7f, 0x9f, 0x02, 0x00, 0x00,
}
