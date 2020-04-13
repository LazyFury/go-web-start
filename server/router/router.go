package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"EK-Server/config"
	"EK-Server/router/admin"
	"EK-Server/router/api"
	"EK-Server/router/home"
	"EK-Server/router/product"
	"EK-Server/router/wechat"
	"EK-Server/router/ws"
	"EK-Server/util"

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
	// 入口
	index := g

	index.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world！")
	})
	index.POST("/upload", upload)
	index.GET("/video", func(c echo.Context) (err error) {
		video, err := os.Open("./static/hello.m3u8")
		if err != nil {
			return util.JSONErr(c, nil, "未找到文件")
		}
		defer video.Close()
		return c.Stream(http.StatusOK, "application/x-mpegURL", video)
	})

	index.GET("/👌", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "")
	})

}

func upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return util.JSONErr(c, err, "上传错误")
	}

	pathExt := path.Ext(file.Filename)
	acceptsExt := []string{"jpg", "png"}
	index := -1
	for i, item := range acceptsExt {
		if pathExt == "."+item {
			index = i
		}
	}
	if index == -1 {
		return util.JSONErr(c, nil, "文件不允许")
	}

	src, err := file.Open()
	if err != nil {
		return util.JSONErr(c, err, "打开文件失败")
	}
	defer src.Close()

	dir, err := getDir("./static/upload/", time.Now().Format("2006_01_02"))
	if err != nil {
		return util.JSONErr(c, err, "创建文件夹失败")
	}

	randName := util.RandStringBytes(32) + pathExt
	// Destination
	fileName := filepath.Join(dir, randName)
	dst, err := os.Create(fileName)
	if err != nil {
		return util.JSONErr(c, err, "创建文件失败")
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return util.JSONErr(c, err, "拷贝文件至目标失败")
	}
	return util.JSON(c, fmt.Sprintf("//%s/%s", c.Request().Host, fileName), "上传成功", 200)
}

func getDir(path string, foderName string) (dir string, err error) {
	foder := filepath.Join(path, foderName)
	if _, err = os.Stat(foder); os.IsNotExist(err) {
		err = os.MkdirAll(foder, os.ModePerm)
		if err != nil {
			return
		}
	}

	dir = foder
	return
}
