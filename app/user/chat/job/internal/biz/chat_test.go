package biz

import (
	do "douyin/app/user/chat/common/entity"
	mock_biz "douyin/app/user/chat/job/internal/biz/mock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestChatUsecase_CreateMessage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_biz.NewMockChatRepo(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateMessage(gomock.Any(), gomock.Any()).Return(nil),
	)

	uc := NewChatUsecase(repo, nil)
	err := uc.CreateMessage(nil, &do.Message{
		ID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
}
