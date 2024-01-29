package service

import (
	"context"

	v1 "douyin/api/user/passport/service/v1"
	do "douyin/app/user/passport/common/entity"
	"douyin/app/user/passport/common/mapper"
	"douyin/app/user/passport/service/internal/biz"
	"douyin/common/ecode"
	"douyin/common/jwt"
)

type PassportService struct {
	v1.UnimplementedPassportServer

	uc *biz.PassportUsecase
}

func NewPassportService(uc *biz.PassportUsecase) *PassportService {
	return &PassportService{uc: uc}
}

// Register 用户注册
func (s *PassportService) Register(ctx context.Context, req *v1.DouyinUserRegisterRequest) (*v1.DouyinUserRegisterResponse, error) {
	uid, err := s.uc.CreateUser(ctx, &do.User{
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
	token, err := jwt.CreateTokenByID(uid)
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

// Login 用户登陆
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
		token, err := jwt.CreateTokenByID(uid)
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

// GetInfo 获取用户信息
func (s *PassportService) GetInfo(ctx context.Context, req *v1.DouyinGetUserInfoRequest) (*v1.DouyinGetUserInfoResponse, error) {
	user, err := s.uc.GetUserByID(ctx, req.GetUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinGetUserInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	userInfo, err := mapper.UserToDTO(user)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinGetUserInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.DouyinGetUserInfoResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Info:       userInfo,
	}, nil
}

// MGetInfo 批量获取用户信息
func (s *PassportService) MGetInfo(ctx context.Context, req *v1.DouyinMultipleGetUserInfoRequest) (*v1.DouyinMultipleGetUserInfoResponse, error) {
	users, err := s.uc.MGetUserByID(ctx, req.GetUserIds())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinMultipleGetUserInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	userInfos, err := mapper.UserToDTOs(users)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinMultipleGetUserInfoResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.DouyinMultipleGetUserInfoResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Infos:      userInfos,
	}, nil
}
