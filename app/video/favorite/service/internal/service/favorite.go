package service

import (
	"context"
	"time"

	event "douyin/app/video/favorite/common/event"

	v1 "douyin/api/video/favorite/service/v1"
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

// FavoriteAction 视频点赞
func (s *FavoriteService) FavoriteAction(ctx context.Context, req *v1.DouyinFavoriteActionRequest) (*v1.DouyinFavoriteActionResponse, error) {
	err := s.uc.FavoriteAction(ctx, &event.FavoriteAction{
		Type:      event.FavoriteActionType(req.GetActionType()),
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

// GetUserFavoriteVideoIdList 获取用户点赞视频列表
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

// GetFavoriteStatusByUserIdAndVideoIds 获取用户是否点赞视频
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

// CountUserFavoriteByUserId 获取用户点赞数
func (s *FavoriteService) CountUserFavoriteByUserId(ctx context.Context, req *v1.CountUserFavoriteByUserIdRequest) (*v1.CountUserFavoriteByUserIdResponse, error) {
	count, err := s.uc.CountUserFavoriteByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountUserFavoriteByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountUserFavoriteByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      count,
	}, nil
}

// CountUserFavoritedByUserId 获取用户被点赞数
func (s *FavoriteService) CountUserFavoritedByUserId(ctx context.Context, req *v1.CountUserFavoritedByUserIdRequest) (*v1.CountUserFavoritedByUserIdResponse, error) {
	count, err := s.uc.CountUserFavoritedByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountUserFavoritedByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountUserFavoritedByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      count,
	}, nil
}

// CountVideoFavoritedByVideoId 获取视频被点赞数
func (s *FavoriteService) CountVideoFavoritedByVideoId(ctx context.Context, req *v1.CountVideoFavoritedByVideoIdRequest) (*v1.CountVideoFavoritedByVideoIdResponse, error) {
	count, err := s.uc.CountVideoFavoritedByVideoId(ctx, req.GetVideoId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountVideoFavoritedByVideoIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountVideoFavoritedByVideoIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      count,
	}, nil
}

// MCountVideoFavoritedByVideoIds 批量获取用户点赞数
func (s *FavoriteService) MCountVideoFavoritedByVideoIds(ctx context.Context, req *v1.MCountVideoFavoritedByVideoIdsRequest) (*v1.MCountVideoFavoritedByVideoIdsResponse, error) {
	countList, err := s.uc.MCountVideoFavoritedByVideoId(ctx, req.GetVideoIds())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.MCountVideoFavoritedByVideoIdsResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.MCountVideoFavoritedByVideoIdsResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		CountList:  countList,
	}, nil
}
