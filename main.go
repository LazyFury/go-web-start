package main

import (
	"fmt"
	"html/template"
	"main/router"
	"main/util"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New() //echo实例

	// fmt.Println(time.Now().In(cstZone).Format("01-02-2006 15:04:05"))

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Gzip())   //gzip压缩
	e.Use(middleware.Logger()) //日志
	e.Use(middleware.CORS())   //跨域

	// 模版
	// filenames, _ := filepath.Glob("template/error/*.html")
	// e.Logger.Fatal(filenames)
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
	e.GET("requestInfo", requestInfo)

	// 注册路由
	router.Start(e)

	// 启动服务
	e.Logger.Fatal(e.Start(":8080"))

	defer util.DB.Close()
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
		c.Logger().Error(c.Render(http.StatusOK, "error.html", map[string]interface{}{
			"msg":  msg,
			"code": code,
		}))
		return
	}
	// 如果是ajax
	c.Logger().Error(util.JSONBase(c, "", msg, code, code))
}
