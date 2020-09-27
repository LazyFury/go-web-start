package api

import (
	"github.com/Treblex/go-echo-demo/server/model"

	"github.com/labstack/echo/v4"
)

var modelUser model.User

// 用户API
func user(g *echo.Group) {
	modelUser.BaseControll.Model = &modelUser
	modelUser.Install(g, "/users")
}
