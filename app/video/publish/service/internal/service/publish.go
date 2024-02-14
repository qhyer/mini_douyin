package service

import (
	"context"

	"douyin/common/constants"

	v1 "douyin/api/video/publish/service/v1"
	"douyin/app/video/publish/common/mapper"
	"douyin/app/video/publish/service/internal/biz"
	"douyin/common/ecode"
)

type PublishService struct {
	v1.UnimplementedPublishServer

	uc *biz.VideoUsecase
}

func NewPublishService(uc *biz.VideoUsecase) *PublishService {
	return &PublishService{uc: uc}
}

// PublishVideo 发布视频
func (s *PublishService) PublishVideo(ctx context.Context, req *v1.PublishActionRequest) (*v1.PublishActionResponse, error) {
	err := s.uc.PublishVideo(ctx, req.GetData(), req.GetUserId(), req.GetTitle())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.PublishActionResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.PublishActionResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
	}, nil
}

// GetUserPublishedVideoList 获取用户发布的视频列表
func (s *PublishService) GetUserPublishedVideoList(ctx context.Context, req *v1.GetUserPublishedVideoListRequest) (*v1.GetUserPublishedVideoListResponse, error) {
	videos, err := s.uc.GetPublishedVideosByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserPublishedVideoListResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	vs, err := mapper.VideoToDTOs(videos)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserPublishedVideoListResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetUserPublishedVideoListResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		VideoList:  vs,
	}, nil
}

// GetPublishedVideoByLatestTime 获取小于某个时间的视频列表
func (s *PublishService) GetPublishedVideoByLatestTime(ctx context.Context, req *v1.GetPublishedVideoByLatestTimeRequest) (*v1.GetPublishedVideoByLatestTimeResponse, error) {
	videos, err := s.uc.GetPublishedVideosByLatestTime(ctx, req.GetLatestTime(), constants.VideoQueryLimit)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetPublishedVideoByLatestTimeResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	vs, err := mapper.VideoToDTOs(videos)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetPublishedVideoByLatestTimeResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetPublishedVideoByLatestTimeResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		VideoList:  vs,
	}, nil
}

// GetVideoById 获取视频信息
func (s *PublishService) GetVideoById(ctx context.Context, req *v1.GetVideoInfoRequest) (*v1.GetVideoInfoResponse, error) {
	video, err := s.uc.GetVideoById(ctx, req.GetVideoId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetVideoInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	v, err := mapper.VideoToDTO(video)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetVideoInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetVideoInfoResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		VideoList:  v,
	}, nil
}

// MGetVideoByIds 批量获取视频信息
func (s *PublishService) MGetVideoByIds(ctx context.Context, req *v1.MGetVideoInfoRequest) (*v1.MGetVideoInfoResponse, error) {
	videos, err := s.uc.MGetVideoByIds(ctx, req.GetVideoIds())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.MGetVideoInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	vs, err := mapper.VideoToDTOs(videos)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.MGetVideoInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.MGetVideoInfoResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		VideoList:  vs,
	}, nil
}

// CountUserPublishedVideoByUserId 获取用户发布视频数量
func (s *PublishService) CountUserPublishedVideoByUserId(ctx context.Context, req *v1.CountUserPublishedVideoByUserIdRequest) (*v1.CountUserPublishedVideoByUserIdResponse, error) {
	count, err := s.uc.CountUserPublishedVideoByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountUserPublishedVideoByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.CountUserPublishedVideoByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      count,
	}, nil
}
