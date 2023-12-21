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
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
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
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "PublishService.GetUserPublishedVideoList failed")
	}
	videos, err = mapper.VideoFromPublishDTOs(res.GetVideoList())
	if err != nil {
		r.log.Errorf("VideoFromPublishDTOs failed, err: %v", err)
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
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "PublishService.GetPublishedVideoByLatestTime failed")
	}
	videos, err := mapper.VideoFromPublishDTOs(res.GetVideoList())
	if err != nil {
		r.log.Errorf("VideoFromPublishDTOs failed, err: %v", err)
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
	res, err := r.data.favoriteCli.GetFavoriteStatusByUserIdAndVideoIds(ctx, &favRPC.GetFavoriteStatusByUserIdAndVideoIdsRequest{
		UserId:   userId,
		VideoIds: videoIds,
	})
	isFavoriteList := make([]bool, 0, len(videoIds))
	// 这里不返回错误，直接返回未点赞
	if err != nil {
		r.log.Errorf("MGetIsVideoFavoritedByVideoIdAndUserId failed, err: %v", err)
		return isFavoriteList, nil
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("MGetIsVideoFavoritedByVideoIdAndUserId failed, err: %v", res.GetStatusMsg())
		return isFavoriteList, nil
	}
	isFavoriteList = res.GetIsFavoriteList()
	return isFavoriteList, nil
}

func (r *feedRepo) GetUserFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res, err := r.data.favoriteCli.GetUserFavoriteVideoIdList(ctx, &favRPC.GetUserFavoriteListRequest{
		UserId: userId,
	})
	videoIds := make([]int64, 0)
	// 这里不返回错误，直接返回空列表
	if err != nil {
		r.log.Errorf("GetUserFavoriteVideoIdListByUserId failed, err: %v", err)
		return videoIds, nil
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("GetUserFavoriteVideoIdListByUserId failed, err: %v", res.GetStatusMsg())
		return videoIds, nil
	}
	videoIds = res.GetVideoIdList()
	return videoIds, nil
}

func (r *feedRepo) MGetVideoInfoByVideoId(ctx context.Context, videoIds []int64) ([]*do.Video, error) {
	videos := make([]*do.Video, 0, len(videoIds))
	// 先从缓存中获取
	videos, missVideoIds := r.batchGetVideoInfoFromCache(ctx, videoIds)
	if len(missVideoIds) == 0 {
		return videos, nil
	}
	// 从rpc获取
	tmpVideos, err := r.batchGetVideoInfoFromRPC(ctx, missVideoIds)
	if err != nil {
		r.log.Errorf("batchGetVideoInfoFromRPC failed, err: %v", err)
		return nil, err
	}
	// 热门视频缓存到redis
	for _, video := range tmpVideos {
		if video.FavoriteCount >= constants.LeastVideoFavoritedCount {
			r.setVideoInfoCache(ctx, video)
		}
	}
	vm := make(map[int64]*do.Video, len(videoIds))
	for _, video := range videos {
		vm[video.ID] = video
	}
	for _, video := range tmpVideos {
		vm[video.ID] = video
	}
	videos = make([]*do.Video, 0, len(videoIds))
	for _, videoId := range videoIds {
		videos = append(videos, vm[videoId])
	}
	return videos, nil
}

func (r *feedRepo) setVideoInfoCache(ctx context.Context, video *do.Video) {
	b, err := video.MarshalJson()
	if err != nil {
		r.log.Errorf("video MarshalJson failed, err: %v", err)
		return
	}
	r.data.redis.Set(ctx, constants.VideoInfoCacheKey(video.ID), b, constants.VideoInfoCacheExpiration)
}

func (r *feedRepo) getVideoInfoFromCache(ctx context.Context, videoId int64) (*do.Video, error) {
	res, err := r.data.redis.Get(ctx, constants.VideoInfoCacheKey(videoId)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		r.log.Errorf("redis get failed, err: %v", err)
		return nil, err
	}
	video := &do.Video{}
	err = video.UnmarshalJson(res)
	if err != nil {
		r.log.Errorf("video UnmarshalJson failed, err: %v", err)
		return nil, err
	}
	return video, nil
}

func (r *feedRepo) batchGetVideoInfoFromCache(ctx context.Context, videoIds []int64) ([]*do.Video, []int64) {
	videos := make([]*do.Video, 0, len(videoIds))
	missVideoIds := make([]int64, 0)
	pipe := r.data.redis.Pipeline()
	for _, videoId := range videoIds {
		pipe.Get(ctx, constants.VideoInfoCacheKey(videoId))
	}
	res, err := pipe.Exec(ctx)
	if err != nil {
		r.log.Errorf("redis pipeline exec failed, err: %v", err)
		return nil, videoIds
	}
	for i, re := range res {
		if re.Err() != nil {
			missVideoIds = append(missVideoIds, videoIds[i])
			if !errors.Is(re.Err(), redis.Nil) {
				r.log.Errorf("redis pipeline exec failed, err: %v", re.Err())
			}
			continue
		}
		b, err := re.(*redis.StringCmd).Bytes()
		if err != nil {
			r.log.Errorf("redis pipeline exec failed, err: %v", err)
			missVideoIds = append(missVideoIds, videoIds[i])
			continue
		}
		video := &do.Video{}
		err = video.UnmarshalJson(b)
		if err != nil {
			r.log.Errorf("video UnmarshalJson failed, err: %v", err)
			missVideoIds = append(missVideoIds, videoIds[i])
			continue
		}
		videos = append(videos, video)
	}
	return videos, missVideoIds
}

func (r *feedRepo) batchGetVideoInfoFromRPC(ctx context.Context, videoIds []int64) ([]*do.Video, error) {
	res, err := r.data.publishCli.MGetVideoInfoByVideoIds(ctx, &publishRPC.MGetVideoInfoRequest{
		VideoIds: videoIds,
	})
	if err != nil {
		r.log.Errorf("MGetVideoInfoByVideoId failed, err: %v", err)
		return nil, err
	}
	if res.GetStatusCode() != ecode.Success.ErrCode {
		r.log.Errorf("MGetVideoInfoByVideoId failed, err: %v", res.GetStatusMsg())
		return nil, ecode.NewErrNo(res.GetStatusCode(), res.GetStatusMsg(), "PublishService.MGetVideoInfoByVideoId failed")
	}
	videos, err := mapper.VideoFromPublishDTOs(res.GetVideoList())
	if err != nil {
		r.log.Errorf("VideoFromPublishDTOs failed, err: %v", err)
		return nil, err
	}
	return videos, nil
}
