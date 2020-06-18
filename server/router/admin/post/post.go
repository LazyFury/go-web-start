package post

import (
	"EK-Server/model"
	"EK-Server/util"
	"net/http"

	"github.com/labstack/echo"
)

var modelPost model.Post

// Init 初始化
func Init(g *echo.Group) {
	post := g.Group("/post")

	post.GET("", func(c echo.Context) error {
		posts, _ := modelPost.List(c)
		page := model.GeneratePaging(posts.PageCount, posts.PageNow, "/admin/post?page=")
		return c.Render(http.StatusOK, "admin/post/post.html", map[string]interface{}{
			"posts": posts,
			"page":  page,
		})
	})
	// 添加文章
	post.GET("/add", func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin/post/add.html", map[string]interface{}{})
	})
	post.POST("/add", addArticle)
}

// 添加文章
func addArticle(c echo.Context) error {
	article := &model.Post{}
	if err := c.Bind(article); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}
	db := model.DB
	db.NewRecord(article)

	row := db.Create(article)

	if row.Error != nil {
		return util.JSONErr(c, row.Error, "保存失败")
	}

	if row.RowsAffected <= 0 {
		return util.JSONErr(c, nil, "保存失败 没有更改")
	}

	return util.JSONSuccess(c, nil, "提交成功")
}
