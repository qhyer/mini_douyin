package data

import (
	"context"
	"douyin/app/user/relation/common/event"
	"douyin/common/ecode"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"douyin/app/user/relation/common/constants"
	do "douyin/app/user/relation/common/entity"
	"douyin/app/user/relation/common/mapper"
	po "douyin/app/user/relation/common/model"
	"douyin/app/user/relation/job/internal/biz"
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

func (r *relationRepo) CreateRelation(ctx context.Context, relation *event.RelationAction) error {
	rel, err := mapper.RelationActionToPO(relation)
	if err != nil {
		r.log.Errorf("mapper.RelationToPO error(%v)", err)
		return err
	}
	relationActionBytes, err := relation.MarshalJson()
	if err != nil {
		r.log.Errorf("relation.MarshalJson error(%v)", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建关注，如果已经存在则更新
		followRel := po.Relation{}
		followRes := tx.WithContext(ctx).Table(constants.FollowRecordTable(rel.FromUserId)).Where(po.Relation{
			FromUserId: rel.FromUserId,
			ToUserId:   rel.ToUserId,
		}).Assign(po.Relation{ID: rel.ID}).FirstOrCreate(&followRel)
		if followRes.Error != nil {
			return followRes.Error
		}

		if do.SetAttr(followRel.Attr, do.RelationAttrFollowing) == followRel.Attr {
			return ecode.RelationNotModifyErr
		} else {
			followRes = followRes.Update("attr", do.SetAttr(followRel.Attr, do.RelationAttrFollowing))
		}

		if followRes.Error != nil {
			return followRes.Error
		}

		// 创建粉丝
		followerRel := po.Relation{}
		followerRes := tx.WithContext(ctx).Table(constants.FollowerRecordTable(rel.ToUserId)).Where(po.Relation{
			FromUserId: rel.ToUserId,
			ToUserId:   rel.FromUserId,
		}).Assign(po.Relation{ID: rel.ID}).FirstOrCreate(&followerRel)
		if followerRes.Error != nil {
			return followerRes.Error
		}
		if do.SetAttr(followerRel.Attr, do.RelationAttrFollowed) == followerRel.Attr {
			return ecode.RelationNotModifyErr
		} else {
			followerRes = followerRes.Update("attr", do.SetAttr(followerRel.Attr, do.RelationAttrFollowed))
		}
		if followerRes.Error != nil {
			return followerRes.Error
		}

		// 更新关注数和粉丝数
		err = r.data.kafkaProducer.SendMessages([]*sarama.ProducerMessage{
			{
				Topic: constants.UpdateUserFollowCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowCountKafkaKey(rel.FromUserId)),
				Value: sarama.ByteEncoder(relationActionBytes),
			},
			{
				Topic: constants.UpdateUserFollowerCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowerCountKafkaKey(rel.ToUserId)),
				Value: sarama.ByteEncoder(relationActionBytes),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("CreateRelation error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) DeleteRelation(ctx context.Context, relation *event.RelationAction) error {
	rel, err := mapper.RelationActionToPO(relation)
	if err != nil {
		r.log.Errorf("mapper.RelationToPO error(%v)", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除关注，如果已经存在则更新
		followRel := po.Relation{}
		followRes := tx.WithContext(ctx).Table(constants.FollowRecordTable(rel.FromUserId)).Where(po.Relation{
			FromUserId: rel.FromUserId,
			ToUserId:   rel.ToUserId,
		}).Update("attr", do.UnsetAttr(followRel.Attr, do.RelationAttrFollowing))
		if followRes.Error != nil {
			return followRes.Error
		}
		if followRes.RowsAffected == 0 {
			return ecode.RelationNotModifyErr
		}
		if followRes.Error != nil {
			return followRes.Error
		}

		// 删除粉丝
		followerRel := po.Relation{}
		followerRes := tx.WithContext(ctx).Table(constants.FollowerRecordTable(rel.ToUserId)).Where(po.Relation{
			FromUserId: rel.ToUserId,
			ToUserId:   rel.FromUserId,
		}).Update("attr", do.UnsetAttr(followerRel.Attr, do.RelationAttrFollowed))
		if followerRes.Error != nil {
			return followerRes.Error
		}

		if followerRes.RowsAffected == 0 {
			return ecode.RelationNotModifyErr
		}

		relationActionBytes, err := relation.MarshalJson()
		if err != nil {
			r.log.Errorf("relation.MarshalJson error(%v)", err)
			return err
		}

		// 更新关注数和粉丝数
		err = r.data.kafkaProducer.SendMessages([]*sarama.ProducerMessage{
			{
				Topic: constants.UpdateUserFollowCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowCountKafkaKey(rel.FromUserId)),
				Value: sarama.ByteEncoder(relationActionBytes),
			},
			{
				Topic: constants.UpdateUserFollowerCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowerCountKafkaKey(rel.ToUserId)),
				Value: sarama.ByteEncoder(relationActionBytes),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("DeleteRelation error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) UpdateUserFollowCount(ctx context.Context, userId int64, incr int64) error {
	userCnt := po.UserRelationCount{UserId: userId}
	err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).FirstOrCreate(&userCnt, userCnt).Update("follow_count", gorm.Expr("follow_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFollowCount error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) UpdateUserFollowerCount(ctx context.Context, userId int64, incr int64) error {
	userCnt := po.UserRelationCount{UserId: userId}
	err := r.data.db.WithContext(ctx).Table(constants.RelationCountTable(userId)).FirstOrCreate(&userCnt, userCnt).Update("follower_count", gorm.Expr("follower_count + ?", incr)).Error
	if err != nil {
		r.log.Errorf("UpdateUserFollowerCount error(%v)", err)
		return err
	}
	return nil
}
