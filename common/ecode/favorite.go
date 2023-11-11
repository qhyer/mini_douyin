package ecode

var (
	FavoriteRecordAlreadyExistErr = NewErrNo(2001001, "已经点赞过该视频", "Favorite record already exists")
	FavoriteRecordNotExistErr     = NewErrNo(2001002, "未点赞过该视频", "Favorite record not exists")
)
