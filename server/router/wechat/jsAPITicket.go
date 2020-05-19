package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	// JsAPI 全局jsAPITicket
	JsAPI *jsAPITicket = &jsAPITicket{}
	// jsapi_ticket授权请求
	jsAPITicketURL string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

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
	fmt.Printf(">>>>重启请求微信服务器,获取微信 JsAPITicket>>>>\n")
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
