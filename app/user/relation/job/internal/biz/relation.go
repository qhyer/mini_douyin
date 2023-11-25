package biz

import (
	"context"
	do "douyin/app/user/relation/common/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type RelationRepo interface {
	CreateRelation(ctx context.Context, relation *do.Relation) error
	DeleteRelation(ctx context.Context, relation *do.Relation) error
	UpdateUserFollowCount(ctx context.Context, userId int64, incr int64) error
	UpdateUserFollowerCount(ctx context.Context, userId int64, incr int64) error
}

type RelationUsecase struct {
	repo RelationRepo
	log  *log.Helper
}

func NewRelationUsecase(repo RelationRepo, logger log.Logger) *RelationUsecase {
	return &RelationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RelationUsecase) CreateRelation(ctx context.Context, relation *do.Relation) error {
	return uc.repo.CreateRelation(ctx, relation)
}

func (uc *RelationUsecase) DeleteRelation(ctx context.Context, relation *do.Relation) error {
	return uc.repo.DeleteRelation(ctx, relation)
}

func (uc *RelationUsecase) UpdateUserFollowCount(ctx context.Context, userId int64, incr int64) error {
	return uc.repo.UpdateUserFollowCount(ctx, userId, incr)
}

func (uc *RelationUsecase) UpdateUserFollowerCount(ctx context.Context, userId int64, incr int64) error {
	return uc.repo.UpdateUserFollowerCount(ctx, userId, incr)
}
