package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	do "douyin/app/video/favorite/common/entity"
)

type FavoriteRepo interface {
	UpdateUserFavoriteCount(ctx context.Context, userId int64, incr int64) error
	UpdateUserFavoritedCount(ctx context.Context, userId int64, incr int64) error
	UpdateVideoFavoritedCount(ctx context.Context, videoId int64, incr int64) error
	CreateFavorite(ctx context.Context, favorite *do.FavoriteAction) error
	DeleteFavorite(ctx context.Context, favorite *do.FavoriteAction) error
	BatchUpdateUserFavoriteCount(ctx context.Context, userIds []int64, incr []int64) error
	BatchUpdateUserFavoritedCount(ctx context.Context, userIds []int64, incr []int64) error
	BatchUpdateVideoFavoritedCount(ctx context.Context, videoIds []int64, incr []int64) error
	BatchCreateFavorite(ctx context.Context, favorites []*do.FavoriteAction) error
	BatchDeleteFavorite(ctx context.Context, favorites []*do.FavoriteAction) error
}

type FavoriteUsecase struct {
	repo FavoriteRepo
	log  *log.Helper
}

func NewFavoriteUsecase(repo FavoriteRepo, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FavoriteUsecase) CreateFavorite(ctx context.Context, favorite *do.FavoriteAction) error {
	err := uc.repo.CreateFavorite(ctx, favorite)
	if err != nil {
		uc.log.Errorf("CreateFavorite error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) DeleteFavorite(ctx context.Context, favorite *do.FavoriteAction) error {
	err := uc.repo.DeleteFavorite(ctx, favorite)
	if err != nil {
		uc.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchCreateFavorite(ctx context.Context, favorites []*do.FavoriteAction) error {
	err := uc.repo.BatchCreateFavorite(ctx, favorites)
	if err != nil {
		uc.log.Errorf("BatchCreateFavorite error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchDeleteFavorite(ctx context.Context, favorites []*do.FavoriteAction) error {
	err := uc.repo.BatchDeleteFavorite(ctx, favorites)
	if err != nil {
		uc.log.Errorf("BatchDeleteFavorite error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) UpdateUserFavoriteCount(ctx context.Context, userId int64, incr int64) error {
	err := uc.repo.UpdateUserFavoriteCount(ctx, userId, incr)
	if err != nil {
		uc.log.Errorf("UpdateUserFavoriteCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) UpdateUserFavoritedCount(ctx context.Context, userId int64, incr int64) error {
	err := uc.repo.UpdateUserFavoritedCount(ctx, userId, incr)
	if err != nil {
		uc.log.Errorf("UpdateUserFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) UpdateVideoFavoritedCount(ctx context.Context, videoId int64, incr int64) error {
	err := uc.repo.UpdateVideoFavoritedCount(ctx, videoId, incr)
	if err != nil {
		uc.log.Errorf("UpdateVideoFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchUpdateUserFavoriteCount(ctx context.Context, userIds []int64, incr []int64) error {
	err := uc.repo.BatchUpdateUserFavoriteCount(ctx, userIds, incr)
	if err != nil {
		uc.log.Errorf("BatchUpdateUserFavoriteCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchUpdateUserFavoritedCount(ctx context.Context, userIds []int64, incr []int64) error {
	err := uc.repo.BatchUpdateUserFavoritedCount(ctx, userIds, incr)
	if err != nil {
		uc.log.Errorf("BatchUpdateUserFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchUpdateVideoFavoritedCount(ctx context.Context, videoIds []int64, incr []int64) error {
	err := uc.repo.BatchUpdateVideoFavoritedCount(ctx, videoIds, incr)
	if err != nil {
		uc.log.Errorf("BatchUpdateVideoFavoritedCount error(%v)", err)
		return err
	}
	return nil
}
