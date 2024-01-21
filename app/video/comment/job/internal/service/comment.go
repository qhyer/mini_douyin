package service

import (
	"context"
	v1 "douyin/api/video/comment/job"
	"douyin/app/video/comment/common/constants"
	do "douyin/app/video/comment/common/entity"
	"douyin/app/video/comment/job/internal/biz"
	"douyin/app/video/comment/job/internal/conf"
	"douyin/common/queue/kafka"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentService struct {
	v1.UnimplementedCommentServer

	uc    *biz.CommentUsecase
	kafka sarama.Consumer
	log   *log.Helper
}

func NewKafka(c *conf.Data) sarama.Consumer {
	return kafka.NewKafkaConsumer(&kafka.Config{
		Addr: c.GetKafka().GetAddr(),
	})
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
		commentAct := do.CommentAction{}
		err := commentAct.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("CommentAction UnmarshalJson error: %v", err)
			return
		}
		if commentAct.Type == do.CommentActionTypePublish {
			err := s.uc.CreateComment(context.Background(), &do.Comment{
				ID:      commentAct.ID,
				VideoId: commentAct.VideoId,
				User: &do.User{
					ID: commentAct.UserId,
				},
				Content:   commentAct.Content,
				CreatedAt: commentAct.CreatedAt,
			})
			if err != nil {
				s.log.Errorf("PublishComment error: %v", err)
			}
		} else if commentAct.Type == do.CommentActionTypeDelete {
			err := s.uc.DeleteComment(context.Background(), &do.Comment{
				ID: commentAct.ID,
			})
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
	for message := range partitionConsumer.Messages() {
		// todo
	}
}
