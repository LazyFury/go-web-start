package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-start/server/middleware"
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-start/server/router"
	"github.com/lazyfury/go-web-start/server/utils"
	gwt "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/response"
	_template "github.com/lazyfury/go-web-template/tools/template"
)

func main() {
	g := gwt.New()
	g.PreUse(func(c *gin.Context) {
		url := c.Request.URL.Path
		fmt.Printf("%s [%s]%s\n", time.Now().Format("2006-01-02 15:04:05"), c.Request.Method, url)
	})
	g.PreUse(func(c *gin.Context) {
		gwt.Cors(c, &gwt.CorsConfig{
			AllowOrigins:     []string{"*"},
			AllowCredentials: true,
			AllowHeaders:     gwt.DefaultAllowHeaders,
			AllowMethods:     gwt.DefaultAllowMethods,
		})
	})
	g.PreUse(middleware.AuthOrNot)
	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}

	// 静态目录
	g.Use(static.Serve("/", static.LocalFile("wwwroot", false)))

	// 注册html模板
	html := template.Must(_template.ParseGlob(template.New("main"), "templates", "*.html"))
	g.SetHTMLTemplate(html)

	// 注册路由
	router.Start(&g.RouterGroup)

	// 扩展自定义错误码
	response.PushErrCodeTextMap(utils.ErrCodeText)
	response.RecoverRender = func(c *gin.Context, code int, result *response.Result) {
		c.HTML(http.StatusOK, "err/error.html", result)
	}
	// 启动
	err := g.Run(fmt.Sprintf(":%d", config.Global.Port))
	if err != nil {
		panic(err)
	}
}
