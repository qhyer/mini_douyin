package service

import (
	"context"
	v1 "douyin/api/user/chat/job"
	"douyin/app/user/chat/common/constants"
	do "douyin/app/user/chat/common/entity"
	"douyin/app/user/chat/common/event"
	"douyin/app/user/chat/job/internal/biz"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

type ChatService struct {
	v1.UnimplementedChatServer

	uc    *biz.ChatUsecase
	kafka sarama.ConsumerGroup
	log   *log.Helper
}

func NewChatService(uc *biz.ChatUsecase, kafka sarama.ConsumerGroup, logger log.Logger) *ChatService {
	s := &ChatService{uc: uc, kafka: kafka, log: log.NewHelper(logger)}
	go s.SendMessage()
	return s
}

func (s *ChatService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	const concurrency = 10
	for i := 0; i < concurrency; i++ {
		go func() {
			for message := range claim.Messages() {
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

				session.MarkMessage(message, "")
			}
		}()
	}
	<-session.Context().Done()
	return nil
}

type ConsumerGroupHandler struct {
	ChatService *ChatService
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	return h.ChatService.ConsumeClaim(session, claim)
}

func (s *ChatService) SendMessage() {
	ctx := context.Background()
	topics := []string{constants.SendMsgTopic}

	for {
		handler := ConsumerGroupHandler{ChatService: s}
		err := s.kafka.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
