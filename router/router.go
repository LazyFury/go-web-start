package router

import (
	"main/router/admin"
	"net/http"

	"github.com/labstack/echo"
)

var baseURL string = "/api"

// Start 入口
func Start(e *echo.Echo) {
	// 项目首页
	admin.Init(e, baseURL)

	// 入口
	e.GET(baseURL, func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world！")
	})
}
