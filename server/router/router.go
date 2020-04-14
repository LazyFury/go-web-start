package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
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
		return util.JSONErr(c, err, "上传错误") //未获取到文件流
	}
	pathExt := path.Ext(file.Filename)
	acceptsImgExt := []interface{}{"jpg", "png", "jpeg", "webp"} //图片类型
	acceptsVideoExt := []interface{}{"mov", "mp4", "avi"}        //视频类型
	acceptsPdfExt := []interface{}{"pdf", "emu"}                 //其他文件类型
	folder := ""
	// 如果符合类型，设定目录
	if inArray(acceptsImgExt, strings.Trim(pathExt, ".")) > -1 {
		folder = "image"
	}
	if inArray(acceptsVideoExt, strings.Trim(pathExt, ".")) > -1 {
		folder = "video"
	}
	if inArray(acceptsPdfExt, strings.Trim(pathExt, ".")) > -1 {
		folder = "pdf"
	}
	// 如果不符合任何一种类型
	if folder == "" {
		return util.JSONErr(c, nil, "文件不合法")
	}

	// 打开文件流
	src, err := file.Open()
	if err != nil {
		return util.JSONErr(c, err, "打开文件失败")
	}
	defer src.Close() //函数结束时自动关闭文件

	//创建文件夹
	dir, err := getDir("./static/upload/"+folder+"/", time.Now().Format("2006_01_02"))
	if err != nil {
		return util.JSONErr(c, err, "创建文件夹失败")
	}

	// 随机文件名 + 文件后缀
	randName := util.RandStringBytes(32) + pathExt
	// Destination
	fileName := filepath.Join(dir, randName)

	// 创建空文件
	dst, err := os.Create(fileName)
	if err != nil {
		return util.JSONErr(c, err, "创建文件失败")
	}
	defer dst.Close()
	// Copy文件流到新建到文件
	if _, err = io.Copy(dst, src); err != nil {
		return util.JSONErr(c, err, "拷贝文件至目标失败")
	}
	// 拼接文件地址，不带协议头，方便处理http 到https升级 ， 其实也没找到协议头在哪儿，req对象里没有返回到空字符串
	return util.JSON(c, fmt.Sprintf("//%s/%s", c.Request().Host, fileName), "上传成功", 200)
}

// 创建文件夹
func getDir(path string, foderName string) (dir string, err error) {
	folder := filepath.Join(path, foderName)
	if _, err = os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return
		}
	}
	dir = folder
	return
}

// 在数组中
func inArray(arr []interface{}, item interface{}) (index int) {
	index = -1
	for i, x := range arr {
		if item == x {
			index = i
		}
	}
	return index
}
