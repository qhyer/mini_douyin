package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	do "douyin/app/video/comment/common/entity"
)

type CommentRepo interface {
	UpdateVideoCommentCount(ctx context.Context, videoId int64, incr int64) error
	BatchUpdateVideoCommentCount(ctx context.Context, videoIds []int64, incr []int64) error
	CreateComment(ctx context.Context, comment *do.CommentAction) error
	BatchCreateComment(ctx context.Context, comments []*do.CommentAction) error
	DeleteComment(ctx context.Context, comment *do.CommentAction) error
	BatchDeleteComment(ctx context.Context, comments []*do.CommentAction) error
}

type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) CreateComment(ctx context.Context, comment *do.CommentAction) error {
	err := uc.repo.CreateComment(ctx, comment)
	if err != nil {
		uc.log.Errorf("CreateComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) BatchCreateComment(ctx context.Context, comments []*do.CommentAction) error {
	err := uc.repo.BatchCreateComment(ctx, comments)
	if err != nil {
		uc.log.Errorf("BatchCreateComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) DeleteComment(ctx context.Context, comment *do.CommentAction) error {
	err := uc.repo.DeleteComment(ctx, comment)
	if err != nil {
		uc.log.Errorf("DeleteComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) BatchDeleteComment(ctx context.Context, comments []*do.CommentAction) error {
	err := uc.repo.BatchDeleteComment(ctx, comments)
	if err != nil {
		uc.log.Errorf("BatchDeleteComment error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) UpdateVideoCommentCount(ctx context.Context, videoId int64, incr int64) error {
	err := uc.repo.UpdateVideoCommentCount(ctx, videoId, incr)
	if err != nil {
		uc.log.Errorf("UpdateVideoCommentCount error(%v)", err)
		return err
	}
	return nil
}

func (uc *CommentUsecase) BatchUpdateVideoCommentCount(ctx context.Context, videoIds []int64, incr []int64) error {
	err := uc.repo.BatchUpdateVideoCommentCount(ctx, videoIds, incr)
	if err != nil {
		uc.log.Errorf("BatchUpdateVideoCommentCount error(%v)", err)
		return err
	}
	return nil
}
