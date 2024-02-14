package mapper

import (
	dto "douyin/api/user/passport/service/v1"
	do "douyin/app/user/passport/common/entity"
	po "douyin/app/user/passport/common/model"
	"fmt"
)

func UserFromPO(u *po.User) (us *do.User, err error) {
	if u == nil {
		return &do.User{}, fmt.Errorf("user is nil")
	}
	us = &do.User{
		ID:                u.ID,
		Name:              u.Name,
		EncryptedPassword: u.EncryptedPassword,
		Avatar:            "https://p3-pc.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813_442293a2db8f4c719ecbd8fb6e35ae21.jpeg",
		BackgroundImage:   "https://lf6-cdn-tos.douyinstatic.com/obj/venus/pc-theme/spring/20230411-162827.jpeg",
		Signature:         "暂无个性签名",
	}
	return us, nil
}

func UserFromPOs(us []*po.User) (dos []*do.User, err error) {
	for _, u := range us {
		us, err := UserFromPO(u)
		if err != nil {
			return nil, err
		}
		dos = append(dos, us)
	}
	return dos, nil
}

func UserToPO(u *do.User) (us *po.User, err error) {
	if u == nil {
		return &po.User{}, fmt.Errorf("user is nil")
	}
	us = &po.User{
		ID:                u.ID,
		Name:              u.Name,
		EncryptedPassword: u.EncryptedPassword,
	}
	return us, nil
}

func UserToDTO(u *do.User) (us *dto.UserInfo, err error) {
	if u == nil {
		return &dto.UserInfo{}, fmt.Errorf("user is nil")
	}
	us = &dto.UserInfo{
		Id:              u.ID,
		Name:            u.Name,
		Avatar:          &u.Avatar,
		BackgroundImage: &u.BackgroundImage,
		Signature:       &u.Signature,
	}
	return us, nil
}

func UserToDTOs(us []*do.User) (dtos []*dto.UserInfo, err error) {
	for _, u := range us {
		us, err := UserToDTO(u)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, us)
	}
	return dtos, nil
}
