package biz

import (
	"context"
	do "douyin/app/video/feed/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type FeedRepo interface {
	GetPublishedVideoByUserId(ctx context.Context, userId int64) ([]*do.Video, error)
	GetPublishedVideoByLatestTime(ctx context.Context, latestTime int64) ([]*do.Video, error)
	GetUserInfoByUserId(ctx context.Context, userId, toUserId int64) (*do.User, error)
	MGetUserInfoByUserId(ctx context.Context, userId int64, toUserIds []int64) ([]*do.User, error)
	MCountVideoFavoritedByVideoId(ctx context.Context, videoId []int64) ([]int64, error)
	MCountCommentByVideoId(ctx context.Context, videoId []int64) ([]int64, error)
}

type FeedUsecase struct {
	repo FeedRepo
	log  *log.Helper
}

func NewFeedUsecase(repo FeedRepo, logger log.Logger) *FeedUsecase {
	return &FeedUsecase{repo: repo, log: log.NewHelper(logger)}
}
