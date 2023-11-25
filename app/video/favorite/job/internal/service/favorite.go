package service

import (
	"context"
	v1 "douyin/api/video/favorite/job"
	"douyin/app/video/favorite/common/constants"
	do "douyin/app/video/favorite/common/entity"
	"douyin/app/video/favorite/job/internal/biz"
	"douyin/app/video/favorite/job/internal/conf"
	"douyin/common/queue/kafka"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

type FavoriteService struct {
	v1.UnimplementedFavoriteServer

	uc    *biz.FavoriteUsecase
	kafka sarama.Consumer
	log   *log.Helper
}

func NewFavoriteService(uc *biz.FavoriteUsecase, kafka sarama.Consumer, logger log.Logger) *FavoriteService {
	s := &FavoriteService{
		uc:    uc,
		kafka: kafka,
		log:   log.NewHelper(logger),
	}
	go s.FavoriteAction()
	return s
}

func NewKafka(c *conf.Data) sarama.Consumer {
	return kafka.NewKafkaConsumer(&kafka.Config{
		Addr: c.GetKafka().GetAddr(),
	})
}

func (s *FavoriteService) FavoriteAction() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.FavoriteVideoActionTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		fav := do.FavoriteAction{}
		err := fav.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("FavoriteAction UnmarshalJson error: %v", err)
			continue
		}
		if fav.Type == do.FavoriteActionAdd {
			err = s.uc.CreateFavorite(context.Background(), &fav)
			if err != nil {
				s.log.Errorf("CreateFavorite error: %v", err)
			}
		} else if fav.Type == do.FavoriteActionDelete {
			err = s.uc.DeleteFavorite(context.Background(), &fav)
			if err != nil {
				s.log.Errorf("DeleteFavorite error: %v", err)
			}
		} else {
			s.log.Errorf("FavoriteAction type error: %v", err)
		}
	}
}
