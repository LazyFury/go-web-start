package api

import (
	"EK-Server/util"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var appid = "wx2e3ad7f8f3558963"
var secret = "0493d2ec984ba126909ba24e449ddc5b"

var code2SessionKeyURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

type (
	code2SessionKey struct {
		OpenID     string `json:"open_id"`
		SessionKey string `json:"session_key"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	wechatLoginParams struct {
		EncryptedData string
		Iv            string
	}
)

func wehcatMini(g *echo.Group) {
	mini := g.Group("/wechat-mini")
	mini.POST("/login", wechatMiniLogin)
}

func wechatMiniLogin(c echo.Context) error {
	jsCode := c.QueryParam("js_code")
	if jsCode == "" {
		return util.JSONErr(c, nil, "请输入js_code")
	}

	var param wechatLoginParams

	if err := c.Bind(&param); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	if param.EncryptedData == "" || param.Iv == "" {
		return util.JSONErr(c, nil, "请传入用户信息")
	}

	// 请求微信服务器
	url := fmt.Sprintf(code2SessionKeyURL, appid, secret, jsCode)
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return util.JSONErr(c, nil, "获取失败")
	}

	// 解码
	var m code2SessionKey

	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return util.JSONErr(c, err, "获取微信返回内容失败")
	}

	// 获取session失败
	if m.ErrCode != 0 {
		return util.JSONErr(c, nil, m.ErrMsg)
	}

	if m.SessionKey == "" {
		return util.JSONErr(c, nil, "获取session_key失败")
	}

	// baseDecode
	key, _ := base64.StdEncoding.DecodeString(m.SessionKey)
	encryptedData, _ := base64.StdEncoding.DecodeString(param.EncryptedData)
	iv, _ := base64.StdEncoding.DecodeString(param.Iv)
	// 解密数据
	result := wechatMiniDecoder(string(encryptedData), key, string(iv))

	// 解码用户信息
	var info map[string]interface{}

	if err := json.Unmarshal(result, &info); err != nil {
		return util.JSONErr(c, nil, "解码用户信息失败")
	}

	return util.JSONSuccess(c, info, "")
}

func wechatMiniDecoder(str string, key []byte, iv string) []byte {
	c, _ := aes.NewCipher(key)
	strNew := []byte(str)

	cbcDecoder := cipher.NewCBCDecrypter(c, []byte(iv))
	plaintextCopy := make([]byte, len(strNew))

	cbcDecoder.CryptBlocks(plaintextCopy, strNew)
	return plaintextCopy
}
