package data

import (
	"context"
	po "douyin/app/video/favorite/common/model"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"strconv"
	"time"
)

var favoriteTableName = "favorite"

var userFavoriteListCacheKey = func(uid int64) string {
	return fmt.Sprintf("USER_FAV_LIST:%d", uid)
}

var userFavoriteListCacheExpiration = 3 * time.Minute

type FavoriteRepo struct {
	data *Data
	log  *log.Helper
}

func NewFavoriteRepo(data *Data, logger log.Logger) *FavoriteRepo {
	return &FavoriteRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *FavoriteRepo) FavoriteVideo(ctx context.Context, userId int64, videoId int64) error {
	// TODO send favorite to mq
	return nil
}

func (r *FavoriteRepo) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	vids, err := r.getUserFavoriteVideoIdListCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		favs := make([]*po.Favorite, 0)
		if err := r.data.db.WithContext(ctx).Table(favoriteTableName).Where("user_id = ?", userId).Order("updated_at desc").Find(&favs).Error; err != nil {
			return nil, err
		}
		if err := r.setUserFavoriteVideoIdListCache(ctx, userId, favs); err != nil {
			return nil, err
		}
		vids = make([]int64, 0, len(favs))
		for _, v := range favs {
			vids = append(vids, v.VideoId)
		}
	}
	return vids, nil
}

func (r *FavoriteRepo) IsUserFavoriteVideo(ctx context.Context, userId int64, videoId int64) (bool, error) {
	res := r.data.redis.ZScore(ctx, userFavoriteListCacheKey(userId), strconv.FormatInt(videoId, 10))
	if err := res.Err(); err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		var count int64
		if err := r.data.db.WithContext(ctx).Table(favoriteTableName).Where("user_id = ? and video_id = ?", userId, videoId).Count(&count).Error; err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return res.Val() > 0, nil
}

func (r *FavoriteRepo) setUserFavoriteVideoIdListCache(ctx context.Context, userId int64, favs []*po.Favorite) error {
	pipe := r.data.redis.TxPipeline()
	key := userFavoriteListCacheKey(userId)
	for _, v := range favs {
		pipe.ZAdd(ctx, key, redis.Z{
			Score:  float64(v.UpdatedAt.UnixMilli()),
			Member: v.VideoId,
		})
	}
	pipe.Expire(ctx, key, userFavoriteListCacheExpiration)
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *FavoriteRepo) getUserFavoriteVideoIdListCache(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRangeByScore(ctx, userFavoriteListCacheKey(userId), &redis.ZRangeBy{
		Min: "0",
		Max: fmt.Sprintf("%d", time.Now().UnixMilli()),
	}).Result()
	if err != nil {
		return nil, err
	}
	vids := make([]int64, 0, len(res))
	for _, v := range res {
		vid := cast.ToInt64(v)
		vids = append(vids, vid)
	}
	return vids, nil
}
