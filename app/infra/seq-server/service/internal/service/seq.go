package service

import (
	"context"

	"douyin/app/infra/seq-server/service/internal/biz"

	v1 "douyin/api/seq-server/service/v1"
)

type SeqService struct {
	v1.UnimplementedSeqServer
	uc *biz.SeqUsecase
}

func NewSeqService(uc *biz.SeqUsecase) *SeqService {
	return &SeqService{uc: uc}
}

func (s *SeqService) GetID(ctx context.Context, req *v1.GetIDRequest) (*v1.GetIDResponse, error) {
	id, err := s.uc.GetID()
	if err != nil {
		return &v1.GetIDResponse{
			ID:   0,
			IsOk: false,
		}, err
	}
	return &v1.GetIDResponse{
		ID:   id,
		IsOk: true,
	}, nil
}
