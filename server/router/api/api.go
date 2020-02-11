package api

import (
	"suke-go-test/util"

	"github.com/labstack/echo"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "/api"
	api := g.Group(baseURL)

	api.GET("/add", addAPI)
}

// 添加API
func addAPI(c echo.Context) (err error) {
	name := c.QueryParam("name")
	if name == "" {
		return util.JSONErr(c, nil, "API名称不可空")
	}
	data := c.QueryParam("data")
	if data == "" {
		return util.JSONErr(c, nil, "配置内容不可空")
	}

	return util.JSONSuccess(c, nil, "cool! it`s work")
}
