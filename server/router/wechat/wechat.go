package wechat

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"suke-go-test/model"
	"suke-go-test/util"
	"time"

	"github.com/labstack/echo"
)

var (
	appid     string = "wx8bddf23d9228626d"
	appsecret string = "0f28c6ea02973d1719a3312e17f38501"
	// 拼接微信登陆请求
	loginURL string = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	// 跳转微信登陆授权页
	wechatRedirectURL string = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
	// AccessToken 授权请求
	accessTokenURL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	// jsapi_ticket授权请求
	jsAPITicketURL string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	// 微信用户信息
	wechatUserInfo    string = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	wechatUserInfoCgi string = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/wechat"
	wechat := g.Group(baseURL)

	wechat.GET("/jsApiConfig", jsAPIConfig)
	wechat.GET("/wechat_redirect", wechatRedirect)
	wechat.GET("/login", login)
	wechat.GET("/info", userInfo)
}

// 微信登陆
func login(c echo.Context) (err error) {
	code := c.QueryParam("code")
	if code == "" {
		return util.JSONErr(c, nil, "code不可空")
	}

	// 请求微信登陆返回信息
	wechatLogin, msg := sendCodeToWechatServer(code)
	if msg != "" {
		return util.JSONErr(c, nil, msg)
	}
	// 更新过期时间
	wechatLogin.ExpiresIn += time.Now().Unix()
	// 查询微信用户数据库表
	wechatUser, msg, errCode := wechatDoLogin(wechatLogin.Openid, wechatLogin)
	if errCode == 0 {
		errCode = 1 //成功的返回码
	}
	if msg != "" {
		return util.JSON(c, wechatUser, msg, errCode)
	}
	// 如果uid存在则查询用户 返回用户信息 生产jwt token
	// ...

	return util.JSONSuccess(c, wechatUser, "登陆成功")
}

// 拉去微信用户信息绑定数据类型
type wechatUserType struct {
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Headimgurl string `json:"headimgurl"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Unionid    string `json:"unionid"`
}

// 获取微信用户信息  是否关注
func userInfo(c echo.Context) (err error) {
	wechatID := c.QueryParam("uid")
	if wechatID == "" {
		return util.JSONErr(c, nil, "用户id不可空")
	}
	newID, err := strconv.Atoi(wechatID)
	if err != nil {
		return util.JSONErr(c, err, "参数错误")
	}
	user := model.WechatOauth{ID: newID}

	db := model.DB
	if db.Find(&user).RecordNotFound() {
		return util.JSONErr(c, err, "未找到用户")
	}
	// token := user.AccessToken
	openID := user.Openid
	token, err := AccessToken.getAccessToken()
	if err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}
	url := fmt.Sprintf(wechatUserInfoCgi, token, openID)
	// return util.JSONSuccess(c, url, "")
	res, err := http.Get(url)
	if err != nil {
		return util.JSONErr(c, err, "拉去用户信息失败")
	}

	u := wechatUserType{}
	if err = json.NewDecoder(res.Body).Decode(&u); err != nil {
		return util.JSONErr(c, nil, "请求微信服务器失败")
	}
	return util.JSONSuccess(c, u, "获取成功")
}

// 换取微信网页登陆授权链接
func wechatRedirect(c echo.Context) (err error) {
	host := "http://"
	if c.IsTLS() {
		host = "https://"
	}
	host += c.Request().Host
	host = "http://wechat.abadboy.cn"
	redirectURI := host + "/wechat/login"
	redirectURI = url.PathEscape(redirectURI)
	urlStr := fmt.Sprintf(wechatRedirectURL, appid, redirectURI)
	return util.JSONSuccess(c, urlStr, "")
}

// 换取微信分享 jsapi授权
func jsAPIConfig(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return util.JSONErr(c, nil, "url不可空")
	}

	token, err := AccessToken.getAccessToken()
	if err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}
	ticket, err := JsAPI.getJsAPITicket(token)
	if err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}
	noncestr := util.RandStringBytes(16)
	timestamp := time.Now().Unix()

	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, noncestr, timestamp, url)

	h := sha1.New()
	h.Write([]byte(str))

	sign := fmt.Sprintf("%x", h.Sum(nil))
	return util.JSONSuccess(c, map[string]string{
		"noncestr":  noncestr,
		"timestamp": strconv.FormatInt(timestamp, 10),
		"url":       url,
		"rawString": str,
		"signature": sign,
		"appid":     appid,
	}, "请求成功")
}
