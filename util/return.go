package util

import "github.com/labstack/echo"
import "net/http"

// Return 公共返回类型
type returnJSON struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Success int= 1
	Error int= -1
	Logout int =-100
	LogTimeOut int= -101
	
	// 自定义错误码
	errCode map[int]string = map[int]string{
		// 正常
		1:    "请求成功",
		-1:   "请求错误,仅提示类型，应该返回错误原因，需要操作的设置另外的错误码",
		// 登陆
		-100: "用户未登陆",
		-101: "用户登陆超时",

		// 需要客户端指定操作
		-1002: "用户名已存在,请尝试其他",
		// 4开头微信
		-4001: "微信授权登陆失败",
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
	return c.JSON(httpCode, &returnJSON{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// JSON 自定义通用返回方法
func JSON(c echo.Context, data interface{}, msg string, code int) error {
	return JSONBase(c, data, msg, code, http.StatusOK)
}

// JSONErr 默认code -1
func JSONErr(c echo.Context, data interface{}, msg string) error {
	return JSON(c, data, msg, -1)
}

// JSONSuccess 默认code 1
func JSONSuccess(c echo.Context, data interface{}, msg string) error {
	return JSON(c, data, msg, 1)
}
