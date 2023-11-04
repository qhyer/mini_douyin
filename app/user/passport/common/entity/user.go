package entity

import (
	"douyin/common/ecode"
)

type User struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encrypted_password"`
	Avatar            string `json:"avatar"`
	BackgroundImage   string `json:"background_image"`
	Signature         string `json:"signature"`
}

func (u *User) Validate() error {
	if len(u.Name) == 0 {
		// 用户名不能为空
		return ecode.NewErrNo(ecode.ParamErr.ErrCode, "请填写用户名", "Username cannot be empty")
	}
	if len(u.Password) == 0 {
		// 密码不能为空
		return ecode.NewErrNo(ecode.ParamErr.ErrCode, "请填写密码", "Password cannot be empty")
	}
	if len(u.Name) > 32 {
		// 用户名过长
		return ecode.NewErrNo(ecode.ParamErr.ErrCode, "用户名不能超过32个字符", "Username is too long")
	}
	if len(u.Password) > 32 {
		// 密码过长
		return ecode.NewErrNo(ecode.ParamErr.ErrCode, "密码不能超过32个字符", "Password is too long")
	}
	return nil
}
