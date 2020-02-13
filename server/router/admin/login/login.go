package login

import (
	"fmt"
	"github.com/labstack/echo"
	"suke-go-test/model"
	"suke-go-test/util"
	"suke-go-test/util/sha"
	"time"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/admin/login"
	login := g.Group(baseURL)
	login.GET("", doLogin)
	login.GET("/reg", reg)
}
func reg(c echo.Context) (err error) {
	user := new(model.User)

	user.Name = c.QueryParam("username")
	if user.Name == "" {
		return util.JSONErr(c, nil, "用户名不可空")
	}
	user.Password = c.QueryParam("password")
	if user.Password == "" {
		return util.JSONErr(c, nil, "用户密码不可空")
	}
	// 密码
	user.Password = sha.EnCode(user.Password)
	// 获取ip ua
	req := c.Request()
	ua := req.UserAgent()
	ip := util.ClientIP(c)
	user.IP = ip
	user.Ua = ua
	// 更新时间
	user.CreateTime = util.LocalTime{Time: time.Now()}
	user.LoginTime = util.LocalTime{Time: time.Now()}
	// 状态可用
	user.Status = 1

	fmt.Println(user)
	msg, err := user.AddUser()
	if err != nil {
		return util.JSONErr(c, err, msg)
	}
	return util.JSONSuccess(c, nil, msg)
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
			jwtUser := util.UserInfo{ID: float64(user.ID), Name: user.Name}

			str, _ := util.CreateToken(&jwtUser)

			return util.JSONSuccess(c, str, "登陆成功")
		}
		return util.JSONErr(c, nil, "密码错误")
	}
	return util.JSONErr(c, nil, "用户不存在")
}
