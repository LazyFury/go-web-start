package api

import (
	"EK-Server/middleware"
	"EK-Server/model"
	"EK-Server/util"
	"EK-Server/util/sha"
	"strings"

	"github.com/labstack/echo"
)

func login(g *echo.Group) {
	login := g.Group("/login")

	login.POST("", doLogin)

	login.POST("/reg", modelUser.RegController)

}

func doLogin(c echo.Context) error {
	var u = &struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(u); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	u.UserName = strings.Trim(u.UserName, " ")
	if u.UserName == "" {
		return util.JSONErr(c, nil, "用户昵称不可空")
	}

	u.Password = strings.Trim(u.Password, " ")
	if u.Password == "" {
		return util.JSONErr(c, nil, "用户密码不可空")
	}

	user := model.User{Name: u.UserName}

	err := user.HasUser()
	if err != nil {
		return util.JSONErr(c, nil, "用户不存在")
	}
	password := sha.EnCode(u.Password)
	if user.Password == password {
		jwtUser := middleware.UserInfo{ID: float64(user.ID), Name: user.Name, IsAdmin: user.IsAdmin > 0}
		str, _ := middleware.CreateToken(&jwtUser)
		return util.JSONSuccess(c, str, "登陆成功")
	}
	return util.JSONErr(c, nil, "密码错误")
}
