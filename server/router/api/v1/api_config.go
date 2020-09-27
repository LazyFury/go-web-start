package api

import (
	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/Treblex/go-echo-demo/server/util"

	"github.com/labstack/echo/v4"
)

func configRouter(g *echo.Group) {
	conf := g.Group("/config")

	conf.GET("", func(c echo.Context) error {
		return util.JSONSuccess(c, config.Global, "")
	})
	conf.POST("/save", writeConfig)
}

// 写配置 TODO:todo
func writeConfig(c echo.Context) error {
	return util.JSONSuccess(c, nil, "更新设置成功")
}
