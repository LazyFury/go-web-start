package app

import (
	"html/template"
	"net/http"

	"github.com/Treblex/go-web-start/server/utils"
	"github.com/Treblex/go-web-template/tools"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// New 初始化
func New() *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger())
	// g.Use(cors.Default())
	g.Use(tools.DefaultCors)

	// 静态目录
	g.Use(static.Serve("/", static.LocalFile("wwwroot", false)))

	html := template.Must(tools.ParseGlob(template.New("main"), "templates", "*.html"))
	g.SetHTMLTemplate(html)
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
		if c.Request.URL.Path != "/favicon.ico" {
			utils.Error(utils.NoRoute)
		}
	})

	// 移除多余斜杠 /api//v1/doSomething/ => /api/v1/doSomething
	g.RemoveExtraSlash = true
	// 重定向请求
	g.RedirectTrailingSlash = true

	return g
}
