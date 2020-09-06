package api

import (
	"strings"

	"github.com/Treblex/go-echo-demo/server/middleware"
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/util"
	"github.com/Treblex/go-echo-demo/server/util/sha"

	"github.com/labstack/echo/v4"
)

func login(g *echo.Group) {
	login := g.Group("/login")

	login.POST("", doLogin)

	login.POST("/reg", modelUser.RegController)

	login.GET("/init_admin", initAdmin)

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

func initAdmin(c echo.Context) error {

	ip := util.ClientIP(c)
	if ip != "127.0.0.1" {
		return util.JSONErr(c, nil, "")
	}

	secret := c.QueryParam("secret")
	if secret != "fqEeEPlgFECywkwqVMoCEmBzmRmFPZwt" {
		return util.JSONErr(c, nil, "密钥错误")
	}
	db := model.DB

	a := &model.User{Name: "admin", IsAdmin: 1}
	if findAdmin := db.Where(a).Find(a).RowsAffected; findAdmin >= 1 {
		a.Password = sha.DeCode(a.Password)
		return util.JSONSuccess(c, a, "")
	}
	pwd := util.RandStringBytes(32)

	admin := &model.User{
		Name:     "admin",
		Password: pwd,
		IsAdmin:  1,
	}

	if err := db.Save(admin).Error; err != nil {
		return util.JSONErr(c, err, "")
	}
	admin.Password = sha.DeCode(admin.Password)
	return util.JSONErr(c, admin, "注册成功")
}
