package api

import (
	"net/http"

	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/Treblex/go-echo-demo/server/controller"
	"github.com/Treblex/go-echo-demo/server/router/api/wechat"
	"github.com/Treblex/go-echo-demo/server/router/api/ws"
	"github.com/Treblex/go-echo-demo/server/tools/upload"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
)

// var uploader = upload.NewEchoUploader()
var aliUploader = upload.NewAliOssUploader(config.Global.AliOss)

// Init  api Version 1.0 初始化
func Init(g *gin.RouterGroup) {

	apiV1 := g.Group("/v1")

	//常用到资源整理到这里统一到api暴露处理，暂定根据methods get和other来处理权限
	//get 常用于获取列表 详情，不涉及更新和修改数据到方法
	apiV1.GET("/", resources)
	apiV1.POST("/upload", func(c *gin.Context) {
		url, err := aliUploader.Default(c.Request)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, utils.JSONSuccess("", url))
	})
	apiV1.POST("/upload-img", func(c *gin.Context) {
		url, err := aliUploader.OnlyAcceptsExt(c.Request, upload.AcceptsImgExt, "image")
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, utils.JSONSuccess("上传成功", url))
	})
	apiV1.POST("/upload-head-pic", func(c *gin.Context) {
		url, err := aliUploader.Custom(c.Request, upload.AcceptsImgExt, "head_pic")
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, utils.JSONSuccess("上传成功", url))
	})
	// base
	login(apiV1)
	//文章
	controller.NewArticleController().Install(apiV1, "/articles")
	controller.NewArticleRecController().Install(apiV1, "/article-recs")
	controller.NewArticleCategoryController().Install(apiV1, "/article-cates")
	controller.NewArticleTagController().Install(apiV1, "/article-tags")
	//商品
	product(apiV1)
	productCate(apiV1) //商品分类
	// 订单
	order(apiV1)
	// banner 广告位
	controller.NewAdController().Install(apiV1, "/ads")
	controller.NewAdGroupController().Install(apiV1, "/ad-groups")
	controller.NewAdEventController().Install(apiV1, "/ad-events")

	// 用户
	user(apiV1)
	// 意见反馈
	controller.NewFeedbackController().Install(apiV1, "/feedbacks")
	// 用户消息
	controller.NewMessageController().Install(apiV1, "/messages")
	controller.NewMessageTemplateController().Install(apiV1, "/message-templates")

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

func resources(c *gin.Context) {
	res := []resource{
		{"文章", "/api/v1/posts", ""},
		{"商品", "/api/v1/goods", ""},
		{"订单", "/api/v1/orders", ""},
		{"广告", "/api/v1/ads", ""},
	}

	c.JSON(http.StatusOK, utils.JSONSuccess("", map[string]interface{}{
		"resources": res,
	}))
}
