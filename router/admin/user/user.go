package user

import (
	"fmt"
	"main/model"
	"main/util"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// Init chushih
func Init(app *echo.Echo, baseURL string) {
	baseURL += "/user"
	app.GET(baseURL+"/list", allUser)
	app.POST(baseURL+"/addUser", addUser)
	app.POST(baseURL+"/updateUser", updateUser)
	app.POST(baseURL+"/delUser", delUser)
	app.GET(baseURL+"/repeatOfName", repeatOfName)
}

func repeatOfName(c echo.Context) error {
	user := new(model.User)
	name := c.QueryParam("name")
	user.Name = name
	result, err := user.Find()
	if err != nil {
		return util.JSON(c, result, "", -1002)
	}

	return util.JSONSuccess(c, result, "没有重复")
}

func addUser(c echo.Context) error {

	user := new(model.User)

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
	msg, err := user.AddUser()
	if err != nil {
		return util.JSONErr(c, err, msg)
	}
	return util.JSONSuccess(c, 1, msg)
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
	user := model.User{}
	return util.JSONSuccess(c, user.GetAllUser(limit, page), "")
}

func updateUser(c echo.Context) error {
	u := new(model.User)
	req := c.Request()
	fmt.Println(req.Header.Get("Content-Type"))

	if err := c.Bind(u); err != nil {
		return util.JSONErr(c, err, "获取数据失败")
	}

	result, err := u.UpdateUser(u.ID, u)

	if err != nil {
		return util.JSONErr(c, result, fmt.Sprintf("%s", err))
	}

	return util.JSONSuccess(c, result, "保存成功")
}

func delUser(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	result, err := user.DelUser()

	if err != nil {
		return util.JSONErr(c, result, fmt.Sprintf("%s", err))
	}

	return util.JSONSuccess(c, result, "删除成功")
}
