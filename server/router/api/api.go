package api

import "github.com/labstack/echo"

// Init  api Version 1.0 初始化
func Init(g *echo.Group) {
	apiV1 := g.Group("/api/v1")

	post(apiV1)        //文章
	product(apiV1)     //商品
	productCate(apiV1) //商品分类
}
