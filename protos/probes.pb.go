// Code generated by protoc-gen-go. DO NOT EDIT.
// source: probes.proto

package probepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Hello struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_05350e41f3518669, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Hello) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05350e41f3518669, []int{1}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type HelloRequest struct {
	Hello                *Hello   `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05350e41f3518669, []int{2}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

type HelloResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05350e41f3518669, []int{3}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type CommonProbes struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Probes               []*Probe `protobuf:"bytes,2,rep,name=probes,proto3" json:"probes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonProbes) Reset()         { *m = CommonProbes{} }
func (m *CommonProbes) String() string { return proto.CompactTextString(m) }
func (*CommonProbes) ProtoMessage()    {}
func (*CommonProbes) Descriptor() ([]byte, []int) {
	return fileDescriptor_05350e41f3518669, []int{4}
}

func (m *CommonProbes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonProbes.Unmarshal(m, b)
}
func (m *CommonProbes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonProbes.Marshal(b, m, deterministic)
}
func (m *CommonProbes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonProbes.Merge(m, src)
}
func (m *CommonProbes) XXX_Size() int {
	return xxx_messageInfo_CommonProbes.Size(m)
}
func (m *CommonProbes) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonProbes.DiscardUnknown(m)
}

var xxx_messageInfo_CommonProbes proto.InternalMessageInfo

func (m *CommonProbes) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CommonProbes) GetProbes() []*Probe {
	if m != nil {
		return m.Probes
	}
	return nil
}

type Probe struct {
	ProbeID              int32    `protobuf:"varint,1,opt,name=probeID,proto3" json:"probeID,omitempty"`
	IpVer                int32    `protobuf:"varint,2,opt,name=ipVer,proto3" json:"ipVer,omitempty"`
	ProbeType            string   `protobuf:"bytes,3,opt,name=probeType,proto3" json:"probeType,omitempty"`
	TargetAddress        string   `protobuf:"bytes,4,opt,name=targetAddress,proto3" json:"targetAddress,omitempty"`
	NumPackets           int32    `protobuf:"varint,5,opt,name=numPackets,proto3" json:"numPackets,omitempty"`
	Interval             int32    `protobuf:"varint,6,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Probe) Reset()         { *m = Probe{} }
func (m *Probe) String() string { return proto.CompactTextString(m) }
func (*Probe) ProtoMessage()    {}
func (*Probe) Descriptor() ([]byte, []int) {
	return fileDescriptor_05350e41f3518669, []int{5}
}

func (m *Probe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Probe.Unmarshal(m, b)
}
func (m *Probe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Probe.Marshal(b, m, deterministic)
}
func (m *Probe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Probe.Merge(m, src)
}
func (m *Probe) XXX_Size() int {
	return xxx_messageInfo_Probe.Size(m)
}
func (m *Probe) XXX_DiscardUnknown() {
	xxx_messageInfo_Probe.DiscardUnknown(m)
}

var xxx_messageInfo_Probe proto.InternalMessageInfo

func (m *Probe) GetProbeID() int32 {
	if m != nil {
		return m.ProbeID
	}
	return 0
}

func (m *Probe) GetIpVer() int32 {
	if m != nil {
		return m.IpVer
	}
	return 0
}

func (m *Probe) GetProbeType() string {
	if m != nil {
		return m.ProbeType
	}
	return ""
}

func (m *Probe) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func (m *Probe) GetNumPackets() int32 {
	if m != nil {
		return m.NumPackets
	}
	return 0
}

func (m *Probe) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func init() {
	proto.RegisterType((*Hello)(nil), "nm_probes.Hello")
	proto.RegisterType((*VersionRequest)(nil), "nm_probes.VersionRequest")
	proto.RegisterType((*HelloRequest)(nil), "nm_probes.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "nm_probes.HelloResponse")
	proto.RegisterType((*CommonProbes)(nil), "nm_probes.CommonProbes")
	proto.RegisterType((*Probe)(nil), "nm_probes.Probe")
}

func init() { proto.RegisterFile("probes.proto", fileDescriptor_05350e41f3518669) }

var fileDescriptor_05350e41f3518669 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x50, 0xbb, 0x78, 0x0a, 0x15, 0x5a, 0x55, 0xed, 0x96, 0x7e, 0x08, 0xad, 0xaa, 0x86,
	0x13, 0x07, 0x22, 0xe5, 0x94, 0x4b, 0x42, 0xa4, 0x90, 0x4b, 0x84, 0x9c, 0x88, 0x43, 0x2e, 0xc8,
	0xc0, 0x24, 0xb1, 0x62, 0x7b, 0x9d, 0xdd, 0x05, 0x29, 0x7f, 0x24, 0x7f, 0x25, 0x7f, 0x2f, 0xf2,
	0xec, 0x1a, 0x99, 0x70, 0x7c, 0xef, 0xcd, 0xbe, 0x19, 0xbd, 0xb7, 0xd0, 0x29, 0x94, 0x5c, 0xa2,
	0x1e, 0x15, 0x4a, 0x1a, 0xc9, 0xc2, 0x3c, 0x5b, 0x58, 0x42, 0x4c, 0xc0, 0x9f, 0x62, 0x9a, 0x4a,
	0xf6, 0x07, 0xe0, 0x3e, 0x51, 0xda, 0x2c, 0xf2, 0x38, 0x43, 0xee, 0x0d, 0xbc, 0x61, 0x18, 0x85,
	0xc4, 0x5c, 0xc7, 0x19, 0xb2, 0x5f, 0x10, 0xa6, 0x71, 0xa5, 0x36, 0x49, 0x6d, 0x97, 0x44, 0x29,
	0x8a, 0x1e, 0x7c, 0x9d, 0xa3, 0xd2, 0x89, 0xcc, 0x23, 0x7c, 0xde, 0xa0, 0x36, 0xe2, 0x04, 0x3a,
	0x64, 0xeb, 0x30, 0xfb, 0x0f, 0xfe, 0x63, 0x89, 0xc9, 0xf8, 0xcb, 0xb8, 0x37, 0xda, 0x5d, 0x30,
	0xb2, 0x73, 0x56, 0x16, 0x47, 0xd0, 0x75, 0xef, 0x74, 0x21, 0x73, 0x8d, 0xec, 0x3b, 0x04, 0x0a,
	0xf5, 0x26, 0x35, 0xee, 0x24, 0x87, 0x44, 0x04, 0x9d, 0x89, 0xcc, 0x32, 0x99, 0xcf, 0xc8, 0x85,
	0x71, 0xf8, 0xbc, 0xb5, 0x27, 0xd0, 0xa0, 0x1f, 0x55, 0x90, 0x0d, 0x21, 0xb0, 0x9b, 0x78, 0x73,
	0xd0, 0xfa, 0xb0, 0x9b, 0x1e, 0x47, 0x4e, 0x17, 0x6f, 0x1e, 0xf8, 0xc4, 0x94, 0x6e, 0xc4, 0x5d,
	0x5d, 0x54, 0x6e, 0x0e, 0xb2, 0x6f, 0xe0, 0x27, 0xc5, 0x1c, 0x15, 0x65, 0xe0, 0x47, 0x16, 0xb0,
	0xdf, 0x10, 0xd2, 0xc0, 0xed, 0x4b, 0x81, 0xbc, 0x65, 0xb3, 0xdb, 0x11, 0xec, 0x1f, 0x74, 0x4d,
	0xac, 0x1e, 0xd0, 0x9c, 0xad, 0xd7, 0x0a, 0xb5, 0xe6, 0x9f, 0x68, 0x62, 0x9f, 0x64, 0x7f, 0x01,
	0xf2, 0x4d, 0x36, 0x8b, 0x57, 0x4f, 0x68, 0x34, 0xf7, 0xc9, 0xbe, 0xc6, 0xb0, 0x3e, 0xb4, 0x93,
	0xdc, 0xa0, 0xda, 0xc6, 0x29, 0x0f, 0x48, 0xdd, 0xe1, 0xf1, 0xab, 0xe7, 0xf2, 0xbe, 0x41, 0xb5,
	0x4d, 0x56, 0xc8, 0x4e, 0xab, 0x5a, 0x7f, 0x1c, 0x24, 0x6d, 0x1b, 0xe9, 0xf3, 0x43, 0xc1, 0x46,
	0x2e, 0x1a, 0x6c, 0x0a, 0xbd, 0x4b, 0x34, 0x36, 0x59, 0x57, 0x2c, 0xfb, 0x59, 0x9b, 0xdf, 0x2f,
	0xbb, 0x5f, 0xdf, 0x51, 0x2f, 0x45, 0x34, 0xce, 0xc3, 0x3b, 0x9b, 0x5c, 0xb1, 0x5c, 0x06, 0xf4,
	0xf7, 0x8e, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xdd, 0xf9, 0x9c, 0x8b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	GetProbesVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*CommonProbes, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/nm_probes.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) GetProbesVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*CommonProbes, error) {
	out := new(CommonProbes)
	err := c.cc.Invoke(ctx, "/nm_probes.HelloService/GetProbesVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	GetProbesVersion(context.Context, *VersionRequest) (*CommonProbes, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedHelloServiceServer) GetProbesVersion(ctx context.Context, req *VersionRequest) (*CommonProbes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProbesVersion not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nm_probes.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_GetProbesVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).GetProbesVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nm_probes.HelloService/GetProbesVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).GetProbesVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nm_probes.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
		{
			MethodName: "GetProbesVersion",
			Handler:    _HelloService_GetProbesVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "probes.proto",
}