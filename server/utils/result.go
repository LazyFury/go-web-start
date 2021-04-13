package utils

import "github.com/lazyfury/go-web-template/response"

const (
	CustomErrCode response.ErrCode = 123123
)

var ErrCodeText = response.ErrorCodeTextInterface{
	CustomErrCode: "custom err code",
}
