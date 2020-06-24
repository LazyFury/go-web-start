package wechat

import (
	"EK-Server/util"
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo"
)

var (
	templateMsgURL string = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

//服务端验证token绑定
func signatureCheck(c echo.Context) (err error) {
	echostr := c.QueryParam("echostr")
	signature := c.QueryParam("signature")
	timestamp := c.QueryParam("timestamp")
	nonce := c.QueryParam("nonce")

	tmpArr := []string{"hello world", timestamp, nonce}
	sort.Strings(tmpArr)

	b := sha1.Sum([]byte(strings.Join(tmpArr, "")))
	str := fmt.Sprintf("%x", b)
	log.Println(str)
	if signature == str {
		return c.HTML(http.StatusOK, echostr)
	}
	return util.JSONErr(c, nil, "解密失败")
}

type (
	messageXML struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   string
		FromUserName string
		CreateTime   int64
		Content      string
		MsgType      string `xml:"MsgType"`
		XX           string
	}
	templateMsg struct {
		ToUser     string                     `json:"touser"`
		TemplateID string                     `json:"template_id"`
		URL        string                     `json:"url"`
		Data       map[string]templateMsgData `json:"data"`
	}
	templateMsgData struct {
		Value string `json:"value"`
		Color string `json:"color"`
	}
	templateReturn struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		MsgID   string `json:"msgid"`
	}
)

// 处理微信消息
func handleWechatMessage(c echo.Context) (err error) {
	data := messageXML{}
	b, err := ioutil.ReadAll(c.Request().Body)
	if err = xml.Unmarshal(b, &data); err != nil {
		log.Printf("%+v\n", err)
		return util.JSONErr(c, nil, "err")
	}
	log.Println(data)
	result := messageXML{FromUserName: data.ToUserName, ToUserName: data.FromUserName, MsgType: "text", Content: "你好", CreateTime: time.Now().Unix()}
	b, err = xml.Marshal(&result)
	log.Println(string(b))
	return c.Blob(http.StatusOK, "text/xml", b)
}

func sendTemplateMsgHandler(c echo.Context) error {
	post := &templateMsg{
		ToUser:     "oUsta6PmPtCCs-XSuw02Q07p1OB0",
		URL:        "http://blog.abadboy.cn",
		TemplateID: "O1SYftOnqonEL3aWPVd67-bzxiFsCi_msMmxgXZWzLk",
		Data: map[string]templateMsgData{
			"user": {
				Value: "suke",
				Color: "#dd3333",
			},
		}}
	data, err := sendTemplateMsg(post)
	if err != nil {
		return util.JSONErr(c, data, fmt.Sprintf("%s", err.Error()))
	}
	return util.JSONSuccess(c, nil, "发送成功")
}

func sendTemplateMsg(postData *templateMsg) (body *templateReturn, err error) {
	body = &templateReturn{}
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		err = errors.New("获取assessToken失败")
		return
	}

	b, err := json.Marshal(&postData)
	if err != nil {
		err = errors.New("编码失败")
		return
	}
	reader := bytes.NewReader(b)
	url := fmt.Sprintf(templateMsgURL, accessToken)
	// log.Printf(url)
	res, err := http.Post(url, "application/json", reader)
	if err != nil {
		err = errors.New("发送消息失败")
		return
	}
	defer res.Body.Close()

	_ = json.NewDecoder(res.Body).Decode(body)
	if body.ErrMsg != "ok" {
		err = errors.New("发送失败")
		return
	}
	return
}
