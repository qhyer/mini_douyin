package data

import (
	"context"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"douyin/app/video/comment/common/event"
	"douyin/common/ecode"

	"douyin/app/video/comment/common/constants"
	"douyin/app/video/comment/common/mapper"
	po "douyin/app/video/comment/common/model"
	"douyin/app/video/comment/job/internal/biz"
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

func (r *commentRepo) UpdateVideoCommentCount(ctx context.Context, videoId int64, incr int64) error {
	comCnt := po.CommentCount{VideoId: videoId}
	err := r.data.db.WithContext(ctx).Table(constants.CommentCountTableName(videoId)).FirstOrCreate(&comCnt, comCnt).Update("comment_count", gorm.Expr("comment_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateVideoCommentCount error(%v)", err)
		return err
	}
	return nil
}

func (r *commentRepo) BatchUpdateVideoCommentCount(ctx context.Context, stats map[int64]int64) error {
	tx := r.data.db.WithContext(ctx).Begin()
	for vid, cnt := range stats {
		comCnt := po.CommentCount{VideoId: vid}
		err := tx.Table(constants.CommentCountTableName(vid)).FirstOrCreate(&comCnt, comCnt).Update("comment_count", gorm.Expr("comment_count + ?", cnt)).Error
		if err != nil {
			r.log.Errorf("BatchUpdateVideoCommentCount error(%v)", err)
			tx.Rollback()
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		r.log.Errorf("BatchUpdateVideoCommentCount error(%v)", err)
		return err
	}
	return nil
}

func (r *commentRepo) CreateComment(ctx context.Context, commentAct *event.CommentAction) error {
	comment, err := mapper.ParseCommentFromCommentAction(commentAct)
	if err != nil {
		r.log.Errorf("mapper.ParseCommentFromCommentAction error(%v)", err)
		return err
	}
	com, err := mapper.CommentToPO(comment)
	if err != nil {
		r.log.Errorf("mapper.CommentToPO error(%v)", err)
		return err
	}

	delta := calcCommentStatDelta(commentAct)
	comStat := &event.CommentStat{
		VideoId: com.VideoId,
		Delta:   delta,
	}
	comStatByte, err := comStat.MarshalJson()
	if err != nil {
		r.log.Errorf("comStat.MarshalJson error(%v)", err)
		return err
	}

	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table(constants.CommentRecordTableName(com.VideoId)).Create(com).Error
		if err != nil {
			return err
		}
		_, _, err = r.data.kafkaProducer.SendMessage(&sarama.ProducerMessage{
			Topic: constants.UpdateCommentCountTopic,
			Key:   sarama.StringEncoder(constants.UpdateCommentCountKafkaKey(com.VideoId)),
			Value: sarama.ByteEncoder(comStatByte),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("CreateComment error(%v)", err)
		return err
	}
	return nil
}

func (r *commentRepo) BatchCreateComment(ctx context.Context, comments []*event.CommentAction) error {
	// TODO implement me
	panic("implement me")
}

func (r *commentRepo) DeleteComment(ctx context.Context, commentAct *event.CommentAction) error {
	comment, err := mapper.ParseCommentFromCommentAction(commentAct)
	if err != nil {
		r.log.Errorf("mapper.ParseCommentFromCommentAction error(%v)", err)
		return err
	}
	com, err := mapper.CommentToPO(comment)
	if err != nil {
		r.log.Errorf("mapper.CommentToPO error(%v)", err)
		return err
	}
	commentActByte, err := commentAct.MarshalJson()
	if err != nil {
		r.log.Errorf("commentAct.MarshalJson error(%v)", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Table(constants.CommentRecordTableName(com.VideoId)).Where("id = ?", com.ID).Delete(com)
		if err != nil {
			return err
		}

		if res.RowsAffected == 0 {
			return ecode.CommentRecordNotFoundErr
		}

		_, _, err = r.data.kafkaProducer.SendMessage(&sarama.ProducerMessage{
			Topic: constants.UpdateCommentCountTopic,
			Key:   sarama.StringEncoder(constants.UpdateCommentCountKafkaKey(com.VideoId)),
			Value: sarama.ByteEncoder(commentActByte),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("DeleteComment error(%v)", err)
		return err
	}
	return nil
}

func (r *commentRepo) BatchDeleteComment(ctx context.Context, comments []*event.CommentAction) error {
	// TODO implement me
	panic("implement me")
}

func (r *commentRepo) UpdateVideoCommentTempCountCache(ctx context.Context, procId int, stat *event.CommentStat) error {
	err := r.data.redis.HIncrBy(ctx, constants.CommentStatTempCacheKey(procId), strconv.FormatInt(stat.VideoId, 10), stat.Delta).Err()
	if err != nil {
		r.log.Errorf("redis.HIncrBy error(%v)", err)
		return err
	}
	return nil
}

func (r *commentRepo) GetVideoCommentTempCountFromCache(ctx context.Context, procId int) (map[int64]int64, error) {
	m, err := r.data.redis.HGetAll(ctx, constants.CommentStatTempCacheKey(procId)).Result()
	if err != nil {
		r.log.Errorf("redis.HGetAll error(%v)", err)
		return nil, err
	}
	ret := make(map[int64]int64, len(m))
	for k, v := range m {
		vid, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			r.log.Errorf("strconv.ParseInt error(%v)", err)
			return nil, err
		}
		cnt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			r.log.Errorf("strconv.ParseInt error(%v)", err)
			return nil, err
		}
		ret[vid] = cnt
	}
	return ret, nil
}

func (r *commentRepo) PurgeVideoCommentTempCountCache(ctx context.Context, procId int) error {
	err := r.data.redis.Del(ctx, constants.CommentStatTempCacheKey(procId)).Err()
	if err != nil {
		r.log.Errorf("redis.Del error(%v)", err)
		return err
	}
	return nil
}

func calcCommentStatDelta(e *event.CommentAction) int64 {
	if e.Type == event.CommentActionPublish {
		return 1
	} else if e.Type == event.CommentActionDelete {
		return -1
	}
	return 0
}
