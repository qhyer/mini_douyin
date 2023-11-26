package service

import (
	"context"
	"douyin/app/user/relation/service/internal/biz"

	v1 "douyin/api/user/relation/service/v1"
)

type RelationService struct {
	v1.UnimplementedRelationServer

	uc *biz.RelationUsecase
}

func NewRelationService(uc *biz.RelationUsecase) *RelationService {
	return &RelationService{uc: uc}
}

func (s *RelationService) RelationAction(ctx context.Context, req *v1.RelationActionRequest) (*v1.RelationActionResponse, error) {
	return &v1.RelationActionResponse{}, nil
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

func (s *RelationService) CountFollowByUserId(ctx context.Context, req *v1.CountFollowByUserIdRequest) (*v1.CountFollowByUserIdResponse, error) {
	return &v1.CountFollowByUserIdResponse{}, nil
}

func (s *RelationService) CountFollowerByUserId(ctx context.Context, req *v1.CountFollowerByUserIdRequest) (*v1.CountFollowerByUserIdResponse, error) {
	return &v1.CountFollowerByUserIdResponse{}, nil
}

func (s *RelationService) IsFollowByUserId(ctx context.Context, req *v1.IsFollowByUserIdRequest) (*v1.IsFollowByUserIdResponse, error) {
	return &v1.IsFollowByUserIdResponse{}, nil
}

func (s *RelationService) IsFollowByUserIds(ctx context.Context, req *v1.IsFollowByUserIdsRequest) (*v1.IsFollowByUserIdsResponse, error) {
	return &v1.IsFollowByUserIdsResponse{}, nil
}
