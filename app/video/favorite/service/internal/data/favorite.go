package data

import (
	"context"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"

	"douyin/app/video/favorite/common/constants"
	do "douyin/app/video/favorite/common/entity"
	po "douyin/app/video/favorite/common/model"
	"douyin/app/video/favorite/service/internal/biz"
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

// CountVideoFavoritedByVideoId 获取视频获赞数
func (r *favoriteRepo) CountVideoFavoritedByVideoId(ctx context.Context, videoId int64) (int64, error) {
	count, err := r.getVideoFavoritedCountFromCache(ctx, videoId)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.VideoFavoritedCountTableName(videoId)).Where("video_id = ?", videoId).Pluck("favd_count", &count).Error; err != nil {
			r.log.Errorf("db error: %v", err)
			return 0, err
		}
		err := r.data.cacheFan.Do(ctx, func(ctx context.Context) {
			r.setVideoFavoritedCountCache(ctx, videoId, count)
		})
		if err != nil {
			r.log.Errorf("fanout error: %v", err)
		}
	}
	return count, nil
}

// IsUserFavoriteVideoList 批量获取用户是否点赞视频
func (r *favoriteRepo) IsUserFavoriteVideoList(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	favs, err := r.GetFavoriteVideoIdListByUserId(ctx, userId)
	if err != nil {
		r.log.Errorf("get user favorite video id list error: %v", err)
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

// FavoriteVideo 点赞视频
func (r *favoriteRepo) FavoriteVideo(ctx context.Context, fav *do.FavoriteAction) error {
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

// GetFavoriteVideoIdListByUserId 获取用户点赞视频id列表
func (r *favoriteRepo) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	vids, err := r.getUserFavoriteVideoIdListFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		favs := make([]*po.Favorite, 0)
		if err := r.data.db.WithContext(ctx).Table(constants.FavoriteVideoRecordTableName(userId)).Where("user_id = ?", userId).Order("updated_at desc").Find(&favs).Error; err != nil {
			r.log.Errorf("db error: %v", err)
			return nil, err
		}
		err := r.data.cacheFan.Do(ctx, func(ctx context.Context) {
			r.setUserFavoriteVideoIdListCache(ctx, userId, favs)
		})
		if err != nil {
			r.log.Errorf("fanout error: %v", err)
		}
		vids = make([]int64, 0, len(favs))
		for _, v := range favs {
			vids = append(vids, v.VideoId)
		}
	}
	return vids, nil
}

// IsUserFavoriteVideo 判断用户是否点赞视频
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

// CountUserFavoriteByUserId 获取用户点赞数
func (r *favoriteRepo) CountUserFavoriteByUserId(ctx context.Context, userId int64) (int64, error) {
	count, err := r.getUserFavoriteCountFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).Where("user_id = ?", userId).Pluck("fav_count", &count).Error; err != nil {
			return 0, err
		}
		err := r.data.cacheFan.Do(ctx, func(ctx context.Context) {
			r.setUserFavoriteCountCache(ctx, userId, count)
		})
		if err != nil {
			r.log.Errorf("fanout error: %v", err)
		}
	}
	return count, nil
}

// CountUserFavoritedByUserId 获取用户获赞数
func (r *favoriteRepo) CountUserFavoritedByUserId(ctx context.Context, userId int64) (int64, error) {
	count, err := r.getUserFavoritedCountFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.UserFavoriteVideoCountTableName(userId)).Where("user_id = ?", userId).Pluck("favd_count", &count).Error; err != nil {
			return 0, err
		}
		err := r.data.cacheFan.Do(ctx, func(ctx context.Context) {
			r.setUserFavoritedCountCache(ctx, userId, count)
		})
		if err != nil {
			r.log.Errorf("fanout error: %v", err)
		}
	}
	return count, nil
}

// MCountVideoFavoritedByVideoId 批量获取视频获赞数
func (r *favoriteRepo) MCountVideoFavoritedByVideoId(ctx context.Context, videoId []int64) ([]int64, error) {
	favCntMap, missed := r.batchGetVideoFavoritedCountFromCache(ctx, videoId)
	if len(missed) > 0 {
		for _, v := range missed {
			res, err := r.CountVideoFavoritedByVideoId(ctx, v)
			// 如果没查到，设0
			if err == nil {
				favCntMap[v] = res
			}
		}
	}
	res := make([]int64, 0, len(videoId))
	for _, v := range videoId {
		res = append(res, favCntMap[v])
	}
	return res, nil
}

// 设置用户点赞视频列表到缓存
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

// 从缓存中获取用户点赞视频列表
func (r *favoriteRepo) getUserFavoriteVideoIdListFromCache(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRange(ctx, constants.UserFavoriteListCacheKey(userId), 0, -1).Result()
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

// 设置用户获赞数到缓存
func (r *favoriteRepo) setUserFavoritedCountCache(ctx context.Context, userId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.UserFavoritedCountCacheKey(userId), count, constants.UserFavoritedCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 从redis获取用户获赞数
func (r *favoriteRepo) getUserFavoritedCountFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserFavoritedCountCacheKey(userId)).Int64()
}

// 设置用户点赞数到缓存
func (r *favoriteRepo) setUserFavoriteCountCache(ctx context.Context, userId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.UserFavoriteCountCacheKey(userId), count, constants.UserFavoriteCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 从缓存中获取用户点赞数
func (r *favoriteRepo) getUserFavoriteCountFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserFavoriteCountCacheKey(userId)).Int64()
}

// 设置视频获赞数到缓存
func (r *favoriteRepo) setVideoFavoritedCountCache(ctx context.Context, videoId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.VideoFavoritedCountCacheKey(videoId), count, constants.VideoFavoritedCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 从redis中获取视频获赞数
func (r *favoriteRepo) getVideoFavoritedCountFromCache(ctx context.Context, videoId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.VideoFavoritedCountCacheKey(videoId)).Int64()
}

// 批量设置视频获赞数到缓存
func (r *favoriteRepo) batchSetVideoFavoritedCountCache(ctx context.Context, videoId []int64, count []int64) {
	pipe := r.data.redis.TxPipeline()
	for i, v := range videoId {
		pipe.Set(ctx, constants.VideoFavoritedCountCacheKey(v), count[i], constants.VideoFavoritedCountCacheExpiration)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 批量从redis中获取视频获赞数
func (r *favoriteRepo) batchGetVideoFavoritedCountFromCache(ctx context.Context, videoId []int64) (favCntMap map[int64]int64, missed []int64) {
	favCntMap = make(map[int64]int64, len(videoId))
	missed = make([]int64, 0, len(videoId))
	pipe := r.data.redis.TxPipeline()
	for _, v := range videoId {
		pipe.Get(ctx, constants.VideoFavoritedCountCacheKey(v))
	}
	res, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
		return favCntMap, videoId
	}
	for i, v := range res {
		if v.Err() != nil {
			if v.Err() != redis.Nil {
				r.log.Errorf("redis error: %v", v.Err())
			}
			missed = append(missed, videoId[i])
		} else {
			favCntMap[videoId[i]], _ = v.(*redis.StringCmd).Int64()
		}
	}
	return favCntMap, missed
}
