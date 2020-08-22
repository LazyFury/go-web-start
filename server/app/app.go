package app

import (
	"EK-Server/router"
	"EK-Server/util"
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New 初始化
func New() *echo.Echo {
	e := echo.New() //echo实例

	e.Pre(middleware.RemoveTrailingSlash())                                        //删除url反斜杠
	e.Use(middleware.Gzip())                                                       //gzip压缩
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout})) //日志
	// e.Use(midd.LogMiddleware())
	e.Use(middleware.Recover())        //错误处理
	e.Use(middleware.NonWWWRedirect()) //跳转到没有www到顶级域名
	e.Use(middleware.Secure())         //安全

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(util.RandStringBytes(32))))) //session
	//跨域
	e.Use(middleware.CORS())
	// fix：csrf未验证crash之后造成之后的cosr配置未生效，已修改顺序
	e.Use(middleware.CSRF())

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
