package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/h2non/filetype"

	feed "douyin/api/video/feed/service/v1"
	publish "douyin/api/video/publish/service/v1"
	"douyin/common/ecode"
)

type PublishRepo interface {
	PublishVideo(ctx context.Context, userId int64, title string, data []byte) (*publish.PublishActionResponse, error)
	GetUserPublishedVideoList(ctx context.Context, userId, toUserId int64) (*feed.GetPublishedVideoByUserIdResponse, error)
}

type PublishUsecase struct {
	repo PublishRepo
	log  *log.Helper
}

func NewPublishUsecase(repo PublishRepo, logger log.Logger) *PublishUsecase {
	return &PublishUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PublishUsecase) PublishVideo(ctx context.Context, data []byte, userId int64, title string) (*publish.PublishActionResponse, error) {
	if !uc.isVideo(data) {
		return nil, ecode.NotValidVideoFileErr
	}

	res, err := uc.repo.PublishVideo(ctx, userId, title, data)
	if err != nil {
		uc.log.Errorf("PublishVideo error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *PublishUsecase) GetUserPublishedVideoList(ctx context.Context, userId, toUserId int64) (*feed.GetPublishedVideoByUserIdResponse, error) {
	res, err := uc.repo.GetUserPublishedVideoList(ctx, userId, toUserId)
	if err != nil {
		uc.log.Errorf("GetUserPublishedVideoList error: %v", err)
		return nil, err
	}
	return res, err
}

func (uc *PublishUsecase) isVideo(data []byte) bool {
	return filetype.IsVideo(data)
}
