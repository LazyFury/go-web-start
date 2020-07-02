package api

import (
	"EK-Server/model"
	"EK-Server/util/middleware"

	"github.com/labstack/echo"
)

var modelGood model.Goods
var modelGoodsCate model.GoodsCate

// Init 初始化
func product(g *echo.Group) {
	modelGood.BaseControll.Model = &modelGood

	product := g.Group("/goods")

	// 列表
	product.GET("", modelGood.List)
	// 详情
	product.GET("/:id", func(c echo.Context) error {
		return modelGood.BaseControll.GetDetail(c, "商品不存在")
	})
	// del
	product.GET("/:id/del", modelGood.BaseControll.Delete, middleware.AdminJWT)
}

// 商品分类
func productCate(g *echo.Group) {
	modelGoodsCate.BaseControll.Model = &modelGoodsCate

	cate := g.Group("/good-cates")

	cate.GET("", modelGoodsCate.List)
	cate.GET("/:id", modelGoodsCate.Detail)
	cate.POST("", modelGoodsCate.Add)
	cate.DELETE("/:id", modelGoodsCate.Delete)
	cate.PUT("/:id", modelGoodsCate.Update)
}
