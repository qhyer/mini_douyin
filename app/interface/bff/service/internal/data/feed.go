package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	video "douyin/api/video/feed/service/v1"
	"douyin/app/interface/bff/service/internal/biz"
)

type feedRepo struct {
	data *Data
	log  *log.Helper
}

func NewFeedRepo(data *Data, logger log.Logger) biz.FeedRepo {
	return &feedRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *feedRepo) Feed(ctx context.Context, userId, latestTime int64) (*video.FeedResponse, error) {
	res, err := r.data.VideoRPC.Feed(ctx, &video.FeedRequest{
		UserId:     &userId,
		LatestTime: &latestTime,
	})
	return res, err
}
