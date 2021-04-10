package utils

import (
	"net/http"
	"time"

	"github.com/lazyfury/go-web-template/response"
)

// Result Result
type Result struct {
	Code    response.ErrCode `json:"code"`
	Message string           `json:"message"`
	Data    interface{}      `json:"data"`
	BuildBy time.Time        `json:"build_by"`
}

const (
	// Success Success
	Success response.ErrCode = 1
	// Errors 失败
	Errors response.ErrCode = -1
	// NoRoute NoRoute
	NoRoute response.ErrCode = http.StatusNotFound
	// NoMethod NoMethod
	NoMethod response.ErrCode = http.StatusMethodNotAllowed
)
const (
	// LoginSuccess 登陆成功
	LoginSuccess response.ErrCode = iota + 100
)
const (
	// AuthedError 认证失败
	AuthedError response.ErrCode = -iota - 100
	// NotFound 没有数据
	NotFound
	// RepeatEmail 邮箱已存在
	RepeatEmail
	// RepeatUserName 用户名已存在
	RepeatUserName
	// BindJSONErr 绑定json失败
	BindJSONErr
)

// ErrorCodeText 错误提示
var ErrorCodeText = map[response.ErrCode]string{
	// base
	Success: "获取成功",
	Errors:  "遇到错误",

	// business
	LoginSuccess:   "登陆成功",
	AuthedError:    "登陆超时",
	NotFound:       "没有数据",
	RepeatEmail:    "邮箱已存在",
	RepeatUserName: "用户名已存在",
	BindJSONErr:    "绑定失败,请检查参数",

	// system
	NoRoute:  "路由不存在",
	NoMethod: "方法不存在",
}

// BuildBy BuildBy
var BuildBy = time.Now()

// StatusText StatusText
func StatusText(code response.ErrCode) string {
	msg := ErrorCodeText[code]
	if msg == "" {
		msg = http.StatusText(int(code))
	}
	if msg == "" {
		msg = "未知错误码"
	}
	return msg
}

// JSON JSON
func JSON(code response.ErrCode, message string, data interface{}) Result {
	if message == "" {
		message = StatusText(code)
	}
	return Result{
		Code:    code,
		Message: message,
		Data:    data,
		BuildBy: BuildBy,
	}
}

// JSONSuccess 成功
func JSONSuccess(message string, data interface{}) Result {
	return JSON(Success, message, data)
}

// JSONError JSONError
func JSONError(message string, data interface{}) Result {
	return JSON(Errors, message, data)
}
