package data

import (
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	seq "douyin/api/seq-server/service/v1"
	"douyin/app/video/favorite/common/constants"
	do "douyin/app/video/favorite/common/entity"
	"douyin/app/video/favorite/common/mapper"
	po "douyin/app/video/favorite/common/model"
	"douyin/app/video/favorite/job/internal/biz"
	constants2 "douyin/common/constants"
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
	usFav := po.UserFavoriteCount{UserId: userId}
	err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).FirstOrInit(&usFav, usFav).Update("favorite_count", gorm.Expr("favorite_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFavoriteCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) UpdateUserFavoritedCount(ctx context.Context, userId int64, incr int64) error {
	usFav := po.UserFavoriteCount{UserId: userId}
	err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).FirstOrInit(&usFav, usFav).Update("favorited_count", gorm.Expr("favorited_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) UpdateVideoFavoritedCount(ctx context.Context, videoId int64, incr int64) error {
	videoFav := po.VideoFavoritedCount{VideoId: videoId}
	err := r.data.db.WithContext(ctx).Table(constants.VideoFavoritedCountTableName(videoId)).FirstOrInit(&videoFav, videoFav).Update("favorited_count", gorm.Expr("favorited_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateVideoFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) CreateFavorite(ctx context.Context, favorite *do.FavoriteAction) error {
	// 获取点赞ID
	fid, err := r.data.seqRPC.GetID(ctx, &seq.GetIDRequest{
		BusinessId: constants2.FavoriteBusinessId,
	})
	if err != nil || !fid.GetIsOk() {
		r.log.Errorf("Get seq num error(%v)", err)
		return err
	}
	favorite.ID = fid.GetID()
	fav, err := mapper.FavoriteToPO(favorite)
	if err != nil {
		r.log.Errorf("CreateFavorite error(%v)", err)
		return err
	}
	// 删除用户喜欢缓存
	err = r.delUserFavoriteListCache(ctx, favorite.UserId)
	if err != nil {
		r.log.Errorf("CreateFavorite error(%v)", err)
	}
	// 创建喜欢
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建点赞记录
		err := tx.Table(constants.FavoriteVideoRecordTableName(favorite.UserId)).Create(fav).Error
		if err != nil {
			return err
		}

		// 更新用户点赞数
		err = tx.Table(constants.UserFavoriteVideoCountTableName(favorite.UserId)).FirstOrInit(&po.UserFavoriteCount{UserId: favorite.UserId}, po.UserFavoriteCount{UserId: favorite.UserId}).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			return err
		}

		// 异步更新视频获赞数和用户获赞数
		b, err := favorite.MarshalJson()
		_, _, err = r.data.kafkaProducer.SendMessage(&sarama.ProducerMessage{
			Topic: constants.UpdateVideoFavoritedCountTopic,
			Key:   sarama.StringEncoder(constants.UpdateVideoFavoritedCountKafkaKey(favorite.VideoId)),
			Value: sarama.ByteEncoder(b),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("CreateFavorite error(%v)", err)
		return err
	}

	// 延时删除用户喜欢缓存
	err = r.data.cacheFan.Do(ctx, func(ctx context.Context) {
		time.Sleep(100 * time.Millisecond)
		err = r.delUserFavoriteListCache(ctx, favorite.UserId)
		if err != nil {
			r.log.Errorf("DelUserFavoriteCache error(%v)", err)
		}
	})
	if err != nil {
		r.log.Errorf("Fanout error(%v)", err)
	}
	return nil
}

func (r *favoriteRepo) DeleteFavorite(ctx context.Context, favorite *do.FavoriteAction) error {
	fav, err := mapper.FavoriteToPO(favorite)
	if err != nil {
		r.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Table(constants.FavoriteVideoRecordTableName(favorite.UserId)).Where("user_id = ? and video_id = ?", fav.UserId, fav.VideoId).Delete(fav).Error
	if err != nil {
		r.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) BatchUpdateUserFavoriteCount(ctx context.Context, userIds []int64, incr []int64) error {
	// TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchUpdateUserFavoritedCount(ctx context.Context, userIds []int64, incr []int64) error {
	// TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchUpdateVideoFavoritedCount(ctx context.Context, videoIds []int64, incr []int64) error {
	// TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchCreateFavorite(ctx context.Context, favorites []*do.FavoriteAction) error {
	// TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchDeleteFavorite(ctx context.Context, favorites []*do.FavoriteAction) error {
	// TODO implement me
	panic("implement me")
}

// 从redis中删除用户喜欢缓存
func (r *favoriteRepo) delUserFavoriteListCache(ctx context.Context, userId int64) error {
	err := r.data.redis.Del(ctx, constants.UserFavoriteListCacheKey(userId)).Err()
	if err != nil {
		r.log.Errorf("delUserFavoriteListCache error(%v)", err)
		return err
	}
	return nil
}
