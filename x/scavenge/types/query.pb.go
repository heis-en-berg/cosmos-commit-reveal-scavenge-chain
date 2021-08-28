// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("scavenge/query.proto", fileDescriptor_761e1fe4c505c12c) }

var fileDescriptor_761e1fe4c505c12c = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x31, 0x8e, 0x82, 0x40,
	0x14, 0x86, 0xa1, 0xd8, 0xdd, 0x84, 0x72, 0xb3, 0x15, 0x21, 0x73, 0x80, 0x4d, 0xe0, 0x05, 0xbd,
	0x81, 0xad, 0x95, 0xad, 0xdd, 0x0c, 0x79, 0x19, 0x26, 0x91, 0x79, 0x23, 0x6f, 0x20, 0x72, 0x0b,
	0x8f, 0x65, 0x49, 0x69, 0x69, 0xe0, 0x22, 0x06, 0x50, 0xec, 0xfe, 0xe2, 0xfb, 0xbf, 0x7c, 0xd1,
	0x1f, 0x17, 0xb2, 0x45, 0xab, 0x11, 0xce, 0x0d, 0xd6, 0x5d, 0xe6, 0x6a, 0xf2, 0xf4, 0x9b, 0x94,
	0x68, 0x18, 0xad, 0xc2, 0x5a, 0x67, 0x6f, 0x60, 0x1d, 0x71, 0xa2, 0x89, 0xf4, 0x09, 0x41, 0x3a,
	0x03, 0xd2, 0x5a, 0xf2, 0xd2, 0x1b, 0xb2, 0xbc, 0x7c, 0xe3, 0xff, 0x82, 0xb8, 0x22, 0x06, 0x25,
	0xf9, 0x25, 0x85, 0x36, 0x57, 0xe8, 0x65, 0x0e, 0x4e, 0x6a, 0x63, 0x67, 0x78, 0x61, 0x37, 0x3f,
	0xd1, 0xd7, 0x61, 0x22, 0x76, 0xfb, 0xdb, 0x20, 0xc2, 0x7e, 0x10, 0xe1, 0x63, 0x10, 0xe1, 0x75,
	0x14, 0x41, 0x3f, 0x8a, 0xe0, 0x3e, 0x8a, 0xe0, 0x98, 0x6b, 0xe3, 0xcb, 0x46, 0x65, 0x05, 0x55,
	0x30, 0x55, 0xa5, 0x68, 0xd3, 0xa9, 0x0b, 0xd6, 0xf0, 0xcb, 0x67, 0xfa, 0xce, 0x21, 0xab, 0xef,
	0x59, 0xbe, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x86, 0x78, 0xbb, 0x3f, 0xdc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heisenberg.scavenge.scavenge.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "scavenge/query.proto",
}