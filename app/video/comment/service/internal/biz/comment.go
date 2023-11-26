package biz

import (
	"context"
	"douyin/app/video/comment/common/constants"
	do "douyin/app/video/comment/common/entity"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
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
	sf   *singleflight.Group
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		sf:   &singleflight.Group{},
	}
}

// CommentAction 发布/删除评论
func (u *CommentUsecase) CommentAction(ctx context.Context, comment *do.CommentAction) (res *do.Comment, err error) {
	// TODO 发布评论过滤敏感词，返回过滤后的评论
	err = u.repo.CommentAction(ctx, comment)
	if err != nil {
		u.log.Errorf("CommentAction error(%v)", err)
		return res, err
	}
	return res, nil
}

// GetCommentListByVideoId 获取视频的评论列表
func (u *CommentUsecase) GetCommentListByVideoId(ctx context.Context, videoId int64) (comments []*do.Comment, err error) {
	res, err, _ := u.sf.Do(constants.SFVideoCommentList(videoId), func() (interface{}, error) {
		return u.repo.GetCommentListByVideoId(ctx, videoId)
	})
	if err != nil {
		u.log.Errorf("GetCommentListByVideoId error(%v)", err)
		return nil, err
	}
	comments = res.([]*do.Comment)
	return comments, nil
}

// CountCommentByVideoId 获取视频的评论数
func (u *CommentUsecase) CountCommentByVideoId(ctx context.Context, videoId int64) (count int64, err error) {
	res, err, _ := u.sf.Do(constants.SFVideoCommentCount(videoId), func() (interface{}, error) {
		return u.repo.CountCommentByVideoId(ctx, videoId)
	})
	if err != nil {
		u.log.Errorf("CountCommentByVideoId error(%v)", err)
		return 0, err
	}
	count = res.(int64)
	return count, nil
}

// MCountCommentByVideoId 批量获取视频的评论数
func (u *CommentUsecase) MCountCommentByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	countList, err := u.repo.MCountCommentByVideoId(ctx, videoIds)
	if err != nil {
		u.log.Errorf("MCountCommentByVideoId error(%v)", err)
		return nil, err
	}
	return countList, nil
}
