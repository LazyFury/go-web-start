package app

import (
	"net/http"

	"github.com/Treblex/go-echo-demo/server/router"
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/Treblex/go-web-template/tools"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// New 初始化
func New() *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger())
	// g.Use(cors.Default())
	g.Use(tools.Cors)

	// 静态目录
	g.Use(static.Serve("/", static.LocalFile("wwwroot", false)))

	// recover panic
	g.Use(gin.Recovery()) //保证panic时不cash

	g.Use(func(c *gin.Context) {
		defer utils.Recover(c) //panic时处理自定义错误
		c.Next()
	})

	g.HandleMethodNotAllowed = true

	g.NoMethod(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		utils.Error(utils.NoMethod)
	})

	g.NoRoute(func(c *gin.Context) {
		utils.Error(utils.NoRoute)
	})

	g.RemoveExtraSlash = true
	g.RedirectTrailingSlash = true

	// 注册路由
	router.Start(g)

	return g
}
