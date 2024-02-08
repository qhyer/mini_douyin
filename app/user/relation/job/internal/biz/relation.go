package biz

import (
	"context"

	"douyin/app/user/relation/common/event"

	"github.com/go-kratos/kratos/v2/log"
)

type RelationRepo interface {
	CreateRelation(ctx context.Context, relation *event.RelationAction) error
	DeleteRelation(ctx context.Context, relation *event.RelationAction) error
	UpdateUserFollowCount(ctx context.Context, userId int64, incr int64) error
	UpdateUserFollowerCount(ctx context.Context, userId int64, incr int64) error
	UpdateUserFollowTempCount(ctx context.Context, procId int, userId int64, incr int64) error
	UpdateUserFollowerTempCount(ctx context.Context, procId int, userId int64, incr int64) error
	GetUserFollowTempCount(ctx context.Context, procId int) (map[int64]int64, error)
	GetUserFollowerTempCount(ctx context.Context, procId int) (map[int64]int64, error)
	PurgeUserFollowTempCount(ctx context.Context, procId int) error
	PurgeUserFollowerTempCount(ctx context.Context, procId int) error
	BatchUpdateUserRelationStat(ctx context.Context, follow map[int64]int64, follower map[int64]int64) error
}

type RelationUsecase struct {
	repo RelationRepo
	log  *log.Helper
}

func NewRelationUsecase(repo RelationRepo, logger log.Logger) *RelationUsecase {
	return &RelationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RelationUsecase) CreateRelation(ctx context.Context, relation *event.RelationAction) error {
	return uc.repo.CreateRelation(ctx, relation)
}

func (uc *RelationUsecase) DeleteRelation(ctx context.Context, relation *event.RelationAction) error {
	return uc.repo.DeleteRelation(ctx, relation)
}

func (uc *RelationUsecase) UpdateUserFollowCount(ctx context.Context, userId int64, incr int64) error {
	return uc.repo.UpdateUserFollowCount(ctx, userId, incr)
}

func (uc *RelationUsecase) UpdateUserFollowerCount(ctx context.Context, userId int64, incr int64) error {
	return uc.repo.UpdateUserFollowerCount(ctx, userId, incr)
}

func (uc *RelationUsecase) UpdateUserFollowTempCount(ctx context.Context, procId int, userId int64, incr int64) error {
	return uc.repo.UpdateUserFollowTempCount(ctx, procId, userId, incr)
}

func (uc *RelationUsecase) UpdateUserFollowerTempCount(ctx context.Context, procId int, userId int64, incr int64) error {
	return uc.repo.UpdateUserFollowerTempCount(ctx, procId, userId, incr)
}

func (uc *RelationUsecase) GetUserFollowTempCount(ctx context.Context, procId int) (map[int64]int64, error) {
	return uc.repo.GetUserFollowTempCount(ctx, procId)
}

func (uc *RelationUsecase) GetUserFollowerTempCount(ctx context.Context, procId int) (map[int64]int64, error) {
	return uc.repo.GetUserFollowerTempCount(ctx, procId)
}

func (uc *RelationUsecase) PurgeUserFollowTempCount(ctx context.Context, procId int) error {
	return uc.repo.PurgeUserFollowTempCount(ctx, procId)
}

func (uc *RelationUsecase) PurgeUserFollowerTempCount(ctx context.Context, procId int) error {
	return uc.repo.PurgeUserFollowerTempCount(ctx, procId)
}

func (uc *RelationUsecase) BatchUpdateUserRelationStat(ctx context.Context, follow map[int64]int64, follower map[int64]int64) error {
	return uc.repo.BatchUpdateUserRelationStat(ctx, follow, follower)
}
