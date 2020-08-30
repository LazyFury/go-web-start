package app

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// 自己尝试的cors配置实现
func cosr() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			res.Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
			res.Header().Set(echo.HeaderAccessControlAllowMethods, strings.Join([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}, ","))
			res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			res.Header().Set(echo.HeaderAccessControlAllowHeaders, strings.Join([]string{"token", echo.HeaderContentType}, ","))

			if req.Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}
			return next(c)
		}
	}
}
