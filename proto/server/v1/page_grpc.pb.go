// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: server/v1/page.proto

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

const (
	PageService_CreatePage_FullMethodName  = "/v1.PageService/CreatePage"
	PageService_UpdatePage_FullMethodName  = "/v1.PageService/UpdatePage"
	PageService_GetPage_FullMethodName     = "/v1.PageService/GetPage"
	PageService_GetPageList_FullMethodName = "/v1.PageService/GetPageList"
)

// PageServiceClient is the client API for PageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PageServiceClient interface {
	CreatePage(ctx context.Context, in *CreatePageRequest, opts ...grpc.CallOption) (*CreatePageResponse, error)
	UpdatePage(ctx context.Context, in *UpdatePageRequest, opts ...grpc.CallOption) (*UpdatePageResponse, error)
	GetPage(ctx context.Context, in *GetPageRequest, opts ...grpc.CallOption) (*GetPageResponse, error)
	GetPageList(ctx context.Context, in *GetPageListRequest, opts ...grpc.CallOption) (*GetPageListResponse, error)
}

type pageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPageServiceClient(cc grpc.ClientConnInterface) PageServiceClient {
	return &pageServiceClient{cc}
}

func (c *pageServiceClient) CreatePage(ctx context.Context, in *CreatePageRequest, opts ...grpc.CallOption) (*CreatePageResponse, error) {
	out := new(CreatePageResponse)
	err := c.cc.Invoke(ctx, PageService_CreatePage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageServiceClient) UpdatePage(ctx context.Context, in *UpdatePageRequest, opts ...grpc.CallOption) (*UpdatePageResponse, error) {
	out := new(UpdatePageResponse)
	err := c.cc.Invoke(ctx, PageService_UpdatePage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageServiceClient) GetPage(ctx context.Context, in *GetPageRequest, opts ...grpc.CallOption) (*GetPageResponse, error) {
	out := new(GetPageResponse)
	err := c.cc.Invoke(ctx, PageService_GetPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pageServiceClient) GetPageList(ctx context.Context, in *GetPageListRequest, opts ...grpc.CallOption) (*GetPageListResponse, error) {
	out := new(GetPageListResponse)
	err := c.cc.Invoke(ctx, PageService_GetPageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PageServiceServer is the server API for PageService service.
// All implementations must embed UnimplementedPageServiceServer
// for forward compatibility
type PageServiceServer interface {
	CreatePage(context.Context, *CreatePageRequest) (*CreatePageResponse, error)
	UpdatePage(context.Context, *UpdatePageRequest) (*UpdatePageResponse, error)
	GetPage(context.Context, *GetPageRequest) (*GetPageResponse, error)
	GetPageList(context.Context, *GetPageListRequest) (*GetPageListResponse, error)
	mustEmbedUnimplementedPageServiceServer()
}

// UnimplementedPageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPageServiceServer struct {
}

func (UnimplementedPageServiceServer) CreatePage(context.Context, *CreatePageRequest) (*CreatePageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePage not implemented")
}
func (UnimplementedPageServiceServer) UpdatePage(context.Context, *UpdatePageRequest) (*UpdatePageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePage not implemented")
}
func (UnimplementedPageServiceServer) GetPage(context.Context, *GetPageRequest) (*GetPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (UnimplementedPageServiceServer) GetPageList(context.Context, *GetPageListRequest) (*GetPageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPageList not implemented")
}
func (UnimplementedPageServiceServer) mustEmbedUnimplementedPageServiceServer() {}

// UnsafePageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PageServiceServer will
// result in compilation errors.
type UnsafePageServiceServer interface {
	mustEmbedUnimplementedPageServiceServer()
}

func RegisterPageServiceServer(s grpc.ServiceRegistrar, srv PageServiceServer) {
	s.RegisterService(&PageService_ServiceDesc, srv)
}

func _PageService_CreatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageServiceServer).CreatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageService_CreatePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageServiceServer).CreatePage(ctx, req.(*CreatePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageService_UpdatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageServiceServer).UpdatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageService_UpdatePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageServiceServer).UpdatePage(ctx, req.(*UpdatePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageService_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageServiceServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageService_GetPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageServiceServer).GetPage(ctx, req.(*GetPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PageService_GetPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PageServiceServer).GetPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PageService_GetPageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PageServiceServer).GetPageList(ctx, req.(*GetPageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PageService_ServiceDesc is the grpc.ServiceDesc for PageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PageService",
	HandlerType: (*PageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePage",
			Handler:    _PageService_CreatePage_Handler,
		},
		{
			MethodName: "UpdatePage",
			Handler:    _PageService_UpdatePage_Handler,
		},
		{
			MethodName: "GetPage",
			Handler:    _PageService_GetPage_Handler,
		},
		{
			MethodName: "GetPageList",
			Handler:    _PageService_GetPageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/v1/page.proto",
}