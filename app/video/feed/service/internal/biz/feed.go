package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"

	"douyin/app/video/feed/common/constants"
	do "douyin/app/video/feed/common/entity"
)

type FeedRepo interface {
	GetPublishedVideoByUserId(ctx context.Context, userId int64) ([]*do.Video, error)
	GetPublishedVideoByLatestTime(ctx context.Context, latestTime int64) ([]*do.Video, error)
	GetUserInfoByUserId(ctx context.Context, userId, toUserId int64) (*do.User, error)
	GetUserFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error)
	MGetUserInfoByUserId(ctx context.Context, userId int64, toUserIds []int64) ([]*do.User, error)
	MCountVideoFavoritedByVideoId(ctx context.Context, videoId []int64) ([]int64, error)
	MCountCommentByVideoId(ctx context.Context, videoId []int64) ([]int64, error)
	MGetIsVideoFavoritedByVideoIdAndUserId(ctx context.Context, userId int64, videoIds []int64) ([]bool, error)
	MGetVideoInfoByVideoId(ctx context.Context, videoIds []int64) ([]*do.Video, error)
}

type FeedUsecase struct {
	repo FeedRepo
	log  *log.Helper
	sf   *singleflight.Group
}

func NewFeedUsecase(repo FeedRepo, logger log.Logger) *FeedUsecase {
	return &FeedUsecase{repo: repo, log: log.NewHelper(logger)}
}

// GetPublishedVideoByUserId 获取用户发布视频列表
func (uc *FeedUsecase) GetPublishedVideoByUserId(ctx context.Context, userId, toUserId int64) ([]*do.Video, error) {
	videoList, err := uc.repo.GetPublishedVideoByUserId(ctx, toUserId)
	if err != nil {
		uc.log.Errorf("uc.repo.GetPublishedVideoByUserId error(%v)", err)
		return nil, err
	}
	g := errgroup.Group{}
	var isFavoriteList []bool
	var authorList []*do.User
	videoIdList := make([]int64, 0, len(videoList))
	for _, video := range videoList {
		videoIdList = append(videoIdList, video.ID)
	}
	g.Go(func() error {
		// 获取是否点赞
		isFavoriteList, err = uc.repo.MGetIsVideoFavoritedByVideoIdAndUserId(ctx, userId, videoIdList)
		if err != nil {
			uc.log.Errorf("uc.repo.MGetIsVideoFavoritedByVideoIdAndUserId error(%v)", err)
			return err
		}
		return nil
	})
	g.Go(func() error {
		// 获取作者信息
		var uids []int64
		for _, video := range videoList {
			uids = append(uids, video.User.ID)
		}
		authorList, err = uc.repo.MGetUserInfoByUserId(ctx, toUserId, uids)
		if err != nil {
			uc.log.Errorf("uc.repo.MGetUserInfoByUserId error(%v)", err)
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		uc.log.Errorf("g.Wait error(%v)", err)
		return nil, err
	}
	for i, video := range videoList {
		video.IsFavorite = isFavoriteList[i]
		video.User = authorList[i]
	}

	return videoList, nil
}

// GetPublishedVideoByLatestTimeAndUserId 获取用户发布视频列表
func (uc *FeedUsecase) GetPublishedVideoByLatestTimeAndUserId(ctx context.Context, userId int64, latestTime int64) ([]*do.Video, error) {
	videoList, err := uc.repo.GetPublishedVideoByLatestTime(ctx, latestTime)
	if err != nil {
		uc.log.Errorf("uc.repo.GetPublishedVideoByLatestTime error(%v)", err)
		return nil, err
	}
	g := errgroup.Group{}
	var isFavoriteList []bool
	var authorList []*do.User
	videoIdList := make([]int64, 0, len(videoList))
	for _, video := range videoList {
		videoIdList = append(videoIdList, video.ID)
	}
	g.Go(func() error {
		// 获取是否点赞
		isFavoriteList, err = uc.repo.MGetIsVideoFavoritedByVideoIdAndUserId(ctx, userId, videoIdList)
		if err != nil {
			uc.log.Errorf("uc.repo.MGetIsVideoFavoritedByVideoIdAndUserId error(%v)", err)
			return err
		}
		return nil
	})
	g.Go(func() error {
		// 获取作者信息
		var uids []int64
		for _, video := range videoList {
			uids = append(uids, video.User.ID)
		}
		authorList, err = uc.repo.MGetUserInfoByUserId(ctx, userId, uids)
		if err != nil {
			uc.log.Errorf("uc.repo.MGetUserInfoByUserId error(%v)", err)
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		uc.log.Errorf("g.Wait error(%v)", err)
		return nil, err
	}
	for i, video := range videoList {
		video.IsFavorite = isFavoriteList[i]
		video.User = authorList[i]
	}

	return videoList, nil
}

// GetUserFavoriteVideoListByUserId 获取用户点赞视频列表
func (uc *FeedUsecase) GetUserFavoriteVideoListByUserId(ctx context.Context, userId int64, toUserId int64) ([]*do.Video, error) {
	// 获取对方的点赞视频id列表
	favIdList, err, _ := uc.sf.Do(constants.SFUserFavoriteVideoIdListKey(toUserId), func() (interface{}, error) {
		return uc.repo.GetUserFavoriteVideoIdListByUserId(ctx, toUserId)
	})
	if err != nil {
		uc.log.Errorf("uc.repo.GetUserFavoriteVideoIdListByUserId error(%v)", err)
		return nil, err
	}
	g := errgroup.Group{}
	var videoList []*do.Video
	var isFavoriteList []bool
	var authorList []*do.User
	g.Go(func() error {
		// 获取视频信息
		videoList, err = uc.repo.MGetVideoInfoByVideoId(ctx, favIdList.([]int64))
		if err != nil {
			uc.log.Errorf("uc.repo.MGetVideoByIds error(%v)", err)
			return err
		}
		return nil
	})
	g.Go(func() error {
		// 获取是否点赞
		isFavoriteList, err = uc.repo.MGetIsVideoFavoritedByVideoIdAndUserId(ctx, userId, favIdList.([]int64))
		if err != nil {
			uc.log.Errorf("uc.repo.MGetIsVideoFavoritedByVideoIdAndUserId error(%v)", err)
			return err
		}
		return nil
	})
	g.Go(func() error {
		// 获取作者信息
		var uids []int64
		for _, video := range videoList {
			uids = append(uids, video.User.ID)
		}
		authorList, err = uc.repo.MGetUserInfoByUserId(ctx, userId, uids)
		if err != nil {
			uc.log.Errorf("uc.repo.MGetUserInfoByUserId error(%v)", err)
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		uc.log.Errorf("g.Wait error(%v)", err)
		return nil, err
	}
	for i, video := range videoList {
		video.IsFavorite = isFavoriteList[i]
		video.User = authorList[i]
	}

	return videoList, nil
}
