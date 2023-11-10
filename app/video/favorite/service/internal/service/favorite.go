package service

import (
	"context"

	"douyin/api/video/favorite/service/v1"
)

type FavoriteService struct {
	v1.UnimplementedFavoriteServer
}

func NewFavoriteService() *FavoriteService {
	return &FavoriteService{}
}

func (s *FavoriteService) FavoriteAction(ctx context.Context, req *v1.DouyinFavoriteActionRequest) (*v1.DouyinFavoriteActionResponse, error) {
	return &v1.DouyinFavoriteActionResponse{}, nil
}
func (s *FavoriteService) GetUserFavoriteList(ctx context.Context, req *v1.GetUserFavoriteListRequest) (*v1.GetUserFavoriteListResponse, error) {
	return &v1.GetUserFavoriteListResponse{}, nil
}
func (s *FavoriteService) GetFavoriteListByVideoIds(ctx context.Context, req *v1.GetFavoriteListByVideoIdsRequest) (*v1.GetFavoriteListByVideoIdsResponse, error) {
	return &v1.GetFavoriteListByVideoIdsResponse{}, nil
}
