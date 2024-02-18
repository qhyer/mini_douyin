package biz

import (
	do "douyin/app/video/feed/common/entity"
	mock_biz "douyin/app/video/feed/service/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestFeedUsecase_GetPublishedVideoByLatestTimeAndUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFeedRepo(ctl)
	repo.EXPECT().GetPublishedVideoByLatestTime(gomock.Any(), gomock.Any()).Return([]*do.Video{
		&do.Video{
			ID: 1,
		},
	}, nil)
	repo.EXPECT().MGetIsVideoFavoritedByVideoIdAndUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return([]bool{true}, nil)
	repo.EXPECT().MGetUserInfoByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*do.User{
		&do.User{
			ID: 1,
		},
	}, nil)

	uc := NewFeedUsecase(repo, nil)
	_, err := uc.GetPublishedVideoByLatestTimeAndUserId(nil, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFeedUsecase_GetPublishedVideoByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFeedRepo(ctl)
	repo.EXPECT().GetPublishedVideoByUserId(gomock.Any(), gomock.Any()).Return([]*do.Video{
		{
			ID: 1,
		},
	}, nil)
	repo.EXPECT().MGetIsVideoFavoritedByVideoIdAndUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return([]bool{true}, nil)
	repo.EXPECT().MGetUserInfoByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*do.User{
		{
			ID: 1,
		},
	}, nil)

	uc := NewFeedUsecase(repo, nil)
	_, err := uc.GetPublishedVideoByUserId(nil, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}
