// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: bff/bff.proto

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
	BFF_UserRegister_FullMethodName         = "/bff.service.v1.BFF/UserRegister"
	BFF_UserLogin_FullMethodName            = "/bff.service.v1.BFF/UserLogin"
	BFF_GetUserInfo_FullMethodName          = "/bff.service.v1.BFF/GetUserInfo"
	BFF_GetPublishList_FullMethodName       = "/bff.service.v1.BFF/GetPublishList"
	BFF_PublishAction_FullMethodName        = "/bff.service.v1.BFF/PublishAction"
	BFF_Feed_FullMethodName                 = "/bff.service.v1.BFF/Feed"
	BFF_GetFollowerList_FullMethodName      = "/bff.service.v1.BFF/GetFollowerList"
	BFF_GetFollowList_FullMethodName        = "/bff.service.v1.BFF/GetFollowList"
	BFF_RelationAction_FullMethodName       = "/bff.service.v1.BFF/RelationAction"
	BFF_GetFriendList_FullMethodName        = "/bff.service.v1.BFF/GetFriendList"
	BFF_GetMessageList_FullMethodName       = "/bff.service.v1.BFF/GetMessageList"
	BFF_MessageAction_FullMethodName        = "/bff.service.v1.BFF/MessageAction"
	BFF_GetFavoriteVideoList_FullMethodName = "/bff.service.v1.BFF/GetFavoriteVideoList"
	BFF_FavoriteAction_FullMethodName       = "/bff.service.v1.BFF/FavoriteAction"
	BFF_GetCommentList_FullMethodName       = "/bff.service.v1.BFF/GetCommentList"
	BFF_CommentAction_FullMethodName        = "/bff.service.v1.BFF/CommentAction"
)

// BFFClient is the client API for BFF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BFFClient interface {
	// 用户注册
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterReply, error)
	// 用户登陆
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginReply, error)
	// 获取用户信息
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error)
	// 获取用户投稿视频列表
	GetPublishList(ctx context.Context, in *GetPublishListRequest, opts ...grpc.CallOption) (*GetPublishListReply, error)
	// 用户发布视频
	PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionReply, error)
	// 视频流
	Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedReply, error)
	// 获取粉丝列表
	GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListReply, error)
	// 获取关注列表
	GetFollowList(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListReply, error)
	// 关注或取关用户
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionReply, error)
	// 获取好友列表
	GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListReply, error)
	// 获取消息列表
	GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListReply, error)
	// 给好友发送消息
	MessageAction(ctx context.Context, in *MessageActionRequest, opts ...grpc.CallOption) (*MessageActionReply, error)
	// 获取点赞视频列表
	GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListReply, error)
	// 点赞/取消点赞视频
	FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionReply, error)
	// 获取评论列表
	GetCommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListReply, error)
	// 发布评论或者删除评论
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionReply, error)
}

type bFFClient struct {
	cc grpc.ClientConnInterface
}

func NewBFFClient(cc grpc.ClientConnInterface) BFFClient {
	return &bFFClient{cc}
}

func (c *bFFClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterReply, error) {
	out := new(UserRegisterReply)
	err := c.cc.Invoke(ctx, BFF_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginReply, error) {
	out := new(UserLoginReply)
	err := c.cc.Invoke(ctx, BFF_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error) {
	out := new(GetUserInfoReply)
	err := c.cc.Invoke(ctx, BFF_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetPublishList(ctx context.Context, in *GetPublishListRequest, opts ...grpc.CallOption) (*GetPublishListReply, error) {
	out := new(GetPublishListReply)
	err := c.cc.Invoke(ctx, BFF_GetPublishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) PublishAction(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionReply, error) {
	out := new(PublishActionReply)
	err := c.cc.Invoke(ctx, BFF_PublishAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedReply, error) {
	out := new(FeedReply)
	err := c.cc.Invoke(ctx, BFF_Feed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListReply, error) {
	out := new(GetFollowerListReply)
	err := c.cc.Invoke(ctx, BFF_GetFollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetFollowList(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListReply, error) {
	out := new(GetFollowListReply)
	err := c.cc.Invoke(ctx, BFF_GetFollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionReply, error) {
	out := new(RelationActionReply)
	err := c.cc.Invoke(ctx, BFF_RelationAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListReply, error) {
	out := new(GetFriendListReply)
	err := c.cc.Invoke(ctx, BFF_GetFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListReply, error) {
	out := new(GetMessageListReply)
	err := c.cc.Invoke(ctx, BFF_GetMessageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) MessageAction(ctx context.Context, in *MessageActionRequest, opts ...grpc.CallOption) (*MessageActionReply, error) {
	out := new(MessageActionReply)
	err := c.cc.Invoke(ctx, BFF_MessageAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListReply, error) {
	out := new(GetFavoriteVideoListReply)
	err := c.cc.Invoke(ctx, BFF_GetFavoriteVideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionReply, error) {
	out := new(FavoriteActionReply)
	err := c.cc.Invoke(ctx, BFF_FavoriteAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) GetCommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListReply, error) {
	out := new(CommentListReply)
	err := c.cc.Invoke(ctx, BFF_GetCommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFFClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionReply, error) {
	out := new(CommentActionReply)
	err := c.cc.Invoke(ctx, BFF_CommentAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BFFServer is the server API for BFF service.
// All implementations must embed UnimplementedBFFServer
// for forward compatibility
type BFFServer interface {
	// 用户注册
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterReply, error)
	// 用户登陆
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginReply, error)
	// 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error)
	// 获取用户投稿视频列表
	GetPublishList(context.Context, *GetPublishListRequest) (*GetPublishListReply, error)
	// 用户发布视频
	PublishAction(context.Context, *PublishActionRequest) (*PublishActionReply, error)
	// 视频流
	Feed(context.Context, *FeedRequest) (*FeedReply, error)
	// 获取粉丝列表
	GetFollowerList(context.Context, *GetFollowerListRequest) (*GetFollowerListReply, error)
	// 获取关注列表
	GetFollowList(context.Context, *GetFollowListRequest) (*GetFollowListReply, error)
	// 关注或取关用户
	RelationAction(context.Context, *RelationActionRequest) (*RelationActionReply, error)
	// 获取好友列表
	GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListReply, error)
	// 获取消息列表
	GetMessageList(context.Context, *GetMessageListRequest) (*GetMessageListReply, error)
	// 给好友发送消息
	MessageAction(context.Context, *MessageActionRequest) (*MessageActionReply, error)
	// 获取点赞视频列表
	GetFavoriteVideoList(context.Context, *GetFavoriteVideoListRequest) (*GetFavoriteVideoListReply, error)
	// 点赞/取消点赞视频
	FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionReply, error)
	// 获取评论列表
	GetCommentList(context.Context, *CommentListRequest) (*CommentListReply, error)
	// 发布评论或者删除评论
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionReply, error)
	mustEmbedUnimplementedBFFServer()
}

// UnimplementedBFFServer must be embedded to have forward compatible implementations.
type UnimplementedBFFServer struct {
}

func (UnimplementedBFFServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedBFFServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedBFFServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedBFFServer) GetPublishList(context.Context, *GetPublishListRequest) (*GetPublishListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishList not implemented")
}
func (UnimplementedBFFServer) PublishAction(context.Context, *PublishActionRequest) (*PublishActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAction not implemented")
}
func (UnimplementedBFFServer) Feed(context.Context, *FeedRequest) (*FeedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedBFFServer) GetFollowerList(context.Context, *GetFollowerListRequest) (*GetFollowerListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerList not implemented")
}
func (UnimplementedBFFServer) GetFollowList(context.Context, *GetFollowListRequest) (*GetFollowListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowList not implemented")
}
func (UnimplementedBFFServer) RelationAction(context.Context, *RelationActionRequest) (*RelationActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedBFFServer) GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedBFFServer) GetMessageList(context.Context, *GetMessageListRequest) (*GetMessageListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageList not implemented")
}
func (UnimplementedBFFServer) MessageAction(context.Context, *MessageActionRequest) (*MessageActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageAction not implemented")
}
func (UnimplementedBFFServer) GetFavoriteVideoList(context.Context, *GetFavoriteVideoListRequest) (*GetFavoriteVideoListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteVideoList not implemented")
}
func (UnimplementedBFFServer) FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedBFFServer) GetCommentList(context.Context, *CommentListRequest) (*CommentListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedBFFServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedBFFServer) mustEmbedUnimplementedBFFServer() {}

// UnsafeBFFServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BFFServer will
// result in compilation errors.
type UnsafeBFFServer interface {
	mustEmbedUnimplementedBFFServer()
}

func RegisterBFFServer(s grpc.ServiceRegistrar, srv BFFServer) {
	s.RegisterService(&BFF_ServiceDesc, srv)
}

func _BFF_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetPublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetPublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetPublishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetPublishList(ctx, req.(*GetPublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_PublishAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).PublishAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_PublishAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).PublishAction(ctx, req.(*PublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_Feed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).Feed(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetFollowerList(ctx, req.(*GetFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetFollowList(ctx, req.(*GetFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_RelationAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).RelationAction(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetFriendList(ctx, req.(*GetFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetMessageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetMessageList(ctx, req.(*GetMessageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_MessageAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).MessageAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_MessageAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).MessageAction(ctx, req.(*MessageActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetFavoriteVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteVideoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetFavoriteVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetFavoriteVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetFavoriteVideoList(ctx, req.(*GetFavoriteVideoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).FavoriteAction(ctx, req.(*FavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_GetCommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).GetCommentList(ctx, req.(*CommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFF_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFF_CommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BFF_ServiceDesc is the grpc.ServiceDesc for BFF service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BFF_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bff.service.v1.BFF",
	HandlerType: (*BFFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _BFF_UserRegister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _BFF_UserLogin_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _BFF_GetUserInfo_Handler,
		},
		{
			MethodName: "GetPublishList",
			Handler:    _BFF_GetPublishList_Handler,
		},
		{
			MethodName: "PublishAction",
			Handler:    _BFF_PublishAction_Handler,
		},
		{
			MethodName: "Feed",
			Handler:    _BFF_Feed_Handler,
		},
		{
			MethodName: "GetFollowerList",
			Handler:    _BFF_GetFollowerList_Handler,
		},
		{
			MethodName: "GetFollowList",
			Handler:    _BFF_GetFollowList_Handler,
		},
		{
			MethodName: "RelationAction",
			Handler:    _BFF_RelationAction_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _BFF_GetFriendList_Handler,
		},
		{
			MethodName: "GetMessageList",
			Handler:    _BFF_GetMessageList_Handler,
		},
		{
			MethodName: "MessageAction",
			Handler:    _BFF_MessageAction_Handler,
		},
		{
			MethodName: "GetFavoriteVideoList",
			Handler:    _BFF_GetFavoriteVideoList_Handler,
		},
		{
			MethodName: "FavoriteAction",
			Handler:    _BFF_FavoriteAction_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _BFF_GetCommentList_Handler,
		},
		{
			MethodName: "CommentAction",
			Handler:    _BFF_CommentAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bff/bff.proto",
}
