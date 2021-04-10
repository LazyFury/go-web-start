package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lazyfury/go-web-start/server/router/api/v1"
	v2 "github.com/lazyfury/go-web-start/server/router/api/v2"
)

// Init Init
func Init(g *gin.RouterGroup) {
	api := g.Group("/api")
	v1.Init(api)
	v2.Init(api)
}
