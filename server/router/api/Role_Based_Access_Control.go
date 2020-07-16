package api

import (
	"EK-Server/util"
	"EK-Server/util/middleware"

	"github.com/labstack/echo"
)

var (
	rbacAdmin  = middleware.AdminJWT
	rbacUser   = middleware.UserJWT
	rbacAuthor = middleware.UserJWT
)

func login(g *echo.Group) {
	login := g.Group("/login")

	login.POST("", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "登陆成功")
	})

	login.POST("/reg", modelUser.RegController)

}
