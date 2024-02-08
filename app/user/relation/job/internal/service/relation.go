package service

import (
	"context"
	"sync"
	"time"

	"douyin/app/user/relation/common/event"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	v1 "douyin/api/user/relation/job"
	"douyin/app/user/relation/common/constants"
	"douyin/app/user/relation/job/internal/biz"
)

type RelationService struct {
	v1.UnimplementedRelationServer

	uc     *biz.RelationUsecase
	kafka  sarama.Consumer
	log    *log.Helper
	waiter sync.WaitGroup

	statCh []chan *event.RelationStat
}

func NewRelationService(uc *biz.RelationUsecase, kafka sarama.Consumer, logger log.Logger) *RelationService {
	s := &RelationService{
		uc:     uc,
		kafka:  kafka,
		log:    log.NewHelper(logger),
		statCh: make([]chan *event.RelationStat, constants.RelationCountSharding),
	}
	go s.RelationAction()
	go s.FollowStat()
	go s.FollowerStat()
	for i := 0; i < constants.RelationCountSharding; i++ {
		s.statCh[i] = make(chan *event.RelationStat, 1024)
		s.waiter.Add(1)
		go s.RelationStatProc(i)
	}
	return s
}

func (s *RelationService) RelationAction() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.RelationActionTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		relation := &event.RelationAction{}
		err := relation.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		switch relation.Type {
		case event.RelationActionFollow:
			err = s.uc.CreateRelation(context.Background(), relation)
			if err != nil {
				s.log.Errorf("CreateRelation error: %v", err)
			}
		case event.RelationActionUnFollow:
			err = s.uc.DeleteRelation(context.Background(), relation)
			if err != nil {
				s.log.Errorf("DeleteRelation error: %v", err)
			}
		}
	}
}

func (s *RelationService) FollowStat() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.UpdateUserFollowCountTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		stat := &event.RelationStat{}
		err := stat.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		s.statCh[stat.UserId%constants.RelationCountSharding] <- stat
	}
}

func (s *RelationService) FollowerStat() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.UpdateUserFollowerCountTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		stat := &event.RelationStat{}
		err := stat.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		s.statCh[stat.UserId%constants.RelationCountSharding] <- stat
	}
}

func (s *RelationService) RelationStatProc(procId int) {
	defer s.waiter.Done()
	ch := s.statCh[procId]
	ts := time.Now()
	for {
		stat, ok := <-ch
		if !ok {
			s.log.Info("RelationStatProc exit")
			return
		}
		if stat.FollowDelta != 0 {
			err := s.uc.UpdateUserFollowTempCount(context.Background(), procId, stat.UserId, stat.FollowDelta)
			if err != nil {
				s.log.Errorf("UpdateRelationStat error: %v", err)
			}
		}
		if stat.FollowerDelta != 0 {
			err := s.uc.UpdateUserFollowerTempCount(context.Background(), procId, stat.UserId, stat.FollowerDelta)
			if err != nil {
				s.log.Errorf("UpdateRelationStat error: %v", err)
			}
		}
		if time.Since(ts) > time.Minute {
			ts = time.Now()
			follow, err := s.uc.GetUserFollowTempCount(context.Background(), procId)
			if err != nil {
				s.log.Errorf("FlushRelationStat error: %v", err)
				continue
			}
			follower, err := s.uc.GetUserFollowerTempCount(context.Background(), procId)
			if err != nil {
				s.log.Errorf("FlushRelationStat error: %v", err)
				continue
			}
			err = s.uc.BatchUpdateUserRelationStat(context.Background(), follow, follower)
			if err != nil {
				s.log.Errorf("FlushRelationStat error: %v", err)
			}
			err = s.uc.PurgeUserFollowTempCount(context.Background(), procId)
			if err != nil {
				s.log.Errorf("FlushRelationStat error: %v", err)
			}
			err = s.uc.PurgeUserFollowerTempCount(context.Background(), procId)
			if err != nil {
				s.log.Errorf("FlushRelationStat error: %v", err)
			}
		}
	}
}
