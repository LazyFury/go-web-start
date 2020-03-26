package product

import (
	"EK-Server/config"
	"EK-Server/model"
	"EK-Server/util"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/product"
	product := g.Group(baseURL)

	product.GET("", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "hello")
	})

	product.GET("/list", productList)
}

func productList(c echo.Context) error {
	return util.JSONSuccess(c, model.DataBaselimit(10, 1, &model.Goods{}, &[]model.GoodsList{}, config.Global.TablePrefix+"_goods",
		"id desc"), "")
}
