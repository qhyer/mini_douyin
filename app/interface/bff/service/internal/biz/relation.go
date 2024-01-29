package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	account "douyin/api/user/account/service/v1"
	relation "douyin/api/user/relation/service/v1"
	do "douyin/app/user/relation/common/entity"
	"douyin/common/ecode"
)

type RelationRepo interface {
	RelationAction(ctx context.Context, userId, toUserId int64, actionType int32) (*relation.RelationActionResponse, error)
	GetUserFollowList(ctx context.Context, userId, toUserId int64) (*account.GetFollowListByUserIdResponse, error)
	GetUserFollowerList(ctx context.Context, userId, toUserId int64) (*account.GetFollowerListByUserIdResponse, error)
	GetUserFriendList(ctx context.Context, userId int64) (*account.GetFriendListByUserIdResponse, error)
}

type RelationUsecase struct {
	repo RelationRepo
	log  *log.Helper
}

func NewRelationUsecase(repo RelationRepo, logger log.Logger) *RelationUsecase {
	return &RelationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RelationUsecase) RelationAction(ctx context.Context, userId, toUserId int64, actionType int32) (*relation.RelationActionResponse, error) {
	if userId == toUserId {
		return nil, ecode.RelationFollowSelfBannedErr
	}
	if !isRelationActionTypeValid(actionType) {
		return nil, ecode.RelationActionTypeInvalidErr
	}

	res, err := uc.repo.RelationAction(ctx, userId, toUserId, actionType)
	if err != nil {
		uc.log.Errorf("RelationAction error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *RelationUsecase) GetUserFollowList(ctx context.Context, userId, toUserId int64) (*account.GetFollowListByUserIdResponse, error) {
	res, err := uc.repo.GetUserFollowList(ctx, userId, toUserId)
	if err != nil {
		uc.log.Errorf("GetUserFollowList error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *RelationUsecase) GetUserFollowerList(ctx context.Context, userId, toUserId int64) (*account.GetFollowerListByUserIdResponse, error) {
	res, err := uc.repo.GetUserFollowerList(ctx, userId, toUserId)
	if err != nil {
		uc.log.Errorf("GetUserFollowerList error: %v", err)
		return nil, err
	}
	return res, nil
}

func (uc *RelationUsecase) GetUserFriendList(ctx context.Context, userId int64) (*account.GetFriendListByUserIdResponse, error) {
	res, err := uc.repo.GetUserFriendList(ctx, userId)
	if err != nil {
		uc.log.Errorf("GetUserFriendList error: %v", err)
		return nil, err
	}
	return res, nil
}

func isRelationActionTypeValid(actionType int32) bool {
	if actionType == int32(do.RelationActionFollow) || actionType == int32(do.RelationActionUnFollow) {
		return true
	}
	return false
}
