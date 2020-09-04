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
	conf.GET("/reload", reloadConfig)
	conf.POST("/save", writeConfig)
}

// 读取配置 用于页面显示
func reloadConfig(c echo.Context) error {
	//读取配置文件
	if err := config.Global.ReadConfig(); err != nil {
		return util.JSONErr(c, err, "读取配置失败")
	}
	return util.JSONSuccess(c, nil, "刷新成功")
}

// 写配置 TODO:todo
func writeConfig(c echo.Context) error {
	return util.JSONSuccess(c, nil, "更新设置成功")
}
