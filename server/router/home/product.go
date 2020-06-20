package home

import (
	"EK-Server/model"
	"EK-Server/util"

	"github.com/labstack/echo"
)

var (
	goods = model.Goods{}
)

// Init 初始化
func product(g *echo.Group) {
	baseURL := "/product"
	product := g.Group(baseURL)

	product.GET("", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "hello")
	})

	product.GET("/list", goods.List)
	product.GET("/detail", detail)

}

func detail(c echo.Context) error {
	return util.JSONSuccess(c, nil, "")
}
