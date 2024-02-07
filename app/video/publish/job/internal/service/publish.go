package service

import (
	"context"
	v1 "douyin/api/video/publish/job"
	"douyin/app/video/publish/common/constants"
	do "douyin/app/video/publish/common/entity"
	"douyin/app/video/publish/common/event"
	"douyin/app/video/publish/job/internal/biz"
	constants2 "douyin/common/constants"
	"douyin/common/ffmpeg"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

type PublishService struct {
	v1.UnimplementedPublishServer

	uc    *biz.VideoUsecase
	kafka sarama.Consumer
	log   *log.Helper
}

func NewPublishService(uc *biz.VideoUsecase, kafka sarama.Consumer, logger log.Logger) *PublishService {
	s := &PublishService{uc: uc, kafka: kafka, log: log.NewHelper(logger)}
	go s.PublishVideo()
	go s.UploadCover()
	return s
}

func (s *PublishService) PublishVideo() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.PublishVideoTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		video := &event.VideoUpload{}
		err := video.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		err = s.uc.CreateVideo(context.Background(), &do.Video{
			ID:        video.ID,
			AuthorID:  video.AuthorID,
			Title:     video.Title,
			PlayURL:   constants2.VideoOSSURL + video.VideoFileName + ".mp4",
			CoverURL:  constants2.VideoOSSURL + video.VideoFileName + ".jpg",
			CreatedAt: video.CreatedAt,
		})
		if err != nil {
			s.log.Errorf("CreateVideo error: %v", err)
		}
	}
}

func (s *PublishService) UploadCover() {
	partitionConsumer, err := s.kafka.ConsumePartition(constants.GenCoverTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		video := &event.VideoUpload{}
		err := video.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		coverBytes, err := ffmpeg.ReadFrameAsJpeg(fmt.Sprintf("%s%s.mp4", constants2.VideoOSSURL, video.VideoFileName))
		if err != nil {
			s.log.Errorf("readFrameAsJpeg error: %v", err)
			continue
		}
		err = s.uc.UploadCover(context.Background(), coverBytes, video.VideoFileName+".jpg")
		if err != nil {
			s.log.Errorf("UploadCover error: %v", err)
		}
	}
}
