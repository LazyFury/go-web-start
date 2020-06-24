package api

import (
	"EK-Server/model"
	"EK-Server/util"
	"strings"

	"github.com/labstack/echo"
)

var modelPost model.Post

// Init 初始化
func post(g *echo.Group) {

	post := g.Group("/post")

	post.GET("", modelPost.List)

	post.GET("/:id", modelPost.Detail)
	post.DELETE("/:id", modelPost.Delete)

	post.POST("", func(c echo.Context) error {
		empty := func(err interface{}, msg string) error {
			return util.JSONErr(c, err, msg)
		}
		name := c.FormValue("name")
		if strings.Trim(name, " ") == "" {
			return empty(nil, "作者不可以空")
		}
		return util.JSONSuccess(c, name, "添加成功")
	})

}
