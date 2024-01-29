// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.25.2
// source: bff/bff.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBFFCommentAction = "/bff.service.v1.BFF/CommentAction"
const OperationBFFFavoriteAction = "/bff.service.v1.BFF/FavoriteAction"
const OperationBFFFeed = "/bff.service.v1.BFF/Feed"
const OperationBFFGetCommentList = "/bff.service.v1.BFF/GetCommentList"
const OperationBFFGetFavoriteVideoList = "/bff.service.v1.BFF/GetFavoriteVideoList"
const OperationBFFGetFollowList = "/bff.service.v1.BFF/GetFollowList"
const OperationBFFGetFollowerList = "/bff.service.v1.BFF/GetFollowerList"
const OperationBFFGetFriendList = "/bff.service.v1.BFF/GetFriendList"
const OperationBFFGetMessageList = "/bff.service.v1.BFF/GetMessageList"
const OperationBFFGetPublishList = "/bff.service.v1.BFF/GetPublishList"
const OperationBFFGetUserInfo = "/bff.service.v1.BFF/GetUserInfo"
const OperationBFFMessageAction = "/bff.service.v1.BFF/MessageAction"
const OperationBFFPublishAction = "/bff.service.v1.BFF/PublishAction"
const OperationBFFRelationAction = "/bff.service.v1.BFF/RelationAction"
const OperationBFFUserLogin = "/bff.service.v1.BFF/UserLogin"
const OperationBFFUserRegister = "/bff.service.v1.BFF/UserRegister"

type BFFHTTPServer interface {
	// CommentAction 发布评论或者删除评论
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionReply, error)
	// FavoriteAction 点赞/取消点赞视频
	FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionReply, error)
	// Feed 视频流
	Feed(context.Context, *FeedRequest) (*FeedReply, error)
	// GetCommentList 获取评论列表
	GetCommentList(context.Context, *CommentListRequest) (*CommentListReply, error)
	// GetFavoriteVideoList 获取点赞视频列表
	GetFavoriteVideoList(context.Context, *GetFavoriteVideoListRequest) (*GetFavoriteVideoListReply, error)
	// GetFollowList 获取关注列表
	GetFollowList(context.Context, *GetFollowListRequest) (*GetFollowListReply, error)
	// GetFollowerList 获取粉丝列表
	GetFollowerList(context.Context, *GetFollowerListRequest) (*GetFollowerListReply, error)
	// GetFriendList 获取好友列表
	GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListReply, error)
	// GetMessageList 获取消息列表
	GetMessageList(context.Context, *GetMessageListRequest) (*GetMessageListReply, error)
	// GetPublishList 获取用户投稿视频列表
	GetPublishList(context.Context, *GetPublishListRequest) (*GetPublishListReply, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error)
	// MessageAction 给好友发送消息
	MessageAction(context.Context, *MessageActionRequest) (*MessageActionReply, error)
	// PublishAction 用户发布视频
	PublishAction(context.Context, *PublishActionRequest) (*PublishActionReply, error)
	// RelationAction 关注或取关用户
	RelationAction(context.Context, *RelationActionRequest) (*RelationActionReply, error)
	// UserLogin 用户登陆
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginReply, error)
	// UserRegister 用户注册
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterReply, error)
}

func RegisterBFFHTTPServer(s *http.Server, srv BFFHTTPServer) {
	r := s.Route("/")
	r.POST("/douyin/user/register", _BFF_UserRegister0_HTTP_Handler(srv))
	r.POST("/douyin/user/login", _BFF_UserLogin0_HTTP_Handler(srv))
	r.GET("/douyin/user", _BFF_GetUserInfo0_HTTP_Handler(srv))
	r.GET("/douyin/publish/list", _BFF_GetPublishList0_HTTP_Handler(srv))
	r.POST("/douyin/publish/action", _BFF_PublishAction0_HTTP_Handler(srv))
	r.GET("/douyin/feed", _BFF_Feed0_HTTP_Handler(srv))
	r.GET("/douyin/relation/follower/list", _BFF_GetFollowerList0_HTTP_Handler(srv))
	r.GET("/douyin/relation/follow/list", _BFF_GetFollowList0_HTTP_Handler(srv))
	r.POST("/douyin/relation/action", _BFF_RelationAction0_HTTP_Handler(srv))
	r.GET("/douyin/relation/friend/list", _BFF_GetFriendList0_HTTP_Handler(srv))
	r.GET("/douyin/message/chat", _BFF_GetMessageList0_HTTP_Handler(srv))
	r.POST("/douyin/message/action", _BFF_MessageAction0_HTTP_Handler(srv))
	r.GET("/douyin/favorite/list", _BFF_GetFavoriteVideoList0_HTTP_Handler(srv))
	r.POST("/douyin/favorite/action", _BFF_FavoriteAction0_HTTP_Handler(srv))
	r.GET("/douyin/comment/list", _BFF_GetCommentList0_HTTP_Handler(srv))
	r.POST("/douyin/comment/action", _BFF_CommentAction0_HTTP_Handler(srv))
}

func _BFF_UserRegister0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserRegister(ctx, req.(*UserRegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserRegisterReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_UserLogin0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserLogin(ctx, req.(*UserLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserLoginReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetUserInfo0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetPublishList0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPublishListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetPublishList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPublishList(ctx, req.(*GetPublishListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPublishListReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_PublishAction0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PublishActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFPublishAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PublishAction(ctx, req.(*PublishActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PublishActionReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_Feed0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFFeed)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Feed(ctx, req.(*FeedRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeedReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetFollowerList0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFollowerListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetFollowerList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFollowerList(ctx, req.(*GetFollowerListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFollowerListReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetFollowList0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFollowListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetFollowList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFollowList(ctx, req.(*GetFollowListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFollowListReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_RelationAction0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RelationActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFRelationAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RelationAction(ctx, req.(*RelationActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RelationActionReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetFriendList0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFriendListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetFriendList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFriendList(ctx, req.(*GetFriendListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFriendListReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetMessageList0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMessageListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetMessageList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMessageList(ctx, req.(*GetMessageListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMessageListReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_MessageAction0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MessageActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFMessageAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MessageAction(ctx, req.(*MessageActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MessageActionReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetFavoriteVideoList0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFavoriteVideoListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetFavoriteVideoList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFavoriteVideoList(ctx, req.(*GetFavoriteVideoListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFavoriteVideoListReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_FavoriteAction0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFFavoriteAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteAction(ctx, req.(*FavoriteActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FavoriteActionReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_GetCommentList0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFGetCommentList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCommentList(ctx, req.(*CommentListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentListReply)
		return ctx.Result(200, reply)
	}
}

func _BFF_CommentAction0_HTTP_Handler(srv BFFHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBFFCommentAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentAction(ctx, req.(*CommentActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentActionReply)
		return ctx.Result(200, reply)
	}
}

type BFFHTTPClient interface {
	CommentAction(ctx context.Context, req *CommentActionRequest, opts ...http.CallOption) (rsp *CommentActionReply, err error)
	FavoriteAction(ctx context.Context, req *FavoriteActionRequest, opts ...http.CallOption) (rsp *FavoriteActionReply, err error)
	Feed(ctx context.Context, req *FeedRequest, opts ...http.CallOption) (rsp *FeedReply, err error)
	GetCommentList(ctx context.Context, req *CommentListRequest, opts ...http.CallOption) (rsp *CommentListReply, err error)
	GetFavoriteVideoList(ctx context.Context, req *GetFavoriteVideoListRequest, opts ...http.CallOption) (rsp *GetFavoriteVideoListReply, err error)
	GetFollowList(ctx context.Context, req *GetFollowListRequest, opts ...http.CallOption) (rsp *GetFollowListReply, err error)
	GetFollowerList(ctx context.Context, req *GetFollowerListRequest, opts ...http.CallOption) (rsp *GetFollowerListReply, err error)
	GetFriendList(ctx context.Context, req *GetFriendListRequest, opts ...http.CallOption) (rsp *GetFriendListReply, err error)
	GetMessageList(ctx context.Context, req *GetMessageListRequest, opts ...http.CallOption) (rsp *GetMessageListReply, err error)
	GetPublishList(ctx context.Context, req *GetPublishListRequest, opts ...http.CallOption) (rsp *GetPublishListReply, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoReply, err error)
	MessageAction(ctx context.Context, req *MessageActionRequest, opts ...http.CallOption) (rsp *MessageActionReply, err error)
	PublishAction(ctx context.Context, req *PublishActionRequest, opts ...http.CallOption) (rsp *PublishActionReply, err error)
	RelationAction(ctx context.Context, req *RelationActionRequest, opts ...http.CallOption) (rsp *RelationActionReply, err error)
	UserLogin(ctx context.Context, req *UserLoginRequest, opts ...http.CallOption) (rsp *UserLoginReply, err error)
	UserRegister(ctx context.Context, req *UserRegisterRequest, opts ...http.CallOption) (rsp *UserRegisterReply, err error)
}

type BFFHTTPClientImpl struct {
	cc *http.Client
}

func NewBFFHTTPClient(client *http.Client) BFFHTTPClient {
	return &BFFHTTPClientImpl{client}
}

func (c *BFFHTTPClientImpl) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...http.CallOption) (*CommentActionReply, error) {
	var out CommentActionReply
	pattern := "/douyin/comment/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBFFCommentAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...http.CallOption) (*FavoriteActionReply, error) {
	var out FavoriteActionReply
	pattern := "/douyin/favorite/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBFFFavoriteAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) Feed(ctx context.Context, in *FeedRequest, opts ...http.CallOption) (*FeedReply, error) {
	var out FeedReply
	pattern := "/douyin/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFFeed))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetCommentList(ctx context.Context, in *CommentListRequest, opts ...http.CallOption) (*CommentListReply, error) {
	var out CommentListReply
	pattern := "/douyin/comment/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetCommentList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...http.CallOption) (*GetFavoriteVideoListReply, error) {
	var out GetFavoriteVideoListReply
	pattern := "/douyin/favorite/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetFavoriteVideoList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetFollowList(ctx context.Context, in *GetFollowListRequest, opts ...http.CallOption) (*GetFollowListReply, error) {
	var out GetFollowListReply
	pattern := "/douyin/relation/follow/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetFollowList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...http.CallOption) (*GetFollowerListReply, error) {
	var out GetFollowerListReply
	pattern := "/douyin/relation/follower/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetFollowerList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...http.CallOption) (*GetFriendListReply, error) {
	var out GetFriendListReply
	pattern := "/douyin/relation/friend/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetFriendList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...http.CallOption) (*GetMessageListReply, error) {
	var out GetMessageListReply
	pattern := "/douyin/message/chat"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetMessageList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetPublishList(ctx context.Context, in *GetPublishListRequest, opts ...http.CallOption) (*GetPublishListReply, error) {
	var out GetPublishListReply
	pattern := "/douyin/publish/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetPublishList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoReply, error) {
	var out GetUserInfoReply
	pattern := "/douyin/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBFFGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) MessageAction(ctx context.Context, in *MessageActionRequest, opts ...http.CallOption) (*MessageActionReply, error) {
	var out MessageActionReply
	pattern := "/douyin/message/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBFFMessageAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) PublishAction(ctx context.Context, in *PublishActionRequest, opts ...http.CallOption) (*PublishActionReply, error) {
	var out PublishActionReply
	pattern := "/douyin/publish/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBFFPublishAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...http.CallOption) (*RelationActionReply, error) {
	var out RelationActionReply
	pattern := "/douyin/relation/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBFFRelationAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...http.CallOption) (*UserLoginReply, error) {
	var out UserLoginReply
	pattern := "/douyin/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBFFUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BFFHTTPClientImpl) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...http.CallOption) (*UserRegisterReply, error) {
	var out UserRegisterReply
	pattern := "/douyin/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBFFUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
