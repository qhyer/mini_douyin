package biz

import (
	"context"
	favorite "douyin/api/video/favorite/service/v1"
	video "douyin/api/video/feed/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type FavoriteRepo interface {
	FavoriteAction(ctx context.Context, userId, videoId int64, actionType int32) (*favorite.DouyinFavoriteActionResponse, error)
	GetUserFavoriteVideoList(ctx context.Context, userId, toUserId int64) (*video.GetUserFavoriteVideoListByUserIdResponse, error)
}

type FavoriteUsecase struct {
	repo FavoriteRepo
	log  *log.Helper
}

func NewFavoriteUsecase(repo FavoriteRepo, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FavoriteUsecase) FavoriteAction(ctx context.Context, userId, videoId int64, actionType int32) (*favorite.DouyinFavoriteActionResponse, error) {
	res, err := uc.repo.FavoriteAction(ctx, userId, videoId, actionType)
	if err != nil {
		uc.log.Errorf("FavoriteAction error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *FavoriteUsecase) GetUserFavoriteVideoList(ctx context.Context, userId, toUserId int64) (*video.GetUserFavoriteVideoListByUserIdResponse, error) {
	res, err := uc.repo.GetUserFavoriteVideoList(ctx, userId, toUserId)
	if err != nil {
		uc.log.Errorf("GetUserFavoriteVideoList error: %v", err)
		return nil, err
	}
	return res, nil
}
