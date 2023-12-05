// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: account.proto

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
	Account_GetUserInfoByUserId_FullMethodName     = "/account.service.v1.Account/GetUserInfoByUserId"
	Account_MGetUserInfoByUserId_FullMethodName    = "/account.service.v1.Account/MGetUserInfoByUserId"
	Account_GetFollowListByUserId_FullMethodName   = "/account.service.v1.Account/GetFollowListByUserId"
	Account_GetFollowerListByUserId_FullMethodName = "/account.service.v1.Account/GetFollowerListByUserId"
	Account_GetFriendListByUserId_FullMethodName   = "/account.service.v1.Account/GetFriendListByUserId"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	GetUserInfoByUserId(ctx context.Context, in *GetUserInfoByUserIdRequest, opts ...grpc.CallOption) (*GetUserInfoByUserIdResponse, error)
	MGetUserInfoByUserId(ctx context.Context, in *MGetUserInfoByUserIdRequest, opts ...grpc.CallOption) (*MGetUserInfoByUserIdResponse, error)
	GetFollowListByUserId(ctx context.Context, in *GetFollowListByUserIdRequest, opts ...grpc.CallOption) (*GetFollowListByUserIdResponse, error)
	GetFollowerListByUserId(ctx context.Context, in *GetFollowerListByUserIdRequest, opts ...grpc.CallOption) (*GetFollowerListByUserIdResponse, error)
	GetFriendListByUserId(ctx context.Context, in *GetFriendListByUserIdRequest, opts ...grpc.CallOption) (*GetFriendListByUserIdResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetUserInfoByUserId(ctx context.Context, in *GetUserInfoByUserIdRequest, opts ...grpc.CallOption) (*GetUserInfoByUserIdResponse, error) {
	out := new(GetUserInfoByUserIdResponse)
	err := c.cc.Invoke(ctx, Account_GetUserInfoByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) MGetUserInfoByUserId(ctx context.Context, in *MGetUserInfoByUserIdRequest, opts ...grpc.CallOption) (*MGetUserInfoByUserIdResponse, error) {
	out := new(MGetUserInfoByUserIdResponse)
	err := c.cc.Invoke(ctx, Account_MGetUserInfoByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetFollowListByUserId(ctx context.Context, in *GetFollowListByUserIdRequest, opts ...grpc.CallOption) (*GetFollowListByUserIdResponse, error) {
	out := new(GetFollowListByUserIdResponse)
	err := c.cc.Invoke(ctx, Account_GetFollowListByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetFollowerListByUserId(ctx context.Context, in *GetFollowerListByUserIdRequest, opts ...grpc.CallOption) (*GetFollowerListByUserIdResponse, error) {
	out := new(GetFollowerListByUserIdResponse)
	err := c.cc.Invoke(ctx, Account_GetFollowerListByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetFriendListByUserId(ctx context.Context, in *GetFriendListByUserIdRequest, opts ...grpc.CallOption) (*GetFriendListByUserIdResponse, error) {
	out := new(GetFriendListByUserIdResponse)
	err := c.cc.Invoke(ctx, Account_GetFriendListByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	GetUserInfoByUserId(context.Context, *GetUserInfoByUserIdRequest) (*GetUserInfoByUserIdResponse, error)
	MGetUserInfoByUserId(context.Context, *MGetUserInfoByUserIdRequest) (*MGetUserInfoByUserIdResponse, error)
	GetFollowListByUserId(context.Context, *GetFollowListByUserIdRequest) (*GetFollowListByUserIdResponse, error)
	GetFollowerListByUserId(context.Context, *GetFollowerListByUserIdRequest) (*GetFollowerListByUserIdResponse, error)
	GetFriendListByUserId(context.Context, *GetFriendListByUserIdRequest) (*GetFriendListByUserIdResponse, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) GetUserInfoByUserId(context.Context, *GetUserInfoByUserIdRequest) (*GetUserInfoByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUserId not implemented")
}
func (UnimplementedAccountServer) MGetUserInfoByUserId(context.Context, *MGetUserInfoByUserIdRequest) (*MGetUserInfoByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetUserInfoByUserId not implemented")
}
func (UnimplementedAccountServer) GetFollowListByUserId(context.Context, *GetFollowListByUserIdRequest) (*GetFollowListByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowListByUserId not implemented")
}
func (UnimplementedAccountServer) GetFollowerListByUserId(context.Context, *GetFollowerListByUserIdRequest) (*GetFollowerListByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerListByUserId not implemented")
}
func (UnimplementedAccountServer) GetFriendListByUserId(context.Context, *GetFriendListByUserIdRequest) (*GetFriendListByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendListByUserId not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_GetUserInfoByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUserInfoByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetUserInfoByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUserInfoByUserId(ctx, req.(*GetUserInfoByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_MGetUserInfoByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetUserInfoByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).MGetUserInfoByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_MGetUserInfoByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).MGetUserInfoByUserId(ctx, req.(*MGetUserInfoByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetFollowListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowListByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetFollowListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetFollowListByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetFollowListByUserId(ctx, req.(*GetFollowListByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetFollowerListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowerListByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetFollowerListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetFollowerListByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetFollowerListByUserId(ctx, req.(*GetFollowerListByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetFriendListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetFriendListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetFriendListByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetFriendListByUserId(ctx, req.(*GetFriendListByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.service.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfoByUserId",
			Handler:    _Account_GetUserInfoByUserId_Handler,
		},
		{
			MethodName: "MGetUserInfoByUserId",
			Handler:    _Account_MGetUserInfoByUserId_Handler,
		},
		{
			MethodName: "GetFollowListByUserId",
			Handler:    _Account_GetFollowListByUserId_Handler,
		},
		{
			MethodName: "GetFollowerListByUserId",
			Handler:    _Account_GetFollowerListByUserId_Handler,
		},
		{
			MethodName: "GetFriendListByUserId",
			Handler:    _Account_GetFriendListByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}