package biz

import (
	"context"
	do "douyin/app/video/publish/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type VideoRepo interface {
	GetPublishedVideosByUserId(ctx context.Context, userId int64, offset int, limit int) ([]*do.Video, error)
	GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error)
	GetVideoById(ctx context.Context, id int64) (*do.Video, error)
	MGetVideoByIds(ctx context.Context, ids []int64) ([]*do.Video, error)
	UploadVideo(ctx context.Context, data []byte, objectName string) error
	CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error)
}

type VideoUsecase struct {
	repo VideoRepo
	log  *log.Logger
}

func NewVideoUsecase(repo VideoRepo, logger log.Logger) *VideoUsecase {
	return &VideoUsecase{
		repo: repo,
		log:  &logger,
	}
}

func (u *VideoUsecase) PublishVideo(ctx context.Context, video []byte, uid int64, title string) error {
	err := u.repo.UploadVideo(ctx, video, title)
	if err != nil {
		return err
	}
	// TODO push 上传成功的消息到消息队列

	return nil
}

func (u *VideoUsecase) GetPublishedVideosByUserId(ctx context.Context, userId int64, offset int, limit int) ([]*do.Video, error) {
	return u.repo.GetPublishedVideosByUserId(ctx, userId, offset, limit)
}

func (u *VideoUsecase) GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error) {
	return u.repo.GetPublishedVideosByLatestTime(ctx, latestTime, limit)
}

func (u *VideoUsecase) GetVideoById(ctx context.Context, id int64) (*do.Video, error) {
	return u.repo.GetVideoById(ctx, id)
}

func (u *VideoUsecase) MGetVideoByIds(ctx context.Context, ids []int64) ([]*do.Video, error) {
	return u.repo.MGetVideoByIds(ctx, ids)
}

func (u *VideoUsecase) CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error) {
	return u.repo.CountUserPublishedVideoByUserId(ctx, userId)
}
