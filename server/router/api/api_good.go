package api

import (
	"EK-Server/model"

	"github.com/labstack/echo"
)

var modelGood model.Goods
var modelGoodsCate model.GoodsCate

// Init 初始化
func product(g *echo.Group) {
	modelGood.BaseControll.Model = &modelGood
	modelGood.Install(g, "/products")
}

// 商品分类
func productCate(g *echo.Group) {
	modelGoodsCate.BaseControll.Model = &modelGoodsCate
	modelGoodsCate.Install(g, "/good-cates")
}
