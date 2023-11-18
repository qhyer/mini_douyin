package data

import (
	"context"
	"douyin/app/video/favorite/common/constants"
	do "douyin/app/video/favorite/common/entity"
	"douyin/app/video/favorite/job/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type favoriteRepo struct {
	data *Data
	log  *log.Helper
}

func NewFavoriteRepo(data *Data, logger log.Logger) biz.FavoriteRepo {
	return &favoriteRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *favoriteRepo) UpdateUserFavoriteCount(ctx context.Context, userId int64, incr int64) error {
	err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).Where("user_id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFavoriteCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) UpdateUserFavoritedCount(ctx context.Context, userId int64, incr int64) error {
	err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).Where("user_id = ?", userId).Update("favorited_count", gorm.Expr("favorited_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) UpdateVideoFavoritedCount(ctx context.Context, videoId int64, incr int64) error {
	err := r.data.db.WithContext(ctx).Table(constants.VideoFavoritedCountTableName(videoId)).Where("video_id = ?", videoId).Update("favorited_count", gorm.Expr("favorited_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateVideoFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) CreateFavorite(ctx context.Context, favorite *do.Favorite) error {
	// TODO do to po
	err := r.data.db.WithContext(ctx).Table(constants.FavoriteVideoRecordTableName(favorite.UserId)).Create(favorite).Error
	if err != nil {
		r.log.Errorf("CreateFavorite error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) DeleteFavorite(ctx context.Context, favorite *do.Favorite) error {
	// TODO do to po
	err := r.data.db.WithContext(ctx).Table(constants.FavoriteVideoRecordTableName(favorite.UserId)).Where("user_id = ? and video_id = ?", favorite.UserId, favorite.VideoId).Delete(favorite).Error
	if err != nil {
		r.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) BatchUpdateUserFavoriteCount(ctx context.Context, userIds []int64, incr []int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchUpdateUserFavoritedCount(ctx context.Context, userIds []int64, incr []int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchUpdateVideoFavoritedCount(ctx context.Context, videoIds []int64, incr []int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchCreateFavorite(ctx context.Context, favorites []*do.Favorite) error {
	//TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchDeleteFavorite(ctx context.Context, favorites []*do.Favorite) error {
	//TODO implement me
	panic("implement me")
}
