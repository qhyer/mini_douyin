package service

import (
	"context"

	"douyin/api/user/chat/service/v1"
)

type ChatService struct {
	v1.UnimplementedChatServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) ChatAction(ctx context.Context, req *v1.DouyinChatActionRequest) (*v1.DouyinChatActionResponse, error) {
	return &v1.DouyinChatActionResponse{}, nil
}

func (s *ChatService) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, req *v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) (*v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse, error) {
	return &v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse{}, nil
}

func (s *ChatService) GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, req *v1.GetLatestMsgByMyUserIdAndHisUserIdRequest) (*v1.GetLatestMsgByMyUserIdAndHisUserIdResponse, error) {
	return &v1.GetLatestMsgByMyUserIdAndHisUserIdResponse{}, nil
}
