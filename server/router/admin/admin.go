package admin

import (
	"EK-Server/config"
	"EK-Server/router/admin/login"
	"EK-Server/router/admin/post"
	"EK-Server/router/admin/product"
	"EK-Server/router/admin/user"
	"EK-Server/util"
	"net/http"

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
	admin := app.Group(baseURL) //注册admin的中间价
	user.Init(admin)            // 用户模块
	product.Init(admin)
	post.Init(admin)
	admin.GET("", index) //首页

	//测试
	admin.GET("/test", func(c echo.Context) error {
		return util.JSON(c, "", "", 1)
	})

	admin.GET("/settings", func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin/settings.html", config.Global)
	})
}

func index(c echo.Context) error {
	user := c.Get("user")

	return c.Render(http.StatusOK, "admin/index.html", map[string]interface{}{
		"user": user,
	})
}
