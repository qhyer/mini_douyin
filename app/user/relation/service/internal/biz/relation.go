package biz

import (
	"context"
	do "douyin/app/user/relation/common/entity"
	"douyin/common/ecode"
	"github.com/go-kratos/kratos/v2/log"
)

type RelationRepo interface {
	RelationAction(ctx context.Context, relation *do.Relation) error
	GetFollowListByUserId(ctx context.Context, userId int64) ([]int64, error)
	GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error)
	GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error)
	CountFollowByUserId(ctx context.Context, userId int64) (int64, error)
	CountFollowerByUserId(ctx context.Context, userId int64) (int64, error)
}

type RelationUsecase struct {
	repo RelationRepo
	log  *log.Helper
}

func NewRelationUsecase(repo RelationRepo, logger log.Logger) *RelationUsecase {
	return &RelationUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *RelationUsecase) RelationAction(ctx context.Context, relation *do.Relation) error {
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
	return uc.repo.GetFollowListByUserId(ctx, userId)
}

func (uc *RelationUsecase) GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	return uc.repo.GetFollowerListByUserId(ctx, userId)
}

func (uc *RelationUsecase) GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	return uc.repo.GetFriendListByUserId(ctx, userId)
}

func (uc *RelationUsecase) CountFollowByUserId(ctx context.Context, userId int64) (int64, error) {
	return uc.repo.CountFollowByUserId(ctx, userId)
}

func (uc *RelationUsecase) CountFollowerByUserId(ctx context.Context, userId int64) (int64, error) {
	return uc.repo.CountFollowerByUserId(ctx, userId)
}
