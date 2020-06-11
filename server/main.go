package main

import (
	"EK-Server/config"
	"EK-Server/model"
	"EK-Server/router"
	"EK-Server/util"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	time.LoadLocation("local")
	e := echo.New() //echo实例                                             //日志
	fmt.Println("hello world!")
	model.DB = model.InitDB(config.Global.Mysql) //初始化数据链接
	defer model.DB.Close()                       //退出时释放链接

	e.Pre(middleware.RemoveTrailingSlash())                                        //删除url反斜杠
	e.Use(middleware.Gzip())                                                       //gzip压缩
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout})) //日志
	// e.Use(util.LogMiddleware())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*", "https://labstack.net"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.CONNECT},
		AllowCredentials: true,
		AllowHeaders:     []string{"token"},
	})) //跨域

	// 模版
	renderer := &util.TemplateRenderer{
		Templates: template.Must(util.ParseGlob(template.New("base").Funcs(template.FuncMap{
			"msg": func() string { return "hello this is a msg" },
		}).Funcs(sprig.FuncMap()), "template", "*.html")),
	}
	e.Renderer = renderer

	// 错误处理
	e.HTTPErrorHandler = httpErrorHandler
	e.GET("/hello", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "hello world!")
	})
	// e.Use("requestInfo")
	// 静态目录
	e.Static("/static", "static")
	e.Static("/h5", "h5")
	// 请求信息
	e.GET("requestInfo", requestInfo)
	// 注册路由
	router.Start(e)

	// 启动服务
	e.Logger.Error(e.Start(fmt.Sprintf(":%d", config.Global.Port)))

}

// requestInfo
func requestInfo(c echo.Context) error {
	req := c.Request()
	format := "<pre><strong>Request Information</strong>\n\n<code>Protocol: %s\nHost: %s\nRemote Address: %s\nMethod: %s\nPath: %s\n</code></pre>"
	fmt.Println(strings.Split(req.Header.Get("Accept"), ",")[0])
	fmt.Printf("%+v", req.Header)
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}

// 错误处理
func httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	msg = http.StatusText(code)
	// 如果是浏览器
	req := c.Request()
	reqAccept := strings.Split(req.Header.Get("Accept"), ",")[0]
	if reqAccept == "text/html" {
		c.Logger().Error(c.Render(code, "error.html", map[string]interface{}{
			"msg":  msg,
			"code": code,
		}))
		return
	}
	// 如果是ajax
	c.Logger().Error(util.JSONBase(c, nil, msg, code, code))
	// c.Logger().Error(util.JSONBase)
}
