package service

import (
	"context"

	"douyin/app/user/account/common/mapper"
	"douyin/app/user/account/service/internal/biz"
	"douyin/common/ecode"

	v1 "douyin/api/user/account/service/v1"
)

type AccountService struct {
	v1.UnimplementedAccountServer

	uc *biz.AccountUsecase
}

func NewAccountService(uc *biz.AccountUsecase) *AccountService {
	return &AccountService{uc: uc}
}

func (s *AccountService) GetUserInfoByUserId(ctx context.Context, req *v1.GetUserInfoByUserIdRequest) (*v1.GetUserInfoByUserIdResponse, error) {
	u, err := s.uc.GetUserInfoByUserId(ctx, req.GetUserId(), req.GetToUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserInfoByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	user, err := mapper.UserToDTO(u)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetUserInfoByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.GetUserInfoByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		User:       user,
	}, nil
}

func (s *AccountService) MGetUserInfoByUserId(ctx context.Context, req *v1.MGetUserInfoByUserIdRequest) (*v1.MGetUserInfoByUserIdResponse, error) {
	us, err := s.uc.MGetUserInfoByUserId(ctx, req.GetToUserIds())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.MGetUserInfoByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	users, err := mapper.UserToDTOs(us)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.MGetUserInfoByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.MGetUserInfoByUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Users:      users,
	}, nil
}

func (s *AccountService) GetFollowListByUserId(ctx context.Context, req *v1.GetFollowListByUserIdRequest) (*v1.GetFollowListByUserIdResponse, error) {
	res, err := s.uc.GetFollowListByUserId(ctx, req.GetUserId(), req.GetToUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFollowListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	users, err := mapper.UserToDTOs(res)
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
		Users:      users,
	}, nil
}

func (s *AccountService) GetFollowerListByUserId(ctx context.Context, req *v1.GetFollowerListByUserIdRequest) (*v1.GetFollowerListByUserIdResponse, error) {
	res, err := s.uc.GetFollowerListByUserId(ctx, req.GetUserId(), req.GetToUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFollowerListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	users, err := mapper.UserToDTOs(res)
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
		Users:      users,
	}, nil
}

func (s *AccountService) GetFriendListByUserId(ctx context.Context, req *v1.GetFriendListByUserIdRequest) (*v1.GetFriendListByUserIdResponse, error) {
	res, err := s.uc.GetFriendListByUserId(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetFriendListByUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	users, err := mapper.UserToDTOs(res)
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
		Users:      users,
	}, nil
}
