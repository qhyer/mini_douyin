package data

import (
	"context"
	"douyin/app/user/chat/common/constants"
	do "douyin/app/user/chat/common/entity"
	"douyin/app/user/chat/service/internal/biz"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
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
	b, err := message.MarshalJson()
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
	//TODO implement me
	panic("implement me")
}

func (r *chatRepo) GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, myUserId, hisUserId int64) (*do.Message, error) {
	//TODO implement me
	panic("implement me")
}
