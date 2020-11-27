package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

// Init 初始化
func post(g *gin.RouterGroup) {
	// var modelPost = model.NewArticle()
	// modelPost.Install(g, "/posts") //list,detail,add,update,delete
}

// 文章分类
func postCate(g *gin.RouterGroup) {
	modelPostCate := model.NewArticleCate()
	modelPostCate.Install(g, "/post-cates")
}

func postRec(g *gin.RouterGroup) {
	// modelPostRec := model.NewArticleRec()
	// modelPostRec.Install(g, "/post-rec")
}

func postTag(g *gin.RouterGroup) {
	modelPostTag := model.NewArticleTag()
	modelPostTag.Install(g, "/post-tags")
}
