// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// PrebuildsServiceClient is the client API for PrebuildsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrebuildsServiceClient interface {
	// GetRunningPrebuild returns the prebuild ID of a running prebuild,
	// or NOT_FOUND if there is no prebuild running for the content_url.
	GetRunningPrebuild(ctx context.Context, in *GetRunningPrebuildRequest, opts ...grpc.CallOption) (*GetRunningPrebuildResponse, error)
}

type prebuildsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrebuildsServiceClient(cc grpc.ClientConnInterface) PrebuildsServiceClient {
	return &prebuildsServiceClient{cc}
}

func (c *prebuildsServiceClient) GetRunningPrebuild(ctx context.Context, in *GetRunningPrebuildRequest, opts ...grpc.CallOption) (*GetRunningPrebuildResponse, error) {
	out := new(GetRunningPrebuildResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.PrebuildsService/GetRunningPrebuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrebuildsServiceServer is the server API for PrebuildsService service.
// All implementations must embed UnimplementedPrebuildsServiceServer
// for forward compatibility
type PrebuildsServiceServer interface {
	// GetRunningPrebuild returns the prebuild ID of a running prebuild,
	// or NOT_FOUND if there is no prebuild running for the content_url.
	GetRunningPrebuild(context.Context, *GetRunningPrebuildRequest) (*GetRunningPrebuildResponse, error)
	mustEmbedUnimplementedPrebuildsServiceServer()
}

// UnimplementedPrebuildsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrebuildsServiceServer struct {
}

func (UnimplementedPrebuildsServiceServer) GetRunningPrebuild(context.Context, *GetRunningPrebuildRequest) (*GetRunningPrebuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunningPrebuild not implemented")
}
func (UnimplementedPrebuildsServiceServer) mustEmbedUnimplementedPrebuildsServiceServer() {}

// UnsafePrebuildsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrebuildsServiceServer will
// result in compilation errors.
type UnsafePrebuildsServiceServer interface {
	mustEmbedUnimplementedPrebuildsServiceServer()
}

func RegisterPrebuildsServiceServer(s grpc.ServiceRegistrar, srv PrebuildsServiceServer) {
	s.RegisterService(&PrebuildsService_ServiceDesc, srv)
}

func _PrebuildsService_GetRunningPrebuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunningPrebuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrebuildsServiceServer).GetRunningPrebuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.PrebuildsService/GetRunningPrebuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrebuildsServiceServer).GetRunningPrebuild(ctx, req.(*GetRunningPrebuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrebuildsService_ServiceDesc is the grpc.ServiceDesc for PrebuildsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrebuildsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.v1.PrebuildsService",
	HandlerType: (*PrebuildsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRunningPrebuild",
			Handler:    _PrebuildsService_GetRunningPrebuild_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitpod/v1/prebuilds.proto",
}
