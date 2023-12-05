package data

import (
	"context"
	accountRPC "douyin/api/user/account/service/v1"
	comRPC "douyin/api/video/comment/service/v1"
	favRPC "douyin/api/video/favorite/service/v1"
	publishRPC "douyin/api/video/publish/service/v1"
	"douyin/app/video/feed/common/constants"
	do "douyin/app/video/feed/common/entity"
	"douyin/app/video/feed/common/mapper"
	"douyin/app/video/feed/service/internal/biz"
	"douyin/common/ecode"
	"github.com/go-kratos/kratos/v2/log"
)

type feedRepo struct {
	data *Data
	log  *log.Helper
}

func NewFeedRepo(data *Data, logger log.Logger) biz.FeedRepo {
	return &feedRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *feedRepo) GetPublishedVideoByUserId(ctx context.Context, userId int64) ([]*do.Video, error) {
	videos := make([]*do.Video, 0)
	res, err := r.data.publishCli.GetUserPublishedVideoList(ctx, &publishRPC.GetUserPublishedVideoListRequest{
		UserId: userId,
	})
	if err != nil {
		r.log.Errorf("GetUserPublishedVideoList failed, err: %v", err)
		return videos, nil
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("GetUserPublishedVideoList failed, err: %v", res.GetStatusMsg())
		return videos, nil
	}
	videos, err = mapper.VideoFromDTOs(res.GetVideoList())
	if err != nil {
		r.log.Errorf("VideoFromDTOs failed, err: %v", err)
		return videos, nil
	}
	return videos, nil
}

func (r *feedRepo) GetPublishedVideoByLatestTime(ctx context.Context, latestTime int64) ([]*do.Video, error) {
	res, err := r.data.publishCli.GetPublishedVideoByLatestTime(ctx, &publishRPC.GetPublishedVideoByLatestTimeRequest{
		LatestTime: latestTime,
		Limit:      constants.VideoQueryLimit,
	})
	if err != nil {
		r.log.Errorf("GetPublishedVideoByLatestTime failed, err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("GetPublishedVideoByLatestTime failed, err: %v", res.GetStatusMsg())
		return nil, nil
	}
	videos, err := mapper.VideoFromDTOs(res.GetVideoList())
	if err != nil {
		r.log.Errorf("VideoFromDTOs failed, err: %v", err)
		return nil, err
	}
	return videos, nil
}

func (r *feedRepo) GetUserInfoByUserId(ctx context.Context, userId, toUserId int64) (*do.User, error) {
	user := &do.User{}
	res, err := r.data.accountCli.GetUserInfoByUserId(ctx, &accountRPC.GetUserInfoByUserIdRequest{
		UserId:   userId,
		ToUserId: toUserId,
	})
	// 这里不返回错误，返回空user
	if err != nil {
		r.log.Errorf("GetUserInfoByUserId failed, err: %v", err)
		return user, nil
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("GetUserInfoByUserId failed, err: %v", res.GetStatusMsg())
		return user, nil
	}
	user, err = mapper.UserFromAccountDTO(res.GetUser())
	if err != nil {
		r.log.Errorf("UserFromAccountDTOs failed, err: %v", err)
		return user, nil
	}
	return user, nil
}

func (r *feedRepo) MGetUserInfoByUserId(ctx context.Context, userId int64, toUserIds []int64) ([]*do.User, error) {
	users := make([]*do.User, 0)
	res, err := r.data.accountCli.MGetUserInfoByUserId(ctx, &accountRPC.MGetUserInfoByUserIdRequest{
		UserId:    userId,
		ToUserIds: toUserIds,
	})
	// 这里不返回错误，返回空user
	if err != nil {
		r.log.Errorf("MGetUserInfoByUserId failed, err: %v", err)
		return users, nil
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("MGetUserInfoByUserId failed, err: %v", res.GetStatusMsg())
		return users, nil
	}
	users, err = mapper.UserFromAccountDTOs(res.GetUsers())
	if err != nil {
		r.log.Errorf("UserFromAccountDTOs failed, err: %v", err)
		return users, nil
	}
	return users, nil
}

func (r *feedRepo) MCountVideoFavoritedByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	res, err := r.data.favoriteCli.MCountVideoFavoritedByVideoIds(ctx, &favRPC.MCountVideoFavoritedByVideoIdsRequest{
		VideoIds: videoIds,
	})
	cnt := make([]int64, 0, len(videoIds))
	// 这里不返回错误，直接置空点赞数
	if err != nil {
		r.log.Errorf("MCountVideoFavoritedByVideoId failed, err: %v", err)
		return cnt, nil
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("MCountVideoFavoritedByVideoId failed, err: %v", res.GetStatusMsg())
		return cnt, nil
	}
	cnt = res.GetCountList()
	return cnt, nil
}

func (r *feedRepo) MCountCommentByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	res, err := r.data.commentCli.MCountCommentByVideoId(ctx, &comRPC.MCountCommentByVideoIdRequest{
		VideoIdList: videoIds,
	})
	cnt := make([]int64, 0, len(videoIds))
	// 这里不返回错误，直接置空评论数
	if err != nil {
		r.log.Errorf("MCountCommentByVideoId failed, err: %v", err)
		return cnt, nil
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("MCountCommentByVideoId failed, err: %v", res.GetStatusMsg())
		return cnt, nil
	}
	cnt = res.GetCountList()
	return cnt, nil
}

func (r *feedRepo) MGetIsVideoFavoritedByVideoIdAndUserId(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	// todo
	panic("implement me ")
}
