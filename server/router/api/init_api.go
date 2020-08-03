package api

import (
	"EK-Server/router/api/wechat"
	"EK-Server/router/api/ws"
	"EK-Server/util"

	"github.com/labstack/echo"
)

// Init  api Version 1.0 初始化
func Init(g *echo.Group) {
	apiV1 := g.Group("/api/v1")
	//常用到资源整理到这里统一到api暴露处理，暂定根据methods get和other来处理权限
	//get 常用于获取列表 详情，不涉及更新和修改数据到方法
	apiV1.GET("", resources)
	// base
	login(apiV1)
	//文章
	post(apiV1)
	postCate(apiV1)
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
}

func resources(c echo.Context) error {
	type resource struct {
		Name string `json:"name"`
		Like string `json:"link"`
	}
	res := []resource{
		{"文章", "/api/v1/posts"},
		{"商品", "/api/v1/goods"},
		{"订单", "/api/v1/orders"},
		{"广告", "/api/v1/ads"},
	}

	return util.JSONSuccess(c, map[string]interface{}{
		"resources": res,
	}, "")
}
