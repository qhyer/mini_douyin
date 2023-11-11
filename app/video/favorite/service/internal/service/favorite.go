package service

import (
	"context"
	"douyin/api/video/favorite/service/v1"
	"douyin/app/video/favorite/service/internal/biz"
	"douyin/common/ecode"
)

type FavoriteService struct {
	v1.UnimplementedFavoriteServer

	uc *biz.FavoriteUsecase
}

func NewFavoriteService(uc *biz.FavoriteUsecase) *FavoriteService {
	return &FavoriteService{uc: uc}
}

func (s *FavoriteService) FavoriteAction(ctx context.Context, req *v1.DouyinFavoriteActionRequest) (*v1.DouyinFavoriteActionResponse, error) {
	err := s.uc.FavoriteAction(ctx, req.GetUserId(), req.GetVideoId(), int(req.GetActionType()))
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

func (s *FavoriteService) GetUserFavoriteList(ctx context.Context, req *v1.GetUserFavoriteListRequest) (*v1.GetUserFavoriteListResponse, error) {
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
