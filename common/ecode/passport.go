package ecode

var (
	UserNotExistErr     = NewErrNo(1001001, "用户不存在", "User not exists")
	UserAlreadyExistErr = NewErrNo(1001002, "用户已存在", "User already exists")
	WrongPasswordErr    = NewErrNo(1001003, "密码错误", "Wrong password")
)
