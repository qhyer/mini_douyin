package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "douyin/api/video/favorite/service/v1"
	video "douyin/api/video/feed/service/v1"
	"douyin/app/interface/bff/service/internal/biz"
)

type favoriteRepo struct {
	data *Data
	log  *log.Helper
}

func NewFavoriteRepo(data *Data, logger log.Logger) biz.FavoriteRepo {
	return &favoriteRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *favoriteRepo) FavoriteAction(ctx context.Context, userId, videoId int64, actionType int32) (*v1.DouyinFavoriteActionResponse, error) {
	res, err := r.data.FavoriteRPC.FavoriteAction(ctx, &v1.DouyinFavoriteActionRequest{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
	})
	return res, err
}

func (r *favoriteRepo) GetUserFavoriteVideoList(ctx context.Context, userId, toUserId int64) (*video.GetUserFavoriteVideoListByUserIdResponse, error) {
	res, err := r.data.VideoRPC.GetUserFavoriteVideoListByUserId(ctx, &video.GetUserFavoriteVideoListByUserIdRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	return res, err
}
