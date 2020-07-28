package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// MP 微信公众号
type MP struct {
	Appid       string `json:"appid"`
	Appsecret   string `json:"appsecret"`
	JsAPITicket jsAPITicket
	AccessToken tokenType
}

//LoginRedirect 获取登陆链接
func (m *MP) LoginRedirect(redirectURI string) (url string) {
	return fmt.Sprintf(RedirectURL, m.Appid, redirectURI)
}

// GetAccessToken 获取token
func (m *MP) GetAccessToken() (token string, err error) {
	// 如果accesstken超时 60为timeout 防止请求超时
	if m.AccessToken.ExpiresIn < time.Now().Unix()-60 {
		err = m.AccessToken.sendAccessTokenReq(m.Appid, m.Appsecret)
	}
	token = m.AccessToken.AccessToken
	return
}

// GetJsAPITicket 获取JsAPITicket
func (m *MP) GetJsAPITicket() (ticket string, err error) {
	// 如果accesstken超时 60为timeout 防止请求超时
	token, err := m.GetAccessToken()
	if err != nil {
		return
	}
	if m.JsAPITicket.ExpiresIn < time.Now().Unix()-60 {
		err = m.JsAPITicket.sendJsAPITicketReq(token)
	}
	ticket = m.JsAPITicket.Ticket
	return
}

//GetUserInfo 发送code换取微信登陆信息
func (m *MP) GetUserInfo(code string) (result *UserInfo, err error) {
	url := fmt.Sprintf(LoginURL, m.Appid, m.Appsecret, code)
	result = &UserInfo{}
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("获取微信用户信息失败：%v\n", err)
		err = errors.New("请求微信服务器失败，请联系管理员")
		return
	}
	if err = json.NewDecoder(res.Body).Decode(result); err != nil {
		err = errors.New("解码处理微信返回result错误")
		return
	}
	// 数据获取失败
	if result.ExpiresIn == 0 {
		err = errors.New("微信Code失效，请尝试重新获取")
		fmt.Printf("失效的微信code: %+v\n", result)
		return
	}
	return
}
