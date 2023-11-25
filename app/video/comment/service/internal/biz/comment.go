package biz

import (
	"context"
	do "douyin/app/video/comment/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentRepo interface {
	CommentAction(ctx context.Context, comment *do.CommentAction) error
	GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*do.Comment, error)
	CountCommentByVideoId(ctx context.Context, videoId int64) (int64, error)
	MCountCommentByVideoId(ctx context.Context, videoIds []int64) ([]int64, error)
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

// CommentAction 发布/删除评论
func (u *CommentUsecase) CommentAction(ctx context.Context, comment *do.CommentAction) (res *do.Comment, err error) {
	// TODO 发布评论过滤敏感词，返回过滤后的评论
	err = u.repo.CommentAction(ctx, comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

// GetCommentListByVideoId 获取视频的评论列表
func (u *CommentUsecase) GetCommentListByVideoId(ctx context.Context, videoId int64) (comments []*do.Comment, err error) {
	comments, err = u.repo.GetCommentListByVideoId(ctx, videoId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// CountCommentByVideoId 获取视频的评论数
func (u *CommentUsecase) CountCommentByVideoId(ctx context.Context, videoId int64) (count int64, err error) {
	count, err = u.repo.CountCommentByVideoId(ctx, videoId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// MCountCommentByVideoId 批量获取视频的评论数
func (u *CommentUsecase) MCountCommentByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	countList, err := u.repo.MCountCommentByVideoId(ctx, videoIds)
	if err != nil {
		return nil, err
	}
	return countList, nil
}
