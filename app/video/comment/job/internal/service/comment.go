package service

import (
	"context"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	"douyin/app/video/comment/common/event"

	v1 "douyin/api/video/comment/job"
	"douyin/app/video/comment/common/constants"
	"douyin/app/video/comment/job/internal/biz"
)

type CommentService struct {
	v1.UnimplementedCommentServer

	uc     *biz.CommentUsecase
	kafka  sarama.Consumer
	log    *log.Helper
	waiter sync.WaitGroup

	statCh []chan *event.CommentStat
}

func NewCommentService(uc *biz.CommentUsecase, kafka sarama.Consumer, logger log.Logger) *CommentService {
	s := &CommentService{
		uc:     uc,
		kafka:  kafka,
		log:    log.NewHelper(logger),
		statCh: make([]chan *event.CommentStat, constants.CommentCountSharding),
	}
	go s.CommentAction()
	go s.CommentStat()
	for i := 0; i < constants.CommentCountSharding; i++ {
		s.statCh[i] = make(chan *event.CommentStat, 1024)
		s.waiter.Add(1)
		go s.CommentStatProc(i)
	}
	return s
}

func (s *CommentService) CommentAction() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.CommentActionTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		commentAct := &event.CommentAction{}
		err := commentAct.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("CommentAction UnmarshalJson error: %v", err)
			return
		}
		if commentAct.Type == event.CommentActionPublish {
			err := s.uc.CreateComment(context.Background(), commentAct)
			if err != nil {
				s.log.Errorf("PublishComment error: %v", err)
			}
		} else if commentAct.Type == event.CommentActionDelete {
			err := s.uc.DeleteComment(context.Background(), commentAct)
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
		commentStst := event.CommentStat{}
		err := commentStst.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("CommentStat UnmarshalJson error: %v", err)
			return
		}
		s.statCh[commentStst.VideoId%constants.CommentCountSharding] <- &commentStst
	}
}

func (s *CommentService) CommentStatProc(procId int) {
	defer s.waiter.Done()
	ch := s.statCh[procId]
	ts := time.Now()
	for {
		m, ok := <-ch
		if !ok {
			s.log.Info("CommentStatProc exit")
			return
		}
		err := s.uc.UpdateVideoCommentTempCountCache(context.Background(), procId, m)
		if err != nil {
			s.log.Errorf("UpdateCommentCount error: %v", err)
		}
		if time.Now().Sub(ts) > time.Second*10 {
			ts = time.Now()
			res, err := s.uc.GetVideoCommentTempCountCache(context.Background(), procId)
			if err != nil {
				s.log.Errorf("GetCommentCount error: %v", err)
				continue
			}
			err = s.uc.BatchUpdateVideoCommentCount(context.Background(), res)
			if err != nil {
				s.log.Errorf("BatchUpdateCommentCount error: %v", err)
				continue
			}
			err = s.uc.PurgeVideoCommentTempCountCache(context.Background(), procId)
			if err != nil {
				s.log.Errorf("PurgeCommentCount error: %v", err)
			}
		}
	}
}
