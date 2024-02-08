package biz

import (
	"context"

	"douyin/app/video/favorite/common/event"

	"github.com/go-kratos/kratos/v2/log"
)

type FavoriteRepo interface {
	UpdateUserFavoriteCount(ctx context.Context, userId int64, incr int64) error
	UpdateUserFavoritedCount(ctx context.Context, userId int64, incr int64) error
	UpdateVideoFavoritedCount(ctx context.Context, videoId int64, incr int64) error
	UpdateVideoFavoritedTempCount(ctx context.Context, procId int, videoId int64, incr int64) error
	PurgeVideoFavoritedTempCount(ctx context.Context, procId int) error
	GetVideoFavoritedCountFromCache(ctx context.Context, procId int) (map[int64]int64, error)
	CreateFavorite(ctx context.Context, favorite *event.FavoriteAction) error
	DeleteFavorite(ctx context.Context, favorite *event.FavoriteAction) error
	BatchUpdateUserFavoriteCount(ctx context.Context, stats map[int64]int64) error
	BatchUpdateUserFavoritedCount(ctx context.Context, stats map[int64]int64) error
	BatchUpdateVideoFavoritedCount(ctx context.Context, stats map[int64]int64) error
	BatchCreateFavorite(ctx context.Context, favorites []*event.FavoriteAction) error
	BatchDeleteFavorite(ctx context.Context, favorites []*event.FavoriteAction) error
}

type FavoriteUsecase struct {
	repo FavoriteRepo
	log  *log.Helper
}

func NewFavoriteUsecase(repo FavoriteRepo, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FavoriteUsecase) CreateFavorite(ctx context.Context, favorite *event.FavoriteAction) error {
	err := uc.repo.CreateFavorite(ctx, favorite)
	if err != nil {
		uc.log.Errorf("CreateFavorite error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) DeleteFavorite(ctx context.Context, favorite *event.FavoriteAction) error {
	err := uc.repo.DeleteFavorite(ctx, favorite)
	if err != nil {
		uc.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchCreateFavorite(ctx context.Context, favorites []*event.FavoriteAction) error {
	err := uc.repo.BatchCreateFavorite(ctx, favorites)
	if err != nil {
		uc.log.Errorf("BatchCreateFavorite error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchDeleteFavorite(ctx context.Context, favorites []*event.FavoriteAction) error {
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

func (uc *FavoriteUsecase) BatchUpdateUserFavoriteCount(ctx context.Context, stats map[int64]int64) error {
	err := uc.repo.BatchUpdateUserFavoriteCount(ctx, stats)
	if err != nil {
		uc.log.Errorf("BatchUpdateUserFavoriteCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchUpdateUserFavoritedCount(ctx context.Context, stats map[int64]int64) error {
	err := uc.repo.BatchUpdateUserFavoritedCount(ctx, stats)
	if err != nil {
		uc.log.Errorf("BatchUpdateUserFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) BatchUpdateVideoFavoritedCount(ctx context.Context, stats map[int64]int64) error {
	err := uc.repo.BatchUpdateVideoFavoritedCount(ctx, stats)
	if err != nil {
		uc.log.Errorf("BatchUpdateVideoFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) UpdateVideoFavoritedTempCount(ctx context.Context, procId int, videoId int64, incr int64) error {
	err := uc.repo.UpdateVideoFavoritedTempCount(ctx, procId, videoId, incr)
	if err != nil {
		uc.log.Errorf("UpdateVideoFavoritedTempCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) PurgeVideoFavoritedTempCount(ctx context.Context, procId int) error {
	err := uc.repo.PurgeVideoFavoritedTempCount(ctx, procId)
	if err != nil {
		uc.log.Errorf("PurgeVideoFavoritedTempCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *FavoriteUsecase) GetVideoFavoritedCountFromCache(ctx context.Context, procId int) (map[int64]int64, error) {
	stats, err := uc.repo.GetVideoFavoritedCountFromCache(ctx, procId)
	if err != nil {
		uc.log.Errorf("GetVideoFavoritedCountFromCache error(%v)", err)
		return nil, err
	}
	return stats, nil
}
