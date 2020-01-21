package admin

import (
	"strconv"
	"suke-go-test/router/admin/login"
	"suke-go-test/router/admin/user"
	"suke-go-test/util"

	"github.com/labstack/echo"
)

// User User
type User struct {
	Name string
}

// Init 初始化
func Init(app *echo.Group) {
	baseURL := "/admin"

	login.Init(app) //登陆页面

	//admin之下 检测登陆权限
	admin := app.Group(baseURL, util.AdminJWT) //注册admin的中间价
	user.Init(admin)                           // 用户模块

	admin.GET("", index) //首页
	//json 测试
	admin.GET("/err", func(c echo.Context) error {
		code := c.QueryParam("code")
		newCode, err := strconv.Atoi(code)
		if err != nil {
			return util.JSONErr(c, err, "")
		}
		return util.JSON(c, "hello", "", newCode)
	})
	//测试
	admin.GET("/test", func(c echo.Context) error {
		return util.JSON(c, "", "", 1)
	})
}

func index(c echo.Context) error {
	user := c.Get("user")
	return util.JSONSuccess(c, user, "管理后台")
}
