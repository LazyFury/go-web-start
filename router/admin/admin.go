package admin

import (
	"main/modal"
	"main/router/admin/user"
	"main/util"
	"strconv"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(app *echo.Echo, baseURL string) {
	baseURL += "/admin"
	user.Init(app, baseURL)

	app.GET(baseURL+"", index)

	app.GET(baseURL+"/err", func(c echo.Context) error {
		code := c.QueryParam("code")
		newCode, err := strconv.Atoi(code)
		if err != nil {
			return util.JSONErr(c, err, "")
		}
		return util.JSON(c, "hello", "", newCode)
	})

	app.POST(baseURL+"/delUser", func(c echo.Context) error {
		user := new(modal.User)

		if err := c.Bind(user); err != nil {
			return util.JSONErr(c, err, "参数错误")
		}

		user.DelUser()
		return util.JSONSuccess(c, "", "删除成功")
	})
}

func index(c echo.Context) error {
	return util.JSONSuccess(c, "", "管理后台")
}
