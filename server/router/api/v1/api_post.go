package api

import (
	"github.com/Treblex/go-echo-demo/server/model"

	"github.com/labstack/echo/v4"
)

// Init 初始化
func post(g *echo.Group) {
	var modelPost = model.NewArticle()
	modelPost.Install(g, "/posts") //list,detail,add,update,delete
}

// 文章分类
func postCate(g *echo.Group) {
	modelPostCate := model.NewArticleCate()
	modelPostCate.Install(g, "/post-cates")
}

func postRec(g *echo.Group) {
	modelPostRec := model.NewArticleRec()
	modelPostRec.Install(g, "/post-rec")
}

func postTag(g *echo.Group) {
	modelPostTag := model.NewArticleTag()
	modelPostTag.Install(g, "/post-tags")
}
