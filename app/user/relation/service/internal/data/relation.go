package data

import (
	"context"
	"douyin/app/user/relation/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type RelationRepo struct {
	data *Data
	log  *log.Helper
}

func NewRelationRepo(data *Data, logger log.Logger) biz.RelationRepo {
	return &RelationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RelationRepo) GetFollowListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RelationRepo) GetFollowerListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RelationRepo) GetFriendListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}
