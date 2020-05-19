package wechat

import (
	"EK-Server/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	// 拼接微信登陆请求
	loginURL string = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	// 跳转微信登陆授权页
	wechatRedirectURL string = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"

	// 微信用户信息
	wechatUserInfo    string = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	wechatUserInfoCgi string = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

//发送code换取微信登陆信息
func sendCodeToWechatServer(code string) (result *model.WechatOauth, msg string) {
	url := fmt.Sprintf(loginURL, appid, appsecret, code)
	result = &model.WechatOauth{}
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

func wechatDoLogin(wechatInfo *model.WechatOauth) (wechatUser *model.WechatOauth, msg string, code int) {
	code = 1
	err := errors.New("")
	//创建微信用户
	wechatUser = &model.WechatOauth{Openid: wechatInfo.Openid}
	// 数据库
	db := model.DB
	//开启自动迁移模式
	db.AutoMigrate(&model.WechatOauth{})

	//使用唯一openid查询用户
	if db.Find(wechatUser).RecordNotFound() {
		fmt.Printf("没有微信用户账户，准备新建...")
		//如不存在，新建用户
		info := &model.WechatOauth{}
		info, msg = updateWechatInfo(wechatInfo, true)
		if msg != "" {
			code = -1
			return
		}
		wechatInfo.Nickname = info.Nickname
		wechatInfo.Sex = info.Sex
		wechatInfo.Headimgurl = info.Headimgurl
		wechatInfo.Province = info.Province
		wechatInfo.City = info.City
		wechatInfo.Country = info.Country
		wechatInfo.Unionid = info.Unionid

		err = db.Create(wechatInfo).Error
		if err != nil {
			code = -1
			msg = "创建微信账号失败"
			return
		}
		db.Find(wechatUser)
	} else {
		fmt.Printf("存在微信用户账户(openid:%s),更新状态...", wechatInfo.Openid)
		// 如果账号存在则更新微信 token 信息
		err = db.Model(wechatUser).Updates(wechatInfo).Error
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

func updateWechatInfo(user *model.WechatOauth, isReg bool) (info *model.WechatOauth, msg string) {
	fmt.Printf("请求微信服务器更新用户信息\n")

	url := ""
	if isReg {
		url = fmt.Sprintf(wechatUserInfo, user.AccessToken, user.Openid) //相对通用
	} else {
		//可以获取到是否关注公众号 为关注到情况下无法获取其他信息
		token, err := AccessToken.getAccessToken()
		if err != nil {
			msg = "获取微信token失败"
			return
		}
		url = fmt.Sprintf(wechatUserInfoCgi, token, user.Openid)
	}

	// fmt.Printf(wechatUserInfoCgi, token, user.Openid)
	// fmt.Printf("\n")
	// fmt.Printf(wechatUserInfo, user.AccessToken, user.Openid)
	// fmt.Printf("\n")
	res, err := http.Get(url)
	if err != nil {
		msg = "拉去用户信息失败"
		return
	}
	fmt.Println("获取用户信息url：", url)
	info = &model.WechatOauth{}
	if err = json.NewDecoder(res.Body).Decode(&info); err != nil {
		msg = "请求微信服务器失败\n"
		return
	}
	return
}
