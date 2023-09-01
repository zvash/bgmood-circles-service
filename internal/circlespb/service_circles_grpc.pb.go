// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_circles.proto

package circlespb

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// CirclesClient is the client API for Circles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CirclesClient interface {
}

type circlesClient struct {
	cc grpc.ClientConnInterface
}

func NewCirclesClient(cc grpc.ClientConnInterface) CirclesClient {
	return &circlesClient{cc}
}

// CirclesServer is the server API for Circles service.
// All implementations must embed UnimplementedCirclesServer
// for forward compatibility
type CirclesServer interface {
	mustEmbedUnimplementedCirclesServer()
}

// UnimplementedCirclesServer must be embedded to have forward compatible implementations.
type UnimplementedCirclesServer struct {
}

func (UnimplementedCirclesServer) mustEmbedUnimplementedCirclesServer() {}

// UnsafeCirclesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CirclesServer will
// result in compilation errors.
type UnsafeCirclesServer interface {
	mustEmbedUnimplementedCirclesServer()
}

func RegisterCirclesServer(s grpc.ServiceRegistrar, srv CirclesServer) {
	s.RegisterService(&Circles_ServiceDesc, srv)
}

// Circles_ServiceDesc is the grpc.ServiceDesc for Circles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Circles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "circlespb.Circles",
	HandlerType: (*CirclesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "service_circles.proto",
}
