package api

import (
	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/Treblex/go-echo-demo/server/router/api/wechat"
	"github.com/Treblex/go-echo-demo/server/router/api/ws"
	"github.com/Treblex/go-echo-demo/server/util"
	"github.com/Treblex/go-echo-demo/server/util/upload"

	"github.com/labstack/echo/v4"
)

// var uploader = upload.NewEchoUploader()
var aliUploader = upload.NewAliOssUploader(config.Global.AliOss)

// Init  api Version 1.0 初始化
func Init(g *echo.Group) {

	apiV1 := g.Group("/v1")
	//常用到资源整理到这里统一到api暴露处理，暂定根据methods get和other来处理权限
	//get 常用于获取列表 详情，不涉及更新和修改数据到方法
	apiV1.GET("", resources)
	apiV1.POST("/upload", func(c echo.Context) error {
		url, err := aliUploader.Default(c.Request())
		if err != nil {
			return util.JSONErr(c, nil, err.Error())
		}
		return util.JSONSuccess(c, url, "上传成功")
	})
	apiV1.POST("/upload-head-pic", func(c echo.Context) error {
		url, err := aliUploader.Custom(c.Request(), upload.AcceptsImgExt, "head_pic")
		if err != nil {
			return util.JSONErr(c, nil, err.Error())
		}
		return util.JSONSuccess(c, url, "上传成功")
	})
	// base
	login(apiV1)
	//文章
	post(apiV1)
	postCate(apiV1)
	postRec(apiV1)
	//商品
	product(apiV1)
	productCate(apiV1) //商品分类
	// 订单
	order(apiV1)
	// banner 广告位
	ad(apiV1)
	adEvent(apiV1)
	adGroup(apiV1)

	// 用户
	user(apiV1)
	// 意见反馈
	feedback(apiV1)

	// 用户消息
	messages(apiV1)

	wechat.Init(apiV1)

	ws.Init(apiV1)

	wehcatMini(apiV1)

	configRouter(apiV1)

	// xml 播客解析测试
	podcastRouter(apiV1)

	// 预约
	appointment(apiV1)

}

type resource struct {
	Name string `json:"name"`
	Like string `json:"link"`
	Doc  string `json:"doc"`
}

func resources(c echo.Context) error {

	res := []resource{
		{"文章", "/api/v1/posts", ""},
		{"商品", "/api/v1/goods", ""},
		{"订单", "/api/v1/orders", ""},
		{"广告", "/api/v1/ads", ""},
	}

	return util.JSONSuccess(c, map[string]interface{}{
		"resources": res,
	}, "")
}
