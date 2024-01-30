package ecode

var (
	FavoriteActionTypeInvalidErr = NewErrNo(2001001, "点赞类型错误", "Favorite action type invalid")
	FavoriteRecordNotFoundErr    = NewErrNo(2001002, "点赞记录不存在", "Favorite record not found")
)
