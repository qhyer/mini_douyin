package service

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	"douyin/app/video/comment/common/event"

	v1 "douyin/api/video/comment/job"
	"douyin/app/video/comment/common/constants"
	"douyin/app/video/comment/job/internal/biz"
)

type CommentService struct {
	v1.UnimplementedCommentServer

	uc    *biz.CommentUsecase
	kafka sarama.Consumer
	log   *log.Helper
}

func NewCommentService(uc *biz.CommentUsecase, kafka sarama.Consumer, logger log.Logger) *CommentService {
	s := &CommentService{uc: uc, kafka: kafka, log: log.NewHelper(logger)}
	go s.CommentAction()
	go s.CommentStat()
	return s
}

func (s *CommentService) CommentAction() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.CommentActionTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		commentAct := event.CommentAction{}
		err := commentAct.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("CommentAction UnmarshalJson error: %v", err)
			return
		}
		if commentAct.Type == event.CommentActionPublish {
			err := s.uc.CreateComment(context.Background(), &commentAct)
			if err != nil {
				s.log.Errorf("PublishComment error: %v", err)
			}
		} else if commentAct.Type == event.CommentActionDelete {
			err := s.uc.DeleteComment(context.Background(), &commentAct)
			if err != nil {
				s.log.Errorf("DeleteComment error: %v", err)
			}
		} else {
			s.log.Errorf("CommentAction type not found: %v", commentAct.Type)
		}
	}
}

func (s *CommentService) CommentStat() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.UpdateCommentCountTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	//for message := range partitionConsumer.Messages() {
	// todo
	//}
}
