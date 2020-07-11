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
	post := modelPost.Install(g, "/posts") //list,detail,add,update,delete

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
	modelPostCate.Install(g, "/post-cates")
}
