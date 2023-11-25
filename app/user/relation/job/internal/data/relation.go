package data

import (
	"context"
	do "douyin/app/user/relation/common/entity"
	"douyin/app/user/relation/job/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type relationRepo struct {
	data *Data
	log  *log.Helper
}

func NewRelationRepo(data *Data, logger log.Logger) biz.RelationRepo {
	return &relationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *relationRepo) CreateRelation(ctx context.Context, relation *do.Relation) error {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepo) DeleteRelation(ctx context.Context, relation *do.Relation) error {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepo) UpdateUserFollowCount(ctx context.Context, userId int64, incr int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepo) UpdateUserFollowerCount(ctx context.Context, userId int64, incr int64) error {
	//TODO implement me
	panic("implement me")
}
