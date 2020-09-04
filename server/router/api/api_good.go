package api

import (
	"github.com/treblex/go-echo-demo/server/model"

	"github.com/labstack/echo/v4"
)

var modelGood model.Goods
var modelGoodCate model.GoodCate
var modelGoodSku model.GoodSku

// Init 初始化
func product(g *echo.Group) {
	modelGood.BaseControll.Model = &modelGood
	modelGood.Install(g, "/goods")
}

// 商品分类
func productCate(g *echo.Group) {
	modelGoodCate.BaseControll.Model = &modelGoodCate
	modelGoodCate.Install(g, "/good-cates")
}

// 商品库存
func productSku(g *echo.Group) {
	modelGoodSku.BaseControll.Model = &modelGoodSku
	modelGoodSku.Install(g, "/good-skus")
}
