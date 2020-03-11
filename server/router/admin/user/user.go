package user

import (
	"fmt"
	"strconv"
	"suke-go-test/model"
	"suke-go-test/util"
	"suke-go-test/util/sha"
	"time"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/user"
	user := g.Group(baseURL)
	user.GET("/list", allUser)
	user.POST("/addUser", addUser)
	user.POST("/updateUser", updateUser)
	user.POST("/delUser", delUser)
	user.GET("/repeatOfName", repeatOfName)
	user.GET("/repeatOfEmail", repeatOfEmail)
	user.POST("/frozen", frozen)
}

func frozen(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(&u); err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}

	db := model.DB
	row := db.Model(&model.User{ID: u.ID}).Update("status", u.Status)
	if row.Error != nil {
		return util.JSONErr(c, nil, "操作失败")
	}

	if u.Status == 0 {
		return util.JSONSuccess(c, nil, "冻结用户")
	}

	return util.JSONSuccess(c, nil, "解冻用户")
}

func repeatOfEmail(c echo.Context) error {
	user := new(model.User)
	email := c.QueryParam("email")
	user.Email = email
	err := user.Find()
	if err != nil {
		return util.JSONSuccess(c, nil, "没有重复")

	}
	return util.JSON(c, nil, "邮箱已被使用,尝试找回密码或者使用其他邮箱", -1)

}

func repeatOfName(c echo.Context) error {
	user := new(model.User)
	name := c.QueryParam("name")
	user.Name = name
	err := user.Find()
	if err != nil {
		return util.JSONSuccess(c, nil, "没有重复")
	}
	return util.JSON(c, nil, "用户名已存在", -1002)
}

func addUser(c echo.Context) error {

	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	user.Password = sha.EnCode(user.Password)

	req := c.Request()
	ua := req.UserAgent()
	ip := util.ClientIP(c)
	user.IP = ip
	user.Ua = ua
	user.CreateTime = util.LocalTime{Time: time.Now()}
	user.LoginTime = util.LocalTime{Time: time.Now()}
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
		limit = 10
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	user := model.User{}
	return util.JSONSuccess(c, user.GetAllUser(limit, page), "")
}

func updateUser(c echo.Context) error {
	u := new(model.User)
	req := c.Request()
	fmt.Println(req.Header.Get("Content-Type"))

	if err := c.Bind(u); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	err := u.UpdateUser(u.ID, u)

	if err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}

	return util.JSONSuccess(c, nil, "保存成功")
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
