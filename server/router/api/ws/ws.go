package ws

import (
	"github.com/Treblex/go-echo-demo/server/router/api/ws/chat"

	"github.com/labstack/echo/v4"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "/ws"
	app := g.Group(baseURL)
	app.GET("", chat.WsServer)
}
