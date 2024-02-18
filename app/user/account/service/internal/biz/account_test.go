package biz

import (
	do "douyin/app/user/account/common/entity"
	mock_biz "douyin/app/user/account/service/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestAccountUsecase_GetFollowListByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockAccountRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetFollowListByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*do.User{
			{
				ID: 1,
			},
		}, nil),
	)

	uc := NewAccountUsecase(repo, nil)
	_, err := uc.GetFollowListByUserId(nil, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccountUsecase_GetFollowerListByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockAccountRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetFollowerListByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*do.User{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil),
	)

	uc := NewAccountUsecase(repo, nil)
	_, err := uc.GetFollowerListByUserId(nil, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccountUsecase_GetFriendListByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockAccountRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetFriendListByUserId(gomock.Any(), gomock.Any()).Return([]*do.User{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil),
	)

	uc := NewAccountUsecase(repo, nil)
	_, err := uc.GetFriendListByUserId(nil, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccountUsecase_GetUserInfoByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockAccountRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetUserInfoByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return(&do.User{
			ID: 1,
		}, nil),
	)

	uc := NewAccountUsecase(repo, nil)
	_, err := uc.GetUserInfoByUserId(nil, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccountUsecase_MGetUserInfoByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockAccountRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().MGetUserInfoByUserId(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewAccountUsecase(repo, nil)
	_, err := uc.MGetUserInfoByUserId(nil, []int64{1})
	if err != nil {
		t.Fatal(err)
	}
}
