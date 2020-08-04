package app

import (
	"EK-Server/util"
	"net/http"

	"github.com/labstack/echo"
)

// 错误处理
func httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	msg = http.StatusText(code)
	// 如果是浏览器
	// req := c.Request()
	// reqAccept := strings.Split(req.Header.Get("Accept"), ",")[0]
	// if reqAccept == "text/html" {
	// 	c.Logger().Error(c.Render(code, "error.html", map[string]interface{}{
	// 		"msg":  msg,
	// 		"code": code,
	// 	}))
	// 	return
	// }
	// 如果是ajax
	c.Logger().Error(util.JSONBase(c, nil, msg, code, code))
}
