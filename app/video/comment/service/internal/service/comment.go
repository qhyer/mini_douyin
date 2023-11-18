package service

import (
	"context"
	"douyin/app/video/comment/service/internal/biz"

	"douyin/api/video/comment/service/v1"
)

type CommentService struct {
	v1.UnimplementedCommentServer

	uc *biz.CommentUsecase
}

func NewCommentService(uc *biz.CommentUsecase) *CommentService {
	return &CommentService{uc: uc}
}

func (s *CommentService) CommentAction(ctx context.Context, req *v1.CommentActionRequest) (*v1.CommentActionResponse, error) {
	return &v1.CommentActionResponse{}, nil
}
func (s *CommentService) GetCommentListByVideoId(ctx context.Context, req *v1.GetCommentListByVideoIdRequest) (*v1.GetCommentListByVideoIdResponse, error) {
	return &v1.GetCommentListByVideoIdResponse{}, nil
}
func (s *CommentService) CountCommentByVideoId(ctx context.Context, req *v1.CountCommentByVideoIdRequest) (*v1.CountCommentByVideoIdResponse, error) {
	return &v1.CountCommentByVideoIdResponse{}, nil
}
