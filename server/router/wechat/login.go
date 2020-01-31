package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"suke-go-test/model"
)

//发送code换取微信登陆信息
func sendCodeToWechatServer(code string, wechat *model.WechatOauth) (result *model.WechatOauth, msg string) {
	url := fmt.Sprintf(loginURL, appid, appsecret, code)
	result = wechat
	res, err := http.Get(url)
	if err != nil {
		msg = "请求微信服务器失败，请联系管理员"
		return
	}
	if err = json.NewDecoder(res.Body).Decode(result); err != nil {
		msg = "解码处理微信返回result错误"
		return
	}
	// 数据获取失败
	if result.ExpiresIn == 0 {
		msg = "微信Code失效，请尝试重新获取"
		fmt.Printf("失效的微信code: %+v\n", result)
		return
	}
	return
}

func wechatDoLogin(Openid string, wechatInfo *model.WechatOauth) (wechatUser *model.WechatOauth, msg string, code int) {
	err := errors.New("")
	//创建微信用户
	wechatUser = &model.WechatOauth{Openid: Openid}
	// 数据库
	db := model.DB
	//开启自动迁移模式
	db.AutoMigrate(&model.WechatOauth{})

	//使用唯一openid查询用户
	if db.Find(wechatUser).RecordNotFound() {
		fmt.Printf("没有微信用户账户，准备新建...")
		//如不存在，新建用户
		err = db.Create(wechatInfo).Error
		if err != nil {
			msg = "创建微信账号失败"
			return
		}
	} else {
		fmt.Printf("存在微信用户账户(openid:%s),更新状态...", wechatInfo.Openid)
		// 如果账号存在则更新微信 token 信息
		err = db.First(wechatUser).Updates(wechatInfo).Error
		if err != nil {
			msg = "更新状态失败"
			return
		}
	}

	// 如果用户未绑定则通知绑定
	if wechatUser.UID == 0 {
		msg = "请先绑定账号，未找到uid"
		code = -102
		return
	}

	return
}
