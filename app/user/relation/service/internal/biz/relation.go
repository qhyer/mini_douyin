package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RelationRepo interface {
	GetFollowListByUserId(ctx context.Context, userId int64) ([]int64, error)
	GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error)
	GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error)
}

type RelationUsecase struct {
	repo *RelationRepo
	log  *log.Helper
}

func NewRelationUsecase(repo *RelationRepo, logger log.Logger) *RelationUsecase {
	return &RelationUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
