package home

import (
	"EK-Server/router/home/post"
	"EK-Server/router/home/product"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := ""
	home := g.Group(baseURL)
	product.Init(home)
	post.Init(home)
}
