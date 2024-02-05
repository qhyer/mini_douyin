package biz

import (
	do "douyin/app/user/chat/common/entity"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
)

type ChatRepo interface {
	CreateMessage(ctx context.Context, message *do.Message) error
}

type ChatUsecase struct {
	repo ChatRepo
	log  *log.Helper
}

func NewChatUsecase(repo ChatRepo, logger log.Logger) *ChatUsecase {
	return &ChatUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ChatUsecase) CreateMessage(ctx context.Context, message *do.Message) error {
	return uc.repo.CreateMessage(ctx, message)
}
