package router

import (
	"net/http"
	"os"
	"suke-go-test/router/admin"
	"suke-go-test/router/wechat"
	"suke-go-test/util"

	"github.com/labstack/echo"
)

// Start 入口
func Start(e *echo.Group) {
	// 入口
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world！")
	}, util.UserJWT)
	// 项目首页
	admin.Init(e)
	wechat.Init(e)

	e.GET("/video", func(c echo.Context) (err error) {
		video, err := os.Open("./static/hello.m3u8")
		if err != nil {
			return util.JSONErr(c, err, "未找到文件")
		}
		defer video.Close()
		return c.Stream(http.StatusOK, "application/x-mpegURL", video)
	})

}
