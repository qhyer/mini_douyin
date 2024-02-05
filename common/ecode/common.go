package ecode

var (
	Success           = NewErrNo(0, "成功", "Success")
	ParamErr          = NewErrNo(-400, "参数错误", "Wrong Parameter has been given")
	AuthorizeErr      = NewErrNo(-401, "请重新登陆账号", "Authorize failed")
	NothingFound      = NewErrNo(-404, "什么都没找到", "Nothing found")
	ServiceErr        = NewErrNo(-500, "服务器内部错误", "Service internal error")
	ChanFullErr       = NewErrNo(-510, "服务器内部错误", "Channel is full")
	GetSeqIdFailedErr = NewErrNo(-511, "服务器内部错误", "Get sequence id failed")
)
