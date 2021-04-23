package router

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-start/server/router/api"
	"github.com/lazyfury/go-web-template/response"
)

// Start 入口
func Start(e *gin.Engine) {
	baseURL := config.Global.BaseURL
	g := e.Group(baseURL)

	// 项目页面
	api.Init(g)

	// 入口
	index := g

	index.GET("/sendMail", func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			response.Error("发送邮箱不可空")
		}

		re, err := regexp.Compile(`^.+?@[a-zA-Z0-9-_+=]{1,}\.[a-zA-Z]{2,}$`)
		if err != nil {
			response.Error(err)
		}
		if !re.MatchString(email) {
			response.Error("不符合email格式")
		}

		err = config.Global.Mail.SendMail("消息通知", []string{email}, "madaksdjadsl<h1>测试邮件</h1>il")
		if err != nil {
			response.Error(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "发送成功"})
	})

}
