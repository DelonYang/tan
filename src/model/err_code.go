package model

import "fmt"

type ErrCode uint32

const (
	ErrCodeNo      ErrCode = 0
	ErrCodeInvalid ErrCode = 11
	ErrCodePanic   ErrCode = 9999
)

var ErrCodeName = map[ErrCode]string{
	ErrCodeNo:      "请求成功!",
	ErrCodeInvalid: "invalid request",
	ErrCodePanic:   "服务panic， 请联系管理员",
}

func (e ErrCode) String() string {
	if s, ok := ErrCodeName[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown error code %d", uint32(e))
}
