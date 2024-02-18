package biz

import (
	mock_biz "douyin/app/video/comment/service/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCommentUsecase_CommentAction(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockCommentRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CommentAction(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewCommentUsecase(repo, nil)
	_, err := uc.CommentAction(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommentUsecase_CountCommentByVideoId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockCommentRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CountCommentByVideoId(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	uc := NewCommentUsecase(repo, nil)
	_, err := uc.CountCommentByVideoId(nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommentUsecase_GetCommentListByVideoId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockCommentRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().GetCommentListByVideoId(gomock.Any(), gomock.Any()).Return(nil, nil),
	)

	uc := NewCommentUsecase(repo, nil)
	_, err := uc.GetCommentListByVideoId(nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommentUsecase_MCountCommentByVideoId(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockCommentRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().MCountCommentByVideoId(gomock.Any(), gomock.Any()).Return([]int64{1}, nil),
	)

	uc := NewCommentUsecase(repo, nil)
	_, err := uc.MCountCommentByVideoId(nil, []int64{1})
	if err != nil {
		t.Fatal(err)
	}
}
