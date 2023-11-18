package data

import (
	"context"
	"douyin/app/user/relation/common/constants"
	do "douyin/app/user/relation/common/entity"
	"douyin/app/user/relation/service/internal/biz"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type relationRepo struct {
	data *Data
	log  *log.Helper
}

func NewRelationRepo(data *Data, logger log.Logger) biz.RelationRepo {
	return &relationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *relationRepo) RelationAction(ctx context.Context, relation *do.Relation) error {
	// TODO get seq-num
	b, err := json.Marshal(relation)
	if err != nil {
		r.log.Errorf("json marshal error: %v", err)
		return err
	}
	_, _, err = r.data.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: constants.RelationActionTopic,
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		r.log.Errorf("send kafka error: %v", err)
		return err
	}
	return nil
}

func (r *relationRepo) GetFollowListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepo) GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepo) GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepo) CountFollowByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err := r.getFollowCountByUserIdFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).Where("user_id = ?", userId).Pluck("follow_count", &res).Error; err != nil {
			return 0, err
		}
		r.setFollowCountByUserId(ctx, userId, res)
	}
	return res, nil
}

func (r *relationRepo) CountFollowerByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err := r.getFollowerCountByUserIdFromCache(ctx, userId)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("redis error: %v", err)
		}
		if err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).Where("user_id = ?", userId).Pluck("follower_count", &res).Error; err != nil {
			return 0, err
		}
		r.setFollowerCountByUserId(ctx, userId, res)
	}
	return res, nil
}

func (r *relationRepo) setFollowCountByUserId(ctx context.Context, userId int64, count int64) {
	r.data.redis.Set(ctx, constants.UserFollowCountCacheKey(userId), count, constants.UserFollowCountCacheExpiration)
}

func (r *relationRepo) setFollowerCountByUserId(ctx context.Context, userId int64, count int64) {
	r.data.redis.Set(ctx, constants.UserFollowerCountCacheKey(userId), count, constants.UserFollowerCountCacheExpiration)
}

func (r *relationRepo) getFollowCountByUserIdFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserFollowCountCacheKey(userId)).Int64()
}

func (r *relationRepo) getFollowerCountByUserIdFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserFollowerCountCacheKey(userId)).Int64()
}
