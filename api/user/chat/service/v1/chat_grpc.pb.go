// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: chat.proto

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
	Chat_ChatAction_FullMethodName                                        = "/api.chat.v1.Chat/ChatAction"
	Chat_GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime_FullMethodName = "/api.chat.v1.Chat/GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime"
	Chat_GetLatestMsgByMyUserIdAndHisUserId_FullMethodName                = "/api.chat.v1.Chat/GetLatestMsgByMyUserIdAndHisUserId"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	ChatAction(ctx context.Context, in *DouyinChatActionRequest, opts ...grpc.CallOption) (*DouyinChatActionResponse, error)
	GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, in *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest, opts ...grpc.CallOption) (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse, error)
	GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, in *GetLatestMsgByMyUserIdAndHisUserIdRequest, opts ...grpc.CallOption) (*GetLatestMsgByMyUserIdAndHisUserIdResponse, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) ChatAction(ctx context.Context, in *DouyinChatActionRequest, opts ...grpc.CallOption) (*DouyinChatActionResponse, error) {
	out := new(DouyinChatActionResponse)
	err := c.cc.Invoke(ctx, Chat_ChatAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, in *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest, opts ...grpc.CallOption) (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse, error) {
	out := new(GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse)
	err := c.cc.Invoke(ctx, Chat_GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, in *GetLatestMsgByMyUserIdAndHisUserIdRequest, opts ...grpc.CallOption) (*GetLatestMsgByMyUserIdAndHisUserIdResponse, error) {
	out := new(GetLatestMsgByMyUserIdAndHisUserIdResponse)
	err := c.cc.Invoke(ctx, Chat_GetLatestMsgByMyUserIdAndHisUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	ChatAction(context.Context, *DouyinChatActionRequest) (*DouyinChatActionResponse, error)
	GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(context.Context, *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse, error)
	GetLatestMsgByMyUserIdAndHisUserId(context.Context, *GetLatestMsgByMyUserIdAndHisUserIdRequest) (*GetLatestMsgByMyUserIdAndHisUserIdResponse, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) ChatAction(context.Context, *DouyinChatActionRequest) (*DouyinChatActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatAction not implemented")
}
func (UnimplementedChatServer) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(context.Context, *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime not implemented")
}
func (UnimplementedChatServer) GetLatestMsgByMyUserIdAndHisUserId(context.Context, *GetLatestMsgByMyUserIdAndHisUserIdRequest) (*GetLatestMsgByMyUserIdAndHisUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestMsgByMyUserIdAndHisUserId not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_ChatAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinChatActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ChatAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_ChatAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ChatAction(ctx, req.(*DouyinChatActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx, req.(*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetLatestMsgByMyUserIdAndHisUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestMsgByMyUserIdAndHisUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetLatestMsgByMyUserIdAndHisUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetLatestMsgByMyUserIdAndHisUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetLatestMsgByMyUserIdAndHisUserId(ctx, req.(*GetLatestMsgByMyUserIdAndHisUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.chat.v1.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChatAction",
			Handler:    _Chat_ChatAction_Handler,
		},
		{
			MethodName: "GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime",
			Handler:    _Chat_GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime_Handler,
		},
		{
			MethodName: "GetLatestMsgByMyUserIdAndHisUserId",
			Handler:    _Chat_GetLatestMsgByMyUserIdAndHisUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
