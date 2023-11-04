package service

import (
	"context"
	do "douyin/app/user/passport/common/entity"
	"douyin/app/user/passport/common/mapper"
	"douyin/app/user/passport/service/internal/biz"
	"douyin/common/ecode"
	"douyin/common/jwt"

	v1 "douyin/app/user/passport/service/api/v1"
)

type PassportService struct {
	v1.UnimplementedPassportServer

	uc    *biz.PassportUsecase
	jwtuc jwt.JWT
}

func NewPassportService(uc *biz.PassportUsecase) *PassportService {
	return &PassportService{uc: uc}
}

func (s *PassportService) Register(ctx context.Context, req *v1.DouyinUserRegisterRequest) (*v1.DouyinUserRegisterResponse, error) {
	uid := int64(0)
	err := s.uc.CreateUser(ctx, &do.User{
		ID:       uid,
		Name:     req.GetUsername(),
		Password: req.GetPassword(),
	})

	// 创建用户失败
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinUserRegisterResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	// 创建成功 生成token
	token, err := s.jwtuc.CreateTokenByID(uid)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinUserRegisterResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  err.ErrMsg,
		}, nil
	}

	return &v1.DouyinUserRegisterResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  ecode.Success.ErrMsg,
		UserId:     uid,
		Token:      token,
	}, nil
}

func (s *PassportService) Login(ctx context.Context, req *v1.DouyinUserLoginRequest) (*v1.DouyinUserLoginResponse, error) {
	verified, uid, err := s.uc.VerifyPassword(ctx, &do.User{
		Name:     req.GetUsername(),
		Password: req.GetPassword(),
	})
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinUserLoginResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	if verified {
		token, err := s.jwtuc.CreateTokenByID(uid)
		if err != nil {
			err := ecode.ConvertErr(err)
			return &v1.DouyinUserLoginResponse{
				StatusCode: err.ErrCode,
				StatusMsg:  &err.ErrMsg,
			}, nil
		}
		return &v1.DouyinUserLoginResponse{
			StatusCode: ecode.Success.ErrCode,
			StatusMsg:  &ecode.Success.ErrMsg,
			UserId:     uid,
			Token:      token,
		}, nil
	} else {
		return &v1.DouyinUserLoginResponse{
			StatusCode: ecode.WrongPasswordErr.ErrCode,
			StatusMsg:  &ecode.WrongPasswordErr.ErrMsg,
		}, nil
	}
}

func (s *PassportService) GetInfo(ctx context.Context, req *v1.DouyinGetUserInfoRequest) (*v1.DouyinGetUserInfoResponse, error) {
	user, err := s.uc.GetUserByID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	userInfo, err := mapper.UserToVO(user)
	if err != nil {
		return nil, err
	}
	return &v1.DouyinGetUserInfoResponse{
		Info: userInfo,
	}, nil
}

func (s *PassportService) MGetInfo(ctx context.Context, req *v1.DouyinMultipleGetUserInfoRequest) (*v1.DouyinMultipleGetUserInfoResponse, error) {
	users, err := s.uc.MGetUserByID(ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}
	userInfos, err := mapper.UsersToVO(users)
	if err != nil {
		return nil, err
	}
	return &v1.DouyinMultipleGetUserInfoResponse{
		Infos: userInfos,
	}, nil
}
