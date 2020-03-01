package router

import (
	"net/http"
	"os"
	"suke-go-test/config"
	"suke-go-test/router/admin"
	"suke-go-test/router/api"
	"suke-go-test/router/wechat"
	"suke-go-test/router/ws"
	"suke-go-test/util"

	"github.com/labstack/echo"
)

// Start 入口
func Start(e *echo.Echo) {
	// baseURl 默认值 / Group的url末尾有斜杠时 get post绑定路由时不要加斜杠  无法识别 //xx 类似 传递下一季group时没有这个问题
	baseURL := config.Global.BaseURL
	if baseURL == "/" {
		baseURL = ""
	}
	g := e.Group(baseURL)

	// 项目页面
	admin.Init(g)
	wechat.Init(g)
	api.Init(g)
	ws.Init(g)

	// 入口
	index := g

	index.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world！")
	}, util.UserJWT)

	index.GET("/video", func(c echo.Context) (err error) {
		video, err := os.Open("./static/hello.m3u8")
		if err != nil {
			return util.JSONErr(c, err, "未找到文件")
		}
		defer video.Close()
		return c.Stream(http.StatusOK, "application/x-mpegURL", video)
	})

	index.GET("/👌", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "")
	})

}
