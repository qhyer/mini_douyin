package data

import (
	"context"
	account "douyin/api/user/account/service/v1"
	relation "douyin/api/user/relation/service/v1"
	"douyin/app/interface/bff/service/internal/biz"
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

func (r *relationRepo) RelationAction(ctx context.Context, userId, toUserId int64, actionType int32) (*relation.RelationActionResponse, error) {
	res, err := r.data.RelationRPC.RelationAction(ctx, &relation.RelationActionRequest{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: actionType,
	})
	return res, err
}

func (r *relationRepo) GetUserFollowList(ctx context.Context, userId, toUserId int64) (*account.GetFollowListByUserIdResponse, error) {
	res, err := r.data.AccountRPC.GetFollowListByUserId(ctx, &account.GetFollowListByUserIdRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	return res, err
}

func (r *relationRepo) GetUserFollowerList(ctx context.Context, userId, toUserId int64) (*account.GetFollowerListByUserIdResponse, error) {
	res, err := r.data.AccountRPC.GetFollowerListByUserId(ctx, &account.GetFollowerListByUserIdRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	return res, err
}

func (r *relationRepo) GetUserFriendList(ctx context.Context, userId int64) (*account.GetFriendListByUserIdResponse, error) {
	res, err := r.data.AccountRPC.GetFriendListByUserId(ctx, &account.GetFriendListByUserIdRequest{
		UserId: userId,
	})
	return res, err
}
