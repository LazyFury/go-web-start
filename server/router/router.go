package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-start/server/router/api"
	"github.com/lazyfury/go-web-start/server/utils"
)

// Start 入口
func Start(e *gin.Engine) {
	// baseURl 默认值 / Group的url末尾有斜杠时 get post绑定路由时不要加斜杠  无法识别 //xx 类似 传递下一级group时没有这个问题
	baseURL := config.Global.BaseURL
	if baseURL == "/" {
		baseURL = ""
	}
	g := e.Group(baseURL)

	// 项目页面
	api.Init(g)

	// 入口
	index := g

	index.GET("/sendMail", func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			utils.Error("发送邮箱不可空")
		}
		err := config.Global.Mail.SendMail("消息通知", []string{email}, "madaksdjadsl<h1>测试邮件</h1>il")
		if err != nil {
			utils.Error(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "发送成功"})
	})

}
