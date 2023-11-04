package mapper

import (
	do "douyin/app/user/passport/common/entity"
	po "douyin/app/user/passport/common/model"
	vo "douyin/app/user/passport/service/api/v1"
)

func UserFromPO(u *po.User) (us *do.User, err error) {
	us = &do.User{
		ID:              u.ID,
		Name:            u.Name,
		Avatar:          "https://p3-pc.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-i-0813_442293a2db8f4c719ecbd8fb6e35ae21.jpeg",
		BackgroundImage: "https://lf6-cdn-tos.douyinstatic.com/obj/venus/pc-theme/spring/20230411-162827.jpeg",
		Signature:       "暂无个性签名",
	}
	err = us.Validate()
	if err != nil {
		return nil, err
	}
	return us, nil
}

func UsersFromPO(us []*po.User) (dos []*do.User, err error) {
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
	us = &po.User{
		ID:                u.ID,
		Name:              u.Name,
		EncryptedPassword: u.EncryptedPassword,
	}
	return us, nil
}

func UserToVO(u *do.User) (us *vo.UserInfo, err error) {
	us = &vo.UserInfo{
		Id:              u.ID,
		Name:            u.Name,
		Avatar:          &u.Avatar,
		BackgroundImage: &u.BackgroundImage,
		Signature:       &u.Signature,
	}
	return us, nil
}

func UsersToVO(us []*do.User) (vos []*vo.UserInfo, err error) {
	for _, u := range us {
		us, err := UserToVO(u)
		if err != nil {
			return nil, err
		}
		vos = append(vos, us)
	}
	return vos, nil

}
