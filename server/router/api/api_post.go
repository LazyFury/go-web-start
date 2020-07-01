package api

import (
	"EK-Server/model"
	"EK-Server/util"
	"EK-Server/util/middleware"

	"github.com/labstack/echo"
)

var modelPost model.Articles

// Init 初始化
func post(g *echo.Group) {
	modelPost.BaseControll.Model = &modelPost
	post := g.Group("/posts")

	//list
	post.GET("", modelPost.List)
	//detail
	post.GET("/:id", modelPost.Detail)
	//添加内容
	post.POST("", modelPost.Add)
	//del
	post.DELETE("/:id", modelPost.Delete)
	// Update 更新内容
	post.PUT("/:id", modelPost.Update)

	// Actions
	// 点赞文章
	post.GET("/:id/actions/like", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "点赞成功")
	}, middleware.AdminJWT)

}
