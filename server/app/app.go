package app

import (
	"EK-Server/router"
	"EK-Server/util"
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New 初始化
func New() *echo.Echo {
	e := echo.New() //echo实例

	e.Pre(middleware.RemoveTrailingSlash())                                        //删除url反斜杠
	e.Use(middleware.Gzip())                                                       //gzip压缩
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout})) //日志
	// e.Use(midd.LogMiddleware())
	// e.Use(middleware.Recover())

	//跨域
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*", "https://labstack.net"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.CONNECT},
		AllowCredentials: true,
		AllowHeaders:     []string{"token", "Content-Type"},
	}))

	// 模版
	renderer := &util.TemplateRenderer{
		Templates: template.Must(util.ParseGlob(template.New("base").Funcs(util.TemplateFuns).Funcs(sprig.FuncMap()), "template", "*.html")),
	}
	// 绑定渲染模版方法
	e.Renderer = renderer

	// 错误处理
	e.HTTPErrorHandler = httpErrorHandler

	// 静态目录
	e.Static("/static", "static")
	e.Static("/h5", "h5")

	// 请求信息
	e.GET("requestInfo", requestInfo)

	// 注册路由
	router.Start(e)

	// 启动服务
	// e.Logger.Error(e.Start(fmt.Sprintf(":%d", config.Global.Port)))
	return e
}
