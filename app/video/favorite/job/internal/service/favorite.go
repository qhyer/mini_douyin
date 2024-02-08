package service

import (
	"context"
	"sync"
	"time"

	"douyin/app/video/favorite/common/event"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	v1 "douyin/api/video/favorite/job"
	"douyin/app/video/favorite/common/constants"
	"douyin/app/video/favorite/job/internal/biz"
)

type FavoriteService struct {
	v1.UnimplementedFavoriteServer

	uc     *biz.FavoriteUsecase
	kafka  sarama.Consumer
	log    *log.Helper
	waiter sync.WaitGroup

	statCh []chan *event.VideoFavoritedStat
}

func NewFavoriteService(uc *biz.FavoriteUsecase, kafka sarama.Consumer, logger log.Logger) *FavoriteService {
	s := &FavoriteService{
		uc:    uc,
		kafka: kafka,
		log:   log.NewHelper(logger),
	}
	go s.FavoriteAction()
	go s.FavoriteStat()
	for i := 0; i < int(constants.VideoFavoritedCountSharding); i++ {
		s.statCh[i] = make(chan *event.VideoFavoritedStat, 1024)
		s.waiter.Add(1)
		go s.FavoriteStatProc(i)
	}
	return s
}

func (s *FavoriteService) FavoriteAction() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.FavoriteVideoActionTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		fav := event.FavoriteAction{}
		err := fav.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("FavoriteAction UnmarshalJson error: %v", err)
			continue
		}
		if fav.Type == event.FavoriteActionAdd {
			err = s.uc.CreateFavorite(context.Background(), &fav)
			if err != nil {
				s.log.Errorf("CreateFavorite error: %v", err)
			}
		} else if fav.Type == event.FavoriteActionDelete {
			err = s.uc.DeleteFavorite(context.Background(), &fav)
			if err != nil {
				s.log.Errorf("DeleteFavorite error: %v", err)
			}
		} else {
			s.log.Errorf("FavoriteAction type error: %v", err)
		}
	}
}

func (s *FavoriteService) FavoriteStat() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.UpdateVideoFavoritedCountTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		fav := event.VideoFavoritedStat{}
		err := fav.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("FavoriteStat UnmarshalJson error: %v", err)
			continue
		}
		s.statCh[fav.VideoId%constants.VideoFavoritedCountSharding] <- &fav
	}
}

func (s *FavoriteService) FavoriteStatProc(procId int) {
	defer s.waiter.Done()
	ch := s.statCh[procId]
	ts := time.Now()
	for {
		stat, ok := <-ch
		if !ok {
			s.log.Info("FavoriteStatProc exit")
			return
		}
		err := s.uc.UpdateVideoFavoritedTempCount(context.Background(), procId, stat.VideoId, stat.Delta)
		if err != nil {
			s.log.Errorf("UpdateFavoriteStat error: %v", err)

		}
		if time.Since(ts) > time.Second*10 {
			ts = time.Now()
			res, err := s.uc.GetVideoFavoritedCountFromCache(context.Background(), procId)
			if err != nil {
				s.log.Errorf("GetFavoriteCount error: %v", err)
				continue
			}
			err = s.uc.BatchUpdateVideoFavoritedCount(context.Background(), res)
			if err != nil {
				s.log.Errorf("UpdateFavoriteCount error: %v", err)
			}
		}
	}
}
