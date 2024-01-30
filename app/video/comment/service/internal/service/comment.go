package service

import (
	"context"
	"douyin/app/video/comment/common/event"
	"time"

	"douyin/app/video/comment/common/mapper"
	"douyin/app/video/comment/service/internal/biz"
	"douyin/common/ecode"

	v1 "douyin/api/video/comment/service/v1"
)

type CommentService struct {
	v1.UnimplementedCommentServer

	uc *biz.CommentUsecase
}

func NewCommentService(uc *biz.CommentUsecase) *CommentService {
	return &CommentService{uc: uc}
}

// CommentAction 发布/删除评论
func (s *CommentService) CommentAction(ctx context.Context, req *v1.CommentActionRequest) (*v1.CommentActionResponse, error) {
	comment, err := s.uc.CommentAction(ctx, &event.CommentAction{
		Type:      event.CommentActionType(req.GetActionType()),
		UserId:    req.GetUserId(),
		VideoId:   req.GetVideoId(),
		Content:   req.GetCommentText(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CommentActionResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	com, err := mapper.CommentToDTO(comment)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CommentActionResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CommentActionResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Comment:    com,
	}, nil
}

// GetCommentListByVideoId 获取视频的评论列表
func (s *CommentService) GetCommentListByVideoId(ctx context.Context, req *v1.GetCommentListByVideoIdRequest) (*v1.GetCommentListByVideoIdResponse, error) {
	res, err := s.uc.GetCommentListByVideoId(ctx, req.GetVideoId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetCommentListByVideoIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	comments, err := mapper.CommentToDTOs(res)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetCommentListByVideoIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetCommentListByVideoIdResponse{
		StatusCode:  ecode.Success.ErrCode,
		StatusMsg:   &ecode.Success.ErrMsg,
		CommentList: comments,
	}, nil
}

func (s *CommentService) CountCommentByVideoId(ctx context.Context, req *v1.CountCommentByVideoIdRequest) (*v1.CountCommentByVideoIdResponse, error) {
	res, err := s.uc.CountCommentByVideoId(ctx, req.GetVideoId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountCommentByVideoIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountCommentByVideoIdResponse{
		StatusCode:   ecode.Success.ErrCode,
		StatusMsg:    &ecode.Success.ErrMsg,
		CommentCount: res,
	}, nil
}

func (s *CommentService) MCountCommentByVideoId(ctx context.Context, req *v1.MCountCommentByVideoIdRequest) (*v1.MCountCommentByVideoIdResponse, error) {
	res, err := s.uc.MCountCommentByVideoId(ctx, req.GetVideoIdList())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.MCountCommentByVideoIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.MCountCommentByVideoIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		CountList:  res,
	}, nil
}
