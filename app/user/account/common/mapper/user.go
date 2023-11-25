package mapper

import (
	passport "douyin/api/user/passport/service/v1"
	do "douyin/app/user/account/common/entity"
)

func UserFromPassportDTO(u *passport.UserInfo) (*do.User, error) {
	return &do.User{
		ID:              u.GetId(),
		Name:            u.GetName(),
		Avatar:          u.GetAvatar(),
		BackgroundImage: u.GetBackgroundImage(),
		Signature:       u.GetSignature(),
	}, nil
}
