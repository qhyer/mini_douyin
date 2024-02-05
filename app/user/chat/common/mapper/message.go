package mapper

import (
	dto "douyin/api/user/chat/service/v1"
	do "douyin/app/user/chat/common/entity"
)

func MessageToDTO(msg *do.Message) (*dto.Message, error) {
	if msg == nil {
		return &dto.Message{}, nil
	}
	return &dto.Message{
		Id:         msg.ID,
		ToUserId:   msg.ToUserId,
		FromUserId: msg.FromUserId,
		Content:    msg.Content,
		CreateTime: msg.CreatedAt.Unix(),
	}, nil
}

func MessageToDTOs(msgs []*do.Message) ([]*dto.Message, error) {
	var res []*dto.Message
	for _, msg := range msgs {
		d, err := MessageToDTO(msg)
		if err != nil {
			return nil, err
		}
		res = append(res, d)
	}
	return res, nil
}
