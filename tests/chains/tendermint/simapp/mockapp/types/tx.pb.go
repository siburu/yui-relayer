// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/mockapp/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgTransfer defines a msg to transfer fungible tokens (i.e Coins) between
// ICS20 enabled chains. See ICS Spec here:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#data-structures
type MsgSendPacket struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// the message
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty" yaml:"message"`
	// the sender address
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,5,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp in absolute nanoseconds since unix epoch.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,6,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *MsgSendPacket) Reset()         { *m = MsgSendPacket{} }
func (m *MsgSendPacket) String() string { return proto.CompactTextString(m) }
func (*MsgSendPacket) ProtoMessage()    {}
func (*MsgSendPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_25430e89edfab170, []int{0}
}
func (m *MsgSendPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendPacket.Merge(m, src)
}
func (m *MsgSendPacket) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendPacket.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendPacket proto.InternalMessageInfo

// MsgTransferResponse defines the Msg/Transfer response type.
type MsgSendPacketResponse struct {
}

func (m *MsgSendPacketResponse) Reset()         { *m = MsgSendPacketResponse{} }
func (m *MsgSendPacketResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendPacketResponse) ProtoMessage()    {}
func (*MsgSendPacketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25430e89edfab170, []int{1}
}
func (m *MsgSendPacketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendPacketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendPacketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendPacketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendPacketResponse.Merge(m, src)
}
func (m *MsgSendPacketResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendPacketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendPacketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendPacketResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendPacket)(nil), "ibc.applications.mockapp.v1.MsgSendPacket")
	proto.RegisterType((*MsgSendPacketResponse)(nil), "ibc.applications.mockapp.v1.MsgSendPacketResponse")
}

func init() {
	proto.RegisterFile("ibc/applications/mockapp/v1/tx.proto", fileDescriptor_25430e89edfab170)
}

var fileDescriptor_25430e89edfab170 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x6d, 0xd2, 0x06, 0x71, 0x55, 0x22, 0x6a, 0x28, 0x35, 0x01, 0xec, 0xc8, 0x62, 0x08,
	0x11, 0xf1, 0x29, 0x61, 0x40, 0xea, 0x84, 0xc2, 0x02, 0x43, 0xa5, 0xca, 0x30, 0xb1, 0xb4, 0xe7,
	0xcb, 0x93, 0x7d, 0xaa, 0xcf, 0x67, 0xf9, 0x2e, 0x11, 0xed, 0x84, 0x98, 0x18, 0xf9, 0x08, 0x5d,
	0xd8, 0xfb, 0x31, 0x3a, 0x76, 0x64, 0xb2, 0x50, 0x32, 0x94, 0x39, 0x9f, 0x00, 0xd9, 0xbe, 0x40,
	0xb2, 0xa0, 0x4e, 0xbe, 0x7b, 0xff, 0xdf, 0xff, 0x3d, 0xdf, 0xbb, 0x77, 0xe8, 0x39, 0x0b, 0x29,
	0x26, 0x59, 0x96, 0x30, 0x4a, 0x14, 0x13, 0xa9, 0xc4, 0x5c, 0xd0, 0x53, 0x92, 0x65, 0x78, 0x36,
	0xc4, 0xea, 0xb3, 0x9f, 0xe5, 0x42, 0x09, 0xeb, 0x09, 0x0b, 0xa9, 0xbf, 0x4e, 0xf9, 0x9a, 0xf2,
	0x67, 0xc3, 0xce, 0xc3, 0x48, 0x44, 0xa2, 0xe2, 0x70, 0xb9, 0xaa, 0x2d, 0x9d, 0x7d, 0x2a, 0x24,
	0x17, 0x12, 0x73, 0x19, 0x95, 0xa9, 0xb8, 0x8c, 0xb4, 0xe0, 0x96, 0x15, 0xa9, 0xc8, 0x01, 0xd3,
	0x84, 0x41, 0xaa, 0x4a, 0xb5, 0x5e, 0xd5, 0x80, 0xf7, 0xa3, 0x81, 0x5a, 0x87, 0x32, 0xfa, 0x00,
	0xe9, 0xe4, 0x88, 0xd0, 0x53, 0x50, 0xd6, 0x6b, 0xb4, 0x23, 0xc5, 0x34, 0xa7, 0x70, 0x9c, 0x89,
	0x5c, 0xd9, 0x66, 0xd7, 0xec, 0xdd, 0x1b, 0x3f, 0x5a, 0x16, 0xae, 0x75, 0x46, 0x78, 0x72, 0xe0,
	0xad, 0x89, 0x5e, 0x80, 0xea, 0xdd, 0x91, 0xc8, 0x95, 0xf5, 0x06, 0xb5, 0xb5, 0x46, 0x63, 0x92,
	0xa6, 0x90, 0xd8, 0x77, 0x2a, 0xef, 0xe3, 0x65, 0xe1, 0xee, 0x6d, 0x78, 0xb5, 0xee, 0x05, 0xad,
	0x3a, 0xf0, 0xb6, 0xde, 0x5b, 0x2f, 0xd1, 0x5d, 0x0e, 0x52, 0x92, 0x08, 0xec, 0x46, 0x65, 0xb5,
	0x96, 0x85, 0xdb, 0xae, 0xad, 0x5a, 0xf0, 0x82, 0x15, 0x62, 0xbd, 0x40, 0x4d, 0x09, 0xe9, 0x04,
	0x72, 0x7b, 0xab, 0x82, 0x77, 0x97, 0x85, 0xdb, 0xd2, 0x75, 0xaa, 0xb8, 0x17, 0x68, 0xc0, 0x3a,
	0x41, 0x6d, 0xc5, 0x38, 0x88, 0xa9, 0x3a, 0x8e, 0x81, 0x45, 0xb1, 0xb2, 0xb7, 0xbb, 0x66, 0x6f,
	0x67, 0xd4, 0xf1, 0xcb, 0x5e, 0x97, 0xfd, 0xf1, 0x75, 0x57, 0x66, 0x43, 0xff, 0x5d, 0x45, 0x8c,
	0x9f, 0x5d, 0x15, 0xae, 0xf1, 0xef, 0xd7, 0x37, 0xfd, 0x5e, 0xd0, 0xd2, 0x81, 0x9a, 0xb6, 0xde,
	0xa3, 0xdd, 0x15, 0x51, 0x7e, 0xa5, 0x22, 0x3c, 0xb3, 0x9b, 0x5d, 0xb3, 0xb7, 0x35, 0x7e, 0xba,
	0x2c, 0x5c, 0x7b, 0x33, 0xc9, 0x5f, 0xc4, 0x0b, 0xee, 0xeb, 0xd8, 0xc7, 0x55, 0xe8, 0xe0, 0xc1,
	0xb7, 0x0b, 0xd7, 0xf8, 0x7d, 0xe1, 0x1a, 0x5f, 0x6f, 0x2e, 0xfb, 0xfa, 0x04, 0xde, 0x3e, 0xda,
	0xdb, 0xb8, 0xa6, 0x00, 0x64, 0x26, 0x52, 0x09, 0xa3, 0x73, 0xd4, 0x38, 0x94, 0x91, 0x95, 0x20,
	0xb4, 0x76, 0x87, 0x7d, 0xff, 0x3f, 0x33, 0xe4, 0x6f, 0x24, 0xea, 0x8c, 0x6e, 0xcf, 0xae, 0x8a,
	0x76, 0xb6, 0xbf, 0xdc, 0x5c, 0xf6, 0xcd, 0xf1, 0xf9, 0xd5, 0xdc, 0x31, 0xaf, 0xe7, 0x8e, 0xf9,
	0x6b, 0xee, 0x98, 0xdf, 0x17, 0x8e, 0x71, 0xbd, 0x70, 0x8c, 0x9f, 0x0b, 0xc7, 0xf8, 0x74, 0x12,
	0x31, 0x15, 0x4f, 0x43, 0x9f, 0x0a, 0x8e, 0x27, 0x44, 0x11, 0x1a, 0x13, 0x96, 0x26, 0x24, 0xc4,
	0xf5, 0xa0, 0x0e, 0x40, 0xc5, 0x90, 0xc3, 0x94, 0x0f, 0x58, 0x48, 0x07, 0x09, 0xcd, 0xb0, 0x02,
	0xa9, 0x24, 0x86, 0x11, 0xe0, 0x0a, 0x95, 0x58, 0x55, 0xe7, 0xe6, 0x2c, 0x55, 0x58, 0x32, 0x5e,
	0xbe, 0x91, 0xd5, 0x5b, 0x51, 0x67, 0x19, 0xc8, 0xb0, 0x59, 0xcd, 0xef, 0xab, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x6d, 0xb0, 0xdf, 0x59, 0x54, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Transfer defines a rpc handler method for MsgTransfer.
	SendPacket(ctx context.Context, in *MsgSendPacket, opts ...grpc.CallOption) (*MsgSendPacketResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendPacket(ctx context.Context, in *MsgSendPacket, opts ...grpc.CallOption) (*MsgSendPacketResponse, error) {
	out := new(MsgSendPacketResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.mockapp.v1.Msg/SendPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Transfer defines a rpc handler method for MsgTransfer.
	SendPacket(context.Context, *MsgSendPacket) (*MsgSendPacketResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendPacket(ctx context.Context, req *MsgSendPacket) (*MsgSendPacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPacket not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.mockapp.v1.Msg/SendPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendPacket(ctx, req.(*MsgSendPacket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.mockapp.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPacket",
			Handler:    _Msg_SendPacket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/mockapp/v1/tx.proto",
}

func (m *MsgSendPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendPacketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendPacketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendPacketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *MsgSendPacketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSendPacketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendPacketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendPacketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
