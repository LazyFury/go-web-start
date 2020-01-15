package user

import (
	"fmt"
	"main/modal"
	"main/util"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func Init(app *echo.Echo, baseURL string) {
	baseURL += "/user"
	app.GET(baseURL+"/list", allUser)
	app.POST(baseURL+"/addUser", addUser)
	app.POST(baseURL+"/updateUser", updateUser)
}
func addUser(c echo.Context) error {

	user := new(modal.User)

	if err := c.Bind(user); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	req := c.Request()
	ua := req.UserAgent()
	ip := req.RemoteAddr
	user.IP = ip
	user.Ua = ua
	user.CreateTime = util.LocalTime{time.Now()}
	user.LoginTime = util.LocalTime{time.Now()}
	user.Status = 1

	fmt.Println(user)
	user.AddUser()

	return util.JSONSuccess(c, 1, "添加成功")
}
func allUser(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return util.JSONErr(c, "", "分页参数不正确")
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return util.JSONErr(c, "", "分页页码不正确")
	}
	user := modal.User{}
	return util.JSONSuccess(c, user.GetAllUser(limit, page), "")
}

func updateUser(c echo.Context) error {
	u := new(modal.User)
	req := c.Request()
	fmt.Println(req.Header.Get("Content-Type"))

	if err := c.Bind(u); err != nil {
		return util.JSONErr(c, err, "获取数据失败")
	}

	row, err := u.UpdateUser(u.ID, u)

	if err != nil {
		return util.JSONErr(c, err, "更新失败")
	}

	return util.JSONSuccess(c, "", row)
}
