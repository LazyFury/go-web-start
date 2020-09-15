package app

import (
	"html/template"
	"os"

	"github.com/Treblex/go-echo-demo/server/router"
	"github.com/Treblex/go-echo-demo/server/util"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New 初始化
func New() *echo.Echo {
	e := echo.New() //echo实例
	//跨域
	e.Pre(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))
	e.Pre(middleware.NonWWWRedirect()) //跳转到没有www到顶级域名
	e.Pre(middleware.Secure())         //安全
	e.Pre(slash)                       //反斜杠处理

	// response
	e.Use(middleware.Gzip())                                                       //gzip压缩
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout})) //日志
	// e.Use(midd.LogMiddleware())
	e.Use(middleware.Recover()) //错误处理

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(util.RandStringBytes(32))))) //session
	e.Use(sessionInit)

	// fix：csrf未验证crash之后造成之后的cosr配置未生效，已修改顺序
	// fix：跨域的情况下不建议使用csrf
	// tips:csrf使用场景，服务端渲染模版的时候，将csrf key自动渲染到页面表单中随数据提交，再跨域到情况下没有比较安全的方案获取到csrf key
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	CookieDomain: "127.0.0.1:8080",
	// 	// CookieHTTPOnly: true,
	// 	// CookieSecure:   true,
	// }))

	// 绑定渲染模版方法
	e.Renderer = &util.TemplateRenderer{
		Templates: template.Must(util.ParseGlob(template.New("base").Funcs(util.TemplateFuns).Funcs(sprig.FuncMap()), "template", "*.html")),
	}

	// 错误处理
	e.HTTPErrorHandler = httpErrorHandler

	// 请求信息
	e.GET("requestInfo", requestInfo)

	// 注册路由
	router.Start(e)

	// 启动服务
	// e.Logger.Error(e.Start(fmt.Sprintf(":%d", config.Global.Port)))
	return e
}
