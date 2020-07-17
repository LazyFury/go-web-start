package api

import (
	"EK-Server/model"
	"EK-Server/model/message"
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
		message.AddUserActionLog(map[string]interface{}{
			"fromID":    uint(1),
			"articleID": uint(1),
			"action":    message.LIKE,
			"remark":    "文章",
		})
		return util.JSONSuccess(c, nil, "点赞成功")
	})

}

// 文章分类
func postCate(g *echo.Group) {
	modelPostCate.BaseControll.Model = &modelPostCate
	modelPostCate.Install(g, "/post-cates")
}
