package api

import (
	"EK-Server/model"
	"EK-Server/util"
	"EK-Server/util/middleware"
	"strings"

	"github.com/labstack/echo"
)

var modelPost model.Post

// Init 初始化
func post(g *echo.Group) {
	post := g.Group("/posts")
	//list
	post.GET("", modelPost.List)
	//detail
	post.GET("/:id", modelPost.Detail)
	//del
	post.DELETE("/:id", modelPost.Delete, middleware.AdminJWT)
	//add todo:整理新建和更新到同一个方法
	post.POST("", func(c echo.Context) error {
		empty := func(err interface{}, msg string) error {
			return util.JSONErr(c, err, msg)
		}
		name := c.FormValue("name")
		if strings.Trim(name, " ") == "" {
			return empty(nil, "作者不可以空")
		}
		return util.JSONSuccess(c, name, "添加成功")
	}, middleware.AdminJWT)
	// Update
	post.PUT("/:id", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "")
	})

	// Actions
	// 点赞文章
	post.GET("/:id/actions/like", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "点赞成功")
	}, middleware.AdminJWT)

	post.GET("/:id/actions/unlike", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "")
	}, middleware.AdminJWT)

}
