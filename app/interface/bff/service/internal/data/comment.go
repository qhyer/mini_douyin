package data

import (
	"context"
	v1 "douyin/api/video/comment/service/v1"
	"douyin/app/interface/bff/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentRepo) CommentAction(ctx context.Context, userId, videoId, commentId int64, actionType int32, commentText string) (*v1.CommentActionResponse, error) {
	res, err := r.data.CommentRPC.CommentAction(ctx, &v1.CommentActionRequest{
		UserId:      userId,
		VideoId:     videoId,
		ActionType:  actionType,
		CommentText: &commentText,
		CommentId:   &commentId,
	})
	return res, err
}

func (r *commentRepo) GetCommentList(ctx context.Context, videoId int64) (*v1.GetCommentListByVideoIdResponse, error) {
	res, err := r.data.CommentRPC.GetCommentListByVideoId(ctx, &v1.GetCommentListByVideoIdRequest{
		VideoId: videoId,
	})
	return res, err
}
