package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"bytes"
)

// Mini Mini
type Mini struct {
	Appid       string `json:"appid"`
	Appsecret   string `json:"appsecret"`
	AssessToken tokenType
}
type result struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// SendSubscribeMessage 	发送订阅消息
func (m *Mini) SendSubscribeMessage(touser string, templateID string, page string, data map[string]interface{}) (err error) {
	token, err := m.GetAccessToken()
	if err != nil {
		return
	}
	msg := map[string]interface{}{
		"access_token": token,
		"touser":       touser,
		"template_id":  templateID,
		"page":         page,
		"data":         data,
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", token)
	body, err := json.Marshal(msg)
	if err != nil {
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return
	}

	var res result
	err = json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
	return
}

// GetAccessToken 获取token
func (m *Mini) GetAccessToken() (token string, err error) {
	// 如果accesstken超时 60为timeout 防止请求超时
	if m.AssessToken.ExpiresIn < time.Now().Unix()-60 {
		err = m.getAccessToken(m.Appid, m.Appsecret)
	}
	token = m.AssessToken.AccessToken
	return
}

func (m *Mini) getAccessToken(appid string, appsecret string) (err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, appsecret)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	var result tokenType
	err = json.NewDecoder(resp.Body).Decode(&result)

	if result.ExpiresIn != 0 {
		result.ExpiresIn += time.Now().Unix() //7200加当前时间为过期时间
	} else {
		err = errors.New("获取微信token授权失败")
	}

	m.AssessToken = result

	return
}
