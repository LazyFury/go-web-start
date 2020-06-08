package post

import (
	"EK-Server/model"
	"EK-Server/util"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/post"
	post := g.Group(baseURL)

	post.GET("", func(c echo.Context) error {
		articles := model.DataBaselimit(10, 1, map[string]interface{}{}, &[]model.User{}, "users", "id desc")
		return c.Render(http.StatusOK, "home/post/list.html", map[string]interface{}{
			"articles": articles,
		})
	})

	post.POST("/add", func(c echo.Context) error {
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
