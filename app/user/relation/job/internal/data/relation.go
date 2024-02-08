package data

import (
	"context"
	"strconv"

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

func (r *relationRepo) UpdateUserFollowTempCount(ctx context.Context, prodId int, userId int64, incr int64) error {
	err := r.data.redis.HIncrBy(ctx, constants.UserFollowTempCountKey(prodId), strconv.FormatInt(userId, 10), incr).Err()
	if err != nil {
		r.log.Errorf("UpdateUserFollowTempCount error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) UpdateUserFollowerTempCount(ctx context.Context, prodId int, userId int64, incr int64) error {
	err := r.data.redis.HIncrBy(ctx, constants.UserFollowerTempCountKey(prodId), strconv.FormatInt(userId, 10), incr).Err()
	if err != nil {
		r.log.Errorf("UpdateUserFollowerTempCount error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) GetUserFollowTempCount(ctx context.Context, procId int) (map[int64]int64, error) {
	countMap, err := r.data.redis.HGetAll(ctx, constants.UserFollowTempCountKey(procId)).Result()
	if err != nil {
		r.log.Errorf("GetUserFollowTempCount error(%v)", err)
		return nil, err
	}
	ret := make(map[int64]int64, len(countMap))
	for k, v := range countMap {
		userId, _ := strconv.ParseInt(k, 10, 64)
		count, _ := strconv.ParseInt(v, 10, 64)
		ret[userId] = count
	}
	return ret, nil
}

func (r *relationRepo) GetUserFollowerTempCount(ctx context.Context, procId int) (map[int64]int64, error) {
	countMap, err := r.data.redis.HGetAll(ctx, constants.UserFollowerTempCountKey(procId)).Result()
	if err != nil {
		r.log.Errorf("GetUserFollowerTempCount error(%v)", err)
		return nil, err
	}
	ret := make(map[int64]int64, len(countMap))
	for k, v := range countMap {
		userId, _ := strconv.ParseInt(k, 10, 64)
		count, _ := strconv.ParseInt(v, 10, 64)
		ret[userId] = count
	}
	return ret, nil
}

func (r *relationRepo) PurgeUserFollowTempCount(ctx context.Context, procId int) error {
	err := r.data.redis.Del(ctx, constants.UserFollowTempCountKey(procId)).Err()
	if err != nil {
		r.log.Errorf("PurgeUserFollowTempCount error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) PurgeUserFollowerTempCount(ctx context.Context, procId int) error {
	err := r.data.redis.Del(ctx, constants.UserFollowerTempCountKey(procId)).Err()
	if err != nil {
		r.log.Errorf("PurgeUserFollowerTempCount error(%v)", err)
		return err
	}
	return nil
}

func (r *relationRepo) BatchUpdateUserRelationStat(ctx context.Context, follow map[int64]int64, follower map[int64]int64) error {
	err := r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for userId, incr := range follow {
			err := tx.Table(constants.RelationCountTable(userId)).Update("follow_count", gorm.Expr("follow_count + ?", incr)).Error
			if err != nil {
				r.log.Errorf("BatchUpdateUserRelationStat error(%v)", err)
				return err
			}
		}
		for userId, incr := range follower {
			err := tx.Table(constants.RelationCountTable(userId)).Update("follower_count", gorm.Expr("follower_count + ?", incr)).Error
			if err != nil {
				r.log.Errorf("BatchUpdateUserRelationStat error(%v)", err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("BatchUpdateUserRelationStat error(%v)", err)
		return err
	}
	return nil
}
