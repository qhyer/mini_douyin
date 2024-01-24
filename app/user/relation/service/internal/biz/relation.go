package biz

import (
	"context"
	"douyin/app/user/relation/common/constants"
	do "douyin/app/user/relation/common/entity"
	"douyin/common/ecode"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
)

type RelationRepo interface {
	RelationAction(ctx context.Context, relation *do.RelationAction) error
	GetFollowListByUserId(ctx context.Context, userId int64) ([]int64, error)
	GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error)
	GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error)
	CountFollowByUserId(ctx context.Context, userId int64) (int64, error)
	CountFollowerByUserId(ctx context.Context, userId int64) (int64, error)
	IsFollowByUserId(ctx context.Context, userId, toUserId int64) (bool, error)
	IsFollowByUserIds(ctx context.Context, userId int64, toUserIds []int64) ([]bool, error)
}

type RelationUsecase struct {
	repo RelationRepo
	log  *log.Helper
	sf   *singleflight.Group
}

func NewRelationUsecase(repo RelationRepo, logger log.Logger) *RelationUsecase {
	return &RelationUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
		sf:   &singleflight.Group{},
	}
}

func (uc *RelationUsecase) RelationAction(ctx context.Context, relation *do.RelationAction) error {
	if relation.FromUserId == relation.ToUserId {
		return ecode.RelationFollowSelfBannedErr
	}
	if relation.Type != do.RelationActionFollow && relation.Type != do.RelationActionUnFollow {
		return ecode.ParamErr
	}
	err := uc.repo.RelationAction(ctx, relation)
	if err != nil {
		return err
	}
	return nil
}

func (uc *RelationUsecase) GetFollowListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err, _ := uc.sf.Do(constants.SFFollowListKey(userId), func() (interface{}, error) {
		return uc.repo.GetFollowListByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("uc.repo.GetFollowListByUserId(%d) error(%v)", userId, err)
		return nil, err
	}
	return res.([]int64), nil
}

func (uc *RelationUsecase) GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err, _ := uc.sf.Do(constants.SFFollowerListKey(userId), func() (interface{}, error) {
		return uc.repo.GetFollowerListByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("uc.repo.GetFollowerListByUserId(%d) error(%v)", userId, err)
		return nil, err
	}
	return res.([]int64), nil
}

func (uc *RelationUsecase) GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err, _ := uc.sf.Do(constants.SFFriendListKey(userId), func() (interface{}, error) {
		return uc.repo.GetFriendListByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("uc.repo.GetFriendListByUserId(%d) error(%v)", userId, err)
		return nil, err
	}
	return res.([]int64), nil
}

func (uc *RelationUsecase) CountFollowByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err, _ := uc.sf.Do(constants.SFFollowCountKey(userId), func() (interface{}, error) {
		return uc.repo.CountFollowByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("uc.repo.CountFollowByUserId(%d) error(%v)", userId, err)
		return 0, err
	}
	return res.(int64), nil
}

func (uc *RelationUsecase) CountFollowerByUserId(ctx context.Context, userId int64) (int64, error) {
	res, err, _ := uc.sf.Do(constants.SFFollowerCountKey(userId), func() (interface{}, error) {
		return uc.repo.CountFollowerByUserId(ctx, userId)
	})
	if err != nil {
		uc.log.Errorf("uc.repo.CountFollowerByUserId(%d) error(%v)", userId, err)
		return 0, err
	}
	return res.(int64), nil
}

func (uc *RelationUsecase) IsFollowByUserId(ctx context.Context, userId, toUserId int64) (bool, error) {
	res, err, _ := uc.sf.Do(constants.SFIsFollowKey(userId, toUserId), func() (interface{}, error) {
		return uc.repo.IsFollowByUserId(ctx, userId, toUserId)
	})
	if err != nil {
		uc.log.Errorf("uc.repo.GetRalationByUserId(%d, %d) error(%v)", userId, toUserId, err)
		return false, err
	}
	return res.(bool), nil
}

func (uc *RelationUsecase) IsFollowByUserIds(ctx context.Context, userId int64, toUserIds []int64) ([]bool, error) {
	return uc.repo.IsFollowByUserIds(ctx, userId, toUserIds)
}
