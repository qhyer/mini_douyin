package service

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	v1 "douyin/api/user/chat/job"
	"douyin/app/user/chat/common/constants"
	do "douyin/app/user/chat/common/entity"
	"douyin/app/user/chat/common/event"
	"douyin/app/user/chat/job/internal/biz"
)

type ChatService struct {
	v1.UnimplementedChatServer

	uc    *biz.ChatUsecase
	kafka sarama.Consumer
	log   *log.Helper
}

func NewChatService(uc *biz.ChatUsecase, kafka sarama.Consumer, logger log.Logger) *ChatService {
	s := &ChatService{uc: uc, kafka: kafka, log: log.NewHelper(logger)}
	go s.SendMessage()
	return s
}

func (s *ChatService) SendMessage() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.SendMsgTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		msg := &event.SendMessage{}
		if err := msg.UnmarshalJson(message.Value); err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}

		err := s.uc.CreateMessage(context.Background(), &do.Message{
			ID:         msg.ID,
			FromUserId: msg.FromUserId,
			ToUserId:   msg.ToUserId,
			Content:    msg.Content,
			CreatedAt:  msg.CreatedAt,
		})
		if err != nil {
			s.log.Errorf("CreateMessage error: %v", err)
		}
	}
}
