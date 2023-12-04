package mapper

import (
	account "douyin/api/user/account/service/v1"
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

func UserFromPassportDTOs(us []*passport.UserInfo) ([]*do.User, error) {
	var users []*do.User
	for _, u := range us {
		user, err := UserFromPassportDTO(u)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func UserToDTO(u *do.User) (*account.User, error) {
	return &account.User{
		Id:              u.ID,
		Name:            u.Name,
		FollowCount:     &u.FollowCount,
		FollowerCount:   &u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          &u.Avatar,
		BackgroundImage: &u.BackgroundImage,
		Signature:       &u.Signature,
		TotalFavorited:  &u.TotalFavorited,
		WorkCount:       &u.WorkCount,
		FavoriteCount:   &u.FavoriteCount,
		Message:         &u.Message,
		MsgType:         &u.MsgType,
	}, nil
}

func UserToDTOs(us []*do.User) ([]*account.User, error) {
	var users []*account.User
	for _, u := range us {
		user, err := UserToDTO(u)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
