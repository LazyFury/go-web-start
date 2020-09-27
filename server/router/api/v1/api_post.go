package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/util"
	"github.com/Treblex/go-echo-demo/server/util/customtype/message"

	"github.com/labstack/echo/v4"
)

var modelPost model.Articles
var modelPostCate model.ArticlesCate
var modelPostRec model.ArticlesRec

// Init 初始化
func post(g *echo.Group) {
	modelPost.BaseControll.Model = &modelPost
	post := modelPost.Install(g, "/posts") //list,detail,add,update,delete

	// Actions
	post.GET("/:id/actions/like", likeArticle)

}

// 文章分类
func postCate(g *echo.Group) {
	modelPostCate.BaseControll.Model = &modelPostCate
	modelPostCate.Install(g, "/post-cates")
}

func postRec(g *echo.Group) {
	modelPostRec.BaseControll.Model = &modelPostRec
	modelPostRec.Install(g, "/post-rec")
}

// 点赞文章
func likeArticle(c echo.Context) error {
	post := &struct {
		*model.Articles
	}{}

	db := model.DB

	if db.Model(&model.Articles{}).Where(map[string]interface{}{
		"id": c.Param("id"),
	}).First(post).RecordNotFound() {
		return util.JSONSuccess(c, nil, "文章不存在")
	}

	// 添加文章消息
	modelMessage.AddArticleLog(uint(1), post.ID, message.LIKE)

	return util.JSONSuccess(c, nil, "点赞成功")
}
