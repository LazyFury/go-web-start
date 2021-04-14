package main

import (
	"fmt"
	"html/template"

	"github.com/gin-contrib/static"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-start/server/middleware"
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-start/server/router"
	"github.com/lazyfury/go-web-start/server/utils"
	gowebtemplate "github.com/lazyfury/go-web-template"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools"
)

func main() {

	g := gowebtemplate.New()
	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}

	// 静态目录
	g.Use(static.Serve("/", static.LocalFile("wwwroot", false)))

	// 注册路由
	g.Use(middleware.AuthOrNot)
	router.Start(g)

	// 注册html模板
	html := template.Must(tools.ParseGlob(template.New("main"), "templates", "*.html"))
	g.SetHTMLTemplate(html)

	// 扩展自定义错误码
	response.PushErrCodeTextMap(utils.ErrCodeText)

	// 启动
	err := g.Run(fmt.Sprintf(":%d", config.Global.Port))
	if err != nil {
		panic(err)
	}
}
