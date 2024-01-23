package data

import (
	"context"
	video "douyin/api/video/feed/service/v1"
	publish "douyin/api/video/publish/service/v1"
	"douyin/app/interface/bff/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type publishRepo struct {
	data *Data
	log  *log.Helper
}

func NewPublishRepo(data *Data, logger log.Logger) biz.PublishRepo {
	return &publishRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *publishRepo) PublishVideo(ctx context.Context, userId int64, title string, data []byte) (*publish.PublishActionResponse, error) {
	res, err := r.data.PublishRPC.PublishVideo(ctx, &publish.PublishActionRequest{
		UserId: userId,
		Data:   data,
		Title:  title,
	})
	return res, err
}

func (r *publishRepo) GetUserPublishedVideoList(ctx context.Context, userId, toUserId int64) (*video.GetPublishedVideoByUserIdResponse, error) {
	res, err := r.data.VideoRPC.GetPublishedVideoByUserId(ctx, &video.GetPublishedVideoByUserIdRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	return res, err
}
