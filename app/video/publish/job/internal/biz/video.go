package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	do "douyin/app/video/publish/common/entity"
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

func (uc *VideoUsecase) CreateVideo(ctx context.Context, video *do.Video) error {
	return uc.repo.CreateVideo(ctx, video)
}

func (uc *VideoUsecase) BatchCreateVideo(ctx context.Context, videos []*do.Video) error {
	return uc.repo.BatchCreateVideo(ctx, videos)
}

func (uc *VideoUsecase) GetVideo(ctx context.Context, objectName string) ([]byte, error) {
	return uc.repo.GetVideo(ctx, objectName)
}

func (uc *VideoUsecase) UploadCover(ctx context.Context, data []byte, objectName string) error {
	return uc.repo.UploadCover(ctx, data, objectName)
}
