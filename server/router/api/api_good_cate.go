package api

import (
	"EK-Server/util"

	"github.com/labstack/echo"
)

func productCate(g *echo.Group) {
	cate := g.Group("/good-cates")
	cate.GET("", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "")
	})
}
