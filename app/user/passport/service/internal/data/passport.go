package data

import (
	"context"
	"douyin/common/ecode"
	"encoding/json"
	"errors"
	"gorm.io/gorm"

	seq "douyin/api/seq-server/service/v1"
	"douyin/app/user/passport/common/constants"
	do "douyin/app/user/passport/common/entity"
	"douyin/app/user/passport/common/mapper"
	po "douyin/app/user/passport/common/model"
	"douyin/app/user/passport/service/internal/biz"
	constants2 "douyin/common/constants"

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

// CreateUser 创建用户
func (r *passportRepo) CreateUser(ctx context.Context, user *do.User) error {
	// 获取用户ID
	uid, err := r.data.seqRPC.GetID(ctx, &seq.GetIDRequest{
		BusinessId: constants2.PassportBusinessId,
	})
	if err != nil || !uid.GetIsOk() {
		r.log.Errorf("get user id err: %v", err)
		return err
	}
	user.ID = uid.GetID()
	poUser, err := mapper.UserToPO(user)
	if err != nil {
		r.log.Errorf("user to po err: %v", err)
		return err
	}
	err = r.data.db.WithContext(ctx).Table(constants.PassportTableName).Create(poUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ecode.UserAlreadyExistErr
		}
		r.log.Errorf("create user err: %v", err)
	}
	return err
}

// GetUserByName 通过用户名获取用户
func (r *passportRepo) GetUserByName(ctx context.Context, name string) (*do.User, error) {
	user := &po.User{}
	if err := r.data.db.WithContext(ctx).Table(constants.PassportTableName).Where("name = ?", name).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ecode.UserNotExistErr
		}
		r.log.Errorf("get user from db err: %v", err)
		return nil, err
	}
	us, err := mapper.UserFromPO(user)
	if err != nil {
		return nil, err
	}
	return us, nil
}

// GetUserById 通过id获取用户
func (r *passportRepo) GetUserById(ctx context.Context, id int64) (*do.User, error) {
	key := constants.UserCacheKey(id)
	user, err := r.getUserFromCache(ctx, key)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
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
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.setUserCache(ctx, user, key)
	})
	if err != nil {
		r.log.Errorf("Fanout error: %v", err)
	}
	us, err := mapper.UserFromPO(user)
	if err != nil {
		r.log.Errorf("user from po err: %v", err)
		return nil, err
	}
	return us, nil
}

// MGetUserById 批量通过id获取用户
func (r *passportRepo) MGetUserById(ctx context.Context, ids []int64) ([]*do.User, error) {
	users, missed, err := r.batchGetUserFromCache(ctx, ids)
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
	err = r.data.cacheFan.Do(context.Background(), func(ctx context.Context) {
		r.batchSetUserCache(ctx, missedUsers)
	})
	if err != nil {
		r.log.Errorf("Fanout error: %v", err)
	}
	users = append(users, missedUsers...)
	us, err := mapper.UserFromPOs(users)
	if err != nil {
		r.log.Errorf("users from po err: %v", err)
		return nil, err
	}
	return us, nil
}

// 从缓存中获取用户
func (r *passportRepo) getUserFromCache(ctx context.Context, key string) (*po.User, error) {
	result, err := r.data.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	user := &po.User{}
	if err := json.Unmarshal([]byte(result), user); err != nil {
		return nil, err
	}
	return user, nil
}

// 批量从缓存中获取用户
func (r *passportRepo) batchGetUserFromCache(ctx context.Context, ids []int64) (res []*po.User, missed []int64, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, constants.UserCacheKey(id))
	}
	pipe := r.data.redis.Pipeline()
	for _, key := range keys {
		pipe.Get(ctx, key)
	}
	results, err := pipe.Exec(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, ids, err
	}
	res = make([]*po.User, 0, len(keys))
	missed = make([]int64, 0, len(keys))
	for i, result := range results {
		if errors.Is(result.Err(), redis.Nil) {
			missed = append(missed, ids[i])
			continue
		}
		user := &po.User{}
		if err := json.Unmarshal([]byte(result.(*redis.StringCmd).Val()), user); err != nil {
			missed = append(missed, ids[i])
			continue
		}
		res = append(res, user)
	}
	return res, missed, nil
}

// 设置用户缓存
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

// 批量设置用户缓存
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
