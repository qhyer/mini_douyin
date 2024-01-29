package service

import (
	"context"

	"douyin/app/video/feed/common/mapper"
	"douyin/app/video/feed/service/internal/biz"
	"douyin/common/ecode"

	v1 "douyin/api/video/feed/service/v1"
)

type FeedService struct {
	v1.UnimplementedFeedServer

	uc *biz.FeedUsecase
}

func NewFeedService(uc *biz.FeedUsecase) *FeedService {
	return &FeedService{uc: uc}
}

// Feed 获取视频流
func (s *FeedService) Feed(ctx context.Context, req *v1.FeedRequest) (*v1.FeedResponse, error) {
	res, err := s.uc.GetPublishedVideoByLatestTimeAndUserId(ctx, req.GetUserId(), req.GetLatestTime())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.FeedResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	nextTime := int64(0)
	if len(res) > 0 {
		nextTime = res[len(res)-1].CreatedAt.UnixMilli()
	}

	videos, err := mapper.VideoToFeedDTOs(res)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.FeedResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.FeedResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		VideoList:  videos,
		NextTime:   &nextTime,
	}, nil
}

// GetPublishedVideoByUserId 获取用户发布视频列表
func (s *FeedService) GetPublishedVideoByUserId(ctx context.Context, req *v1.GetPublishedVideoByUserIdRequest) (*v1.GetPublishedVideoByUserIdResponse, error) {
	res, err := s.uc.GetPublishedVideoByUserId(ctx, req.GetUserId(), req.GetToUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetPublishedVideoByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	videos, err := mapper.VideoToFeedDTOs(res)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetPublishedVideoByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetPublishedVideoByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		VideoList:  videos,
	}, nil
}

// GetUserFavoriteVideoListByUserId 获取用户收藏视频列表
func (s *FeedService) GetUserFavoriteVideoListByUserId(ctx context.Context, req *v1.GetUserFavoriteVideoListByUserIdRequest) (*v1.GetUserFavoriteVideoListByUserIdResponse, error) {
	res, err := s.uc.GetUserFavoriteVideoListByUserId(ctx, req.GetUserId(), req.GetToUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserFavoriteVideoListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	videos, err := mapper.VideoToFeedDTOs(res)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserFavoriteVideoListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetUserFavoriteVideoListByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		VideoList:  videos,
	}, nil
}
