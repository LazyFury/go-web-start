package api

import (
	"net/http"

	"github.com/lazyfury/go-web-start/server/controller"
	"github.com/lazyfury/go-web-start/server/middleware"
	"github.com/lazyfury/go-web-start/server/router/api/wechat"
	"github.com/lazyfury/go-web-start/server/router/api/ws"
	"github.com/lazyfury/go-web-start/server/utils"
	"github.com/lazyfury/go-web-template/response"

	"github.com/gin-gonic/gin"
)

// Init  api Version 1.0 初始化
func Init(g *gin.RouterGroup) {

	apiV1 := g.Group("/v1", middleware.AuthOrNot)

	//常用到资源整理到这里统一到api暴露处理，暂定根据methods get和other来处理权限
	//get 常用于获取列表 详情，不涉及更新和修改数据到方法
	apiV1.GET("/", resources)

	// base
	// login(apiV1)
	controller.NewUserController().Install(apiV1, "/users")
	//文章
	controller.NewArticleController().Install(apiV1, "/articles")
	controller.NewArticleRecController().Install(apiV1, "/article-recs")
	controller.NewArticleCategoryController().Install(apiV1, "/article-cates")
	controller.NewArticleTagController().Install(apiV1, "/article-tags")

	// banner 广告位
	controller.NewAdController().Install(apiV1, "/ads")
	controller.NewAdGroupController().Install(apiV1, "/ad-groups")
	controller.NewAdEventController().Install(apiV1, "/ad-events")

	// 意见反馈
	controller.NewFeedbackController().Install(apiV1, "/feedbacks")

	// 用户消息
	controller.NewMessageController().Install(apiV1, "/messages")
	controller.NewMessageTemplateController().Install(apiV1, "/message-templates")

	wechat.Init(apiV1)
	ws.Init(apiV1)

	InitUpload(apiV1)

	wehcatMini(apiV1)

	configRouter(apiV1)

	// xml 播客解析测试
	podcastRouter(apiV1)

}

type resource struct {
	Name string `json:"name"`
	Like string `json:"link"`
	Doc  string `json:"doc"`
}

func resources(c *gin.Context) {
	res := []resource{
		{"文章", "/api/v1/articles", ""},
		{"消息", "/api/v1/messages", ""},
		{"商品", "/api/v1/goods", ""},
		{"订单", "/api/v1/orders", ""},
		{"广告", "/api/v1/ads", ""},
	}

	c.JSON(http.StatusOK, response.JSON(utils.CustomErrCode, "", map[string]interface{}{
		"resources": res,
	}))
}
