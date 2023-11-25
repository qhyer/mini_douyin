package data

import (
	"context"
	passportRPC "douyin/api/user/passport/service/v1"
	relationRPC "douyin/api/user/relation/service/v1"
	favoriteRPC "douyin/api/video/favorite/service/v1"
	publishRPC "douyin/api/video/publish/service/v1"
	"douyin/app/user/account/common/constants"
	do "douyin/app/user/account/common/entity"
	"douyin/app/user/account/common/mapper"
	"douyin/app/user/account/service/internal/biz"
	"douyin/common/ecode"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
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

// GetUserInfoByUserId 通过用户ID获取用户信息
func (r *passportRepo) GetUserInfoByUserId(ctx context.Context, userId int64, toUserId int64) (*do.User, error) {
	user, err := r.getUserInfoFromCache(ctx, constants.AccountInfoCacheKey(toUserId))
	if err != nil {
		if err != redis.Nil {
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
		r.setUserInfoCache(ctx, user)
	}
	return user, nil
}

// MGetUserInfoByUserId 批量获取用户信息（只获取基本信息）
func (r *passportRepo) MGetUserInfoByUserId(ctx context.Context, userId int64, userIds []int64) ([]*do.User, error) {
	//TODO implement me
	panic("implement me")
}

// 设置用户信息缓存
func (r *passportRepo) setUserInfoCache(ctx context.Context, user *do.User) {
	b, err := user.MarshalJson()
	if err != nil {
		r.log.Errorf("redis marshal user info err: %v", err)
		return
	}
	err = r.data.redis.Set(ctx, constants.AccountInfoCacheKey(user.ID), b, constants.AccountInfoCacheExpiration).Err()
	if err != nil {
		r.log.Errorf("redis set user info err: %v", err)
	}
}

// 从缓存中获取用户信息
func (r *passportRepo) getUserInfoFromCache(ctx context.Context, userId string) (*do.User, error) {
	res, err := r.data.redis.Get(ctx, userId).Bytes()
	if err != nil {
		if err != redis.Nil {
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
func (r *passportRepo) getUserInfoFromPassportRPC(ctx context.Context, userId int64) (*do.User, error) {
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
func (r *passportRepo) getUserFavoriteCountFromFavoriteRPC(ctx context.Context, userId int64) (favCnt int64, favdCnt int64, err error) {
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
func (r *passportRepo) getUserWorkCountFromPublishRPC(ctx context.Context, userId int64) (workCnt int64, err error) {
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
func (r *passportRepo) getUserFollowCountAndFollowerCountFromRelationRPC(ctx context.Context, userId int64) (followCnt int64, followerCnt int64, err error) {
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
func (r *passportRepo) getIsFollowFromRelationRPC(ctx context.Context, userId int64, toUserId int64) (bool, error) {
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
