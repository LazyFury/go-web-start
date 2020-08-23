package app

import (
	"EK-Server/util"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func sessionInit() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			if _, ok := sess.Values["session_key"]; ok {
				return next(c)
			}
			sess.Values["session_key"] = util.RandStringBytes(32)
			sess.Save(c.Request(), c.Response())
			return next(c)
		}
	}
}
