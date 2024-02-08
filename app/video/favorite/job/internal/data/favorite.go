package data

import (
	"context"
	"strconv"
	"time"

	"douyin/app/video/favorite/common/event"
	"douyin/common/ecode"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"douyin/app/video/favorite/common/constants"
	"douyin/app/video/favorite/common/mapper"
	po "douyin/app/video/favorite/common/model"
	"douyin/app/video/favorite/job/internal/biz"
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
	err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).FirstOrCreate(&usFav, usFav).Update("favorite_count", gorm.Expr("favorite_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFavoriteCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) UpdateUserFavoritedCount(ctx context.Context, userId int64, incr int64) error {
	usFav := po.UserFavoriteCount{UserId: userId}
	err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).FirstOrCreate(&usFav, usFav).Update("favorited_count", gorm.Expr("favorited_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) UpdateVideoFavoritedCount(ctx context.Context, videoId int64, incr int64) error {
	videoFav := po.VideoFavoritedCount{VideoId: videoId}
	err := r.data.db.WithContext(ctx).Table(constants.VideoFavoritedCountTableName(videoId)).FirstOrCreate(&videoFav, videoFav).Update("favorited_count", gorm.Expr("favorited_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateVideoFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) CreateFavorite(ctx context.Context, favoriteAction *event.FavoriteAction) error {
	favorite, err := mapper.ParseFavoriteFromFavoriteAction(favoriteAction)
	if err != nil {
		r.log.Errorf("CreateFavorite error(%v)", err)
		return err
	}
	fav, err := mapper.FavoriteToPO(favorite)
	if err != nil {
		r.log.Errorf("CreateFavorite error(%v)", err)
		return err
	}
	// 删除用户喜欢缓存
	err = r.delUserFavoriteListCache(ctx, favoriteAction.UserId)
	if err != nil {
		r.log.Errorf("CreateFavorite error(%v)", err)
	}
	// 创建喜欢
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建点赞记录
		err := tx.Table(constants.FavoriteVideoRecordTableName(favoriteAction.UserId)).Create(fav).Error
		if err != nil {
			return err
		}

		// 更新用户点赞数
		err = tx.Table(constants.UserFavoriteVideoCountTableName(favoriteAction.UserId)).FirstOrCreate(&po.UserFavoriteCount{UserId: favoriteAction.UserId}, po.UserFavoriteCount{UserId: favoriteAction.UserId}).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			return err
		}

		// 异步更新视频获赞数和用户获赞数
		b, err := favoriteAction.MarshalJson()
		_, _, err = r.data.kafkaProducer.SendMessage(&sarama.ProducerMessage{
			Topic: constants.UpdateVideoFavoritedCountTopic,
			Key:   sarama.StringEncoder(constants.UpdateVideoFavoritedCountKafkaKey(favoriteAction.VideoId)),
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
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		time.Sleep(100 * time.Millisecond)
		err = r.delUserFavoriteListCache(ctx, favoriteAction.UserId)
		if err != nil {
			r.log.Errorf("DelUserFavoriteCache error(%v)", err)
		}
	})
	if err != nil {
		r.log.Errorf("Fanout error(%v)", err)
	}
	return nil
}

func (r *favoriteRepo) DeleteFavorite(ctx context.Context, favoriteAction *event.FavoriteAction) error {
	favorite, err := mapper.ParseFavoriteFromFavoriteAction(favoriteAction)
	if err != nil {
		r.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	fav, err := mapper.FavoriteToPO(favorite)
	if err != nil {
		r.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	// 删除用户喜欢缓存
	err = r.delUserFavoriteListCache(ctx, favoriteAction.UserId)
	if err != nil {
		r.log.Errorf("DeleteFavorite error(%v)", err)
	}
	// 删除喜欢
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除点赞记录
		res := tx.Table(constants.FavoriteVideoRecordTableName(favoriteAction.UserId)).Where("user_id = ? and video_id = ?", fav.UserId, fav.VideoId).Delete(fav)
		if err != nil {
			return err
		}

		if res.RowsAffected == 0 {
			return ecode.FavoriteRecordNotFoundErr
		}

		// 更新用户点赞数
		err := tx.Table(constants.UserFavoriteVideoCountTableName(favoriteAction.UserId)).FirstOrCreate(&po.UserFavoriteCount{UserId: favoriteAction.UserId}, po.UserFavoriteCount{UserId: favoriteAction.UserId}).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			return err
		}

		// 异步更新视频获赞数和用户获赞数
		delta := calcVideoFavoritedCountDelta(favoriteAction.Type)
		stat := &event.VideoFavoritedStat{
			VideoId: favoriteAction.VideoId,
			Delta:   delta,
		}
		b, err := stat.MarshalJson()
		if err != nil {
			return err
		}
		_, _, err = r.data.kafkaProducer.SendMessage(&sarama.ProducerMessage{
			Topic: constants.UpdateVideoFavoritedCountTopic,
			Key:   sarama.StringEncoder(constants.UpdateVideoFavoritedCountKafkaKey(favoriteAction.VideoId)),
			Value: sarama.ByteEncoder(b),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("DeleteFavorite error(%v)", err)
		return err
	}
	// 延时删除用户喜欢缓存
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		time.Sleep(100 * time.Millisecond)
		err = r.delUserFavoriteListCache(ctx, favoriteAction.UserId)
		if err != nil {
			r.log.Errorf("DelUserFavoriteCache error(%v)", err)
		}
	})
	if err != nil {
		r.log.Errorf("Fanout error(%v)", err)
	}
	return nil
}

func (r *favoriteRepo) BatchUpdateUserFavoriteCount(ctx context.Context, stats map[int64]int64) error {
	tx := r.data.db.WithContext(ctx).Begin()
	for userId, count := range stats {
		err := tx.Table(constants.UserFavoriteVideoCountTableName(userId)).FirstOrCreate(&po.UserFavoriteCount{UserId: userId}, po.UserFavoriteCount{UserId: userId}).Update("favorite_count", gorm.Expr("favorite_count + ?", count)).Error
		if err != nil {
			tx.Rollback()
			r.log.Errorf("BatchUpdateUserFavoriteCount error(%v)", err)
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		r.log.Errorf("BatchUpdateUserFavoriteCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) BatchUpdateUserFavoritedCount(ctx context.Context, stats map[int64]int64) error {
	tx := r.data.db.WithContext(ctx).Begin()
	for userId, count := range stats {
		err := tx.Table(constants.UserFavoriteVideoCountTableName(userId)).FirstOrCreate(&po.UserFavoriteCount{UserId: userId}, po.UserFavoriteCount{UserId: userId}).Update("favorited_count", gorm.Expr("favorited_count + ?", count)).Error
		if err != nil {
			tx.Rollback()
			r.log.Errorf("BatchUpdateUserFavoritedCount error(%v)", err)
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		r.log.Errorf("BatchUpdateUserFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) BatchUpdateVideoFavoritedCount(ctx context.Context, stats map[int64]int64) error {
	tx := r.data.db.WithContext(ctx).Begin()
	for videoId, count := range stats {
		err := tx.Table(constants.VideoFavoritedCountTableName(videoId)).FirstOrCreate(&po.VideoFavoritedCount{VideoId: videoId}, po.VideoFavoritedCount{VideoId: videoId}).Update("favorited_count", gorm.Expr("favorited_count + ?", count)).Error
		if err != nil {
			tx.Rollback()
			r.log.Errorf("BatchUpdateVideoFavoritedCount error(%v)", err)
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		r.log.Errorf("BatchUpdateVideoFavoritedCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) BatchCreateFavorite(ctx context.Context, favorites []*event.FavoriteAction) error {
	// TODO implement me
	panic("implement me")
}

func (r *favoriteRepo) BatchDeleteFavorite(ctx context.Context, favorites []*event.FavoriteAction) error {
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

func calcVideoFavoritedCountDelta(actionType event.FavoriteActionType) int64 {
	if actionType == event.FavoriteActionAdd {
		return 1
	} else if actionType == event.FavoriteActionDelete {
		return -1
	}
	return 0
}

func (r *favoriteRepo) UpdateVideoFavoritedTempCount(ctx context.Context, procId int, videoId int64, incr int64) error {
	err := r.data.redis.HIncrBy(ctx, constants.VideoFavoritedStatTempCacheKey(procId), strconv.FormatInt(videoId, 10), incr).Err()
	if err != nil {
		r.log.Errorf("UpdateVideoFavoritedTempCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) PurgeVideoFavoritedTempCount(ctx context.Context, procId int) error {
	err := r.data.redis.Del(ctx, constants.VideoFavoritedStatTempCacheKey(procId)).Err()
	if err != nil {
		r.log.Errorf("PurgeVideoFavoritedTempCount error(%v)", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) GetVideoFavoritedCountFromCache(ctx context.Context, procId int) (map[int64]int64, error) {
	m, err := r.data.redis.HGetAll(ctx, constants.VideoFavoritedStatTempCacheKey(procId)).Result()
	if err != nil {
		r.log.Errorf("GetVideoFavoritedCountFromCache error(%v)", err)
		return nil, err
	}
	ret := make(map[int64]int64, len(m))
	for k, v := range m {
		vid, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			r.log.Errorf("GetVideoFavoritedCountFromCache error(%v)", err)
			return nil, err
		}
		count, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			r.log.Errorf("GetVideoFavoritedCountFromCache error(%v)", err)
			return nil, err
		}
		ret[vid] = count
	}
	return ret, nil
}
