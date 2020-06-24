package wechat

import (
	"fmt"
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
