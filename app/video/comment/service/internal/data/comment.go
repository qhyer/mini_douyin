package data

import (
	"context"
	do "douyin/app/video/comment/common/entity"
	"douyin/app/video/comment/common/event"
	"encoding/json"
	"errors"

	"douyin/app/video/comment/common/constants"
	"douyin/app/video/comment/common/mapper"
	po "douyin/app/video/comment/common/model"
	"douyin/app/video/comment/service/internal/biz"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CommentAction 发布/删除评论
func (r *commentRepo) CommentAction(ctx context.Context, comment *event.CommentAction) error {
	b, err := comment.MarshalJson()
	if err != nil {
		r.log.Errorf("CommentAction err:%v", err)
		return err
	}
	_, _, err = r.data.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: constants.CommentActionTopic,
		Key:   sarama.StringEncoder(constants.CommentActionKafkaKey(comment.UserId)),
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		r.log.Errorf("CommentAction err:%v", err)
		return err
	}
	return nil
}

// GetCommentListByVideoId 获取视频的评论列表
func (r *commentRepo) GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*do.Comment, error) {
	cacheComments, err := r.getCommentIdListFromCache(ctx, videoId)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("GetCommentListByVideoId err:%v", err)
		}
		dbComments := make([]*po.Comment, 0, constants.CommentQueryLimit)
		if err := r.data.db.WithContext(ctx).Table(constants.CommentRecordTableName(videoId)).Where("video_id = ?", videoId).Limit(constants.CommentQueryLimit).Find(&dbComments).Error; err != nil {
			r.log.Errorf("GetCommentListByVideoId err:%v", err)
			return nil, err
		}
		err := r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
			r.setCommentIdListCache(ctx, videoId, dbComments)
			r.batchSetCommentInfoCache(ctx, dbComments)
		})
		if err != nil {
			r.log.Errorf("Fanout err:%v", err)
		}
		return mapper.CommentFromPOs(dbComments)
	}
	return r.batchGetCommentInfoByVideoIdAndCommentIds(ctx, videoId, cacheComments)
}

// 批量获取评论信息
func (r *commentRepo) batchGetCommentInfoByVideoIdAndCommentIds(ctx context.Context, videoId int64, commentIds []int64) ([]*do.Comment, error) {
	cacheComments, missed, err := r.batchGetCommentInfoFromCache(ctx, commentIds)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("batchGetCommentInfoByVideoIdAndCommentIds err:%v", err)
		}
		var comments []*po.Comment
		if err := r.data.db.WithContext(ctx).Table(constants.CommentRecordTableName(videoId)).Where("video_id = ? and id in (?)", videoId, missed).Find(&comments).Error; err != nil {
			r.log.Errorf("batchGetCommentInfoByVideoIdAndCommentIds err:%v", err)
			return nil, err
		}
		err := r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
			r.batchSetCommentInfoCache(ctx, comments)
		})
		if err != nil {
			r.log.Errorf("Fanout err:%v", err)
		}
		cacheComments = append(cacheComments, comments...)
	}
	comments, err := mapper.CommentFromPOs(cacheComments)
	return comments, err
}

// CountCommentByVideoId 获取视频的评论数
func (r *commentRepo) CountCommentByVideoId(ctx context.Context, videoId int64) (int64, error) {
	res, err := r.getCommentCountFromCache(ctx, videoId)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("CountCommentByVideoId err:%v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.CommentCountTableName(videoId)).Where("video_id = ?", videoId).Pluck("comment_count", &res).Error; err != nil {
			r.log.Errorf("CountCommentByVideoId err:%v", err)
			return 0, err
		}
		err := r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
			r.setCommentCountCache(ctx, videoId, res)
		})
		if err != nil {
			r.log.Errorf("Fanout err:%v", err)
		}
	}
	return res, nil
}

// 设置视频的评论数到缓存
func (r *commentRepo) setCommentCountCache(ctx context.Context, videoId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.CommentCountCacheKey(videoId), count, constants.CommentCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("setCommentCountCache err:%v", err)
	}
}

// 从缓存中获取视频的评论数
func (r *commentRepo) getCommentCountFromCache(ctx context.Context, videoId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.CommentCountCacheKey(videoId)).Int64()
}

// MCountCommentByVideoId 批量获取视频的评论数
func (r *commentRepo) MCountCommentByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	comCntMap, missed := r.batchGetCommentCountFromCache(ctx, videoIds)
	if len(missed) > 0 {
		for _, v := range missed {
			res, err := r.CountCommentByVideoId(ctx, v)
			// 如果没查到，设0
			if err == nil {
				comCntMap[v] = res
			}
		}
	}
	res := make([]int64, 0, len(videoIds))
	for _, videoId := range videoIds {
		res = append(res, comCntMap[videoId])
	}
	return res, nil
}

// 批量设置视频的评论数到缓存
func (r *commentRepo) batchSetCommentCountCache(ctx context.Context, videoIds []int64, counts []int64) {
	pipe := r.data.redis.Pipeline()
	for i, videoId := range videoIds {
		pipe.Set(ctx, constants.CommentCountCacheKey(videoId), counts[i], constants.CommentCountCacheExpiration)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("batchSetCommentCountCache err:%v", err)
	}
}

// 批量从缓存中获取视频的评论数
func (r *commentRepo) batchGetCommentCountFromCache(ctx context.Context, videoIds []int64) (comCntMap map[int64]int64, missed []int64) {
	comCntMap = make(map[int64]int64, len(videoIds))
	missed = make([]int64, 0, len(videoIds))
	pipe := r.data.redis.Pipeline()
	for _, videoId := range videoIds {
		pipe.Get(ctx, constants.CommentCountCacheKey(videoId))
	}
	results, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("batchGetCommentCountFromCache err:%v", err)
		return comCntMap, videoIds
	}
	for i, result := range results {
		if result.Err() != nil {
			missed = append(missed, videoIds[i])
			continue
		}
		comCntMap[videoIds[i]] = cast.ToInt64(result.(*redis.StringCmd).Val())
	}
	return comCntMap, missed
}

// 从缓存中获取评论信息
func (r *commentRepo) getCommentInfoFromCache(ctx context.Context, commentId int64) (*po.Comment, error) {
	bytes, err := r.data.redis.Get(ctx, constants.CommentInfoCacheKey(commentId)).Bytes()
	if err != nil {
		r.log.Errorf("getCommentInfoFromCache err:%v", err)
		return nil, err
	}
	comment := &po.Comment{}
	err = json.Unmarshal(bytes, comment)
	if err != nil {
		r.log.Errorf("getCommentInfoFromCache err:%v", err)
		return nil, err
	}
	return comment, nil
}

// 批量从缓存中获取评论信息
func (r *commentRepo) batchGetCommentInfoFromCache(ctx context.Context, commentIds []int64) (comments []*po.Comment, missed []int64, err error) {
	pipe := r.data.redis.Pipeline()
	for _, commentId := range commentIds {
		pipe.Get(ctx, constants.CommentInfoCacheKey(commentId))
	}
	results, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("batchGetCommentInfoFromCache err:%v", err)
		return nil, commentIds, err
	}
	comments = make([]*po.Comment, 0, len(commentIds))
	missed = make([]int64, 0, len(commentIds))
	for i, result := range results {
		if result.Err() != nil {
			missed = append(missed, commentIds[i])
			continue
		}
		comment := &po.Comment{}
		err := json.Unmarshal([]byte(result.(*redis.StringCmd).Val()), comment)
		if err != nil {
			missed = append(missed, commentIds[i])
			continue
		}
		comments = append(comments, comment)
	}
	return comments, missed, nil
}

// 批量设置评论信息到缓存
func (r *commentRepo) batchSetCommentInfoCache(ctx context.Context, comments []*po.Comment) {
	pipe := r.data.redis.Pipeline()
	for _, comment := range comments {
		bytes, err := json.Marshal(comment)
		if err != nil {
			r.log.Errorf("batchSetCommentInfoCache err:%v", err)
			continue
		}
		pipe.Set(ctx, constants.CommentInfoCacheKey(comment.ID), bytes, constants.CommentInfoCacheExpiration)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("batchSetCommentInfoCache err:%v", err)
	}
}

// 设置视频的评论id列表到缓存
func (r *commentRepo) setCommentIdListCache(ctx context.Context, videoId int64, comments []*po.Comment) {
	pipe := r.data.redis.Pipeline()
	for _, comment := range comments {
		pipe.ZAdd(ctx, constants.CommentListCacheKey(videoId), redis.Z{
			Score:  float64(comment.CreatedAt.Unix()),
			Member: comment.ID,
		})
	}
	pipe.Expire(ctx, constants.CommentListCacheKey(videoId), constants.CommentListCacheExpiration)
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("setCommentIdListCache err:%v", err)
	}
}

// 从缓存中获取评论id列表
func (r *commentRepo) getCommentIdListFromCache(ctx context.Context, videoId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRange(ctx, constants.CommentListCacheKey(videoId), 0, -1).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("getCommentListCache err:%v", err)
		}
		return nil, err
	}
	commentIds := make([]int64, 0, len(res))
	for _, id := range res {
		commentIds = append(commentIds, cast.ToInt64(id))
	}
	return commentIds, nil
}
