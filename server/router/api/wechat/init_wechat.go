package wechat

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-start/server/utils"
)

var mp = &config.Global.WechatMP

// Init 初始化
func Init(g *gin.RouterGroup) {
	baseURL := "/wechat"
	wechat := g.Group(baseURL)

	wechat.GET("/js-api-config", jsAPIConfig)
	wechat.GET("/redirect", wechatRedirect)
	wechat.GET("/login", login)

	wechat.GET("/info", userInfo)
	wechat.GET("/signature", signatureCheck)                 //配置接口token验证
	wechat.POST("/signature", handleWechatMessage)           //服务token验证，验证成功之后微信会post用户消息 和 事件到这个接口
	wechat.GET("/send-template-msg", sendTemplateMsgHandler) //发送模版消息

}

// 微信登陆
func login(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		utils.Error("code不可空")
	}

	// 请求微信登陆返回信息
	WeChatLogin, err := mp.GetUserInfo(code)
	if err != nil {
		utils.Error(err)
	}
	// 更新过期时间
	WeChatLogin.ExpiresIn += time.Now().Unix()
	// 查询微信用户数据库表
	WeChatUser, err := wechatDoLogin(&model.WechatOauth{UserInfo: *WeChatLogin})
	if err != nil {
		utils.Error(utils.JSONError(err.Error(), WeChatUser))
	}
	// 如果uid存在则查询用户 返回用户信息 生产jwt token
	// TODO: ...

	// 如果用户未绑定则通知绑定
	// if WeChatUser.UID == 0 {
	// 	return util.JSON(c, nil, "", util.BindWeChat)
	// }

	c.JSON(http.StatusOK, utils.JSONSuccess(err.Error(), WeChatUser))
}

// 获取微信用户信息  是否关注
func userInfo(c *gin.Context) {
	wechatID := c.Query("id")
	if wechatID == "" {
		utils.Error("用户id不可空")
	}
	newID, err := strconv.Atoi(wechatID)
	if err != nil {
		utils.Error(err)
	}
	if newID < 1 {
		utils.Error("用户id不可为空")
	}
	user := &model.WechatOauth{BaseControll: model.BaseControll{ID: uint(newID)}}

	db := model.DB
	if db.Where(user).Find(user).Error != nil {
		utils.Error("未找到用户")
	}
	// token := user.AccessToken
	fmt.Println(time.Now().Unix())
	if user.CreatedAt.Unix()-time.Now().Unix() > 3600*24*10 || user.Nickname == "" || user.Headimgurl == "" {
		info, err := updateWechatInfo(user, false)
		if err != nil {
			utils.Error(err)
		}

		db.Model(&user).Updates(&info)
	}

	c.JSON(http.StatusOK, utils.JSONSuccess("获取成功", user))
}

// 换取微信网页登陆授权链接
func wechatRedirect(c *gin.Context) {
	host := "http://"
	if config.Global.SupportTls {
		host = "https://"
	}
	host += c.Request.Host
	redirectURI := host + "/api/v1/wechat/login"

	callbackURL := c.Query("callback")
	if callbackURL != "" {
		redirectURI = callbackURL
	}

	redirectURI = url.PathEscape(redirectURI)
	urlStr := mp.LoginRedirect(redirectURI)
	c.Redirect(http.StatusFound, urlStr)
}

// 换取微信分享 jsapi授权
func jsAPIConfig(c *gin.Context) {
	urlStr := c.Query("url")
	if urlStr == "" {
		utils.Error("url不可空")
	}
	urlStr, _ = url.QueryUnescape(urlStr)

	conf, err := mp.JsAPIConfig(urlStr)
	if err != nil {
		utils.Error(err)
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("请求成功", conf))
}
