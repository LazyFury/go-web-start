package v2

import (
	"github.com/Treblex/go-echo-demo/server/util"
	"github.com/labstack/echo/v4"
)

// Init 初始化
func Init(g *echo.Group) {
	api := g.Group("/v2")

	api.GET("", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "wellcome!")
	})
}
