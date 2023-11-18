package data

import (
	"context"
	"douyin/app/user/passport/common/constants"
	do "douyin/app/user/passport/common/entity"
	"douyin/app/user/passport/common/mapper"
	po "douyin/app/user/passport/common/model"
	"douyin/app/user/passport/service/internal/biz"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type passportRepo struct {
	data *Data
	log  *log.Helper
}

func NewPassportRepo(data *Data, logger log.Logger) biz.PassportRepo {
	return &passportRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *passportRepo) CreateUser(ctx context.Context, user *do.User) error {
	poUser, err := mapper.UserToPO(user)
	if err != nil {
		return err
	}
	err = r.data.db.WithContext(ctx).Table(constants.PassportTableName).Create(poUser).Error
	if err != nil {
		r.log.Errorf("create user err: %v", err)
	}
	return err
}

func (r *passportRepo) GetUserByName(ctx context.Context, name string) (*do.User, error) {
	user := &po.User{}
	if err := r.data.db.WithContext(ctx).Table(constants.PassportTableName).Where("name = ?", name).First(user).Error; err != nil {
		r.log.Errorf("get user from db err: %v", err)
		return nil, err
	}
	us, err := mapper.UserFromPO(user)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (r *passportRepo) GetUserById(ctx context.Context, id int64) (*do.User, error) {
	key := constants.UserCacheKey(id)
	user, err := r.getUserFromCache(ctx, key)
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("get user from cache err: %v", err)
		}
	} else {
		us, err := mapper.UserFromPO(user)
		if err != nil {
			r.log.Errorf("user from po err: %v", err)
			return nil, err
		}
		return us, nil
	}
	user = &po.User{}
	if err := r.data.db.WithContext(ctx).Table(constants.PassportTableName).Where("id = ?", id).First(user).Error; err != nil {
		r.log.Errorf("get user from db err: %v", err)
		return nil, err
	}
	r.setUserCache(ctx, user, key)
	us, err := mapper.UserFromPO(user)
	if err != nil {
		r.log.Errorf("user from po err: %v", err)
		return nil, err
	}
	return us, nil
}

func (r *passportRepo) MGetUserById(ctx context.Context, ids []int64) ([]*do.User, error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, constants.UserCacheKey(id))
	}
	users, missed, err := r.batchGetUserCache(ctx, keys)
	if err != nil {
		r.log.Errorf("batch get user cache err: %v", err)
		return nil, err
	}
	if len(missed) == 0 {
		us, err := mapper.UserFromPOs(users)
		if err != nil {
			r.log.Errorf("users from po err: %v", err)
			return nil, err
		}
		return us, nil
	}
	missedUsers := make([]*po.User, 0, len(missed))
	if err := r.data.db.WithContext(ctx).Table(constants.PassportTableName).Where("id in (?)", missed).Find(&missedUsers).Error; err != nil {
		r.log.Errorf("get user from db err: %v", err)
		return nil, err
	}
	r.batchSetUserCache(ctx, missedUsers)
	users = append(users, missedUsers...)
	us, err := mapper.UserFromPOs(users)
	if err != nil {
		r.log.Errorf("users from po err: %v", err)
		return nil, err
	}
	return us, nil
}

func (r *passportRepo) getUserFromCache(ctx context.Context, key string) (*po.User, error) {
	result, err := r.data.redis.Get(ctx, key).Result()
	if err != nil {
		r.log.Errorf("get user from cache err: %v", err)
		return nil, err
	}
	user := &po.User{}
	if err := json.Unmarshal([]byte(result), user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *passportRepo) batchGetUserCache(ctx context.Context, keys []string) (res []*po.User, missed []string, err error) {
	pipe := r.data.redis.Pipeline()
	for _, key := range keys {
		pipe.Get(ctx, key)
	}
	results, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("batch get user cache err: %v", err)
		return nil, keys, err
	}
	res = make([]*po.User, 0, len(keys))
	missed = make([]string, 0, len(keys))
	for i, result := range results {
		if result.Err() == redis.Nil {
			missed = append(missed, keys[i])
			continue
		}
		user := &po.User{}
		if err := json.Unmarshal([]byte(result.(*redis.StringCmd).Val()), user); err != nil {
			missed = append(missed, keys[i])
			continue
		}
		res = append(res, user)
	}
	return res, missed, nil
}

func (r *passportRepo) setUserCache(ctx context.Context, user *po.User, key string) {
	bytes, err := json.Marshal(user)
	if err != nil {
		r.log.Errorf("marshal user err: %v", err)
		return
	}
	if err := r.data.redis.Set(ctx, key, bytes, constants.UserCacheExpiration).Err(); err != nil {
		r.log.Errorf("set user cache err: %v", err)
		return
	}
}

func (r *passportRepo) batchSetUserCache(ctx context.Context, users []*po.User) {
	pipe := r.data.redis.Pipeline()
	for _, user := range users {
		key := constants.UserCacheKey(user.ID)
		bytes, err := json.Marshal(user)
		if err != nil {
			r.log.Errorf("marshal user err: %v", err)
			continue
		}
		pipe.Set(ctx, key, bytes, constants.UserCacheExpiration)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		r.log.Errorf("batch set user cache err: %v", err)
		return
	}
}
