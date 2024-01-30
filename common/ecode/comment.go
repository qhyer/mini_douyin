package ecode

var (
	CommentActionTypeInvalidErr = NewErrNo(2003001, "评论类型错误", "Comment action type invalid")
	CommentRecordNotFoundErr    = NewErrNo(2003002, "评论记录不存在", "Comment record not found")
)
