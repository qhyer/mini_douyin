package biz

import (
	mock_biz "douyin/app/video/publish/service/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestVideoUsecase_CountUserPublishedVideoByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockVideoRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CountUserPublishedVideoByUserId(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	uc := NewVideoUsecase(repo, nil)
	_, err := uc.CountUserPublishedVideoByUserId(nil, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVideoUsecase_GetPublishedVideosByLatestTime(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockVideoRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetPublishedVideosByLatestTime(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewVideoUsecase(repo, nil)
	_, err := uc.GetPublishedVideosByLatestTime(nil, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVideoUsecase_GetPublishedVideosByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockVideoRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetPublishedVideosByUserId(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewVideoUsecase(repo, nil)
	_, err := uc.GetPublishedVideosByUserId(nil, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVideoUsecase_PublishVideo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockVideoRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().UploadVideo(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
		repo.EXPECT().PublishVideo(gomock.Any(), gomock.Any()).Return(nil),
	)

	uc := NewVideoUsecase(repo, nil)
	err := uc.PublishVideo(nil, []byte{}, 1, "test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestVideoUsecase_GetVideoById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockVideoRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetVideoById(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewVideoUsecase(repo, nil)
	_, err := uc.GetVideoById(nil, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVideoUsecase_MGetVideoByIds(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockVideoRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().MGetVideoByIds(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewVideoUsecase(repo, nil)
	_, err := uc.MGetVideoByIds(nil, []int64{1})
	if err != nil {
		t.Fatal(err)
	}
}
