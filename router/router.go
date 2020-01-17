package router

import (
	"main/router/admin"
	"main/util"
	"net/http"

	"github.com/labstack/echo"
)

// Start 入口
func Start(e *echo.Group) {
	// 项目首页
	admin.Init(e)

	// 入口
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world！")
	}, util.UserJWT)
}
