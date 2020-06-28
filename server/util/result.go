package util

import (
	"net/http"

	"github.com/labstack/echo"
)

// Return 公共返回类型
type returnJSON struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	// Success 返回码 成功
	Success int = 1
	// Error 返回码 失败
	Error int = -1

	// Logout 返回码 为登陆
	Logout int = -101
	// LogTimeOut 返回码 登陆超时
	LogTimeOut int = -102
	// BindWeChat 绑定微信
	BindWeChat int = -103
	// BingPhone 绑定手机号
	BingPhone int = -104
)

var (

	// 自定义错误码
	errCode map[int]string = map[int]string{
		// 正常
		Success: "请求成功",
		Error:   "请求错误",
		// 登陆
		Logout:     "请先登录～",
		LogTimeOut: "登陆超时失效～",
		BindWeChat: "请先绑定微信账号",
		BingPhone:  "请先绑定手机号",

		// 无权限
		http.StatusUnauthorized: "无权限操作!",
	}
)

// JSONBase 增加了httpcode参数
func JSONBase(c echo.Context, data interface{}, msg string, code int, httpCode int) error {
	if msg == "" {
		if errCode[code] != "" {
			msg = errCode[code]
		} else {
			msg = "未知错误码"
		}
	}
	return c.JSONPretty(httpCode, &returnJSON{
		Code: code,
		Msg:  msg,
		Data: data,
	}, " ")
}

// JSON 自定义通用返回方法
func JSON(c echo.Context, data interface{}, msg string, code int) error {
	return JSONBase(c, data, msg, code, http.StatusOK)
}

// JSONErr 默认code -1
func JSONErr(c echo.Context, data interface{}, msg string) error {
	return JSON(c, data, msg, Error)
}

// JSONErrJustCode err
func JSONErrJustCode(c echo.Context, code int) error {
	return JSON(c, nil, "", code)
}

// JSONSuccess 默认code 1
func JSONSuccess(c echo.Context, data interface{}, msg string) error {
	return JSON(c, data, msg, Success)
}
