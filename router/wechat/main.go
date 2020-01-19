package wechat

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"suke-go-test/model"
	"suke-go-test/util"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

var (
	appid     string = "wx8bddf23d9228626d"
	appsecret string = "0f28c6ea02973d1719a3312e17f38501"
	// 拼接微信登陆请求
	loginURL string = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	// 跳转微信登陆授权页
	wechatRedirectURL string = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
	// AccessToken 授权请求
	accessTokenURL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	// jsapi_ticket授权请求
	jsAPITicketURL string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	// AccessToken 全局 AccessToken
	AccessToken *tokenType = &tokenType{}
	// JsAPI 全局jsAPITicket
	JsAPI *jsAPITicket = &jsAPITicket{}
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/wechat"
	wechat := g.Group(baseURL)
	wechat.GET("/jsApiConfig", jsAPIConfig)
	wechat.GET("/wechat_redirect", wechatRedirect)
	wechat.GET("/login", login)
}
func login(c echo.Context) (err error) {
	code := c.QueryParam("code")
	if code == "" {
		return util.JSONErr(c, nil, "code不可空")
	}

	url := fmt.Sprintf(loginURL, appid, appsecret, code)
	res, err := http.Get(url)
	if err != nil {
		return util.JSONErr(c, err, "登陆失败,请求微信服务器失败")
	}
	user := &model.WechatOauth{}
	if err = json.NewDecoder(res.Body).Decode(user); err != nil {
		return util.JSONErr(c, err, "登陆失败,解码失败")
	}
	if user.ExpiresIn == 0 {
		return util.JSONErr(c, nil, "code无效，请尝试重新获取")
	}
	db := util.DB
	if db.Find(&model.WechatOauth{Openid: user.Openid}).RecordNotFound() {
		return util.JSON(c, user, "请先绑定账号", -102)
	}

	return util.JSONSuccess(c, user, "登陆成功")
}
func wechatRedirect(c echo.Context) (err error) {
	host := "http://"
	if c.IsTLS() {
		host = "https://"
	}
	host += c.Request().Host
	// host := c.Request().Host
	redirectURI := host + "/wechat/login"
	redirectURI = url.PathEscape(redirectURI)
	urlStr := fmt.Sprintf(wechatRedirectURL, appid, redirectURI)
	return util.JSONSuccess(c, urlStr, "")
}
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

type tokenType struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

// getAccessToken 获取token
func (t *tokenType) getAccessToken() (token string, err error) {
	// 如果accesstken超时 60为timeout 防止请求超时
	if t.ExpiresIn < time.Now().Unix()-60 {
		err = t.sendAccessTokenReq()
	}
	token = t.AccessToken
	return
}

// access_token 向微信服务器发送请求
func (t *tokenType) sendAccessTokenReq() (err error) {
	fmt.Printf(">>>>重启请求微信服务器,获取微信 access_token>>>>")
	url := fmt.Sprintf(accessTokenURL, appid, appsecret)
	res, err := http.Get(url)

	if err != nil {
		err = errors.New("获取微信授权失败")
		return
	}

	err = json.NewDecoder(res.Body).Decode(t)
	if t.ExpiresIn != 0 {
		t.ExpiresIn += time.Now().Unix() //7200加当前时间为过期时间
	} else {
		err = errors.New("获取微信授权失败")
	}
	return
}

// JsAPITicket c
type jsAPITicket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

func (j *jsAPITicket) getJsAPITicket(token string) (ticket string, err error) {
	// 如果accesstken超时 60为timeout 防止请求超时
	if j.ExpiresIn < time.Now().Unix()-60 {
		err = j.sendJsAPITicketReq(token)
	}
	ticket = j.Ticket
	return
}
func (j *jsAPITicket) sendJsAPITicketReq(token string) (err error) {
	fmt.Printf(">>>>重启请求微信服务器,获取微信 JsAPITicket>>>>")
	url := fmt.Sprintf(jsAPITicketURL, token)
	res, err := http.Get(url)

	if err != nil {
		err = errors.New("api_ticket获取微信授权失败")
		return
	}

	err = json.NewDecoder(res.Body).Decode(j)
	if j.ExpiresIn != 0 {
		j.ExpiresIn += time.Now().Unix() //7200加当前时间为过期时间
	} else {
		err = errors.New("api_ticket获取微信授权失败")
	}
	return
}
