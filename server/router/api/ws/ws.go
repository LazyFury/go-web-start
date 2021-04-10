package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/router/api/ws/chat"
)

// Init Init
func Init(g *gin.RouterGroup) {
	baseURL := "/ws"
	app := g.Group(baseURL)
	app.GET("", chat.WsServer)
}
