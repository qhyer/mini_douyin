package ecode

var (
	RelationFollowSelfBannedErr  = NewErrNo(1002001, "不能关注自己", "Relation follow self banned")
	RelationActionTypeInvalidErr = NewErrNo(1002002, "关注类型错误", "Relation action type invalid")
)
