package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// 获取全局 accesstoken 类型
type tokenType struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

// access_token 向微信服务器发送请求
func (t *tokenType) sendAccessTokenReq(appid string, appsecret string) (err error) {
	fmt.Printf(">>>>请求微信服务器,获取微信 access_token>>>>\n")
	url := fmt.Sprintf(accessTokenURL, appid, appsecret)
	res, err := http.Get(url)

	if err != nil {
		err = errors.New("请求微信授权服务器失败")
		return
	}

	err = json.NewDecoder(res.Body).Decode(t)
	if t.ExpiresIn != 0 {
		t.ExpiresIn += time.Now().Unix() //7200加当前时间为过期时间
	} else {
		fmt.Println(url)
		fmt.Println(t)
		err = errors.New("获取微信token授权失败")
	}
	return
}
