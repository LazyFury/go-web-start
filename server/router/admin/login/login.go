package login

import (
	"EK-Server/model"
	"EK-Server/util"
	"EK-Server/util/middleware"
	"EK-Server/util/sha"

	"github.com/labstack/echo"
)

var (
	user = &model.User{}
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/admin/login"
	login := g.Group(baseURL)
	login.GET("", doLogin)
	login.POST("/reg", user.RegController)

}

func doLogin(c echo.Context) error {

	username := c.QueryParam("username")
	if username == "" {
		return util.JSONErr(c, nil, "用户名不可空")
	}
	password := c.QueryParam("password")
	if password == "" {
		return util.JSONErr(c, nil, "用户密码不可空")
	}

	user := model.User{Name: username}

	err := user.Find()
	if err == nil {
		password := sha.EnCode(password)
		if user.Password == password {
			jwtUser := middleware.UserInfo{ID: float64(user.ID), Name: user.Name, IsAdmin: user.IsAdmin}

			str, _ := middleware.CreateToken(&jwtUser)

			return util.JSONSuccess(c, str, "登陆成功")
		}
		return util.JSONErr(c, nil, "密码错误")
	}
	return util.JSONErr(c, nil, "用户不存在")
}
