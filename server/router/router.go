package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-start/server/router/api"
	"github.com/lazyfury/go-web-template/response"
)

// Start 入口
func Start(e *gin.RouterGroup) {
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

	index.GET("/wechat_mini_qrcode", func(c *gin.Context) {
		mini := config.Global.WechatMini
		// 获取领跑
		token, err := mini.GetAccessToken()
		if err != nil {
			c.JSON(http.StatusBadGateway, err)
			return
		}

		// 组织请求参数
		data := map[string]interface{}{
			"scene":      1,
			"page":       "pages/index/index",
			"is_hyaline": true,
		}
		body, err := json.Marshal(data)
		if err != nil {
			c.JSON(http.StatusBadGateway, err)
			return
		}
		// 发送请求
		res, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", token),
			"content-type:application/json",
			bytes.NewBuffer(body))
		if err != nil {
			c.JSON(http.StatusBadGateway, err)
			return
		}
		defer res.Body.Close()

		// 获取response body
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusOK, err)
			return
		}
		// 如果返回了json
		_map := &map[string]interface{}{}
		err = json.Unmarshal(b, _map)
		if err == nil {
			c.JSON(http.StatusOK, _map)
			return
		}
		// 否则返回文件流
		c.Data(http.StatusOK, "image/png;", b)
	})
}
