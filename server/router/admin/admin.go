package admin

import (
	"EK-Server/config"
	"EK-Server/router/admin/login"
	"EK-Server/router/admin/user"
	"EK-Server/util"
	"net/http"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(app *echo.Group) {

	login.Init(app) //登陆页面

	//admin之下 检测登陆权限
	admin := app.Group("/admin") //注册admin的中间价

	user.Init(admin)     // 用户模块
	admin.GET("", index) //首页

	//测试
	admin.GET("/test", func(c echo.Context) error {
		return util.JSON(c, "", "", 1)
	})

	admin.GET("/settings", func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin/settings.html", map[string]interface{}{
			"g": config.Global,
		})
	})

	admin.GET("/logout", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "退出登陆")
	})

}

func index(c echo.Context) error {
	user := c.Get("user")

	return c.Render(http.StatusOK, "admin/index.html", map[string]interface{}{
		"user": user,
	})
}
