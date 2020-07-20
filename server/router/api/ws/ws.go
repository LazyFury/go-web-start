package ws

import (
	"EK-Server/router/ws/game"

	"github.com/labstack/echo"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "ws"
	app := g.Group(baseURL)
	app.GET("", game.WsServer)
	go game.Push()
}
