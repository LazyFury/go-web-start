package config

import (
	"github.com/lazyfury/go-web-template/config"
	"github.com/lazyfury/go-web-template/tools/mail"
	"github.com/lazyfury/go-web-template/tools/mysql"
	"github.com/lazyfury/go-web-template/tools/sha"
	"github.com/lazyfury/go-web-template/tools/upload"
	"github.com/lazyfury/go-web-template/tools/wechat"
)

// Global 全局配置
var Global *configType = config.ReadConfig(&configType{}, "./config/config.json").(*configType)

type configType struct {
	config.BaseConfig
	Mysql      mysql.Mysql       `json:"mysql"` // 数据库链接
	Mail       mail.Mail         `json:"mail"`
	WechatMP   wechat.MP         `json:"wechat"`
	WechatMini wechat.Mini       `json:"wechat_mini"`
	AliOss     upload.AliOssConf `json:"ali_oss"` //阿里云oss
	Sha1       sha.Sha1          `json:"sha1"`
}
