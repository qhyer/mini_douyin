package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	do "douyin/app/user/account/common/entity"
)

type AccountRepo interface {
	GetUserInfoByUserId(ctx context.Context, userId int64, toUserId int64) (*do.User, error)
	MGetUserInfoByUserId(ctx context.Context, toUserIds []int64) ([]*do.User, error)
	GetFollowListByUserId(ctx context.Context, userId int64, toUserId int64) ([]*do.User, error)
	GetFollowerListByUserId(ctx context.Context, userId int64, toUserId int64) ([]*do.User, error)
	GetFriendListByUserId(ctx context.Context, userId int64) ([]*do.User, error)
}

type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

// GetUserInfoByUserId 获取用户的详细信息
func (uc *AccountUsecase) GetUserInfoByUserId(ctx context.Context, userId int64, toUserId int64) (*do.User, error) {
	return uc.repo.GetUserInfoByUserId(ctx, userId, toUserId)
}

// MGetUserInfoByUserId 获取用户的基本信息
func (uc *AccountUsecase) MGetUserInfoByUserId(ctx context.Context, toUserIds []int64) ([]*do.User, error) {
	return uc.repo.MGetUserInfoByUserId(ctx, toUserIds)
}

// GetFollowListByUserId 获取用户的关注列表
func (uc *AccountUsecase) GetFollowListByUserId(ctx context.Context, userId int64, toUserId int64) ([]*do.User, error) {
	return uc.repo.GetFollowListByUserId(ctx, userId, toUserId)
}

// GetFollowerListByUserId 获取用户的粉丝列表
func (uc *AccountUsecase) GetFollowerListByUserId(ctx context.Context, userId int64, toUserId int64) ([]*do.User, error) {
	return uc.repo.GetFollowerListByUserId(ctx, userId, toUserId)
}

// GetFriendListByUserId 获取用户的好友列表
func (uc *AccountUsecase) GetFriendListByUserId(ctx context.Context, userId int64) ([]*do.User, error) {
	return uc.repo.GetFriendListByUserId(ctx, userId)
}
