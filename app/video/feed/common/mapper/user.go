package mapper

import (
	account "douyin/api/user/account/service/v1"
	do "douyin/app/video/feed/common/entity"
)

func UserFromAccountDTO(dto *account.User) (*do.User, error) {
	return &do.User{
		ID:       dto.GetId(),
		Name:     dto.GetName(),
		IsFollow: dto.GetIsFollow(),
		Avatar:   dto.GetAvatar(),
	}, nil
}

func UserFromAccountDTOs(dtos []*account.User) ([]*do.User, error) {
	users := make([]*do.User, 0, len(dtos))
	for _, dto := range dtos {
		user, err := UserFromAccountDTO(dto)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
