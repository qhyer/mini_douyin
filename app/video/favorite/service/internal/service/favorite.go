package service

import (
	"context"
	"douyin/api/video/favorite/service/v1"
	do "douyin/app/video/favorite/common/entity"
	"douyin/app/video/favorite/service/internal/biz"
	"douyin/common/ecode"
	"time"
)

type FavoriteService struct {
	v1.UnimplementedFavoriteServer

	uc *biz.FavoriteUsecase
}

func NewFavoriteService(uc *biz.FavoriteUsecase) *FavoriteService {
	return &FavoriteService{uc: uc}
}

func (s *FavoriteService) FavoriteAction(ctx context.Context, req *v1.DouyinFavoriteActionRequest) (*v1.DouyinFavoriteActionResponse, error) {
	err := s.uc.FavoriteAction(ctx, &do.Favorite{
		Type:      do.FavoriteActionType(req.GetActionType()),
		UserId:    req.GetUserId(),
		VideoId:   req.GetVideoId(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinFavoriteActionResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.DouyinFavoriteActionResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
	}, nil
}

func (s *FavoriteService) GetUserFavoriteVideoIdList(ctx context.Context, req *v1.GetUserFavoriteListRequest) (*v1.GetUserFavoriteListResponse, error) {
	favoriteVideoIdList, err := s.uc.GetFavoriteVideoIdListByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserFavoriteListResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetUserFavoriteListResponse{
		StatusCode:  ecode.Success.ErrCode,
		StatusMsg:   &ecode.Success.ErrMsg,
		VideoIdList: favoriteVideoIdList,
	}, nil
}

func (s *FavoriteService) GetFavoriteStatusByUserIdAndVideoIds(ctx context.Context, req *v1.GetFavoriteStatusByUserIdAndVideoIdsRequest) (*v1.GetFavoriteStatusByUserIdAndVideoIdsResponse, error) {
	isFavoriteList, err := s.uc.GetFavoriteStatusByUserIdAndVideoIds(ctx, req.GetUserId(), req.GetVideoIds())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFavoriteStatusByUserIdAndVideoIdsResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetFavoriteStatusByUserIdAndVideoIdsResponse{
		StatusCode:     ecode.Success.ErrCode,
		StatusMsg:      &ecode.Success.ErrMsg,
		IsFavoriteList: isFavoriteList,
	}, nil
}

func (s *FavoriteService) CountVideoFavoriteByUserId(ctx context.Context, req *v1.CountVideoFavoriteByUserIdRequest) (*v1.CountVideoFavoriteByUserIdResponse, error) {
	count, err := s.uc.CountVideoFavoriteByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountVideoFavoriteByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountVideoFavoriteByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      count,
	}, nil
}

func (s *FavoriteService) CountVideoFavoritedByUserId(ctx context.Context, req *v1.CountVideoFavoritedByUserIdRequest) (*v1.CountVideoFavoritedByUserIdResponse, error) {
	count, err := s.uc.CountVideoFavoritedByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountVideoFavoritedByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountVideoFavoritedByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      count,
	}, nil
}

func (s *FavoriteService) CountFavoritedByVideoId(ctx context.Context, req *v1.CountFavoritedByVideoIdRequest) (*v1.CountFavoritedByVideoIdResponse, error) {
	count, err := s.uc.CountVideoFavoritedByVideoId(ctx, req.GetVideoId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountFavoritedByVideoIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountFavoritedByVideoIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      count,
	}, nil
}
