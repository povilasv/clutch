// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: quicklinks/v1/quicklinks.proto

package quicklinksv1

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

// QuickLinksAPIClient is the client API for QuickLinksAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuickLinksAPIClient interface {
	GetQuickLinks(ctx context.Context, in *GetQuickLinksRequest, opts ...grpc.CallOption) (*GetQuickLinksResponse, error)
}

type quickLinksAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewQuickLinksAPIClient(cc grpc.ClientConnInterface) QuickLinksAPIClient {
	return &quickLinksAPIClient{cc}
}

func (c *quickLinksAPIClient) GetQuickLinks(ctx context.Context, in *GetQuickLinksRequest, opts ...grpc.CallOption) (*GetQuickLinksResponse, error) {
	out := new(GetQuickLinksResponse)
	err := c.cc.Invoke(ctx, "/clutch.quicklinks.v1.QuickLinksAPI/GetQuickLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuickLinksAPIServer is the server API for QuickLinksAPI service.
// All implementations should embed UnimplementedQuickLinksAPIServer
// for forward compatibility
type QuickLinksAPIServer interface {
	GetQuickLinks(context.Context, *GetQuickLinksRequest) (*GetQuickLinksResponse, error)
}

// UnimplementedQuickLinksAPIServer should be embedded to have forward compatible implementations.
type UnimplementedQuickLinksAPIServer struct {
}

func (UnimplementedQuickLinksAPIServer) GetQuickLinks(context.Context, *GetQuickLinksRequest) (*GetQuickLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuickLinks not implemented")
}

// UnsafeQuickLinksAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuickLinksAPIServer will
// result in compilation errors.
type UnsafeQuickLinksAPIServer interface {
	mustEmbedUnimplementedQuickLinksAPIServer()
}

func RegisterQuickLinksAPIServer(s grpc.ServiceRegistrar, srv QuickLinksAPIServer) {
	s.RegisterService(&QuickLinksAPI_ServiceDesc, srv)
}

func _QuickLinksAPI_GetQuickLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuickLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuickLinksAPIServer).GetQuickLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.quicklinks.v1.QuickLinksAPI/GetQuickLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuickLinksAPIServer).GetQuickLinks(ctx, req.(*GetQuickLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuickLinksAPI_ServiceDesc is the grpc.ServiceDesc for QuickLinksAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuickLinksAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.quicklinks.v1.QuickLinksAPI",
	HandlerType: (*QuickLinksAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuickLinks",
			Handler:    _QuickLinksAPI_GetQuickLinks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicklinks/v1/quicklinks.proto",
}
