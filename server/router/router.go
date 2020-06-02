package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"EK-Server/config"
	"EK-Server/router/admin"
	"EK-Server/router/api"
	"EK-Server/router/home"
	"EK-Server/router/product"
	"EK-Server/router/tg"
	"EK-Server/router/wechat"
	"EK-Server/router/ws"
	"EK-Server/util"
	"EK-Server/util/upload"

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
	home.Init(g)
	product.Init(g)
	tg.Init(g)
	// 入口
	index := g

	index.GET("/", func(c echo.Context) error {

		fmt.Printf("hello world!")
		return c.String(http.StatusOK, "hello world！docker got it")
	})
	index.GET("/svg", func(c echo.Context) error {
		color := c.QueryParam("color")
		svgStr := `<svg xmlns="http://www.w3.org/2000/svg"  width="500" height="200">
				<path id="形状 1" fill='%s' d="
				M 0 10 
				l 500 10 
				v 100 
				h -500
				Z" />
			</svg>`
		log.Println("color:"+color == "")
		if color == "" {
			color = "#000"
		}
		return c.Blob(http.StatusOK, "image/svg+xml", []byte(fmt.Sprintf(svgStr, color)))
	})
	// index.POST("/upload", func(c echo.Context) error {
	// 	return util.UploadCustom(c, util.AcceptsImgExt, "pic")
	// })
	index.POST("/upload", func(c echo.Context) error {
		return upload.Default(c)
	})

	index.GET("/video", func(c echo.Context) (err error) {
		video, err := os.Open("./static/hello.m3u8")
		if err != nil {
			return util.JSONErr(c, nil, "未找到文件")
		}
		defer video.Close()
		return c.Stream(http.StatusOK, "application/x-mpegURL", video)
	})

	index.GET("/👌", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "👌")
	})

}
