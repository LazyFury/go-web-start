package wechat

import (
	"EK-Server/util"
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
)

//JsAPIConfig 换取微信分享 jsapi授权
func (m *MP) JsAPIConfig(url string) (conf map[string]string, err error) {

	ticket, err := m.GetJsAPITicket()

	if err != nil {
		return
	}

	noncestr := util.RandStringBytes(32)
	timestamp := time.Now().Unix()

	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, noncestr, timestamp, url)

	h := sha1.New()
	h.Write([]byte(str))

	sign := fmt.Sprintf("%x", h.Sum(nil))
	conf = map[string]string{
		"nonceStr":  noncestr,
		"timestamp": strconv.FormatInt(timestamp, 10),
		"url":       url,
		"rawString": str,
		"signature": sign,
		"appId":     m.Appid,
	}
	return
}
