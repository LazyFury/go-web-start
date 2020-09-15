package app

import (
	"fmt"

	"github.com/Treblex/go-echo-demo/server/util"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func sessionInit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session_key", c)
		if _, ok := sess.Values["id"]; ok {
			return next(c)
		}

		fmt.Println(c.Request().URL)

		sess.Values["id"] = util.RandStringBytes(32)
		sess.Save(c.Request(), c.Response())
		return next(c)
	}
}
