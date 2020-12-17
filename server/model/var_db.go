package model

import (
	"fmt"
	"time"

	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/Treblex/go-echo-demo/server/utils/customtype"
	"github.com/Treblex/go-web-template/xmodel"
	// _ 数据库驱动
	// _ "gorm.io/gorm/dialects/sqlite"
)

// 需要自动迁移的表
var autoMigrate = []interface{}{
	//user
	&User{},
	&WechatOauth{},
	&WechatMiniUser{},
	//goods
	&Goods{},
	&GoodCate{},
	&GoodSku{},
	&GoodStock{},
	//article
	&Articles{},
	&ArticlesCate{},
	&ArticlesRec{},
	&ArticlesTag{},
	//ad
	&AdEvent{},
	&AdGroup{},
	&Ad{},
	//feedback
	&Feedback{},
	// message
	&Message{},
	&MessageTemplate{},

	// 订单
	&Order{},
}

// DB DB
var DB *xmodel.GormDB = &xmodel.GormDB{}

// MysqlConn InitDB
func MysqlConn(dsn string) (err error) {
	fmt.Printf("数据库链接>>>准备>> %s \n", time.Now().Format(customtype.DefaultTimeLayout))
	err = DB.ConnectMysql(dsn)
	// db, err := gorm.Open("sqlite3", "config/database.db")
	if err != nil {
		return
	}

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return TableName(defaultTableName)
	// }

	DB.AutoMigrate(autoMigrate...)

	fmt.Printf("数据库链接>>>成功>> %s \n", time.Now().Format(customtype.DefaultTimeLayout))
	return nil
}

//TableName 拼接默认表名
func TableName(str string) (result string) {
	result = config.Global.TablePrefix + "_" + str
	return result
}
