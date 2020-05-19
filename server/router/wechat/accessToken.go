package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	// AccessToken 全局 AccessToken
	AccessToken *tokenType = &tokenType{}
	// AccessToken 授权请求
	accessTokenURL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

// 获取全局 accesstoken 类型
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
	fmt.Printf(">>>>重启请求微信服务器,获取微信 access_token>>>>\n")
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
