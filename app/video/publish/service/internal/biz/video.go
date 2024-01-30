package biz

import (
	"context"
	"douyin/app/video/publish/common/event"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"

	"douyin/app/video/publish/common/constants"
	do "douyin/app/video/publish/common/entity"
)

type VideoRepo interface {
	PublishVideo(ctx context.Context, video *event.VideoUpload) error
	GetPublishedVideosByUserId(ctx context.Context, userId int64) ([]*do.Video, error)
	GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error)
	GetVideoById(ctx context.Context, id int64) (*do.Video, error)
	MGetVideoByIds(ctx context.Context, ids []int64) ([]*do.Video, error)
	UploadVideo(ctx context.Context, data []byte, objectName string) error
	CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error)
}

type VideoUsecase struct {
	repo VideoRepo
	log  *log.Helper
	sf   *singleflight.Group
}

func NewVideoUsecase(repo VideoRepo, logger log.Logger) *VideoUsecase {
	return &VideoUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		sf:   &singleflight.Group{},
	}
}

// PublishVideo 发布视频
func (u *VideoUsecase) PublishVideo(ctx context.Context, video []byte, uid int64, title string) error {
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), uuid.New().String())
	err := u.repo.UploadVideo(ctx, video, filename+".mp4")
	if err != nil {
		u.log.Errorf("upload video error: %v", err)
		return err
	}
	err = u.repo.PublishVideo(ctx, &event.VideoUpload{
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
func (u *VideoUsecase) GetPublishedVideosByUserId(ctx context.Context, userId int64) ([]*do.Video, error) {
	res, err, _ := u.sf.Do(constants.SFUserPublishedVideoKey(userId), func() (interface{}, error) {
		return u.repo.GetPublishedVideosByUserId(ctx, userId)
	})
	if err != nil {
		u.log.Errorf("get published videos by user id error: %v", err)
		return nil, err
	}
	return res.([]*do.Video), nil
}

// GetPublishedVideosByLatestTime 获取小于某个时间的视频
func (u *VideoUsecase) GetPublishedVideosByLatestTime(ctx context.Context, latestTime int64, limit int) ([]*do.Video, error) {
	res, err, _ := u.sf.Do(constants.SFLatestPublishedVideoKey(latestTime, limit), func() (interface{}, error) {
		return u.repo.GetPublishedVideosByLatestTime(ctx, latestTime, limit)
	})
	if err != nil {
		u.log.Errorf("get published videos by latest time error: %v", err)
		return nil, err
	}
	return res.([]*do.Video), nil
}

// GetVideoById 获取视频信息
func (u *VideoUsecase) GetVideoById(ctx context.Context, videoId int64) (*do.Video, error) {
	res, err, _ := u.sf.Do(constants.SFVideoKey(videoId), func() (interface{}, error) {
		return u.repo.GetVideoById(ctx, videoId)
	})
	if err != nil {
		u.log.Errorf("get video by id error: %v", err)
		return nil, err
	}
	return res.(*do.Video), nil
}

// MGetVideoByIds 批量获取视频信息
func (u *VideoUsecase) MGetVideoByIds(ctx context.Context, videoIds []int64) ([]*do.Video, error) {
	return u.repo.MGetVideoByIds(ctx, videoIds)
}

// CountUserPublishedVideoByUserId 获取用户发布视频数量
func (u *VideoUsecase) CountUserPublishedVideoByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err, _ := u.sf.Do(constants.SFVideoCountKey(userId), func() (interface{}, error) {
		return u.repo.CountUserPublishedVideoByUserId(ctx, userId)
	})
	if err != nil {
		u.log.Errorf("count user published video by user id error: %v", err)
		return 0, err
	}
	return res.(int64), nil
}
