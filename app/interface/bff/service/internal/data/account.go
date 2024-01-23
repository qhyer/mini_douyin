package data

import (
	"context"
	account "douyin/api/user/account/service/v1"
	passport "douyin/api/user/passport/service/v1"
	"douyin/app/interface/bff/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) Login(ctx context.Context, username, password string) (*passport.DouyinUserLoginResponse, error) {
	res, err := r.data.PassportRPC.Login(ctx, &passport.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})
	return res, err
}

func (r *accountRepo) Register(ctx context.Context, username, password string) (*passport.DouyinUserRegisterResponse, error) {
	res, err := r.data.PassportRPC.Register(ctx, &passport.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	return res, err
}

func (r *accountRepo) GetUserInfo(ctx context.Context, userId, toUserId int64) (*account.GetUserInfoByUserIdResponse, error) {
	res, err := r.data.AccountRPC.GetUserInfoByUserId(ctx, &account.GetUserInfoByUserIdRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	return res, err
}
