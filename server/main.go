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

	"github.com/Masterminds/sprig"
	"github.com/fvbock/endless"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New() //echo实例

	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql); err != nil {
		panic(err)
	}
	defer model.DB.Close() //退ß出时释放链接

	e.Pre(middleware.RemoveTrailingSlash())                                        //删除url反斜杠
	e.Use(middleware.Gzip())                                                       //gzip压缩
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout})) //日志
	// e.Use(midd.LogMiddleware())
	e.Use(middleware.Recover())

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

	endless.ListenAndServe(fmt.Sprintf(":%d", config.Global.Port), e)
}

// requestInfo
func requestInfo(c echo.Context) error {
	req := c.Request()
	format := "<pre><strong>Request Information test auto build</strong>\n\n<code>Protocol: %s\nHost: %s\nRemote Address: %s\nMethod: %s\nPath: %s\n</code></pre>"
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
}
