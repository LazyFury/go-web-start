package api

import (
	v1 "github.com/Treblex/go-echo-demo/server/router/api/v1"
	v2 "github.com/Treblex/go-echo-demo/server/router/api/v2"
	"github.com/labstack/echo/v4"
)

// Init Init
func Init(g *echo.Group) {
	api := g.Group("/api")
	v1.Init(api)
	v2.Init(api)
}
