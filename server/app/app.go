package app

import (
	"github.com/Treblex/go-echo-demo/server/router"
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// New 初始化
func New() *gin.Engine {
	g := gin.Default()

	// 静态目录
	g.Use(static.Serve("/static", static.LocalFile("static", false)))

	g.Use(utils.Recover)
	// 注册路由
	router.Start(g)

	return g
}
