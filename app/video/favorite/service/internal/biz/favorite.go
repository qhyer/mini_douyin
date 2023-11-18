package biz

import (
	"context"
	do "douyin/app/video/favorite/common/entity"
	"douyin/common/ecode"
	"github.com/go-kratos/kratos/v2/log"
)

type FavoriteRepo interface {
	GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error)
	IsUserFavoriteVideo(ctx context.Context, userId int64, videoId int64) (bool, error)
	IsUserFavoriteVideoList(ctx context.Context, userId int64, videoIds []int64) ([]bool, error)
	FavoriteVideo(ctx context.Context, fav *do.Favorite) error
	CountVideoFavoriteByUserId(ctx context.Context, userId int64) (int64, error)
	CountVideoFavoritedByUserId(ctx context.Context, userId int64) (int64, error)
	CountFavoritedByVideoId(ctx context.Context, videoId int64) (int64, error)
}

type FavoriteUsecase struct {
	repo FavoriteRepo
	log  *log.Helper
}

func NewFavoriteUsecase(repo FavoriteRepo, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FavoriteUsecase) FavoriteAction(ctx context.Context, fav *do.Favorite) error {
	if fav.Type != do.FavoriteActionAdd && fav.Type != do.FavoriteActionDelete {
		return ecode.ParamErr
	}
	err := uc.repo.FavoriteVideo(ctx, fav)
	if err != nil {
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	return uc.repo.GetFavoriteVideoIdListByUserId(ctx, userId)
}

func (uc *FavoriteUsecase) GetFavoriteStatusByUserIdAndVideoIds(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	isFavoriteList, err := uc.repo.IsUserFavoriteVideoList(ctx, userId, videoIds)
	if err != nil {
		return nil, err
	}
	return isFavoriteList, nil
}

func (uc *FavoriteUsecase) CountVideoFavoriteByUserId(ctx context.Context, userId int64) (int64, error) {
	return uc.repo.CountVideoFavoriteByUserId(ctx, userId)
}

func (uc *FavoriteUsecase) CountVideoFavoritedByUserId(ctx context.Context, userId int64) (int64, error) {
	return uc.repo.CountVideoFavoritedByUserId(ctx, userId)
}

func (uc *FavoriteUsecase) CountVideoFavoritedByVideoId(ctx context.Context, videoId int64) (int64, error) {
	return uc.repo.CountFavoritedByVideoId(ctx, videoId)
}
