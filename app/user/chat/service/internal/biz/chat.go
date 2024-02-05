package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	do "douyin/app/user/chat/common/entity"
)

type ChatRepo interface {
	SendMessage(ctx context.Context, message *do.Message) error
	GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, myUserId, hisUserId, preMsgTime int64, limit int) ([]*do.Message, error)
	GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, myUserId, hisUserId int64) (*do.Message, error)
}

type ChatUsecase struct {
	repo ChatRepo
	log  *log.Helper
}

func NewChatUsecase(repo ChatRepo, logger log.Logger) *ChatUsecase {
	return &ChatUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ChatUsecase) SendMessage(ctx context.Context, message *do.Message) error {
	return uc.repo.SendMessage(ctx, message)
}

func (uc *ChatUsecase) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, myUserId, hisUserId, preMsgTime int64, limit int) ([]*do.Message, error) {
	return uc.repo.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx, myUserId, hisUserId, preMsgTime, limit)
}

func (uc *ChatUsecase) GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, myUserId, hisUserId int64) (*do.Message, error) {
	return uc.repo.GetLatestMsgByMyUserIdAndHisUserId(ctx, myUserId, hisUserId)
}
