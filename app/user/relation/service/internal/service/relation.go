package service

import (
	"context"

	do "douyin/app/user/relation/common/event"

	v1 "douyin/api/user/relation/service/v1"
	"douyin/app/user/relation/service/internal/biz"
	"douyin/common/ecode"
)

type RelationService struct {
	v1.UnimplementedRelationServer

	uc *biz.RelationUsecase
}

func NewRelationService(uc *biz.RelationUsecase) *RelationService {
	return &RelationService{uc: uc}
}

func (s *RelationService) RelationAction(ctx context.Context, req *v1.RelationActionRequest) (*v1.RelationActionResponse, error) {
	err := s.uc.RelationAction(ctx, &do.RelationAction{
		Type:       do.RelationActionType(req.GetActionType()),
		FromUserId: req.GetUserId(),
		ToUserId:   req.GetToUserId(),
	})
	e := ecode.ConvertErr(err)
	return &v1.RelationActionResponse{
		StatusCode: e.ErrCode,
		StatusMsg:  &e.ErrMsg,
	}, nil
}

func (s *RelationService) GetFollowListByUserId(ctx context.Context, req *v1.GetFollowListByUserIdRequest) (*v1.GetFollowListByUserIdResponse, error) {
	res, err := s.uc.GetFollowListByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFollowListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetFollowListByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		UserIdList: res,
	}, nil
}

func (s *RelationService) GetFollowerListByUserId(ctx context.Context, req *v1.GetFollowerListByUserIdRequest) (*v1.GetFollowerListByUserIdResponse, error) {
	res, err := s.uc.GetFollowerListByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFollowerListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetFollowerListByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		UserIdList: res,
	}, nil
}

func (s *RelationService) GetUserFriendListByUserId(ctx context.Context, req *v1.GetFriendListByUserIdRequest) (*v1.GetFriendListByUserIdResponse, error) {
	res, err := s.uc.GetFriendListByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFriendListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetFriendListByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		UserIdList: res,
	}, nil
}

func (s *RelationService) CountFollowByUserId(ctx context.Context, req *v1.CountFollowByUserIdRequest) (*v1.CountFollowByUserIdResponse, error) {
	res, err := s.uc.CountFollowByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountFollowByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountFollowByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      res,
	}, nil
}

func (s *RelationService) CountFollowerByUserId(ctx context.Context, req *v1.CountFollowerByUserIdRequest) (*v1.CountFollowerByUserIdResponse, error) {
	res, err := s.uc.CountFollowerByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.CountFollowerByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.CountFollowerByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Count:      res,
	}, nil
}

func (s *RelationService) IsFollowByUserId(ctx context.Context, req *v1.IsFollowByUserIdRequest) (*v1.IsFollowByUserIdResponse, error) {
	res, err := s.uc.IsFollowByUserId(ctx, req.GetUserId(), req.GetToUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.IsFollowByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.IsFollowByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		IsFollow:   res,
	}, nil
}

func (s *RelationService) IsFollowByUserIds(ctx context.Context, req *v1.IsFollowByUserIdsRequest) (*v1.IsFollowByUserIdsResponse, error) {
	res, err := s.uc.IsFollowByUserIds(ctx, req.GetUserId(), req.GetToUserIdList())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.IsFollowByUserIdsResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.IsFollowByUserIdsResponse{
		StatusCode:   ecode.Success.ErrCode,
		StatusMsg:    &ecode.Success.ErrMsg,
		IsFollowList: res,
	}, nil
}
