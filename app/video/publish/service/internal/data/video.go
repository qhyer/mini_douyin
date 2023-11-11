package data

import (
	"bytes"
	"context"
	do "douyin/app/video/publish/common/entity"
	"douyin/app/video/publish/common/mapper"
	po "douyin/app/video/publish/common/model"
	"douyin/app/video/publish/service/internal/biz"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"time"
)

var videoCacheKey = func(vid int64) string {
	return fmt.Sprintf("VIDEO_INFO_%d", vid)
}

var userPublishedVidListCacheKey = func(uid int64) string {
	return fmt.Sprintf("USER_PUB_VID_LIST_%d", uid)
}

var userPublishedVidCountCacheKey = func(uid int64) string {
	return fmt.Sprintf("USER_PUB_VID_COUNT_%d", uid)
}

var videoCacheExpiration = 3 * time.Minute

var publishTableName = "publish"

var videoBucketName = "video"

type VideoRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoRepo(data *Data, logger log.Logger) biz.VideoRepo {
	return &VideoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *VideoRepo) GetPublishedVideosByUserId(ctx context.Context, userId int64, offset int, limit int) ([]*do.Video, error) {
	vids, err := r.getUserPublishedVidListFromCache(ctx, userId, 0, 0)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		vs := make([]*po.Video, 0)
		if err := r.data.db.WithContext(ctx).Table(publishTableName).Where("user_id = ?", userId).Offset(offset).Limit(limit).Find(&vs).Error; err != nil {
			return nil, err
		}
		if err := r.setUserPublishedVidListCache(ctx, userId, vs); err != nil {
			return nil, err
		}
		vids = make([]int64, 0, len(vs))
		for _, v := range vs {
			vids = append(vids, v.ID)
		}
	}
	videos, err := r.MGetVideoByIds(ctx, vids)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *VideoRepo) GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error) {
	vids := make([]int64, limit)
	if err := r.data.db.WithContext(ctx).Table(publishTableName).Where("created_at < ?", time.Unix(latestTime, 0)).Order("created_at desc").Limit(limit).Pluck("id", &vids).Error; err != nil {
		return nil, err
	}
	videos, err := r.MGetVideoByIds(ctx, vids)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *VideoRepo) GetVideoById(ctx context.Context, id int64) (*do.Video, error) {
	video, err := r.getVideoFromCache(ctx, id)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		video = &po.Video{}
		if err := r.data.db.WithContext(ctx).Table(publishTableName).Where("id = ?", id).First(video).Error; err != nil {
			return nil, err
		}
		if err := r.setVideoCache(ctx, video); err != nil {
			return nil, err
		}
	}
	vs, err := mapper.VideoFromPO(video)
	if err != nil {
		return nil, err
	}
	return vs, nil
}

func (r *VideoRepo) MGetVideoByIds(ctx context.Context, ids []int64) ([]*do.Video, error) {
	videos, missed, err := r.batchGetVideoFromCache(ctx, ids)
	if err != nil {
		log.Errorf("redis error: %v", err)
	}
	if len(missed) > 0 {
		missedVideos := make([]*po.Video, 0, len(missed))
		if err := r.data.db.WithContext(ctx).Table(publishTableName).Where("id in (?)", missed).Find(&missedVideos).Error; err != nil {
			return nil, err
		}
		if err := r.batchSetVideoCache(ctx, missedVideos); err != nil {
			return nil, err
		}
		videos = append(videos, missedVideos...)
	}
	result := make([]*do.Video, 0, len(videos))
	for _, video := range videos {
		vs, err := mapper.VideoFromPO(video)
		if err != nil {
			return nil, err
		}
		result = append(result, vs)
	}
	return result, nil
}

func (r *VideoRepo) CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err := r.getUserPublishedVidCountFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		var count int64
		if err := r.data.db.WithContext(ctx).Table(publishTableName).Where("user_id = ?", userId).Count(&count).Error; err != nil {
			return 0, err
		}
		if err := r.setUserPublishedVidCountCache(ctx, userId, count); err != nil {
			return 0, err
		}
		return count, nil
	}
	return res, nil
}

func (r *VideoRepo) setVideoCache(ctx context.Context, video *po.Video) error {
	return r.data.redis.Set(ctx, videoCacheKey(video.ID), video, videoCacheExpiration).Err()
}

func (r *VideoRepo) batchSetVideoCache(ctx context.Context, videos []*po.Video) error {
	pipe := r.data.redis.Pipeline()
	for _, video := range videos {
		pipe.Set(ctx, videoCacheKey(video.ID), video, videoCacheExpiration)
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (r *VideoRepo) getVideoFromCache(ctx context.Context, vid int64) (*po.Video, error) {
	video := &po.Video{}
	err := r.data.redis.Get(ctx, videoCacheKey(vid)).Scan(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (r *VideoRepo) batchGetVideoFromCache(ctx context.Context, vids []int64) (videos []*po.Video, missed []int64, err error) {
	pipe := r.data.redis.Pipeline()
	for _, vid := range vids {
		pipe.Get(ctx, videoCacheKey(vid))
	}
	results, err := pipe.Exec(ctx)
	if err != nil {
		return nil, vids, err
	}
	videos = make([]*po.Video, 0, len(vids))
	missed = make([]int64, 0, len(vids))
	for i, result := range results {
		if result.Err() == redis.Nil {
			missed = append(missed, vids[i])
			continue
		}
		video := &po.Video{}
		if err := result.(*redis.StringCmd).Scan(video); err != nil {
			return nil, vids, err
		}
		videos = append(videos, video)
	}
	return videos, missed, nil
}

func (r *VideoRepo) setUserPublishedVidListCache(ctx context.Context, uid int64, videos []*po.Video) error {
	vids := make([]redis.Z, 0, len(videos))
	for _, video := range videos {
		vids = append(vids, redis.Z{
			Score:  float64(video.CreatedAt.Unix()),
			Member: video.ID,
		})
	}
	return r.data.redis.ZAdd(ctx, userPublishedVidListCacheKey(uid), vids...).Err()
}

func (r *VideoRepo) getUserPublishedVidListFromCache(ctx context.Context, uid int64, offset int, limit int) ([]int64, error) {
	data, err := r.data.redis.ZRevRangeWithScores(ctx, userPublishedVidListCacheKey(uid), int64(offset), int64(offset+limit-1)).Result()
	if err != nil {
		return nil, err
	}
	vids := make([]int64, 0, len(data))
	for _, item := range data {
		vids = append(vids, item.Member.(int64))
	}
	return vids, nil
}

func (r *VideoRepo) getUserPublishedVidCountFromCache(ctx context.Context, uid int64) (int64, error) {
	return r.data.redis.Get(ctx, userPublishedVidCountCacheKey(uid)).Int64()
}

func (r *VideoRepo) setUserPublishedVidCountCache(ctx context.Context, uid int64, count int64) error {
	return r.data.redis.Set(ctx, userPublishedVidCountCacheKey(uid), count, videoCacheExpiration).Err()
}

func (r *VideoRepo) UploadVideo(ctx context.Context, data []byte, objectName string) error {
	_, err := r.putObject(ctx, videoBucketName, objectName, data)
	return err
}

func (r *VideoRepo) putObject(ctx context.Context, bucketName, objectName string, data []byte) (info minio.UploadInfo, err error) {
	reader := bytes.NewReader(data)
	return r.data.minio.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
}
