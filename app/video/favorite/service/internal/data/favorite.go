package data

import (
	"context"
	"douyin/app/video/favorite/common/constants"
	do "douyin/app/video/favorite/common/entity"
	po "douyin/app/video/favorite/common/model"
	"douyin/app/video/favorite/service/internal/biz"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"strconv"
	"time"
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

func (r *favoriteRepo) CountFavoritedByVideoId(ctx context.Context, videoId int64) (int64, error) {
	count, err := r.getVideoFavoritedCountFromCache(ctx, videoId)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.VideoFavoritedCountTableName(videoId)).Where("video_id = ?", videoId).Pluck("favd_count", &count).Error; err != nil {
			r.log.Errorf("db error: %v", err)
			return 0, err
		}
		r.setVideoFavoritedCountCache(ctx, videoId, count)
	}
	return count, nil
}

func (r *favoriteRepo) IsUserFavoriteVideoList(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	favs, err := r.GetFavoriteVideoIdListByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	favMap := make(map[int64]bool, len(favs))
	for _, v := range favs {
		favMap[v] = true
	}
	res := make([]bool, 0, len(videoIds))
	for _, v := range videoIds {
		res = append(res, favMap[v])
	}
	return res, nil
}

func (r *favoriteRepo) FavoriteVideo(ctx context.Context, fav *do.Favorite) error {
	// TODO get seq-number
	b, err := fav.MarshalJson()
	if err != nil {
		r.log.Errorf("json marshal error: %v", err)
		return err
	}
	_, _, err = r.data.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: constants.FavoriteVideoActionTopic,
		Key:   sarama.StringEncoder(constants.FavoriteActionKafkaKey(fav.UserId)),
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		r.log.Errorf("kafka error: %v", err)
		return err
	}
	return nil
}

func (r *favoriteRepo) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	vids, err := r.getUserFavoriteVideoIdListCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		favs := make([]*po.Favorite, 0)
		if err := r.data.db.WithContext(ctx).Table(constants.FavoriteVideoRecordTableName(userId)).Where("user_id = ?", userId).Order("updated_at desc").Find(&favs).Error; err != nil {
			r.log.Errorf("db error: %v", err)
			return nil, err
		}
		r.setUserFavoriteVideoIdListCache(ctx, userId, favs)
		vids = make([]int64, 0, len(favs))
		for _, v := range favs {
			vids = append(vids, v.VideoId)
		}
	}
	return vids, nil
}

func (r *favoriteRepo) IsUserFavoriteVideo(ctx context.Context, userId int64, videoId int64) (bool, error) {
	res := r.data.redis.ZScore(ctx, constants.UserFavoriteListCacheKey(userId), strconv.FormatInt(videoId, 10))
	if err := res.Err(); err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		var count int64
		if err := r.data.db.WithContext(ctx).Table(constants.FavoriteVideoRecordTableName(userId)).Where("user_id = ? and video_id = ?", userId, videoId).Count(&count).Error; err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return res.Val() > 0, nil
}

func (r *favoriteRepo) CountVideoFavoriteByUserId(ctx context.Context, userId int64) (int64, error) {
	count, err := r.getUserFavoriteCountFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).Where("user_id = ?", userId).Pluck("fav_count", &count).Error; err != nil {
			return 0, err
		}
		r.setUserFavoriteCountCache(ctx, userId, count)
	}
	return count, nil
}

func (r *favoriteRepo) CountVideoFavoritedByUserId(ctx context.Context, userId int64) (int64, error) {
	count, err := r.getUserFavoritedCountFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).Where("user_id = ?", userId).Pluck("favd_count", &count).Error; err != nil {
			return 0, err
		}
		r.setUserFavoritedCountCache(ctx, userId, count)
	}
	return count, nil
}

func (r *favoriteRepo) setUserFavoriteVideoIdListCache(ctx context.Context, userId int64, favs []*po.Favorite) {
	pipe := r.data.redis.TxPipeline()
	key := constants.UserFavoriteListCacheKey(userId)
	for _, v := range favs {
		pipe.ZAdd(ctx, key, redis.Z{
			Score:  float64(v.UpdatedAt.UnixMilli()),
			Member: v.VideoId,
		})
	}
	pipe.Expire(ctx, key, constants.UserFavoriteListCacheExpiration)
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

func (r *favoriteRepo) getUserFavoriteVideoIdListCache(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRangeByScore(ctx, constants.UserFavoriteListCacheKey(userId), &redis.ZRangeBy{
		Min: "0",
		Max: fmt.Sprintf("%d", time.Now().UnixMilli()),
	}).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		return nil, err
	}
	vids := make([]int64, 0, len(res))
	for _, v := range res {
		vid := cast.ToInt64(v)
		vids = append(vids, vid)
	}
	return vids, nil
}

func (r *favoriteRepo) setUserFavoritedCountCache(ctx context.Context, userId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.UserVideoFavoritedCountCacheKey(userId), count, constants.UserVideoFavoritedCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

func (r *favoriteRepo) getUserFavoritedCountFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserVideoFavoritedCountCacheKey(userId)).Int64()
}

func (r *favoriteRepo) setUserFavoriteCountCache(ctx context.Context, userId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.UserVideoFavoriteCountCacheKey(userId), count, constants.UserVideoFavoriteCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

func (r *favoriteRepo) getUserFavoriteCountFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserVideoFavoriteCountCacheKey(userId)).Int64()
}

func (r *favoriteRepo) setVideoFavoritedCountCache(ctx context.Context, videoId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.VideoFavoritedCountCacheKey(videoId), count, constants.VideoFavoritedCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

func (r *favoriteRepo) getVideoFavoritedCountFromCache(ctx context.Context, videoId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.VideoFavoritedCountCacheKey(videoId)).Int64()
}
