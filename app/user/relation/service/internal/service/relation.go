package service

import (
	"context"

	v1 "douyin/api/user/relation/service/v1"
)

type RelationService struct {
	v1.UnimplementedRelationServer
}

func NewRelationService() *RelationService {
	return &RelationService{}
}

func (s *RelationService) RelationAction(ctx context.Context, req *v1.DouyinRelationActionRequest) (*v1.DouyinRelationActionResponse, error) {
	return &v1.DouyinRelationActionResponse{}, nil
}
func (s *RelationService) GetFollowListByUserId(ctx context.Context, req *v1.GetFollowListByUserIdRequest) (*v1.GetFollowListByUserIdResponse, error) {
	return &v1.GetFollowListByUserIdResponse{}, nil
}
func (s *RelationService) GetFollowerListByUserId(ctx context.Context, req *v1.GetFollowerListByUserIdRequest) (*v1.GetFollowerListByUserIdResponse, error) {
	return &v1.GetFollowerListByUserIdResponse{}, nil
}
func (s *RelationService) GetUserFriendListByUserId(ctx context.Context, req *v1.GetFriendListByUserIdRequest) (*v1.GetFriendListByUserIdResponse, error) {
	return &v1.GetFriendListByUserIdResponse{}, nil
}
