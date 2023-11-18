package biz

import (
	"context"
	do "douyin/app/video/publish/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type VideoRepo interface {
	CreateVideo(ctx context.Context, video *do.Video) error
	BatchCreateVideo(ctx context.Context, videos []*do.Video) error
	GetVideo(ctx context.Context, objectName string) ([]byte, error)
	UploadCover(ctx context.Context, data []byte, objectName string) error
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
