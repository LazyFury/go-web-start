package main

import (
	"fmt"
	"suke-go-test/router"
	"suke-go-test/util"
	"html/template"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New() //echo实例

	util.DB = util.InitDB() //初始化数据链接 不知道为什么 main 大写暴露的变量不能全局调用
	defer util.DB.Close()   //退出时释放链接

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Gzip())   //gzip压缩
	e.Use(middleware.Logger()) //日志
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "https://labstack.net"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	})) //跨域

	// e.Use(util.JWT())
	// 模版
	renderer := &util.TemplateRenderer{
		Templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = renderer
	// 错误处理
	e.HTTPErrorHandler = httpErrorHandler

	// 静态目录
	e.Static("/static", "static")
	e.Static("/h5", "h5")
	// 请求信息
	// e.GET("requestInfo", requestInfo)

	// 注册路由
	g := e.Group("")
	router.Start(g)

	// 启动服务
	e.Logger.Fatal(e.Start(":8080"))

}

// requestInfo
func requestInfo(c echo.Context) error {
	req := c.Request()
	format := "<pre><strong>Request Information</strong>\n\n<code>Protocol: %s\nHost: %s\nRemote Address: %s\nMethod: %s\nPath: %s\n</code></pre>"
	fmt.Println(strings.Split(req.Header.Get("Accept"), ",")[0])
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
	c.Logger().Error(util.JSONBase(c, "", msg, code, code))
}
