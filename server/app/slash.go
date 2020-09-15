package app

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// 处理反斜杠的中间件
func slash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		url := req.URL
		path := url.Path
		uri := path

		// 是否需要重定向url
		redirect := false

		// 重复的斜杠，前端拼接url可能造成的异常
		if strings.Contains(path, "//") {
			uri = strings.ReplaceAll(path, "//", "/")
			redirect = true
		}

		// 删除末尾的斜杠，限制api前缀
		// 静态文件夹会自动添加后缀，与这里造成重复，造成无限♾️301重定向
		if strings.HasPrefix(uri, "/api") && strings.HasSuffix(uri, "/") {
			uri = strings.TrimRight(uri, "/")
		}

		// redirect
		if redirect && req.Method == http.MethodGet {
			// 拼接query
			qs := c.QueryString()
			if qs != "" {
				uri += "?"
				uri += qs
			}
			return c.Redirect(http.StatusMovedPermanently, uri)
		}
		// Forward
		url.Path = uri
		req.URL = url
		req.RequestURI = uri

		return next(c)
	}
}
