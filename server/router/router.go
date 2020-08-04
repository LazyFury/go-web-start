package router

import (
	"net/http"
	"os"

	"EK-Server/config"
	"EK-Server/router/admin"
	"EK-Server/router/api"
	"EK-Server/util"
	"EK-Server/util/upload"

	"github.com/labstack/echo/v4"
)

// Start 入口
func Start(e *echo.Echo) {
	// baseURl 默认值 / Group的url末尾有斜杠时 get post绑定路由时不要加斜杠  无法识别 //xx 类似 传递下一级group时没有这个问题
	baseURL := config.Global.BaseURL
	if baseURL == "/" {
		baseURL = ""
	}
	g := e.Group(baseURL)

	// 项目页面
	admin.Init(g)
	api.Init(g)

	// 入口
	index := g

	// index.POST("/upload", func(c echo.Context) error {
	// 	return util.UploadCustom(c, util.AcceptsImgExt, "pic")
	// })

	index.POST("/upload", func(c echo.Context) error {
		return upload.Default(c)
	})

	index.GET("/sendMail", func(c echo.Context) error {
		email := c.QueryParam("email")
		if email == "" {
			return util.JSONErr(c, nil, "发送邮箱不可空")
		}
		err := config.Global.Mail.SendMail("消息通知", []string{email}, "madaksdjadsl<h1>测试邮件</h1>il")
		if err != nil {
			return util.JSONErr(c, err, "发送失败")
		}
		return util.JSONSuccess(c, nil, "发送成功")
	})

	index.GET("/video", func(c echo.Context) (err error) {
		video, err := os.Open("./static/hello.m3u8")
		if err != nil {
			return util.JSONErr(c, nil, "未找到文件")
		}
		defer video.Close()
		return c.Stream(http.StatusOK, "application/x-mpegURL", video)
	})

	index.GET("/reload", func(c echo.Context) error {
		//读取配置文件
		if err := config.Global.ReadConfig(); err != nil {
			return util.JSONErr(c, err, "读取配置失败")
		}
		return util.JSONSuccess(c, nil, "reload")
	})

}
