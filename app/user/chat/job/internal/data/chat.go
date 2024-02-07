package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"douyin/app/user/chat/common/constants"
	do "douyin/app/user/chat/common/entity"
	"douyin/app/user/chat/job/internal/biz"
)

type chatRepo struct {
	data *Data
	log  *log.Helper
}

func NewChatRepo(data *Data, logger log.Logger) biz.ChatRepo {
	return &chatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *chatRepo) CreateMessage(ctx context.Context, message *do.Message) error {
	err := r.data.db.Table(constants.MessageRecordTable(message.FromUserId, message.ToUserId)).Create(message).Error
	if err != nil {
		r.log.Errorf("CreateMessage error: %v", err)
		return err
	}
	err = r.updateConversationLatestMsg(ctx, message)
	if err != nil {
		r.log.Errorf("updateConversationLatestMsg error: %v", err)
	}

	return nil
}

func (r *chatRepo) updateConversationLatestMsg(ctx context.Context, message *do.Message) error {
	b, err := message.MarshalJson()
	if err != nil {
		r.log.Errorf("message.MarshalJson() error(%v)", err)
		return err
	}
	err = r.data.redis.Set(ctx, constants.ChatConversationLatestMsgCacheKey(message.FromUserId, message.ToUserId), b, constants.ChatLatestMsgCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis.Set error(%v)", err)
		return err
	}
	return nil
}

func (r *chatRepo) addMsgToConversationCache(ctx context.Context, message *do.Message) error {
	key := constants.ChatConversationCacheKey(message.FromUserId, message.ToUserId)
	isExist, err := r.data.redis.Exists(ctx, key).Result()
	if err != nil {
		r.log.Errorf("redis.Exists error(%v)", err)
		return err
	}
	if isExist == 0 {
		return nil
	}

	b, err := message.MarshalJson()
	if err != nil {
		r.log.Errorf("message.MarshalJson() error(%v)", err)
		return err
	}
	err = r.data.redis.ZAdd(ctx, key, redis.Z{
		Score:  float64(message.CreatedAt.Unix()),
		Member: b,
	}).Err()
	if err != nil {
		r.log.Errorf("redis.ZAdd error(%v)", err)
		return err
	}
	return nil
}
