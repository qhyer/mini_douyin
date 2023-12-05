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
	MGetIsVideoFavoritedByVideoIdAndUserId(ctx context.Context, userId int64, videoIds []int64) ([]bool, error)
}

type FeedUsecase struct {
	repo FeedRepo
	log  *log.Helper
}

func NewFeedUsecase(repo FeedRepo, logger log.Logger) *FeedUsecase {
	return &FeedUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FeedUsecase) GetPublishedVideoByUserId(ctx context.Context, userId int64) ([]*do.Video, error) {
	return uc.repo.GetPublishedVideoByUserId(ctx, userId)
}

func (uc *FeedUsecase) GetPublishedVideoByLatestTimeAndUserId(ctx context.Context, userId int64, latestTime int64) ([]*do.Video, error) {
	return uc.repo.GetPublishedVideoByLatestTime(ctx, latestTime)
}
