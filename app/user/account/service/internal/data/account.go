package data

import (
	"context"
	"encoding/json"
	"errors"

	passportRPC "douyin/api/user/passport/service/v1"
	relationRPC "douyin/api/user/relation/service/v1"
	favoriteRPC "douyin/api/video/favorite/service/v1"
	publishRPC "douyin/api/video/publish/service/v1"
	"douyin/app/user/account/common/constants"
	do "douyin/app/user/account/common/entity"
	"douyin/app/user/account/common/mapper"
	"douyin/app/user/account/service/internal/biz"
	constants2 "douyin/common/constants"
	"douyin/common/ecode"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetUserInfoByUserId 通过用户ID获取用户信息
func (r *accountRepo) GetUserInfoByUserId(ctx context.Context, userId int64, toUserId int64) (*do.User, error) {
	user, err := r.getUserInfoFromCache(ctx, constants.AccountInfoCacheKey(toUserId))
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("redis get user info err: %v", err)
			return nil, err
		}

		g := errgroup.Group{}
		usFavCnt, usFavdCnt := int64(0), int64(0)
		usWorkCnt := int64(0)
		usFollowCnt, usFollowerCnt := int64(0), int64(0)
		isFollow := false
		// 获取用户信息
		g.Go(func() error {
			user, err = r.getUserInfoFromPassportRPC(ctx, toUserId)
			return err
		})
		// 获取点赞数和被点赞数
		g.Go(func() error {
			usFavCnt, usFavdCnt, err = r.getUserFavoriteCountFromFavoriteRPC(ctx, toUserId)
			return err
		})
		// 获取作品数量
		g.Go(func() error {
			usWorkCnt, err = r.getUserWorkCountFromPublishRPC(ctx, toUserId)
			return err
		})
		// 获取关注数和粉丝数
		g.Go(func() error {
			usFollowCnt, usFollowerCnt, err = r.getUserFollowCountAndFollowerCountFromRelationRPC(ctx, toUserId)
			return err
		})
		// 获取是否关注
		g.Go(func() error {
			isFollow, err = r.getIsFollowFromRelationRPC(ctx, userId, toUserId)
			return err
		})

		err = g.Wait()
		if err != nil {
			r.log.Errorf("get user info err: %v", err)
			return nil, err
		}
		user.FavoriteCount = usFavCnt
		user.TotalFavorited = usFavdCnt
		user.FollowCount = usFollowCnt
		user.FollowerCount = usFollowerCnt
		user.WorkCount = usWorkCnt
		user.IsFollow = isFollow
		err := r.data.cacheFan.Do(ctx, func(ctx context.Context) {
			r.setUserInfoCache(ctx, user)
		})
		if err != nil {
			r.log.Errorf("Fanout err: %v", err)
		}
	}
	return user, nil
}

// MGetUserInfoByUserId 批量获取用户信息（只获取基本信息）
func (r *accountRepo) MGetUserInfoByUserId(ctx context.Context, userId int64, userIds []int64) ([]*do.User, error) {
	// 目前只转发请求到passport和relation，不做缓存
	g := errgroup.Group{}
	users := make([]*do.User, 0, len(userIds))
	isFollows := make([]bool, 0, len(userIds))
	// 获取用户信息
	g.Go(func() error {
		us, err := r.batchGetUserInfoByUserIdFromPassportRPC(ctx, userIds)
		if err != nil {
			return err
		}
		users = append(users, us...)
		return nil
	})
	// 如果是登陆用户，获取是否关注
	if userId != constants2.GuestUserID {
		g.Go(func() error {
			isFol, err := r.batchGetIsFollowByUserIdFromRelationRPC(ctx, userId, userIds)
			if err != nil {
				return err
			}
			isFollows = append(isFollows, isFol...)
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		r.log.Errorf("batch get user info err: %v", err)
		return nil, err
	}
	for i, isFollow := range isFollows {
		users[i].IsFollow = isFollow
	}
	return users, nil
}

// 设置用户信息缓存
func (r *accountRepo) setUserInfoCache(ctx context.Context, user *do.User) {
	b, err := user.MarshalJson()
	if err != nil {
		r.log.Errorf("redis marshal user info err: %v", err)
		return
	}
	// 如果粉丝数大于最小粉丝数，设置本地缓存
	if user.FollowerCount > constants.LeastFollowerCount {
		err := r.data.localCache.SetWithExpire(constants.AccountInfoCacheKey(user.ID), user, constants.AccountInfoCacheExpiration)
		if err != nil {
			r.log.Errorf("local cache set user info err: %v", err)
		}
	}
	err = r.data.redis.Set(ctx, constants.AccountInfoCacheKey(user.ID), b, constants.AccountInfoCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis set user info err: %v", err)
	}
}

// 从缓存中获取用户信息
func (r *accountRepo) getUserInfoFromCache(ctx context.Context, userId string) (*do.User, error) {
	// 先从本地缓存读取
	us, err := r.data.localCache.Get(userId)
	if err == nil {
		return us.(*do.User), nil
	}
	// 本地缓存没有，从redis读取
	res, err := r.data.redis.Get(ctx, userId).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("redis get user info err: %v", err)
		}
		return nil, err
	}
	user := &do.User{}
	err = user.UnmarshalJson(res)
	if err != nil {
		r.log.Errorf("redis unmarshal user info err: %v", err)
		return nil, err
	}
	return user, nil
}

// 从passport rpc获取用户信息
func (r *accountRepo) getUserInfoFromPassportRPC(ctx context.Context, userId int64) (*do.User, error) {
	res, err := r.data.passportCli.GetInfo(ctx, &passportRPC.DouyinGetUserInfoRequest{UserId: userId})
	if err != nil {
		r.log.Errorf("passport rpc get user info err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("passport rpc get user info err: %v", res.GetStatusCode())
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Passport rpc get userinfo failed")
	}
	user, err := mapper.UserFromPassportDTO(res.GetInfo())
	if err != nil {
		r.log.Errorf("passport rpc get user info err: %v", err)
		return nil, err
	}
	return user, nil
}

// 从favorite rpc获取用户点赞数和被点赞数
func (r *accountRepo) getUserFavoriteCountFromFavoriteRPC(ctx context.Context, userId int64) (favCnt int64, favdCnt int64, err error) {
	g := errgroup.Group{}
	g.Go(func() error {
		res, err := r.data.favoriteCli.CountUserFavoriteByUserId(ctx, &favoriteRPC.CountUserFavoriteByUserIdRequest{UserId: userId})
		if err != nil {
			return err
		}
		if res.GetStatusCode() != ecode.Success.ErrCode {
			return ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Favorite rpc count user favorite failed")
		}
		favCnt = res.GetCount()
		return nil
	})
	g.Go(func() error {
		res, err := r.data.favoriteCli.CountUserFavoritedByUserId(ctx, &favoriteRPC.CountUserFavoritedByUserIdRequest{UserId: userId})
		if err != nil {
			return err
		}
		if res.GetStatusCode() != ecode.Success.ErrCode {
			return ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Favorite rpc count user favorited failed")
		}
		favdCnt = res.GetCount()
		return nil
	})
	err = g.Wait()
	if err != nil {
		r.log.Errorf("favorite rpc get user favorite count err: %v", err)
		return 0, 0, err
	}
	return favCnt, favdCnt, nil
}

// 从publish rpc获取用户作品数
func (r *accountRepo) getUserWorkCountFromPublishRPC(ctx context.Context, userId int64) (workCnt int64, err error) {
	res, err := r.data.publishCli.CountUserPublishedVideoByUserId(ctx, &publishRPC.CountUserPublishedVideoByUserIdRequest{UserId: userId})
	if err != nil {
		r.log.Errorf("publish rpc count user publish err: %v", err)
		return 0, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("publish rpc count user publish err: %v", res.GetStatusCode())
		return 0, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Publish rpc count user publish failed")
	}
	workCnt = res.GetCount()
	return workCnt, nil
}

// 从relation rpc获取用户关注数和粉丝数
func (r *accountRepo) getUserFollowCountAndFollowerCountFromRelationRPC(ctx context.Context, userId int64) (followCnt int64, followerCnt int64, err error) {
	g := errgroup.Group{}
	g.Go(func() error {
		res, err := r.data.relationCli.CountFollowByUserId(ctx, &relationRPC.CountFollowByUserIdRequest{UserId: userId})
		if err != nil {
			return err
		}
		if res.GetStatusCode() != ecode.Success.ErrCode {
			return ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Relation rpc count user follow failed")
		}
		followCnt = res.GetCount()
		return nil
	})
	g.Go(func() error {
		res, err := r.data.relationCli.CountFollowerByUserId(ctx, &relationRPC.CountFollowerByUserIdRequest{UserId: userId})
		if err != nil {
			return err
		}
		if res.GetStatusCode() != ecode.Success.ErrCode {
			return ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Relation rpc count user follower failed")
		}
		followerCnt = res.GetCount()
		return nil
	})
	err = g.Wait()
	if err != nil {
		r.log.Errorf("relation rpc get user follow count and follower count err: %v", err)
		return 0, 0, err
	}
	return followCnt, followerCnt, nil
}

// 从relation rpc获取是否关注
func (r *accountRepo) getIsFollowFromRelationRPC(ctx context.Context, userId int64, toUserId int64) (bool, error) {
	res, err := r.data.relationCli.IsFollowByUserId(ctx, &relationRPC.IsFollowByUserIdRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	if err != nil {
		r.log.Errorf("relation rpc get is follow err: %v", err)
		return false, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("relation rpc get is follow err: %v", res.GetStatusCode())
		return false, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Relation rpc get is follow failed")
	}
	return res.GetIsFollow(), nil
}

// 批量从passport rpc获取用户信息
func (r *accountRepo) batchGetUserInfoByUserIdFromPassportRPC(ctx context.Context, userIds []int64) ([]*do.User, error) {
	res, err := r.data.passportCli.MGetInfo(ctx, &passportRPC.DouyinMultipleGetUserInfoRequest{UserIds: userIds})
	if err != nil {
		r.log.Errorf("passport rpc batch get user info err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("passport rpc batch get user info err: %v", res.GetStatusCode())
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Passport rpc batch get userinfo failed")
	}
	users, err := mapper.UserFromPassportDTOs(res.GetInfos())
	if err != nil {
		r.log.Errorf("passport rpc batch get user info err: %v", err)
		return nil, err
	}
	return users, nil
}

// 批量从favorite rpc获取用户是否关注
func (r *accountRepo) batchGetIsFollowByUserIdFromRelationRPC(ctx context.Context, userId int64, userIds []int64) ([]bool, error) {
	res, err := r.data.relationCli.IsFollowByUserIds(ctx, &relationRPC.IsFollowByUserIdsRequest{
		UserId:       userId,
		ToUserIdList: userIds,
	})
	if err != nil {
		r.log.Errorf("relation rpc batch get is follow err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("relation rpc batch get is follow err: %v", res.GetStatusCode())
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Relation rpc batch get is follow failed")
	}
	return res.GetIsFollowList(), nil
}

// GetFollowListByUserId 获取关注列表
func (r *accountRepo) GetFollowListByUserId(ctx context.Context, userId int64, toUserId int64) ([]*do.User, error) {
	// 从memcached中获取列表
	users, err := r.getUserFollowListFromCache(ctx, userId)
	if err == nil {
		userIds := make([]int64, 0, len(users))
		for _, user := range users {
			userIds = append(userIds, user.ID)
		}
		// 从relation rpc获取是否关注
		rel, err := r.batchGetIsFollowByUserIdFromRelationRPC(ctx, userId, userIds)
		if err != nil {
			r.log.Errorf("get follow relation err: %v", err)
			return nil, err
		}
		for i, user := range users {
			user.IsFollow = rel[i]
		}
		return users, nil
	}
	// 从relation rpc获取关注列表
	userIds, err := r.getFollowListByUserIdFromRelationRPC(ctx, toUserId)
	if err != nil {
		r.log.Errorf("get follow list err: %v", err)
		return nil, err
	}
	g := errgroup.Group{}
	users = make([]*do.User, 0, len(userIds))
	isFollows := make([]bool, 0, len(userIds))
	// 获取用户信息
	g.Go(func() error {
		us, err := r.MGetUserInfoByUserId(ctx, userId, userIds)
		if err != nil {
			return err
		}
		users = append(users, us...)
		return nil
	})
	// 获取是否关注
	g.Go(func() error {
		rel, err := r.batchGetIsFollowByUserIdFromRelationRPC(ctx, userId, userIds)
		if err != nil {
			return err
		}
		isFollows = append(isFollows, rel...)
		return nil
	})
	err = g.Wait()
	if err != nil {
		r.log.Errorf("get follow list err: %v", err)
		return nil, err
	}
	// 设置缓存
	err = r.data.cacheFan.Do(ctx, func(ctx context.Context) {
		r.setUserFollowListCache(ctx, userId, users)
	})
	if err != nil {
		r.log.Errorf("Fanout err: %v", err)
	}

	// 判断是否关注要放在设置缓存后
	for i, isFollow := range isFollows {
		users[i].IsFollow = isFollow
	}
	return users, nil
}

// GetFollowerListByUserId 获取粉丝列表
func (r *accountRepo) GetFollowerListByUserId(ctx context.Context, userId int64, toUserId int64) ([]*do.User, error) {
	// 从memcached中获取列表
	users, err := r.getUserFollowerListFromCache(ctx, userId)
	if err == nil {
		userIds := make([]int64, 0, len(users))
		for _, user := range users {
			userIds = append(userIds, user.ID)
		}
		// 从relation rpc获取是否关注
		rel, err := r.batchGetIsFollowByUserIdFromRelationRPC(ctx, userId, userIds)
		if err != nil {
			r.log.Errorf("get follower relation err: %v", err)
			return nil, err
		}
		for i, user := range users {
			user.IsFollow = rel[i]
		}
		return users, nil
	}
	// 从relation rpc获取粉丝列表
	userIds, err := r.getFollowerListByUserIdFromRelationRPC(ctx, toUserId)
	if err != nil {
		r.log.Errorf("get follower list err: %v", err)
		return nil, err
	}
	g := errgroup.Group{}
	users = make([]*do.User, 0, len(userIds))
	isFollows := make([]bool, 0, len(userIds))
	// 获取用户信息
	g.Go(func() error {
		us, err := r.MGetUserInfoByUserId(ctx, userId, userIds)
		if err != nil {
			return err
		}
		users = append(users, us...)
		return nil
	})
	// 获取是否关注
	g.Go(func() error {
		rel, err := r.batchGetIsFollowByUserIdFromRelationRPC(ctx, userId, userIds)
		if err != nil {
			return err
		}
		isFollows = append(isFollows, rel...)
		return nil
	})
	err = g.Wait()
	if err != nil {
		r.log.Errorf("get follower list err: %v", err)
		return nil, err
	}
	// 设置缓存
	err = r.data.cacheFan.Do(ctx, func(ctx context.Context) {
		r.setUserFollowerListCache(ctx, userId, users)
	})
	if err != nil {
		r.log.Errorf("Fanout err: %v", err)
	}

	// 判断是否关注要放在设置缓存后
	for i, isFollow := range isFollows {
		users[i].IsFollow = isFollow
	}
	return users, nil
}

// GetFriendListByUserId 获取好友列表
func (r *accountRepo) GetFriendListByUserId(ctx context.Context, userId int64) ([]*do.User, error) {
	// 从memcached中获取列表
	users, err := r.getUserFriendListFromCache(ctx, userId)
	if err == nil {
		return users, nil
	}
	// 从relation rpc获取好友列表
	userIds, err := r.getFriendListByUserIdFromRelationRPC(ctx, userId)
	if err != nil {
		r.log.Errorf("get friend list err: %v", err)
		return nil, err
	}
	// 获取用户信息
	users, err = r.MGetUserInfoByUserId(ctx, userId, userIds)
	if err != nil {
		r.log.Errorf("get friend list err: %v", err)
		return nil, err
	}
	// 设置缓存
	err = r.data.cacheFan.Do(ctx, func(ctx context.Context) {
		r.setUserFriendListCache(ctx, userId, users)
	})
	if err != nil {
		r.log.Errorf("Fanout err: %v", err)
	}
	// 朋友一定是关注的
	for i := range users {
		users[i].IsFollow = true
	}
	return users, nil
}

// 从relation rpc获取关注列表
func (r *accountRepo) getFollowListByUserIdFromRelationRPC(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.relationCli.GetFollowListByUserId(ctx, &relationRPC.GetFollowListByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		r.log.Errorf("relation rpc get follow list err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("relation rpc get follow list err: %v", res.GetStatusCode())
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Relation rpc get follow list failed")
	}
	return res.GetUserIdList(), nil
}

// 从relation rpc获取粉丝列表
func (r *accountRepo) getFollowerListByUserIdFromRelationRPC(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.relationCli.GetFollowerListByUserId(ctx, &relationRPC.GetFollowerListByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		r.log.Errorf("relation rpc get follower list err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("relation rpc get follower list err: %v", res.GetStatusCode())
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Relation rpc get follower list failed")
	}
	return res.GetUserIdList(), nil
}

// 从relation rpc获取好友列表
func (r *accountRepo) getFriendListByUserIdFromRelationRPC(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.relationCli.GetUserFriendListByUserId(ctx, &relationRPC.GetFriendListByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		r.log.Errorf("relation rpc get friend list err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("relation rpc get friend list err: %v", res.GetStatusCode())
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "Relation rpc get friend list failed")
	}
	return res.GetUserIdList(), nil
}

// 设置用户关注列表缓存
func (r *accountRepo) setUserFollowListCache(ctx context.Context, userId int64, users []*do.User) {
	b, err := json.Marshal(users)
	if err != nil {
		r.log.Errorf("marshal user follow list err: %v", err)
		return
	}
	err = r.data.memcached.Set(&memcache.Item{
		Key:        constants.UserFollowListCacheKey(userId),
		Value:      b,
		Expiration: int32(constants.UserFollowListCacheExpiration.Seconds()),
	})
	if err != nil {
		r.log.Errorf("memcached set user follow list err: %v", err)
		return
	}
}

// 从缓存中获取用户关注列表
func (r *accountRepo) getUserFollowListFromCache(ctx context.Context, userId int64) ([]*do.User, error) {
	res, err := r.data.memcached.Get(constants.UserFollowListCacheKey(userId))
	if err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			r.log.Errorf("memcached get user follow list err: %v", err)
		}
		return nil, err
	}
	users := make([]*do.User, 0, len(res.Value))
	err = json.Unmarshal(res.Value, &users)
	if err != nil {
		r.log.Errorf("unmarshal user follow list err: %v", err)
		return nil, err
	}
	return users, nil
}

// 设置用户粉丝列表缓存
func (r *accountRepo) setUserFollowerListCache(ctx context.Context, userId int64, users []*do.User) {
	b, err := json.Marshal(users)
	if err != nil {
		r.log.Errorf("marshal user follower list err: %v", err)
		return
	}
	err = r.data.memcached.Set(&memcache.Item{
		Key:        constants.UserFollowerListCacheKey(userId),
		Value:      b,
		Expiration: int32(constants.UserFollowerListCacheExpiration.Seconds()),
	})
	if err != nil {
		r.log.Errorf("memcached set user follower list err: %v", err)
		return
	}
}

// 从缓存中获取用户粉丝列表
func (r *accountRepo) getUserFollowerListFromCache(ctx context.Context, userId int64) ([]*do.User, error) {
	res, err := r.data.memcached.Get(constants.UserFollowerListCacheKey(userId))
	if err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			r.log.Errorf("memcached get user follower list err: %v", err)
		}
		return nil, err
	}
	users := make([]*do.User, 0, len(res.Value))
	err = json.Unmarshal(res.Value, &users)
	if err != nil {
		r.log.Errorf("unmarshal user follower list err: %v", err)
		return nil, err
	}
	return users, nil
}

// 设置用户好友列表缓存
func (r *accountRepo) setUserFriendListCache(ctx context.Context, userId int64, users []*do.User) {
	b, err := json.Marshal(users)
	if err != nil {
		r.log.Errorf("marshal user friend list err: %v", err)
		return
	}
	err = r.data.memcached.Set(&memcache.Item{
		Key:        constants.UserFriendListCacheKey(userId),
		Value:      b,
		Expiration: int32(constants.UserFriendListCacheExpiration.Seconds()),
	})
	if err != nil {
		r.log.Errorf("memcached set user friend list err: %v", err)
		return
	}
}

// 从缓存中获取用户好友列表
func (r *accountRepo) getUserFriendListFromCache(ctx context.Context, userId int64) ([]*do.User, error) {
	res, err := r.data.memcached.Get(constants.UserFriendListCacheKey(userId))
	if err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			r.log.Errorf("memcached get user friend list err: %v", err)
		}
		return nil, err
	}
	users := make([]*do.User, 0, len(res.Value))
	err = json.Unmarshal(res.Value, &users)
	if err != nil {
		r.log.Errorf("unmarshal user friend list err: %v", err)
		return nil, err
	}
	return users, nil
}
