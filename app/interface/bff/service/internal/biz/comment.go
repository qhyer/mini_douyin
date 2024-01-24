package biz

import (
	"context"
	comment "douyin/api/video/comment/service/v1"
	do "douyin/app/video/comment/common/entity"
	"douyin/common/ecode"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentRepo interface {
	CommentAction(ctx context.Context, userId, videoId, commentId int64, actionType int32, commentText string) (*comment.CommentActionResponse, error)
	GetCommentList(ctx context.Context, videoId int64) (*comment.GetCommentListByVideoIdResponse, error)
}

type CommentUsecase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUsecase) CommentAction(ctx context.Context, userId, videoId, commentId int64, actionType int32, commentText string) (*comment.CommentActionResponse, error) {
	if !isCommentActionTypeValid(actionType) {
		return nil, ecode.CommentActionTypeInvalidErr
	}
	res, err := uc.repo.CommentAction(ctx, userId, videoId, commentId, actionType, commentText)
	if err != nil {
		uc.log.Errorf("CommentAction error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *CommentUsecase) GetCommentList(ctx context.Context, videoId int64) (*comment.GetCommentListByVideoIdResponse, error) {
	res, err := uc.repo.GetCommentList(ctx, videoId)
	if err != nil {
		uc.log.Errorf("GetCommentList error: %v", err)
		return nil, err
	}
	return res, nil
}

func isCommentActionTypeValid(actionType int32) bool {
	if actionType == int32(do.CommentActionPublish) || actionType == int32(do.CommentActionDelete) {
		return true
	}
	return false
}
