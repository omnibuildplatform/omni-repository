// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc1
// source: call_repo.proto

package app

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RepoServerClient is the client API for RepoServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepoServerClient interface {
	// Sends a greeting
	CallLoadFrom(ctx context.Context, in *RepRequest, opts ...grpc.CallOption) (*RepResponse, error)
}

type repoServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRepoServerClient(cc grpc.ClientConnInterface) RepoServerClient {
	return &repoServerClient{cc}
}

func (c *repoServerClient) CallLoadFrom(ctx context.Context, in *RepRequest, opts ...grpc.CallOption) (*RepResponse, error) {
	out := new(RepResponse)
	err := c.cc.Invoke(ctx, "/app.RepoServer/CallLoadFrom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepoServerServer is the server API for RepoServer service.
// All implementations must embed UnimplementedRepoServerServer
// for forward compatibility
type RepoServerServer interface {
	// Sends a greeting
	CallLoadFrom(context.Context, *RepRequest) (*RepResponse, error)
	mustEmbedUnimplementedRepoServerServer()
}

// UnimplementedRepoServerServer must be embedded to have forward compatible implementations.
type UnimplementedRepoServerServer struct {
}

func (UnimplementedRepoServerServer) CallLoadFrom(context.Context, *RepRequest) (*RepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallLoadFrom not implemented")
}
func (UnimplementedRepoServerServer) mustEmbedUnimplementedRepoServerServer() {}

// UnsafeRepoServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepoServerServer will
// result in compilation errors.
type UnsafeRepoServerServer interface {
	mustEmbedUnimplementedRepoServerServer()
}

func RegisterRepoServerServer(s grpc.ServiceRegistrar, srv RepoServerServer) {
	s.RegisterService(&RepoServer_ServiceDesc, srv)
}

func _RepoServer_CallLoadFrom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoServerServer).CallLoadFrom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.RepoServer/CallLoadFrom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoServerServer).CallLoadFrom(ctx, req.(*RepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepoServer_ServiceDesc is the grpc.ServiceDesc for RepoServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepoServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.RepoServer",
	HandlerType: (*RepoServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallLoadFrom",
			Handler:    _RepoServer_CallLoadFrom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "call_repo.proto",
}