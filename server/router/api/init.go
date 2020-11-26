package api

import (
	v1 "github.com/Treblex/go-echo-demo/server/router/api/v1"
	v2 "github.com/Treblex/go-echo-demo/server/router/api/v2"
	"github.com/gin-gonic/gin"
)

// Init Init
func Init(g *gin.RouterGroup) {
	api := g.Group("/api")
	v1.Init(api)
	v2.Init(api)
}
