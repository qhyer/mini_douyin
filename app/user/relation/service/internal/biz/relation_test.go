package biz

import (
	"context"
	"douyin/app/user/relation/common/event"
	mock_biz "douyin/app/user/relation/service/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRelationUsecase_CountFollowByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CountFollowByUserId(gomock.Any(), gomock.Any()).Return(int64(0), nil),
	)

	uc := NewRelationUsecase(repo, nil)
	_, err := uc.CountFollowByUserId(context.Background(), int64(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRelationUsecase_CountFollowerByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CountFollowerByUserId(gomock.Any(), gomock.Any()).Return(int64(0), nil),
	)

	uc := NewRelationUsecase(repo, nil)
	_, err := uc.CountFollowerByUserId(context.Background(), int64(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRelationUsecase_GetFollowListByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetFollowListByUserId(gomock.Any(), gomock.Any()).Return([]int64{}, nil),
	)

	uc := NewRelationUsecase(repo, nil)
	_, err := uc.GetFollowListByUserId(context.Background(), int64(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRelationUsecase_GetFollowerListByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetFollowerListByUserId(gomock.Any(), gomock.Any()).Return([]int64{}, nil),
	)

	uc := NewRelationUsecase(repo, nil)
	_, err := uc.GetFollowerListByUserId(context.Background(), int64(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRelationUsecase_GetFriendListByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetFriendListByUserId(gomock.Any(), gomock.Any()).Return([]int64{}, nil),
	)

	uc := NewRelationUsecase(repo, nil)
	_, err := uc.GetFriendListByUserId(context.Background(), int64(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRelationUsecase_IsFollowByUserId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().IsFollowByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil),
	)

	uc := NewRelationUsecase(repo, nil)
	_, err := uc.IsFollowByUserId(context.Background(), int64(0), int64(0))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRelationUsecase_IsFollowByUserIds(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().IsFollowByUserIds(gomock.Any(), gomock.Any(), gomock.Any()).Return([]bool{false}, nil),
	)

	uc := NewRelationUsecase(repo, nil)
	_, err := uc.IsFollowByUserIds(context.Background(), int64(0), []int64{0})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRelationUsecase_RelationAction(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockRelationRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().RelationAction(gomock.Any(), gomock.Any()).Return(nil),
	)

	uc := NewRelationUsecase(repo, nil)
	err := uc.RelationAction(context.Background(), &event.RelationAction{
		ID:         1,
		Type:       event.RelationActionFollow,
		FromUserId: 1,
		ToUserId:   2,
	})
	if err != nil {
		t.Fatal(err)
	}
}
