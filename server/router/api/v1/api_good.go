package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

var modelGood model.Goods
var modelGoodCate model.GoodCate
var modelGoodSku model.GoodSku

// Init 初始化
func product(g *gin.RouterGroup) {
	modelGood.BaseControll.Model = &modelGood
	modelGood.Install(g, "/goods")
}

// 商品分类
func productCate(g *gin.RouterGroup) {
	modelGoodCate.BaseControll.Model = &modelGoodCate
	modelGoodCate.Install(g, "/good-cates")
}

// 商品库存
func productSku(g *gin.RouterGroup) {
	modelGoodSku.BaseControll.Model = &modelGoodSku
	modelGoodSku.Install(g, "/good-skus")
}
