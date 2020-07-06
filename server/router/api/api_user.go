package api

import (
	"EK-Server/model"
	"EK-Server/util"

	"github.com/labstack/echo"
)

var modelUser model.User

// 用户API
func user(g *echo.Group) {
	modelUser.BaseControll.Model = &modelUser

	user := g.Group("/users")
	user.GET("", modelUser.List)
	user.GET("/:id", modelUser.Detail)

	user.POST("", modelUser.Add)
	user.PUT("/:id", modelUser.Update)
	user.DELETE("/:id", modelUser.Delete)

	// actions   url like: /users-actions/repeat-of-name
	user.GET("-actions/repeat-of-name", repeatOfName)
	// 统计
	user.GET("-actions/count", modelUser.Count)
}

func repeatOfEmail(c echo.Context) error {
	user := new(model.User)
	email := c.QueryParam("email")
	user.Email = email
	err := user.HasUser()
	if err != nil {
		return util.JSONSuccess(c, nil, "没有重复")

	}
	return util.JSON(c, nil, "邮箱已被使用,尝试找回密码或者使用其他邮箱", -1)
}

func repeatOfName(c echo.Context) error {
	user := new(model.User)
	name := c.QueryParam("name")
	user.Name = name
	err := user.HasUser()
	if err != nil {
		return util.JSONSuccess(c, nil, "没有重复")
	}

	return util.JSON(c, nil, "用户名已存在", -1002)
}
