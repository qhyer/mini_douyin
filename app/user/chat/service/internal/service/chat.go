package service

import (
	"context"
	do "douyin/app/user/chat/common/entity"
	"douyin/app/user/chat/common/mapper"
	"douyin/app/user/chat/service/internal/biz"
	"douyin/common/ecode"
	"time"

	"douyin/api/user/chat/service/v1"
)

type ChatService struct {
	v1.UnimplementedChatServer

	uc *biz.ChatUsecase
}

func NewChatService(uc *biz.ChatUsecase) *ChatService {
	return &ChatService{
		uc: uc,
	}
}

func (s *ChatService) ChatAction(ctx context.Context, req *v1.DouyinChatActionRequest) (*v1.DouyinChatActionResponse, error) {
	err := s.uc.SendMessage(ctx, &do.Message{
		FromUserId: req.GetUserId(),
		ToUserId:   req.GetToUserId(),
		Content:    req.GetContent(),
		CreatedAt:  time.Now(),
	})
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.DouyinChatActionResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}
	return &v1.DouyinChatActionResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
	}, nil
}

func (s *ChatService) GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx context.Context, req *v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) (*v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse, error) {
	res, err := s.uc.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime(ctx, req.GetUserId(), req.GetToUserId(), req.GetPreMsgTime(), 0)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	msgs, err := mapper.MessageToDTOs(res)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse{
		StatusCode:  ecode.Success.ErrCode,
		StatusMsg:   &ecode.Success.ErrMsg,
		MessageList: msgs,
	}, nil
}

func (s *ChatService) GetLatestMsgByMyUserIdAndHisUserId(ctx context.Context, req *v1.GetLatestMsgByMyUserIdAndHisUserIdRequest) (*v1.GetLatestMsgByMyUserIdAndHisUserIdResponse, error) {
	res, err := s.uc.GetLatestMsgByMyUserIdAndHisUserId(ctx, req.GetUserId(), req.GetToUserId())
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetLatestMsgByMyUserIdAndHisUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	msg, err := mapper.MessageToDTO(res)
	if err != nil {
		err := ecode.ConvertErr(err)
		return &v1.GetLatestMsgByMyUserIdAndHisUserIdResponse{
			StatusCode: err.ErrCode,
			StatusMsg:  &err.ErrMsg,
		}, nil
	}

	return &v1.GetLatestMsgByMyUserIdAndHisUserIdResponse{
		StatusCode: ecode.Success.ErrCode,
		StatusMsg:  &ecode.Success.ErrMsg,
		Message:    msg,
	}, nil
}
