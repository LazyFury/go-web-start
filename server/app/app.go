package app

import (
	"html/template"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/middleware"
	"github.com/lazyfury/go-web-start/server/router"
	"github.com/lazyfury/go-web-start/server/utils"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools"
)

// New 初始化
func New() *gin.Engine {
	g := gin.New()

	// 跨域配置
	g.Use(tools.DefaultCors)

	// 自定义recover
	g.Use(response.Recover)

	// HandleMethodNotAllowed
	g.HandleMethodNotAllowed = true
	g.NoMethod(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		response.Error(response.NoMethod)
	})

	// handle 路由不存在
	g.NoRoute(func(c *gin.Context) {
		if c.Request.URL.Path != "/favicon.ico" {
			response.Error(response.NoRoute)
		}
	})

	// 移除多余斜杠 /api//v1/doSomething/ => /api/v1/doSomething
	g.RemoveExtraSlash = true
	// 重定向请求移除斜杠请求
	g.RedirectTrailingSlash = true

	// 静态目录
	g.Use(static.Serve("/", static.LocalFile("wwwroot", false)))

	// 注册路由
	g.Use(middleware.AuthOrNot)
	router.Start(g)

	// 注册html模板
	html := template.Must(tools.ParseGlob(template.New("main"), "templates", "*.html"))
	g.SetHTMLTemplate(html)

	// 错误码配置
	response.PushErrCodeTextMap(utils.ErrCodeText)
	response.RecoverErrHtml = true
	response.RecoverErrTemplateName = "err/error.html"

	return g
}
