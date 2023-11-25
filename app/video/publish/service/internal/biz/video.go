package biz

import (
	"context"
	do "douyin/app/video/publish/common/entity"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type VideoRepo interface {
	PublishVideo(ctx context.Context, video *do.Video) error
	GetPublishedVideosByUserId(ctx context.Context, userId int64, offset int, limit int) ([]*do.Video, error)
	GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error)
	GetVideoById(ctx context.Context, id int64) (*do.Video, error)
	MGetVideoByIds(ctx context.Context, ids []int64) ([]*do.Video, error)
	UploadVideo(ctx context.Context, data []byte, objectName string) (string, error)
	CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error)
}

type VideoUsecase struct {
	repo VideoRepo
	log  *log.Helper
}

func NewVideoUsecase(repo VideoRepo, logger log.Logger) *VideoUsecase {
	return &VideoUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// PublishVideo 发布视频
func (u *VideoUsecase) PublishVideo(ctx context.Context, video []byte, uid int64, title string) error {
	// TODO get seq-number
	filename, err := u.repo.UploadVideo(ctx, video, "")
	if err != nil {
		u.log.Errorf("upload video error: %v", err)
		return err
	}
	err = u.repo.PublishVideo(ctx, &do.Video{
		AuthorID:      uid,
		Title:         title,
		VideoFileName: filename,
		CreatedAt:     time.Now(),
	})
	if err != nil {
		u.log.Errorf("create video error: %v", err)
		return err
	}
	return nil
}

// GetPublishedVideosByUserId 获取用户发布的视频列表
func (u *VideoUsecase) GetPublishedVideosByUserId(ctx context.Context, userId int64, offset int, limit int) ([]*do.Video, error) {
	return u.repo.GetPublishedVideosByUserId(ctx, userId, offset, limit)
}

// GetPublishedVideosByLatestTime 获取小于某个时间的视频
func (u *VideoUsecase) GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error) {
	return u.repo.GetPublishedVideosByLatestTime(ctx, latestTime, limit)
}

// GetVideoById 获取视频信息
func (u *VideoUsecase) GetVideoById(ctx context.Context, id int64) (*do.Video, error) {
	return u.repo.GetVideoById(ctx, id)
}

// MGetVideoByIds 批量获取视频信息
func (u *VideoUsecase) MGetVideoByIds(ctx context.Context, ids []int64) ([]*do.Video, error) {
	return u.repo.MGetVideoByIds(ctx, ids)
}

// CountUserPublishedVideoByUserId 获取用户发布视频数量
func (u *VideoUsecase) CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error) {
	return u.repo.CountUserPublishedVideoByUserId(ctx, userId)
}
