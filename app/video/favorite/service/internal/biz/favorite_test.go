package biz

import (
	"douyin/app/video/favorite/common/event"
	mock_biz "douyin/app/video/favorite/service/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestFavoriteUsecase_CountUserFavoriteByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFavoriteRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CountUserFavoriteByUserId(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	uc := NewFavoriteUsecase(repo, nil)
	_, err := uc.CountUserFavoriteByUserId(nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFavoriteUsecase_CountUserFavoritedByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFavoriteRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CountUserFavoritedByUserId(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	uc := NewFavoriteUsecase(repo, nil)
	_, err := uc.CountUserFavoritedByUserId(nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFavoriteUsecase_CountVideoFavoritedByVideoId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFavoriteRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CountVideoFavoritedByVideoId(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	uc := NewFavoriteUsecase(repo, nil)
	_, err := uc.CountVideoFavoritedByVideoId(nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFavoriteUsecase_FavoriteAction(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFavoriteRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().FavoriteVideo(gomock.Any(), gomock.Any()).Return(nil),
	)

	uc := NewFavoriteUsecase(repo, nil)
	err := uc.FavoriteAction(nil, &event.FavoriteAction{ID: 1, Type: event.FavoriteActionAdd})
	if err != nil {
		t.Fatal(err)
	}
}

func TestFavoriteUsecase_GetFavoriteStatusByUserIdAndVideoIds(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFavoriteRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().IsUserFavoriteVideoList(gomock.Any(), gomock.Any(), gomock.Any()).Return([]bool{false}, nil),
	)

	uc := NewFavoriteUsecase(repo, nil)
	_, err := uc.GetFavoriteStatusByUserIdAndVideoIds(nil, 0, []int64{1})
	if err != nil {
		t.Fatal(err)
	}
}

func TestFavoriteUsecase_GetFavoriteVideoIdListByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFavoriteRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetFavoriteVideoIdListByUserId(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewFavoriteUsecase(repo, nil)
	_, err := uc.GetFavoriteVideoIdListByUserId(nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFavoriteUsecase_MCountVideoFavoritedByVideoId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockFavoriteRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().MCountVideoFavoritedByVideoId(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewFavoriteUsecase(repo, nil)
	_, err := uc.MCountVideoFavoritedByVideoId(nil, []int64{1})
	if err != nil {
		t.Fatal(err)
	}
}
