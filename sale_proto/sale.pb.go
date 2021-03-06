// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sale.proto

package sale

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

type Cost struct {
	Cost                 float32  `protobuf:"fixed32,1,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cost) Reset()         { *m = Cost{} }
func (m *Cost) String() string { return proto.CompactTextString(m) }
func (*Cost) ProtoMessage()    {}
func (*Cost) Descriptor() ([]byte, []int) {
	return fileDescriptor_983657353a4bf784, []int{0}
}

func (m *Cost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cost.Unmarshal(m, b)
}
func (m *Cost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cost.Marshal(b, m, deterministic)
}
func (m *Cost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cost.Merge(m, src)
}
func (m *Cost) XXX_Size() int {
	return xxx_messageInfo_Cost.Size(m)
}
func (m *Cost) XXX_DiscardUnknown() {
	xxx_messageInfo_Cost.DiscardUnknown(m)
}

var xxx_messageInfo_Cost proto.InternalMessageInfo

func (m *Cost) GetCost() float32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func init() {
	proto.RegisterType((*Cost)(nil), "Cost")
}

func init() { proto.RegisterFile("sale.proto", fileDescriptor_983657353a4bf784) }

var fileDescriptor_983657353a4bf784 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4e, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe2, 0x62, 0x71, 0xce, 0x2f, 0x2e, 0x11, 0x12,
	0xe2, 0x62, 0x49, 0xce, 0x2f, 0x2e, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0a, 0x02, 0xb3, 0x8d,
	0xb4, 0xb9, 0xf8, 0x9c, 0x13, 0x73, 0x92, 0x4b, 0x73, 0x12, 0x4b, 0xf2, 0x8b, 0x82, 0x13, 0x73,
	0x52, 0x85, 0x24, 0xb9, 0xd8, 0xdd, 0x53, 0x4b, 0xc0, 0x4c, 0x56, 0x3d, 0x90, 0x3e, 0x29, 0x08,
	0xa5, 0xc4, 0x90, 0xc4, 0x06, 0x36, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xb2, 0xee,
	0x4a, 0x5d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorSaleClient is the client API for CalculatorSale service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorSaleClient interface {
	GetSale(ctx context.Context, in *Cost, opts ...grpc.CallOption) (*Cost, error)
}

type calculatorSaleClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorSaleClient(cc *grpc.ClientConn) CalculatorSaleClient {
	return &calculatorSaleClient{cc}
}

func (c *calculatorSaleClient) GetSale(ctx context.Context, in *Cost, opts ...grpc.CallOption) (*Cost, error) {
	out := new(Cost)
	err := c.cc.Invoke(ctx, "/CalculatorSale/GetSale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorSaleServer is the server API for CalculatorSale service.
type CalculatorSaleServer interface {
	GetSale(context.Context, *Cost) (*Cost, error)
}

func RegisterCalculatorSaleServer(s *grpc.Server, srv CalculatorSaleServer) {
	s.RegisterService(&_CalculatorSale_serviceDesc, srv)
}

func _CalculatorSale_GetSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorSaleServer).GetSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalculatorSale/GetSale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorSaleServer).GetSale(ctx, req.(*Cost))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorSale_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CalculatorSale",
	HandlerType: (*CalculatorSaleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSale",
			Handler:    _CalculatorSale_GetSale_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sale.proto",
}
