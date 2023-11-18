package biz

import (
	"context"
	do "douyin/app/video/favorite/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type FavoriteRepo interface {
	UpdateUserFavoriteCount(ctx context.Context, userId int64, incr int64) error
	UpdateUserFavoritedCount(ctx context.Context, userId int64, incr int64) error
	UpdateVideoFavoritedCount(ctx context.Context, videoId int64, incr int64) error
	CreateFavorite(ctx context.Context, favorite *do.Favorite) error
	DeleteFavorite(ctx context.Context, favorite *do.Favorite) error
	BatchUpdateUserFavoriteCount(ctx context.Context, userIds []int64, incr []int64) error
	BatchUpdateUserFavoritedCount(ctx context.Context, userIds []int64, incr []int64) error
	BatchUpdateVideoFavoritedCount(ctx context.Context, videoIds []int64, incr []int64) error
	BatchCreateFavorite(ctx context.Context, favorites []*do.Favorite) error
	BatchDeleteFavorite(ctx context.Context, favorites []*do.Favorite) error
}

type FavoriteUsecase struct {
	repo FavoriteRepo
	log  *log.Helper
}

func NewFavoriteUsecase(repo FavoriteRepo, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{repo: repo, log: log.NewHelper(logger)}
}
