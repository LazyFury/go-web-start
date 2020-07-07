package api

import (
	"EK-Server/model"
	"EK-Server/util"

	"github.com/labstack/echo"
)

var modelPost model.Articles
var modelPostCate model.ArticlesCate

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
	}, rbacAdmin)

	post.GET("-actions/count", modelPost.Count)
}

// 文章分类
func postCate(g *echo.Group) {
	modelPostCate.BaseControll.Model = &modelPostCate

	cate := g.Group("/post-cates")

	cate.GET("", modelPostCate.ListWithOutPaging)
	cate.GET("/:id", modelPostCate.Detail)

	cate.POST("", modelPostCate.Add)
	cate.PUT("/:id", modelPostCate.Update)
	cate.DELETE("/:id", modelPostCate.Delete)
}
