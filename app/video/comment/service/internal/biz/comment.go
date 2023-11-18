package biz

import (
	"context"
	do "douyin/app/video/comment/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentRepo interface {
	CommentAction(ctx context.Context, comment *do.Comment) error
	GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*do.Comment, error)
	CountCommentByVideoId(ctx context.Context, videoId int64) (int64, error)
}

type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *CommentUsecase) CommentAction(ctx context.Context, comment *do.Comment) (err error) {
	err = u.repo.CommentAction(ctx, comment)
	if err != nil {
		return err
	}
	return nil
}

func (u *CommentUsecase) GetCommentListByVideoId(ctx context.Context, videoId int64) (comments []*do.Comment, err error) {
	comments, err = u.repo.GetCommentListByVideoId(ctx, videoId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (u *CommentUsecase) CountCommentByVideoId(ctx context.Context, videoId int64) (count int64, err error) {
	count, err = u.repo.CountCommentByVideoId(ctx, videoId)
	if err != nil {
		return 0, err
	}
	return count, nil
}
