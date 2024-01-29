package service

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	v1 "douyin/api/video/publish/job"
	"douyin/app/video/publish/common/constants"
	do "douyin/app/video/publish/common/entity"
	"douyin/app/video/publish/job/internal/biz"
	"douyin/app/video/publish/job/internal/conf"
	constants2 "douyin/common/constants"
	"douyin/common/queue/kafka"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type PublishService struct {
	v1.UnimplementedPublishServer

	uc    *biz.VideoUsecase
	kafka sarama.Consumer
	log   *log.Helper
}

func NewKafka(c *conf.Data) sarama.Consumer {
	return kafka.NewKafkaConsumer(&kafka.Config{
		Addr: c.GetKafka().GetAddr(),
	})
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
		video := &do.Video{}
		err := video.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		err = s.uc.CreateVideo(context.Background(), video)
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
		video := &do.Video{}
		err := video.UnmarshalJson(message.Value)
		if err != nil {
			s.log.Errorf("UnmarshalJson error: %v", err)
			continue
		}
		coverBytes, err := readFrameAsJpeg(constants2.VideoOSSURL + video.VideoFileName)
		if err != nil {
			s.log.Errorf("readFrameAsJpeg error: %v", err)
			continue
		}
		err = s.uc.UploadCover(context.Background(), coverBytes, video.VideoFileName)
		if err != nil {
			s.log.Errorf("UploadCover error: %v", err)
		}
	}
}

func readFrameAsJpeg(inFileName string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), err
}
