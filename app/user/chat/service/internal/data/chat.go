package data

import (
	"context"
	v1 "douyin/api/seq-server/service/v1"
	"douyin/app/user/chat/common/event"
	constants2 "douyin/common/constants"
	"errors"
	"strconv"
	"time"

	"douyin/app/user/chat/common/constants"
	do "douyin/app/user/chat/common/entity"
	"douyin/app/user/chat/service/internal/biz"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
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

func (r *chatRepo) SendMessage(ctx context.Context, message *do.Message) error {
	mid, err := r.data.seqRPC.GetID(ctx, &v1.GetIDRequest{BusinessId: constants2.ChatBusinessId})
	if err != nil {
		r.log.Errorf("r.data.seqRPC.GetID error(%v)", err)
		return err
	}
	message.ID = mid.GetID()
	messageAct := &event.SendMessage{
		ID:         message.ID,
		FromUserId: message.FromUserId,
		ToUserId:   message.ToUserId,
		Content:    message.Content,
		CreatedAt:  time.Now(),
	}
	b, err := messageAct.MarshalJson()
	if err != nil {
		r.log.Errorf("message.MarshalJson() error(%v)", err)
		return err
	}
	_, _, err = r.data.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: constants.SendMsgTopic,
		Key:   sarama.StringEncoder(constants.SendMsgKafkaKey(message.FromUserId, message.ToUserId)),
		Value: sarama.ByteEncoder(b),
	})
	if err != nil {
		r.log.Errorf("kafka.SendMessage error(%v)", err)
		return err
	}
	return nil
}

func (r *chatRepo) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, myUserId, hisUserId, preMsgTime int64, limit int) ([]*do.Message, error) {
	messages, err := r.getMessageListByMyUserIdAndHisUserIdAndPreMsgTimeFromCache(ctx, myUserId, hisUserId, preMsgTime, limit)
	if err == nil {
		return messages, nil
	}
	if !errors.Is(err, redis.Nil) {
		r.log.Errorf("r.getMessageListByMyUserIdAndHisUserIdAndPreMsgTimeFromCache(%d, %d, %d, %d) error(%v)", myUserId, hisUserId, preMsgTime, limit, err)
	}
	var messagesFromDb []*do.Message
	tableName := constants.MessageRecordTable(myUserId, hisUserId)
	res := r.data.db.WithContext(ctx).Table(tableName).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", myUserId, hisUserId, hisUserId, myUserId).Order("id DESC").Limit(limit).Find(&messagesFromDb)
	if res.Error != nil {
		r.log.Errorf("r.data.db.WithContext(ctx).Table(%s).Where(%s).Order(%s).Limit(%d).Find(&messagesFromDb) error(%v)", tableName, "(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", "id DESC", limit, res.Error)
		return nil, res.Error
	}
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setMessageListByMyUserIdAndHisUserIdAndPreMsgTimeCache(ctx, myUserId, hisUserId, preMsgTime, limit, messagesFromDb)
	})
	if err != nil {
		r.log.Errorf("r.data.cacheFan.Do error(%v)", err)
	}
	return messagesFromDb, nil
}

func (r *chatRepo) GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, myUserId, hisUserId int64) (*do.Message, error) {
	message, err := r.getLatestMsgFromCache(ctx, myUserId, hisUserId)
	if err == nil {
		return message, nil
	}
	if !errors.Is(err, redis.Nil) {
		r.log.Errorf("r.getLatestMsgFromCache(%d, %d) error(%v)", myUserId, hisUserId, err)
	}
	tableName := constants.MessageRecordTable(myUserId, hisUserId)
	res := r.data.db.WithContext(ctx).Table(tableName).Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", myUserId, hisUserId, hisUserId, myUserId).Order("id DESC").First(&message)
	if res.Error != nil {
		r.log.Errorf("r.data.db.WithContext(ctx).Table(%s).Where(%s).Order(%s).First(&message) error(%v)", tableName, "(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", "id DESC", res.Error)
		return nil, res.Error
	}
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setLatestMsgCache(ctx, message)
	})
	if err != nil {
		r.log.Errorf("r.data.cacheFan.Do error(%v)", err)
	}
	return message, nil
}

func (r *chatRepo) setLatestMsgCache(ctx context.Context, message *do.Message) {
	key := constants.ChatConversationLatestMsgCacheKey(message.FromUserId, message.ToUserId)
	b, err := message.MarshalJson()
	if err != nil {
		r.log.Errorf("message.MarshalJson() error(%v)", err)
		return
	}
	err = r.data.redis.Set(ctx, key, b, constants.ChatLatestMsgCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis.Set(%s, %s, %d) error(%v)", key, b, constants.ChatLatestMsgCacheExpiration, err)
		return
	}
}

func (r *chatRepo) getLatestMsgFromCache(ctx context.Context, myUserId, hisUserId int64) (*do.Message, error) {
	key := constants.ChatConversationLatestMsgCacheKey(myUserId, hisUserId)
	val, err := r.data.redis.Get(ctx, key).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("redis.Get(%s) error(%v)", key, err)
		}
		return nil, err
	}
	var message do.Message
	err = message.UnmarshalJson(val)
	if err != nil {
		r.log.Errorf("message.UnmarshalJson(%s) error(%v)", val, err)
		return nil, err
	}
	return &message, nil
}

func (r *chatRepo) getMessageListByMyUserIdAndHisUserIdAndPreMsgTimeFromCache(ctx context.Context, myUserId, hisUserId, preMsgTime int64, limit int) ([]*do.Message, error) {
	var messages []*do.Message
	key := constants.ChatConversationCacheKey(myUserId, hisUserId)
	val, err := r.data.redis.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:   strconv.FormatInt(preMsgTime, 10),
		Count: int64(limit),
	}).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("redis.ZRangeByScore(%s, %d) error(%v)", key, preMsgTime, err)
		}
		return nil, err
	}
	for _, z := range val {
		var message do.Message
		err = message.UnmarshalJson([]byte(z.Member.(string)))
		if err != nil {
			r.log.Errorf("message.UnmarshalJson(%s) error(%v)", z.Member, err)
			return nil, err
		}
		messages = append(messages, &message)
	}
	return messages, nil
}

func (r *chatRepo) setMessageListByMyUserIdAndHisUserIdAndPreMsgTimeCache(ctx context.Context, myUserId, hisUserId, preMsgTime int64, limit int, messages []*do.Message) {
	key := constants.ChatConversationCacheKey(myUserId, hisUserId)
	var z []redis.Z
	for _, message := range messages {
		b, err := message.MarshalJson()
		if err != nil {
			r.log.Errorf("message.MarshalJson() error(%v)", err)
			return
		}
		z = append(z, redis.Z{Score: float64(message.CreatedAt.UnixMicro()), Member: b})
	}
	err := r.data.redis.ZAdd(ctx, key, z...).Err()
	if err != nil {
		r.log.Errorf("redis.ZAdd(%s, %v) error(%v)", key, z, err)
		return
	}
	err = r.data.redis.Expire(ctx, key, constants.ChatLatestMsgCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis.Expire(%s, %d) error(%v)", key, constants.ChatLatestMsgCacheExpiration, err)
		return
	}
}
