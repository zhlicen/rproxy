// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package proto

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

type Req struct {
	Req                  string   `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

type Rsp struct {
	Rsp                  string   `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rsp) Reset()         { *m = Rsp{} }
func (m *Rsp) String() string { return proto.CompactTextString(m) }
func (*Rsp) ProtoMessage()    {}
func (*Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rsp.Unmarshal(m, b)
}
func (m *Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rsp.Marshal(b, m, deterministic)
}
func (m *Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rsp.Merge(m, src)
}
func (m *Rsp) XXX_Size() int {
	return xxx_messageInfo_Rsp.Size(m)
}
func (m *Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_Rsp proto.InternalMessageInfo

func (m *Rsp) GetRsp() string {
	if m != nil {
		return m.Rsp
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "proto.Req")
	proto.RegisterType((*Rsp)(nil), "proto.Rsp")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xe2, 0x5c, 0xcc, 0x41, 0xa9,
	0x85, 0x42, 0x02, 0x5c, 0xcc, 0x45, 0xa9, 0x85, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20,
	0x26, 0x58, 0xa2, 0xb8, 0x00, 0x2c, 0x51, 0x5c, 0x00, 0x97, 0x28, 0x2e, 0x30, 0x32, 0xe3, 0xe2,
	0x0e, 0x49, 0x2d, 0x2e, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x52, 0x87, 0x70, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xb8, 0x20, 0xc6, 0xeb, 0x05, 0xa5, 0x16, 0x4a, 0xc1,
	0xd9, 0xc5, 0x05, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0x89, 0x36, 0xb1, 0x85, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	TestRequest(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestRequest(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := c.cc.Invoke(ctx, "/proto.TestService/TestRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	TestRequest(context.Context, *Req) (*Rsp, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_TestRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TestService/TestRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestRequest(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestRequest",
			Handler:    _TestService_TestRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
