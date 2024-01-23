package biz

import (
	"context"
	feed "douyin/api/video/feed/service/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type FeedRepo interface {
	Feed(ctx context.Context, userId, latestTime int64) (*feed.FeedResponse, error)
}

type FeedUsecase struct {
	repo FeedRepo
	log  *log.Helper
}

func NewFeedUsecase(repo FeedRepo, logger log.Logger) *FeedUsecase {
	return &FeedUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FeedUsecase) Feed(ctx context.Context, userId, latestTime int64) (*feed.FeedResponse, error) {
	res, err := uc.repo.Feed(ctx, userId, latestTime)
	if err != nil {
		uc.log.Errorf("Feed error: %v", err)
		return nil, err
	}
	return res, nil
}
