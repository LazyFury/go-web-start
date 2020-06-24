package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// JsAPITicket c
type jsAPITicket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

func (w *jsAPITicket) sendJsAPITicketReq(token string) (err error) {
	fmt.Printf(">>>>重启请求微信服务器,获取微信 JsAPITicket>>>>\n")
	url := fmt.Sprintf(jsAPITicketURL, token)
	res, err := http.Get(url)

	if err != nil {
		err = errors.New("api_ticket获取微信授权失败")
		return
	}
	// fmt.Println(url)
	err = json.NewDecoder(res.Body).Decode(w)
	if w.ExpiresIn != 0 {
		w.ExpiresIn += time.Now().Unix() //7200加当前时间为过期时间
	} else {
		err = errors.New("api_ticket获取微信授权失败")
	}

	return
}
