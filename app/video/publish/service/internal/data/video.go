package data

import (
	"bytes"
	"context"
	"douyin/app/video/publish/common/constants"
	do "douyin/app/video/publish/common/entity"
	"douyin/app/video/publish/common/mapper"
	po "douyin/app/video/publish/common/model"
	"douyin/app/video/publish/service/internal/biz"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"time"
)

type videoRepo struct {
	data *Data
	log  *log.Helper
}

func NewVideoRepo(data *Data, logger log.Logger) biz.VideoRepo {
	return &videoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// PublishVideo 发布视频
func (r *videoRepo) PublishVideo(ctx context.Context, video *do.Video) error {
	// TODO get seq-number
	b, err := video.MarshalJson()
	if err != nil {
		r.log.Errorf("json marshal error: %v", err)
		return err
	}
	msgs := make([]*sarama.ProducerMessage, 0, 2)
	msgs = append(msgs, &sarama.ProducerMessage{
		Topic: constants.PublishVideoTopic,
		Value: sarama.ByteEncoder(b),
	})
	msgs = append(msgs, &sarama.ProducerMessage{
		Topic: constants.GenCoverTopic,
		Value: sarama.ByteEncoder(b),
	})
	err = r.data.kafka.SendMessages(msgs)
	if err != nil {
		r.log.Errorf("kafka send message error: %v", err)
		return err
	}
	return nil
}

// GetPublishedVideosByUserId 获取用户发布视频列表
func (r *videoRepo) GetPublishedVideosByUserId(ctx context.Context, userId int64) ([]*do.Video, error) {
	vids, err := r.getUserPublishedVidListFromCache(ctx, userId, 0, 0)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		vs := make([]*po.Video, 0)
		if err := r.data.db.WithContext(ctx).Table(constants.PublishRecordTableName).Where("user_id = ?", userId).Find(&vs).Error; err != nil {
			r.log.Errorf("db error: %v", err)
			return nil, err
		}
		err := r.data.cacheFan.Do(ctx, func(ctx context.Context) {
			r.setUserPublishedVidListCache(ctx, userId, vs)
		})
		if err != nil {
			r.log.Errorf("cache fanout error: %v", err)
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

// GetPublishedVideosByLatestTime 获取小于某个时间的视频列表
func (r *videoRepo) GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error) {
	vids := make([]int64, limit)
	if err := r.data.db.WithContext(ctx).Table(constants.PublishRecordTableName).Where("created_at < ?", time.Unix(latestTime, 0)).Order("created_at desc").Limit(limit).Pluck("id", &vids).Error; err != nil {
		r.log.Errorf("db error: %v", err)
		return nil, err
	}
	videos, err := r.MGetVideoByIds(ctx, vids)
	if err != nil {
		r.log.Errorf("MGetVideoByIds err: %v", err)
		return nil, err
	}
	return videos, nil
}

// GetVideoById 获取视频信息
func (r *videoRepo) GetVideoById(ctx context.Context, id int64) (*do.Video, error) {
	video, err := r.getVideoFromCache(ctx, id)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		video = &po.Video{}
		if err := r.data.db.WithContext(ctx).Table(constants.PublishRecordTableName).Where("id = ?", id).First(video).Error; err != nil {
			r.log.Errorf("db error: %v", err)
			return nil, err
		}
		err := r.data.cacheFan.Do(ctx, func(ctx context.Context) {
			r.setVideoCache(ctx, video)
		})
		if err != nil {
			r.log.Errorf("cache fanout error: %v", err)
		}
	}
	vs, err := mapper.VideoFromPO(video)
	if err != nil {
		r.log.Errorf("mapper video from po error: %v", err)
		return nil, err
	}
	return vs, nil
}

// MGetVideoByIds 批量获取视频信息
func (r *videoRepo) MGetVideoByIds(ctx context.Context, ids []int64) ([]*do.Video, error) {
	videos, missed, err := r.batchGetVideoFromCache(ctx, ids)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
	if len(missed) > 0 {
		missedVideos := make([]*po.Video, 0, len(missed))
		if err := r.data.db.WithContext(ctx).Table(constants.PublishRecordTableName).Where("id in (?)", missed).Find(&missedVideos).Error; err != nil {
			r.log.Errorf("db error: %v", err)
			return nil, err
		}
		r.batchSetVideoCache(ctx, missedVideos)
		videos = append(videos, missedVideos...)
	}
	result, err := mapper.VideoFromPOs(videos)
	if err != nil {
		r.log.Errorf("mapper video from po error: %v", err)
		return nil, err
	}
	return result, nil
}

// CountUserPublishedVideoByUserId 获取用户发布视频数量
func (r *videoRepo) CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err := r.getUserPublishedVidCountFromCache(ctx, userId)
	if err == nil {
		return res, nil
	}
	if err != redis.Nil {
		log.Errorf("redis error: %v", err)
	}
	var count int64
	if err := r.data.db.WithContext(ctx).Table(constants.PublishCountTableName(userId)).Where("user_id = ?", userId).Pluck("count", &count).Error; err != nil {
		r.log.Errorf("db error: %v", err)
		return 0, err
	}
	err = r.data.cacheFan.Do(ctx, func(ctx context.Context) {
		r.setUserPublishedVidCountCache(ctx, userId, count)
	})
	if err != nil {
		r.log.Errorf("cache fanout error: %v", err)
	}
	return count, nil
}

// setVideoCache 设置视频信息缓存
func (r *videoRepo) setVideoCache(ctx context.Context, video *po.Video) {
	err := r.data.redis.Set(ctx, constants.VideoCacheKey(video.ID), video, constants.VideoCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// batchGetVideoFromCache 批量获取视频信息缓存
func (r *videoRepo) batchSetVideoCache(ctx context.Context, videos []*po.Video) {
	pipe := r.data.redis.Pipeline()
	for _, video := range videos {
		pipe.Set(ctx, constants.VideoCacheKey(video.ID), video, constants.VideoCacheExpiration)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// getVideoFromCache 从redis获取视频信息
func (r *videoRepo) getVideoFromCache(ctx context.Context, vid int64) (*po.Video, error) {
	video := &po.Video{}
	err := r.data.redis.Get(ctx, constants.VideoCacheKey(vid)).Scan(video)
	return video, err
}

// batchGetVideoFromCache 批量从redis获取视频信息
func (r *videoRepo) batchGetVideoFromCache(ctx context.Context, vids []int64) (videos []*po.Video, missed []int64, err error) {
	pipe := r.data.redis.Pipeline()
	for _, vid := range vids {
		pipe.Get(ctx, constants.VideoCacheKey(vid))
	}
	results, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
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
			missed = append(missed, vids[i])
			continue
		}
		videos = append(videos, video)
	}
	return videos, missed, nil
}

// setUserPublishedVidListCache 设置用户发布视频id列表缓存
func (r *videoRepo) setUserPublishedVidListCache(ctx context.Context, uid int64, videos []*po.Video) {
	vids := make([]redis.Z, 0, len(videos))
	for _, video := range videos {
		vids = append(vids, redis.Z{
			Score:  float64(video.CreatedAt.Unix()),
			Member: video.ID,
		})
	}
	err := r.data.redis.ZAdd(ctx, constants.UserPublishedVidListCacheKey(uid), vids...).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// getUserPublishedVidListFromCache 从redis获取用户发布视频id列表
func (r *videoRepo) getUserPublishedVidListFromCache(ctx context.Context, uid int64, offset int, limit int) ([]int64, error) {
	data, err := r.data.redis.ZRevRangeWithScores(ctx, constants.UserPublishedVidListCacheKey(uid), int64(offset), int64(offset+limit-1)).Result()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
		return nil, err
	}
	vids := make([]int64, 0, len(data))
	for _, item := range data {
		vids = append(vids, item.Member.(int64))
	}
	return vids, nil
}

// getUserPublishedVidCountFromCache 从redis获取用户发布视频数量
func (r *videoRepo) getUserPublishedVidCountFromCache(ctx context.Context, uid int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserPublishedVidCountCacheKey(uid)).Int64()
}

// setUserPublishedVidCountCache 设置用户发布视频数量缓存
func (r *videoRepo) setUserPublishedVidCountCache(ctx context.Context, uid int64, count int64) {
	err := r.data.redis.Set(ctx, constants.UserPublishedVidCountCacheKey(uid), count, constants.VideoCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// UploadVideo minio上传视频
func (r *videoRepo) UploadVideo(ctx context.Context, data []byte, objectName string) (string, error) {
	_, err := r.putObject(ctx, constants.VideoBucketName, objectName, data)
	return objectName, err
}

// minio上传文件
func (r *videoRepo) putObject(ctx context.Context, bucketName, objectName string, data []byte) (info minio.UploadInfo, err error) {
	reader := bytes.NewReader(data)
	return r.data.minio.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
}
