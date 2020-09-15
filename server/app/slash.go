package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func slash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		url := req.URL
		path := url.Path
		fmt.Printf("\n%v\n", c)
		uri := path

		redirect := false

		if strings.Contains(path, "//") {
			uri = strings.ReplaceAll(path, "//", "/")
			redirect = true
		}

		if strings.HasPrefix(uri, "/api") && strings.HasSuffix(uri, "/") {
			uri = strings.TrimRight(uri, "/")
		}

		// redirect
		if redirect {
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
		fmt.Printf("\n%v\n", c)

		return next(c)
	}
}
