package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	account "douyin/api/user/account/service/v1"
	passport "douyin/api/user/passport/service/v1"
)

type AccountRepo interface {
	Login(ctx context.Context, username, password string) (*passport.DouyinUserLoginResponse, error)
	Register(ctx context.Context, username, password string) (*passport.DouyinUserRegisterResponse, error)
	GetUserInfo(ctx context.Context, userId, toUserId int64) (*account.GetUserInfoByUserIdResponse, error)
}

type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUsecase) Login(ctx context.Context, username, password string) (*passport.DouyinUserLoginResponse, error) {
	res, err := uc.repo.Login(ctx, username, password)
	if err != nil {
		uc.log.Errorf("Login error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *AccountUsecase) Register(ctx context.Context, username, password string) (*passport.DouyinUserRegisterResponse, error) {
	res, err := uc.repo.Register(ctx, username, password)
	if err != nil {
		uc.log.Errorf("Register error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *AccountUsecase) GetUserInfo(ctx context.Context, userId, toUserId int64) (*account.GetUserInfoByUserIdResponse, error) {
	res, err := uc.repo.GetUserInfo(ctx, userId, toUserId)
	if err != nil {
		uc.log.Errorf("GetUserInfo error: %v", err)
		return nil, err
	}
	return res, nil
}
