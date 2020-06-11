package home

import (
	"EK-Server/config"
	"EK-Server/router/home/post"
	"EK-Server/router/home/product"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(g *echo.Group) {
	home := g
	product.Init(home)
	post.Init(home)

	indexURL := "/"
	if config.Global.BaseURL == "" || config.Global.BaseURL == "/" {
		indexURL = "/" + strings.TrimLeft(indexURL, "/")
	} else {
		indexURL = strings.TrimLeft(indexURL, "/")
	}
	home.GET(indexURL, func(c echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = "jon"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
	})
}
