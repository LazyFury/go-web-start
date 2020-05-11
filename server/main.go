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

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	filename := "./log/" + time.Now().Format("2006_01_02") + ".log"

	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// os.Stdout = logfile
	os.Stderr = logfile
	time.LoadLocation("local")
	e := echo.New()                                                                //echo实例                                             //日志
	model.DB = model.InitDB(config.Global.Mysql)                                   //初始化数据链接 不知道为什么 main.go 大写暴露的变量不能全局调用
	defer model.DB.Close()                                                         //退出时释放链接
	e.Pre(middleware.RemoveTrailingSlash())                                        //删除url反斜杠
	e.Use(middleware.Gzip())                                                       //gzip压缩
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout})) //日志
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "https://labstack.net"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.CONNECT},
	})) //跨域
	// e.Use(util.JWT())
	
	// 模版
	renderer := &util.TemplateRenderer{
		Templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer =  renderer
	
	// 错误处理
	e.HTTPErrorHandler = httpErrorHandler
	e.GET("/hello", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "hello world!")
	})
	// 静态目录
	e.Static("/static", "static")
	e.Static("/h5", "h5")
	// 请求信息
	e.GET("requestInfo", requestInfo)
	// 注册路由
	router.Start(e)
	// router.Start(e)
	
	// 启动服务
	e.Logger.Error(e.Start(":8080"))
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
