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
func Init(app *echo.Group) {
	baseURL := "/admin"

	admin := app.Group(baseURL, util.AdminJWT)
	user.Init(admin)

	admin.GET("", index)

	admin.GET("/err", func(c echo.Context) error {
		code := c.QueryParam("code")
		newCode, err := strconv.Atoi(code)
		if err != nil {
			return util.JSONErr(c, err, "")
		}
		return util.JSON(c, "hello", "", newCode)
	})
	admin.GET("/test", func(c echo.Context) error {
		return util.JSON(c, "", "", 1)
	})
}

func index(c echo.Context) error {
	return util.JSONSuccess(c, "", "管理后台")
}
