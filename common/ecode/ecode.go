package ecode

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode   int32  `json:"err_code"`
	ErrMsg    string `json:"err_msg"`
	ErrReason string `json:"-"`
}

// Err represents an error
type Err struct {
	ErrCode   int32
	ErrMsg    string
	ErrReason string
	Err       error
}

func NewErrNo(code int32, msg string, reason string) ErrNo {
	return ErrNo{code, msg, reason}
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s, err_reason=%s", e.ErrCode, e.ErrMsg, e.ErrReason)
}

func NewErr(errno *ErrNo, err error) *Err {
	return &Err{ErrCode: errno.ErrCode, ErrMsg: errno.ErrMsg, ErrReason: errno.ErrReason, Err: err}
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, reason: %s, error: %s", err.ErrCode, err.ErrMsg, err.ErrReason, err.Err)
}

func DecodeErr(err error) (int32, string) {
	if err == nil {
		return Success.ErrCode, Success.ErrMsg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.ErrCode, typed.ErrMsg
	case *ErrNo:
		return typed.ErrCode, typed.ErrMsg
	default:
	}

	return ServiceErr.ErrCode, err.Error()
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	if err == nil {
		return Success
	}
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
