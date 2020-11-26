package app

import (
	"net/http"

	"github.com/Treblex/go-echo-demo/server/router"
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// New 初始化
func New() *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger())

	g.Use(cors.Default())

	// 静态目录
	g.Use(static.Serve("/", static.LocalFile("wwwroot", false)))

	// recover panic
	g.Use(gin.Recovery())

	g.Use(func(c *gin.Context) {
		defer utils.Recover(c)
		c.Next()
	})

	g.HandleMethodNotAllowed = true

	g.NoMethod(func(c *gin.Context) {
		panic(utils.JSON(http.StatusMethodNotAllowed, "", nil))
	})

	g.NoRoute(func(c *gin.Context) {
		panic(utils.JSON(http.StatusNotFound, "", nil))
	})

	g.RemoveExtraSlash = true
	g.RedirectTrailingSlash = true

	// 注册路由
	router.Start(g)

	return g
}
