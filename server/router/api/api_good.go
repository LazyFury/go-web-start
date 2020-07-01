package api

import (
	"EK-Server/model"
	"EK-Server/util"
	"EK-Server/util/middleware"

	"github.com/labstack/echo"
)

var modelGoods model.Goods

// Init 初始化
func product(g *echo.Group) {
	modelGoods.BaseControll.Model = &modelGoods

	product := g.Group("/goods")

	// 列表
	product.GET("", modelGoods.List)
	// 详情
	product.GET("/:id", func(c echo.Context) error {
		return modelGoods.BaseControll.GetDetail(c, "商品不存在")
	})
	// del
	product.GET("/:id/del", modelGoods.BaseControll.Delete, middleware.AdminJWT)
}

func detail(c echo.Context) error {
	return util.JSONSuccess(c, nil, "")
}
