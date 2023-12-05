package service

import (
	"context"
	"douyin/app/video/feed/service/internal/biz"

	"douyin/api/video/feed/service/v1"
)

type FeedService struct {
	v1.UnimplementedFeedServer

	uc *biz.FeedUsecase
}

func NewFeedService(uc *biz.FeedUsecase) *FeedService {
	return &FeedService{uc: uc}
}

func (s *FeedService) Feed(ctx context.Context, req *v1.FeedRequest) (*v1.FeedResponse, error) {
	return &v1.FeedResponse{}, nil
}
func (s *FeedService) GetPublishedVideoByUserId(ctx context.Context, req *v1.GetPublishedVideoByUserIdRequest) (*v1.GetPublishedVideoByUserIdResponse, error) {
	return &v1.GetPublishedVideoByUserIdResponse{}, nil
}
func (s *FeedService) GetUserFavoriteVideoListByUserId(ctx context.Context, req *v1.GetUserFavoriteVideoListByUserIdRequest) (*v1.GetUserFavoriteVideoListByUserIdResponse, error) {
	return &v1.GetUserFavoriteVideoListByUserIdResponse{}, nil
}
