// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello-world.proto

package helloworld

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SayHelloRequest struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a62523e9758a7c9, []int{0}
}

func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloRequest.Unmarshal(m, b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
}
func (m *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(m, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return xxx_messageInfo_SayHelloRequest.Size(m)
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (m *SayHelloRequest) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type SayHelloResponse struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a62523e9758a7c9, []int{1}
}

func (m *SayHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloResponse.Unmarshal(m, b)
}
func (m *SayHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloResponse.Marshal(b, m, deterministic)
}
func (m *SayHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResponse.Merge(m, src)
}
func (m *SayHelloResponse) XXX_Size() int {
	return xxx_messageInfo_SayHelloResponse.Size(m)
}
func (m *SayHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResponse proto.InternalMessageInfo

func (m *SayHelloResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "helloworld.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "helloworld.SayHelloResponse")
}

func init() { proto.RegisterFile("hello-world.proto", fileDescriptor_5a62523e9758a7c9) }

var fileDescriptor_5a62523e9758a7c9 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2d, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02,
	0x0b, 0x81, 0x45, 0x94, 0x54, 0xb9, 0xf8, 0x83, 0x13, 0x2b, 0x3d, 0x40, 0x02, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x05, 0x99, 0x79, 0xe9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x1a, 0x97, 0x00, 0x42, 0x59, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0x2a, 0x58, 0x5d, 0x3e, 0x92, 0xba, 0xfc, 0xbc, 0x74, 0xa3, 0x60, 0x2e, 0x2e, 0xb0, 0xa2, 0x70,
	0x90, 0xe1, 0x42, 0xae, 0x5c, 0x1c, 0x30, 0x5d, 0x42, 0xd2, 0x7a, 0x08, 0x5b, 0xf5, 0xd0, 0xac,
	0x94, 0x92, 0xc1, 0x2e, 0x09, 0xb1, 0x28, 0x89, 0x0d, 0xec, 0x6c, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0xbb, 0xd7, 0xf9, 0xcb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type helloWorldClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldClient(cc *grpc.ClientConn) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/helloworld.HelloWorld/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
type HelloWorldServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

func RegisterHelloWorldServer(s *grpc.Server, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorld/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloWorld_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello-world.proto",
}
