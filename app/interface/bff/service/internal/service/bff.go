package service

import (
	"context"

	v1 "douyin/api/bff"
)

type BFFService struct {
	v1.UnimplementedBFFServer
}

func NewBFFService() *BFFService {
	return &BFFService{}
}

func (s *BFFService) UserRegister(ctx context.Context, req *v1.UserRegisterRequest) (*v1.UserRegisterReply, error) {
	return &v1.UserRegisterReply{}, nil
}
func (s *BFFService) UserLogin(ctx context.Context, req *v1.UserLoginRequest) (*v1.UserLoginReply, error) {
	return &v1.UserLoginReply{}, nil
}
func (s *BFFService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoRequest) (*v1.GetUserInfoReply, error) {
	return &v1.GetUserInfoReply{}, nil
}
func (s *BFFService) GetPublishList(ctx context.Context, req *v1.GetPublishListRequest) (*v1.GetPublishListReply, error) {
	return &v1.GetPublishListReply{}, nil
}
func (s *BFFService) PublishAction(ctx context.Context, req *v1.PublishActionRequest) (*v1.PublishActionReply, error) {
	return &v1.PublishActionReply{}, nil
}
func (s *BFFService) Feed(ctx context.Context, req *v1.FeedRequest) (*v1.FeedReply, error) {
	return &v1.FeedReply{}, nil
}
func (s *BFFService) GetFollowerList(ctx context.Context, req *v1.GetFollowerListRequest) (*v1.GetFollowerListReply, error) {
	return &v1.GetFollowerListReply{}, nil
}
func (s *BFFService) GetFollowList(ctx context.Context, req *v1.GetFollowListRequest) (*v1.GetFollowListReply, error) {
	return &v1.GetFollowListReply{}, nil
}
func (s *BFFService) RelationAction(ctx context.Context, req *v1.RelationActionRequest) (*v1.RelationActionReply, error) {
	return &v1.RelationActionReply{}, nil
}
func (s *BFFService) GetFriendList(ctx context.Context, req *v1.GetFriendListRequest) (*v1.GetFriendListReply, error) {
	return &v1.GetFriendListReply{}, nil
}
func (s *BFFService) GetMessageList(ctx context.Context, req *v1.GetMessageListRequest) (*v1.GetMessageListReply, error) {
	return &v1.GetMessageListReply{}, nil
}
func (s *BFFService) MessageAction(ctx context.Context, req *v1.MessageActionRequest) (*v1.MessageActionReply, error) {
	return &v1.MessageActionReply{}, nil
}
func (s *BFFService) GetFavoriteVideoList(ctx context.Context, req *v1.GetFavoriteVideoListRequest) (*v1.GetFavoriteVideoListReply, error) {
	return &v1.GetFavoriteVideoListReply{}, nil
}
func (s *BFFService) FavoriteAction(ctx context.Context, req *v1.FavoriteActionRequest) (*v1.FavoriteActionReply, error) {
	return &v1.FavoriteActionReply{}, nil
}
func (s *BFFService) GetCommentList(ctx context.Context, req *v1.CommentListRequest) (*v1.CommentListReply, error) {
	return &v1.CommentListReply{}, nil
}
func (s *BFFService) CommentAction(ctx context.Context, req *v1.CommentActionRequest) (*v1.CommentActionReply, error) {
	return &v1.CommentActionReply{}, nil
}
