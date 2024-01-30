package data

import (
	"context"
	"douyin/app/user/relation/common/event"
	"errors"
	"strconv"

	"douyin/app/user/relation/common/constants"
	do "douyin/app/user/relation/common/entity"
	po "douyin/app/user/relation/common/model"
	"douyin/app/user/relation/service/internal/biz"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"gorm.io/gorm"
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

// RelationAction 关注/取消关注
func (r *relationRepo) RelationAction(ctx context.Context, relation *event.RelationAction) error {
	b, err := relation.MarshalJson()
	if err != nil {
		r.log.Errorf("json marshal error: %v", err)
		return err
	}
	_, _, err = r.data.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: constants.RelationActionTopic,
		Key:   sarama.StringEncoder(constants.RelationActionKafkaKey(relation.FromUserId)),
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		r.log.Errorf("send kafka error: %v", err)
		return err
	}
	return nil
}

// GetFollowListByUserId 获取关注列表
func (r *relationRepo) GetFollowListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.getFollowListFromCache(ctx, userId)
	if err == nil {
		return res, nil
	}
	if err != redis.Nil {
		log.Errorf("redis error: %v", err)
	}
	var relations []*po.Relation
	// todo
	if err := r.data.db.WithContext(ctx).Table(constants.FollowRecordTable(userId)).Where("from_user_id = ?", userId).Find(&relations).Error; err != nil {
		r.log.Errorf("mysql error: %v", err)
		return nil, err
	}
	ids := make([]int64, 0, len(relations))
	for _, relation := range relations {
		ids = append(ids, relation.ToUserId)
	}
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setUserFollowListCache(ctx, userId, relations)
	})
	if err != nil {
		r.log.Errorf("Fanout error: %v", err)
	}
	return ids, nil
}

// GetFollowerListByUserId 获取粉丝列表
func (r *relationRepo) GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.getFollowerListFromCache(ctx, userId)
	if err == nil {
		return res, nil
	}
	if err != redis.Nil {
		log.Errorf("redis error: %v", err)
	}
	var relations []*po.Relation
	// todo
	if err := r.data.db.WithContext(ctx).Table(constants.FollowerRecordTable(userId)).Where("from_user_id = ?", userId).Find(&relations).Error; err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(relations))
	for _, relation := range relations {
		ids = append(ids, relation.FromUserId)
	}
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setUserFollowerListCache(ctx, userId, relations)
	})
	if err != nil {
		r.log.Errorf("Fanout error: %v", err)
	}
	return ids, nil
}

// GetFriendListByUserId 获取好友列表
func (r *relationRepo) GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.getFriendListFromCache(ctx, userId)
	if err == nil {
		return res, nil
	}
	if err != redis.Nil {
		log.Errorf("redis error: %v", err)
	}
	var relations []*po.Relation
	// todo
	if err := r.data.db.WithContext(ctx).Table(constants.FollowRecordTable(userId)).Where("from_user_id = ? and type = ?", userId, do.RelationFollowed).Find(&relations).Error; err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(relations))
	for _, relation := range relations {
		ids = append(ids, relation.ToUserId)
	}
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setUserFriendListCache(ctx, userId, relations)
	})
	if err != nil {
		r.log.Errorf("Fanout error: %v", err)
	}
	return ids, nil
}

// CountFollowByUserId 获取关注数量
func (r *relationRepo) CountFollowByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err := r.getFollowCountByUserIdFromCache(ctx, userId)
	if err == nil {
		return res, nil
	}
	if err != redis.Nil {
		log.Errorf("redis error: %v", err)
	}
	if err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).Where("user_id = ?", userId).Pluck("follow_count", &res).Error; err != nil {
		return 0, err
	}
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setFollowCountByUserId(ctx, userId, res)
	})
	if err != nil {
		r.log.Errorf("Fanout error: %v", err)
	}
	return res, nil
}

// CountFollowerByUserId 获取粉丝数量
func (r *relationRepo) CountFollowerByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err := r.getFollowerCountByUserIdFromCache(ctx, userId)
	if err == nil {
		return res, nil
	}
	if err != redis.Nil {
		log.Errorf("redis error: %v", err)
	}
	if err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).Where("user_id = ?", userId).Pluck("follower_count", &res).Error; err != nil {
		return 0, err
	}
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setFollowerCountByUserId(ctx, userId, res)
	})
	if err != nil {
		r.log.Errorf("Fanout error: %v", err)
	}
	return res, nil
}

// IsFollowByUserId 是否关注
func (r *relationRepo) IsFollowByUserId(ctx context.Context, userId, toUserId int64) (bool, error) {
	res, err := r.data.redis.BFExists(ctx, constants.UserFollowBloomCacheKey(userId), toUserId).Result()
	// 不在布隆过滤器中，肯定没关注
	if err == nil && !res {
		return false, nil
	}
	if err == redis.Nil {
		err := r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
			r.setUserFollowBloom(ctx, userId)
		})
		if err != nil {
			r.log.Errorf("Fanout error: %v", err)
		}
	} else {
		r.log.Errorf("redis error: %v", err)
	}
	// 在布隆过滤器中或者查询失败，可能关注，需要查询redis
	res, err = r.isUserFollowFromCache(ctx, userId, toUserId)
	if err == nil {
		return res, nil
	}
	if !errors.Is(err, redis.Nil) {
		r.log.Errorf("redis error: %v", err)
	}
	// 不在redis中，可能关注，需要查询mysql
	var relation po.Relation
	if err := r.data.db.WithContext(ctx).Table(constants.FollowRecordTable(userId)).Where("from_user_id = ? and to_user_id = ?", userId, toUserId).First(&relation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		r.log.Errorf("mysql error: %v", err)
		return false, err
	}
	return true, nil
}

// IsFollowByUserIds 批量获取是否关注
func (r *relationRepo) IsFollowByUserIds(ctx context.Context, userId int64, toUserIds []int64) ([]bool, error) {
	// 从布隆过滤器获取
	res, err := r.data.redis.BFMExists(ctx, constants.UserFollowBloomCacheKey(userId), toUserIds).Result()
	if err == nil {
		return res, nil
	}
	if err == redis.Nil {
		err := r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
			r.setUserFollowBloom(ctx, userId)
		})
		if err != nil {
			r.log.Errorf("Fanout error: %v", err)
		}
	} else {
		r.log.Errorf("redis error: %v", err)
	}
	// 从布隆过滤器获取失败，可能关注，需要查询redis
	res, err = r.isUserFollowsFromCache(ctx, userId, toUserIds)
	if err == nil {
		return res, nil
	}
	if !errors.Is(err, redis.Nil) {
		r.log.Errorf("redis error: %v", err)
	}
	// 从redis获取失败，可能关注，需要查询mysql
	var relations []*po.Relation
	if err := r.data.db.WithContext(ctx).Table(constants.FollowRecordTable(userId)).Where("from_user_id = ? and to_user_id in ?", userId, toUserIds).Find(&relations).Error; err != nil {
		r.log.Errorf("mysql error: %v", err)
		return nil, err
	}
	isFollowMap := make(map[int64]bool)
	for _, relation := range relations {
		isFollowMap[relation.ToUserId] = true
	}
	res = make([]bool, len(toUserIds))
	for i, toUserId := range toUserIds {
		res[i] = isFollowMap[toUserId]
	}
	return res, nil
}

// 设置关注数量缓存
func (r *relationRepo) setFollowCountByUserId(ctx context.Context, userId int64, count int64) {
	r.data.redis.Set(ctx, constants.UserFollowCountCacheKey(userId), count, constants.UserFollowCountCacheExpiration)
}

// 设置粉丝数量缓存
func (r *relationRepo) setFollowerCountByUserId(ctx context.Context, userId int64, count int64) {
	r.data.redis.Set(ctx, constants.UserFollowerCountCacheKey(userId), count, constants.UserFollowerCountCacheExpiration)
}

// 从缓存获取关注数量
func (r *relationRepo) getFollowCountByUserIdFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserFollowCountCacheKey(userId)).Int64()
}

// 从缓存获取粉丝数量
func (r *relationRepo) getFollowerCountByUserIdFromCache(ctx context.Context, userId int64) (int64, error) {
	return r.data.redis.Get(ctx, constants.UserFollowerCountCacheKey(userId)).Int64()
}

// 从缓存获取关注列表
func (r *relationRepo) getFollowListFromCache(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRange(ctx, constants.UserFollowListCacheKey(userId), 0, -1).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		return nil, err
	}
	ids := make([]int64, 0, len(res))
	for _, v := range res {
		id := cast.ToInt64(v)
		ids = append(ids, id)
	}
	return ids, nil
}

// 从缓存获取粉丝列表
func (r *relationRepo) getFollowerListFromCache(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRange(ctx, constants.UserFollowerListCacheKey(userId), 0, -1).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		return nil, err
	}
	ids := make([]int64, 0, len(res))
	for _, v := range res {
		id := cast.ToInt64(v)
		ids = append(ids, id)
	}
	return ids, nil
}

// 从缓存获取好友列表
func (r *relationRepo) getFriendListFromCache(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.redis.ZRevRange(ctx, constants.UserFriendListCacheKey(userId), 0, -1).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		return nil, err
	}
	ids := make([]int64, 0, len(res))
	for _, v := range res {
		id := cast.ToInt64(v)
		ids = append(ids, id)
	}
	return ids, nil
}

// 设置关注列表缓存
func (r *relationRepo) setUserFollowListCache(ctx context.Context, userId int64, relations []*po.Relation) {
	pipe := r.data.redis.Pipeline()
	for _, relation := range relations {
		pipe.ZAdd(ctx, constants.UserFollowListCacheKey(userId), redis.Z{
			Score:  float64(relation.CreatedAt.Unix()),
			Member: relation.ToUserId,
		})
	}
	pipe.Expire(ctx, constants.UserFollowListCacheKey(userId), constants.UserFollowListCacheExpiration)
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 设置粉丝列表缓存
func (r *relationRepo) setUserFollowerListCache(ctx context.Context, userId int64, relations []*po.Relation) {
	pipe := r.data.redis.Pipeline()
	for _, relation := range relations {
		pipe.ZAdd(ctx, constants.UserFollowerListCacheKey(userId), redis.Z{
			Score:  float64(relation.CreatedAt.Unix()),
			Member: relation.FromUserId,
		})
	}
	pipe.Expire(ctx, constants.UserFollowerListCacheKey(userId), constants.UserFollowerListCacheExpiration)
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 设置好友列表缓存
func (r *relationRepo) setUserFriendListCache(ctx context.Context, userId int64, relations []*po.Relation) {
	pipe := r.data.redis.Pipeline()
	for _, relation := range relations {
		pipe.ZAdd(ctx, constants.UserFriendListCacheKey(userId), redis.Z{
			Score:  float64(relation.CreatedAt.Unix()),
			Member: relation.ToUserId,
		})
	}
	pipe.Expire(ctx, constants.UserFriendListCacheKey(userId), constants.UserFriendListCacheExpiration)
	_, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 设置用户关注的布隆过滤器
func (r *relationRepo) setUserFollowBloom(ctx context.Context, userId int64) {
	var relations []*po.Relation
	if err := r.data.db.WithContext(ctx).Table(constants.FollowRecordTable(userId)).Where("from_user_id = ?", userId).Find(&relations).Error; err != nil {
		r.log.Errorf("mysql error: %v", err)
		return
	}
	userIds := make([]int64, 0, len(relations))
	for _, relation := range relations {
		userIds = append(userIds, relation.ToUserId)
	}
	err := r.data.redis.BFMAdd(ctx, constants.UserFollowBloomCacheKey(userId), userIds)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
	}
}

// 从缓存的zset获取用户是否关注
func (r *relationRepo) isUserFollowFromCache(ctx context.Context, userId, toUserId int64) (bool, error) {
	res, err := r.data.redis.ZScore(ctx, constants.UserFollowBloomCacheKey(userId), strconv.FormatInt(toUserId, 10)).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("redis error: %v", err)
		}
		return false, err
	}
	return res > 0, nil
}

// 批量从缓存的zset获取用户是否关注
func (r *relationRepo) isUserFollowsFromCache(ctx context.Context, userId int64, toUserIds []int64) ([]bool, error) {
	pipe := r.data.redis.Pipeline()
	for _, toUserId := range toUserIds {
		pipe.ZScore(ctx, constants.UserFollowBloomCacheKey(userId), strconv.FormatInt(toUserId, 10))
	}
	res, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis error: %v", err)
		return nil, err
	}
	results := make([]bool, len(res))
	for i, v := range res {
		if v == nil {
			results[i] = false
			continue
		}
		results[i] = v.(*redis.FloatCmd).Val() > 0
	}
	return results, nil
}
