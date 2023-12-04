package data

import (
	"context"
	"douyin/app/user/relation/common/constants"
	do "douyin/app/user/relation/common/entity"
	"douyin/app/user/relation/common/mapper"
	po "douyin/app/user/relation/common/model"
	"douyin/app/user/relation/job/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	rel, err := mapper.RelationToPO(relation)
	if err != nil {
		r.log.Errorf("mapper.RelationToPO error(%v)", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建关注
		err := tx.WithContext(ctx).Table(constants.FollowRecordTable(rel.FromUserId)).Create(rel).Error
		if err != nil {
			return err
		}
		// 创建粉丝
		err = tx.WithContext(ctx).Table(constants.FollowerRecordTable(rel.ToUserId)).Create(rel).Error
		if err != nil {
			return err
		}
		// todo 更新关注数和粉丝数

		return nil
	})
	if err != nil {
		r.log.Errorf("CreateRelation error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) DeleteRelation(ctx context.Context, relation *do.Relation) error {
	//TODO implement me
	panic("implement me")
}

func (r *relationRepo) UpdateUserFollowCount(ctx context.Context, userId int64, incr int64) error {
	userCnt := po.UserRelationCount{UserId: userId}
	err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).FirstOrInit(&userCnt, userCnt).Update("follow_count", gorm.Expr("follow_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFollowCount error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) UpdateUserFollowerCount(ctx context.Context, userId int64, incr int64) error {
	userCnt := po.UserRelationCount{UserId: userId}
	err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).FirstOrInit(&userCnt, userCnt).Update("follower_count", gorm.Expr("follower_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFollowerCount error(%v)", err)
		return err
	}
	return nil
}
