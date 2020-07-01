package api

import (
	"EK-Server/model"

	"github.com/labstack/echo"
)

var modelGoodsCate model.GoodsCate

func productCate(g *echo.Group) {
	modelGoodsCate.BaseControll.Model = &modelGoodsCate

	cate := g.Group("/good-cates")

	cate.GET("", modelGoodsCate.List)
	cate.GET("/:id", modelGoodsCate.Detail)
	cate.POST("", modelGoodsCate.Add)
	cate.DELETE("/:id", modelGoodsCate.Delete)
	cate.PUT("/:id", modelGoodsCate.Update)
}
