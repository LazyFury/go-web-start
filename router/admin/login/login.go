package login

import "github.com/labstack/echo"

import "main/util"

import "main/model"

import "main/util/sha"

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/admin/login"

	login := g.Group(baseURL)
	login.GET("", doLogin)
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
			jwtUser := util.UserInfo{ID: user.ID, Name: user.Name}

			str, _ := util.CreateToken(&jwtUser)

			return util.JSONSuccess(c, str, "登陆成功")
		}
		return util.JSONErr(c, nil, "密码错误")
	}
	return util.JSONErr(c, nil, "用户不存在")
}
