package data

import (
	"context"
	seq "douyin/api/seq-server/service/v1"
	"douyin/app/user/relation/common/constants"
	do "douyin/app/user/relation/common/entity"
	"douyin/app/user/relation/common/mapper"
	po "douyin/app/user/relation/common/model"
	"douyin/app/user/relation/job/internal/biz"
	constants2 "douyin/common/constants"
	"github.com/IBM/sarama"
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

func (r *relationRepo) CreateRelation(ctx context.Context, relation *do.RelationAction) error {
	// 获取关系ID
	rid, err := r.data.seqRPC.GetID(ctx, &seq.GetIDRequest{
		BusinessId: constants2.RelationBusinessId,
	})
	if err != nil || !rid.GetIsOk() {
		r.log.Errorf("seq rpc error: %v", err)
		return err
	}
	relation.ID = rid.GetID()
	rel, err := mapper.RelationActionToPO(relation)
	if err != nil {
		r.log.Errorf("mapper.RelationToPO error(%v)", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建关注，如果已经存在则更新
		usRel := po.Relation{}
		err := tx.WithContext(ctx).Table(constants.FollowRecordTable(rel.FromUserId)).Where(po.Relation{
			FromUserId: rel.FromUserId,
			ToUserId:   rel.ToUserId,
		}).FirstOrInit(&usRel, po.Relation{
			ID:         rel.ID,
			FromUserId: rel.FromUserId,
			ToUserId:   rel.ToUserId,
			CreatedAt:  rel.CreatedAt,
		}).Update("attr", do.SetAttr(usRel.Attr, do.RelationAttrFollowing)).Error
		if err != nil {
			return err
		}
		// 创建粉丝
		err = tx.WithContext(ctx).Table(constants.FollowerRecordTable(rel.ToUserId)).Where(po.Relation{
			FromUserId: rel.ToUserId,
			ToUserId:   rel.FromUserId,
		}).FirstOrInit(&usRel, po.Relation{
			ID:         rel.ID,
			FromUserId: rel.ToUserId,
			ToUserId:   rel.FromUserId,
			CreatedAt:  rel.CreatedAt,
		}).Update("attr", do.SetAttr(usRel.Attr, do.RelationAttrFollowed)).Error
		if err != nil {
			return err
		}

		// 更新关注数和粉丝数
		err = r.data.kafkaProducer.SendMessages([]*sarama.ProducerMessage{
			{
				Topic: constants.UpdateUserFollowCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowCountKafkaKey(rel.FromUserId)),
				Value: sarama.StringEncoder("1"),
			},
			{
				Topic: constants.UpdateUserFollowerCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowerCountKafkaKey(rel.ToUserId)),
				Value: sarama.StringEncoder("1"),
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

func (r *relationRepo) DeleteRelation(ctx context.Context, relation *do.RelationAction) error {
	rel, err := mapper.RelationActionToPO(relation)
	if err != nil {
		r.log.Errorf("mapper.RelationToPO error(%v)", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除关注，如果已经存在则更新
		usRel := po.Relation{}
		err := tx.WithContext(ctx).Table(constants.FollowRecordTable(rel.FromUserId)).Where(po.Relation{
			FromUserId: rel.FromUserId,
			ToUserId:   rel.ToUserId,
		}).FirstOrInit(&usRel, po.Relation{
			ID:         rel.ID,
			FromUserId: rel.FromUserId,
			ToUserId:   rel.ToUserId,
			CreatedAt:  rel.CreatedAt,
		}).Update("attr", do.UnsetAttr(usRel.Attr, do.RelationAttrFollowing)).Error
		if err != nil {
			return err
		}
		// 删除粉丝
		err = tx.WithContext(ctx).Table(constants.FollowerRecordTable(rel.ToUserId)).Where(po.Relation{
			FromUserId: rel.ToUserId,
			ToUserId:   rel.FromUserId,
		}).FirstOrInit(&usRel, po.Relation{
			ID:         rel.ID,
			FromUserId: rel.ToUserId,
			ToUserId:   rel.FromUserId,
			CreatedAt:  rel.CreatedAt,
		}).Update("attr", do.UnsetAttr(usRel.Attr, do.RelationAttrFollowed)).Error
		if err != nil {
			return err
		}

		// 更新关注数和粉丝数
		err = r.data.kafkaProducer.SendMessages([]*sarama.ProducerMessage{
			{
				Topic: constants.UpdateUserFollowCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowCountKafkaKey(rel.FromUserId)),
				Value: sarama.StringEncoder("-1"),
			},
			{
				Topic: constants.UpdateUserFollowerCountTopic,
				Key:   sarama.StringEncoder(constants.UpdateUserFollowerCountKafkaKey(rel.ToUserId)),
				Value: sarama.StringEncoder("-1"),
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
