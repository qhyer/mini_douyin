package data

import (
	"context"
	"douyin/app/video/comment/common/constants"
	do "douyin/app/video/comment/common/entity"
	"douyin/app/video/comment/common/mapper"
	po "douyin/app/video/comment/common/model"
	"douyin/app/video/comment/service/internal/biz"
	"encoding/json"
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

func (r *commentRepo) CommentAction(ctx context.Context, comment *do.Comment) error {
	// TODO get seq-num
	b, err := json.Marshal(comment)
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

func (r *commentRepo) GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*do.Comment, error) {
	cacheComments, err := r.getCommentIdListFromCache(ctx, videoId)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("GetCommentListByVideoId err:%v", err)
		}
		dbComments := make([]*po.Comment, 0, constants.CommentQueryLimit)
		if err := r.data.db.WithContext(ctx).Table(constants.CommentRecordTableName(videoId)).Where("video_id = ?", videoId).Limit(constants.CommentQueryLimit).Find(&dbComments).Error; err != nil {
			r.log.Errorf("GetCommentListByVideoId err:%v", err)
			return nil, err
		}
		r.setCommentIdListCache(ctx, videoId, dbComments)
		r.batchSetCommentInfoCache(ctx, dbComments)
		return mapper.CommentFromPOs(dbComments)
	}
	return r.batchGetCommentInfoByVideoIdAndCommentIds(ctx, videoId, cacheComments)
}

func (r *commentRepo) batchGetCommentInfoByVideoIdAndCommentIds(ctx context.Context, videoId int64, commentIds []int64) ([]*do.Comment, error) {
	cacheComments, missed, err := r.batchGetCommentInfoFromCache(ctx, commentIds)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("batchGetCommentInfoByVideoIdAndCommentIds err:%v", err)
		}
		var comments []*po.Comment
		if err := r.data.db.WithContext(ctx).Table(constants.CommentRecordTableName(videoId)).Where("video_id = ? and id in (?)", videoId, missed).Find(&comments).Error; err != nil {
			r.log.Errorf("batchGetCommentInfoByVideoIdAndCommentIds err:%v", err)
			return nil, err
		}
		r.batchSetCommentInfoCache(ctx, comments)
		cacheComments = append(cacheComments, comments...)
	}
	comments, err := mapper.CommentFromPOs(cacheComments)
	return comments, err

}

func (r *commentRepo) CountCommentByVideoId(ctx context.Context, videoId int64) (int64, error) {
	res, err := r.getCommentCountFromCache(ctx, videoId)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("CountCommentByVideoId err:%v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.CommentCountTableName(videoId)).Where("video_id = ?", videoId).Pluck("comment_count", &res).Error; err != nil {
			r.log.Errorf("CountCommentByVideoId err:%v", err)
			return 0, err
		}
		r.setCommentCountCache(ctx, videoId, res)
	}
	return res, nil
}

func (r *commentRepo) setCommentCountCache(ctx context.Context, videoId int64, count int64) {
	err := r.data.redis.Set(ctx, constants.CommentCountCacheKey(videoId), count, constants.CommentCountCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("setCommentCountCache err:%v", err)
	}
}

func (r *commentRepo) getCommentCountFromCache(ctx context.Context, videoId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.CommentCountCacheKey(videoId)).Int64()
}

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

func (r *commentRepo) getCommentIdListFromCache(ctx context.Context, videoId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRange(ctx, constants.CommentListCacheKey(videoId), 0, -1).Result()
	if err != nil {
		if err != redis.Nil {
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
