package biz

import (
	"context"

	do "douyin/app/video/favorite/common/event"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"

	"douyin/app/video/favorite/common/constants"
	"douyin/common/ecode"
)

type FavoriteRepo interface {
	GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error)
	IsUserFavoriteVideo(ctx context.Context, userId int64, videoId int64) (bool, error)
	IsUserFavoriteVideoList(ctx context.Context, userId int64, videoIds []int64) ([]bool, error)
	FavoriteVideo(ctx context.Context, fav *do.FavoriteAction) error
	CountUserFavoriteByUserId(ctx context.Context, userId int64) (int64, error)
	CountUserFavoritedByUserId(ctx context.Context, userId int64) (int64, error)
	CountVideoFavoritedByVideoId(ctx context.Context, videoId int64) (int64, error)
	MCountVideoFavoritedByVideoId(ctx context.Context, videoId []int64) ([]int64, error)
}

type FavoriteUsecase struct {
	repo FavoriteRepo
	log  *log.Helper
	sf   *singleflight.Group
}

func NewFavoriteUsecase(repo FavoriteRepo, logger log.Logger) *FavoriteUsecase {
	return &FavoriteUsecase{repo: repo, log: log.NewHelper(logger), sf: &singleflight.Group{}}
}

// FavoriteAction 点赞视频
func (uc *FavoriteUsecase) FavoriteAction(ctx context.Context, fav *do.FavoriteAction) error {
	if fav.Type != do.FavoriteActionAdd && fav.Type != do.FavoriteActionDelete {
		return ecode.ParamErr
	}
	err := uc.repo.FavoriteVideo(ctx, fav)
	if err != nil {
		return err
	}
	return nil
}

// GetFavoriteVideoIdListByUserId 获取用户点赞视频id列表
func (uc *FavoriteUsecase) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err, _ := uc.sf.Do(constants.SFUserFavoriteIdListKey(userId), func() (interface{}, error) {
		return uc.repo.GetFavoriteVideoIdListByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("GetFavoriteVideoIdListByUserId error(%v)", err)
		return nil, err
	}
	return res.([]int64), nil
}

// GetFavoriteStatusByUserIdAndVideoIds 批量获取用户是否点赞视频
func (uc *FavoriteUsecase) GetFavoriteStatusByUserIdAndVideoIds(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	isFavoriteList, err := uc.repo.IsUserFavoriteVideoList(ctx, userId, videoIds)
	if err != nil {
		return nil, err
	}
	return isFavoriteList, nil
}

// CountUserFavoriteByUserId 获取用户点赞数
func (uc *FavoriteUsecase) CountUserFavoriteByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err, _ := uc.sf.Do(constants.SFCountUserFavoriteKey(userId), func() (interface{}, error) {
		return uc.repo.CountUserFavoriteByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("CountUserFavoriteByUserId error(%v)", err)
		return 0, err
	}
	return res.(int64), nil
}

// CountUserFavoritedByUserId 获取用户视频获赞数
func (uc *FavoriteUsecase) CountUserFavoritedByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err, _ := uc.sf.Do(constants.SFCountUserFavoritedKey(userId), func() (interface{}, error) {
		return uc.repo.CountUserFavoritedByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("CountUserFavoritedByUserId error(%v)", err)
		return 0, err
	}
	return res.(int64), nil
}

// CountVideoFavoritedByVideoId 获取视频被点赞数
func (uc *FavoriteUsecase) CountVideoFavoritedByVideoId(ctx context.Context, videoId int64) (int64, error) {
	res, err, _ := uc.sf.Do(constants.SFCountVideoFavoritedKey(videoId), func() (interface{}, error) {
		return uc.repo.CountVideoFavoritedByVideoId(ctx, videoId)
	})
	if err != nil {
		uc.log.Errorf("CountVideoFavoritedByVideoId error(%v)", err)
		return 0, err
	}
	return res.(int64), nil
}

// MCountVideoFavoritedByVideoId 批量获取视频被点赞数
func (uc *FavoriteUsecase) MCountVideoFavoritedByVideoId(ctx context.Context, videoId []int64) ([]int64, error) {
	return uc.repo.MCountVideoFavoritedByVideoId(ctx, videoId)
}
