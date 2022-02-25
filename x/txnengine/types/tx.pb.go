// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: txnengine/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgRequestTransaction struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount  string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee     string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Note    string `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	Sender  string `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgRequestTransaction) Reset()         { *m = MsgRequestTransaction{} }
func (m *MsgRequestTransaction) String() string { return proto.CompactTextString(m) }
func (*MsgRequestTransaction) ProtoMessage()    {}
func (*MsgRequestTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_519f810905f5d8dc, []int{0}
}
func (m *MsgRequestTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestTransaction.Merge(m, src)
}
func (m *MsgRequestTransaction) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestTransaction proto.InternalMessageInfo

func (m *MsgRequestTransaction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRequestTransaction) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *MsgRequestTransaction) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *MsgRequestTransaction) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *MsgRequestTransaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgRequestTransactionResponse struct {
}

func (m *MsgRequestTransactionResponse) Reset()         { *m = MsgRequestTransactionResponse{} }
func (m *MsgRequestTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestTransactionResponse) ProtoMessage()    {}
func (*MsgRequestTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_519f810905f5d8dc, []int{1}
}
func (m *MsgRequestTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestTransactionResponse.Merge(m, src)
}
func (m *MsgRequestTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestTransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestTransaction)(nil), "cosmonaut.brexchain.txnengine.MsgRequestTransaction")
	proto.RegisterType((*MsgRequestTransactionResponse)(nil), "cosmonaut.brexchain.txnengine.MsgRequestTransactionResponse")
}

func init() { proto.RegisterFile("txnengine/tx.proto", fileDescriptor_519f810905f5d8dc) }

var fileDescriptor_519f810905f5d8dc = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x8c, 0x49, 0x29, 0xc2, 0x13, 0xb2, 0x04, 0xb2, 0x90, 0x6a, 0x50, 0x27, 0x26, 0x07, 0x01,
	0x23, 0x13, 0x23, 0x52, 0x97, 0x88, 0x89, 0xcd, 0x09, 0x8f, 0x34, 0x43, 0x9e, 0x83, 0xfd, 0x22,
	0x85, 0x2f, 0x80, 0xb1, 0x9f, 0xc5, 0xd8, 0x91, 0x11, 0x25, 0x3f, 0x82, 0xea, 0xd2, 0x88, 0x21,
	0x62, 0x60, 0xbb, 0x3b, 0xfb, 0x7c, 0x67, 0x1d, 0x17, 0xd4, 0x22, 0x60, 0x51, 0x22, 0x24, 0xd4,
	0xea, 0xda, 0x59, 0xb2, 0x62, 0x96, 0x5b, 0x5f, 0x59, 0x34, 0x0d, 0xe9, 0xcc, 0x41, 0x9b, 0x2f,
	0x4d, 0x89, 0x7a, 0xb8, 0x37, 0x7f, 0x63, 0xfc, 0x78, 0xe1, 0x8b, 0x14, 0x5e, 0x1a, 0xf0, 0xf4,
	0xe0, 0x0c, 0x7a, 0x93, 0x53, 0x69, 0x51, 0x48, 0x7e, 0x90, 0x3b, 0x30, 0x64, 0x9d, 0x64, 0xe7,
	0xec, 0xe2, 0x30, 0xdd, 0x51, 0x71, 0xc2, 0xa7, 0xa6, 0xb2, 0x0d, 0x92, 0xdc, 0x0b, 0x07, 0x3f,
	0x4c, 0x1c, 0xf1, 0xf8, 0x19, 0x40, 0xc6, 0x41, 0xdc, 0x40, 0x21, 0xf8, 0x04, 0x2d, 0x81, 0x9c,
	0x04, 0x29, 0xe0, 0x8d, 0xdb, 0x03, 0x3e, 0x81, 0x93, 0xfb, 0x5b, 0xf7, 0x96, 0xcd, 0xcf, 0xf8,
	0x6c, 0xb4, 0x48, 0x0a, 0xbe, 0xb6, 0xe8, 0xe1, 0x6a, 0xc5, 0x78, 0xbc, 0xf0, 0x85, 0x78, 0x67,
	0x5c, 0x8c, 0xf4, 0xbd, 0xd1, 0x7f, 0xfe, 0x54, 0x8f, 0x3e, 0x7e, 0x7a, 0xfb, 0x1f, 0xd7, 0xae,
	0xd2, 0xdd, 0xfd, 0x47, 0xa7, 0xd8, 0xba, 0x53, 0xec, 0xab, 0x53, 0x6c, 0xd5, 0xab, 0x68, 0xdd,
	0xab, 0xe8, 0xb3, 0x57, 0xd1, 0xe3, 0x65, 0x51, 0xd2, 0xb2, 0xc9, 0x74, 0x6e, 0xab, 0x64, 0x48,
	0x48, 0x86, 0x84, 0xa4, 0x4d, 0x7e, 0x6d, 0xf5, 0x5a, 0x83, 0xcf, 0xa6, 0x61, 0xaf, 0xeb, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x5f, 0xd6, 0xdd, 0xc5, 0x01, 0x00, 0x00,
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
	RequestTransaction(ctx context.Context, in *MsgRequestTransaction, opts ...grpc.CallOption) (*MsgRequestTransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestTransaction(ctx context.Context, in *MsgRequestTransaction, opts ...grpc.CallOption) (*MsgRequestTransactionResponse, error) {
	out := new(MsgRequestTransactionResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.brexchain.txnengine.Msg/RequestTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RequestTransaction(context.Context, *MsgRequestTransaction) (*MsgRequestTransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestTransaction(ctx context.Context, req *MsgRequestTransaction) (*MsgRequestTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestTransaction not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.brexchain.txnengine.Msg/RequestTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestTransaction(ctx, req.(*MsgRequestTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmonaut.brexchain.txnengine.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestTransaction",
			Handler:    _Msg_RequestTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "txnengine/tx.proto",
}

func (m *MsgRequestTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Note) > 0 {
		i -= len(m.Note)
		copy(dAtA[i:], m.Note)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Note)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRequestTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Note)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRequestTransactionResponse) Size() (n int) {
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
func (m *MsgRequestTransaction) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
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
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Note", wireType)
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
			m.Note = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
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
func (m *MsgRequestTransactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
