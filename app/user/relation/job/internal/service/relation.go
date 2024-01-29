package service

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	v1 "douyin/api/user/relation/job"
	"douyin/app/user/relation/common/constants"
	do "douyin/app/user/relation/common/entity"
	"douyin/app/user/relation/job/internal/biz"
)

type RelationService struct {
	v1.UnimplementedRelationServer

	uc    *biz.RelationUsecase
	kafka sarama.Consumer
	log   *log.Helper
}

func NewRelationService(uc *biz.RelationUsecase, kafka sarama.Consumer, logger log.Logger) *RelationService {
	s := &RelationService{uc: uc, kafka: kafka, log: log.NewHelper(logger)}
	go s.RelationAction()
	return s
}

func (s *RelationService) RelationAction() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.RelationActionTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		relation := &do.RelationAction{}
		err := relation.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		switch relation.Type {
		case do.RelationActionFollow:
			err = s.uc.CreateRelation(context.Background(), relation)
			if err != nil {
				s.log.Errorf("CreateRelation error: %v", err)
			}
		case do.RelationActionUnFollow:
			err = s.uc.DeleteRelation(context.Background(), relation)
			if err != nil {
				s.log.Errorf("DeleteRelation error: %v", err)
			}
		}
	}
}
