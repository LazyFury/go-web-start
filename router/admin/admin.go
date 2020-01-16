package admin

import (
	"main/router/admin/user"
	"main/util"
	"strconv"

	"github.com/labstack/echo"
)

// User User
type User struct {
	Name string
}

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
	app.GET(baseURL+"/test", func(c echo.Context) error {
		return util.JSON(c, "", "", 1)
	})
}

func index(c echo.Context) error {
	return util.JSONSuccess(c, "", "管理后台")
}
