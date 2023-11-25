package biz

import (
	"context"
	do "douyin/app/user/account/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type PassportRepo interface {
	GetUserInfoByUserId(ctx context.Context, userId int64, toUserId int64) (*do.User, error)
	MGetUserInfoByUserId(ctx context.Context, userId int64, toUserIds []int64) ([]*do.User, error)
}

type RelationRepo interface {
	GetFollowListByUserId(ctx context.Context, userId int64) ([]*do.User, error)
	GetFollowerListByUserId(ctx context.Context, userId int64) ([]*do.User, error)
	GetFriendListByUserId(ctx context.Context, userId int64) ([]*do.User, error)
	GetRelationByUserId(ctx context.Context, userId int64, toUserIds []int64) (*do.Relation, error)
}

type AccountUsecase struct {
	repo PassportRepo
	log  *log.Helper
}

func NewAccountUsecase(repo PassportRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

// GetUserInfoByUserId 获取用户的详细信息
func (uc *AccountUsecase) GetUserInfoByUserId(ctx context.Context, userId int64, toUserId int64) (*do.User, error) {
	return uc.repo.GetUserInfoByUserId(ctx, userId, toUserId)
}

// MGetUserInfoByUserId 获取用户的基本信息
func (uc *AccountUsecase) MGetUserInfoByUserId(ctx context.Context, userId int64, toUserIds []int64) ([]*do.User, error) {
	return uc.repo.MGetUserInfoByUserId(ctx, userId, toUserIds)
}

func (uc *AccountUsecase) GetFollowListByUserId(ctx context.Context, userId int64) ([]*do.User, error) {
	return uc.repo.GetFollowListByUserId(ctx, userId)
}

func (uc *AccountUsecase) GetFollowerListByUserId(ctx context.Context, userId int64) ([]*do.User, error) {
	return uc.repo.GetFollowerListByUserId(ctx, userId)
}

func (uc *AccountUsecase) GetFriendListByUserId(ctx context.Context, userId int64) ([]*do.User, error) {
	return uc.repo.GetFriendListByUserId(ctx, userId)
}
