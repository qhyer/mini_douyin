// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: favorite.proto

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
	Favorite_FavoriteAction_FullMethodName                       = "/api.favorite.v1.Favorite/FavoriteAction"
	Favorite_GetUserFavoriteVideoIdList_FullMethodName           = "/api.favorite.v1.Favorite/GetUserFavoriteVideoIdList"
	Favorite_GetFavoriteStatusByUserIdAndVideoIds_FullMethodName = "/api.favorite.v1.Favorite/GetFavoriteStatusByUserIdAndVideoIds"
)

// FavoriteClient is the client API for Favorite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteClient interface {
	FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error)
	GetUserFavoriteVideoIdList(ctx context.Context, in *GetUserFavoriteListRequest, opts ...grpc.CallOption) (*GetUserFavoriteListResponse, error)
	GetFavoriteStatusByUserIdAndVideoIds(ctx context.Context, in *GetFavoriteStatusByUserIdAndVideoIdsRequest, opts ...grpc.CallOption) (*GetFavoriteStatusByUserIdAndVideoIdsResponse, error)
}

type favoriteClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteClient(cc grpc.ClientConnInterface) FavoriteClient {
	return &favoriteClient{cc}
}

func (c *favoriteClient) FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error) {
	out := new(DouyinFavoriteActionResponse)
	err := c.cc.Invoke(ctx, Favorite_FavoriteAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) GetUserFavoriteVideoIdList(ctx context.Context, in *GetUserFavoriteListRequest, opts ...grpc.CallOption) (*GetUserFavoriteListResponse, error) {
	out := new(GetUserFavoriteListResponse)
	err := c.cc.Invoke(ctx, Favorite_GetUserFavoriteVideoIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) GetFavoriteStatusByUserIdAndVideoIds(ctx context.Context, in *GetFavoriteStatusByUserIdAndVideoIdsRequest, opts ...grpc.CallOption) (*GetFavoriteStatusByUserIdAndVideoIdsResponse, error) {
	out := new(GetFavoriteStatusByUserIdAndVideoIdsResponse)
	err := c.cc.Invoke(ctx, Favorite_GetFavoriteStatusByUserIdAndVideoIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServer is the server API for Favorite service.
// All implementations must embed UnimplementedFavoriteServer
// for forward compatibility
type FavoriteServer interface {
	FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error)
	GetUserFavoriteVideoIdList(context.Context, *GetUserFavoriteListRequest) (*GetUserFavoriteListResponse, error)
	GetFavoriteStatusByUserIdAndVideoIds(context.Context, *GetFavoriteStatusByUserIdAndVideoIdsRequest) (*GetFavoriteStatusByUserIdAndVideoIdsResponse, error)
	mustEmbedUnimplementedFavoriteServer()
}

// UnimplementedFavoriteServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServer struct {
}

func (UnimplementedFavoriteServer) FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteServer) GetUserFavoriteVideoIdList(context.Context, *GetUserFavoriteListRequest) (*GetUserFavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavoriteVideoIdList not implemented")
}
func (UnimplementedFavoriteServer) GetFavoriteStatusByUserIdAndVideoIds(context.Context, *GetFavoriteStatusByUserIdAndVideoIdsRequest) (*GetFavoriteStatusByUserIdAndVideoIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteStatusByUserIdAndVideoIds not implemented")
}
func (UnimplementedFavoriteServer) mustEmbedUnimplementedFavoriteServer() {}

// UnsafeFavoriteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServer will
// result in compilation errors.
type UnsafeFavoriteServer interface {
	mustEmbedUnimplementedFavoriteServer()
}

func RegisterFavoriteServer(s grpc.ServiceRegistrar, srv FavoriteServer) {
	s.RegisterService(&Favorite_ServiceDesc, srv)
}

func _Favorite_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).FavoriteAction(ctx, req.(*DouyinFavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_GetUserFavoriteVideoIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).GetUserFavoriteVideoIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_GetUserFavoriteVideoIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).GetUserFavoriteVideoIdList(ctx, req.(*GetUserFavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_GetFavoriteStatusByUserIdAndVideoIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteStatusByUserIdAndVideoIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).GetFavoriteStatusByUserIdAndVideoIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_GetFavoriteStatusByUserIdAndVideoIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).GetFavoriteStatusByUserIdAndVideoIds(ctx, req.(*GetFavoriteStatusByUserIdAndVideoIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Favorite_ServiceDesc is the grpc.ServiceDesc for Favorite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Favorite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.favorite.v1.Favorite",
	HandlerType: (*FavoriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _Favorite_FavoriteAction_Handler,
		},
		{
			MethodName: "GetUserFavoriteVideoIdList",
			Handler:    _Favorite_GetUserFavoriteVideoIdList_Handler,
		},
		{
			MethodName: "GetFavoriteStatusByUserIdAndVideoIds",
			Handler:    _Favorite_GetFavoriteStatusByUserIdAndVideoIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
