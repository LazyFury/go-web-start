package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/util/wechat"
)

func wechatDoLogin(wechatInfo *model.WechatOauth) (wechatUser *model.WechatOauth, err error) {
	//创建微信用户
	wechatUser = wechatInfo
	// 数据库
	db := model.DB
	//开启自动迁移模式
	db.AutoMigrate(&model.WechatOauth{})

	//使用唯一openid查询用户
	if db.Model(wechatUser).Where(map[string]interface{}{
		"open_id": wechatInfo.Openid,
	}).RecordNotFound() {
		fmt.Printf("没有微信用户账户，准备新建...")
		//如不存在，新建用户
		info := &model.WechatOauth{}
		info, err = updateWechatInfo(wechatInfo, true)
		if err != nil {
			return
		}
		wechatInfo = info

		err = db.Create(wechatInfo).Error
		if err != nil {
			err = errors.New("创建微信账号失败")
			return
		}
		db.Find(wechatUser)
	} else {
		fmt.Printf("存在微信用户账户(openid:%s),更新状态...", wechatInfo.Openid)
		// 如果账号存在则更新微信 token 信息
		err = db.Model(wechatUser).Updates(wechatInfo).Error
		if err != nil {
			err = errors.New("更新状态失败")
			return
		}
	}

	return
}

func updateWechatInfo(user *model.WechatOauth, isReg bool) (info *model.WechatOauth, err error) {
	fmt.Printf("请求微信服务器更新用户信息\n")
	url := ""
	if isReg {
		url = fmt.Sprintf(wechat.UserInfoURL, user.AccessToken, user.Openid) //相对通用
	} else {
		var token string
		//可以获取到是否关注公众号 为关注到情况下无法获取其他信息
		token, err = mp.GetAccessToken()
		if err != nil {
			return
		}
		url = fmt.Sprintf(wechat.UserInfoCgiURL, token, user.Openid)
	}

	// fmt.Printf(wechatUserInfoCgi, token, user.Openid)
	// fmt.Printf("\n")
	// fmt.Printf(wechatUserInfo, user.AccessToken, user.Openid)
	// fmt.Printf("\n")
	res, err := http.Get(url)
	if err != nil {
		return
	}
	fmt.Println("获取用户信息url：", url)
	info = &model.WechatOauth{}
	if err = json.NewDecoder(res.Body).Decode(&info); err != nil {
		err = errors.New("解码用户信息失败")
		return
	}
	return
}
