package model

import (
	"EK-Server/config"
	"EK-Server/util/customtype"
	"fmt"
	"time"

	// _ 数据库驱动
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB DB
var DB *gorm.DB

// 需要自动迁移的表
var autoMigrate = []interface{}{
	//user
	&User{},
	&WechatOauth{},
	//goods
	&Goods{},
	&GoodCate{},
	&GoodSku{},
	&GoodStock{},
	//article
	&Articles{},
	&ArticlesCate{},
	//ad
	&AdEvent{},
	&AdGroup{},
	&Ad{},
	//feedback
	&Feedback{},
	// message
	&Message{},

	// 订单
	&Order{},
}

// MysqlConn InitDB
func MysqlConn(DataBaseConfig string) (err error) {
	fmt.Printf("数据库链接>>>准备>> %s \n", time.Now().Format(customtype.DefaultTimeLayout))
	db, err := gorm.Open("mysql", DataBaseConfig)
	// db, err := gorm.Open("sqlite3", "config/database.db")
	if err != nil {
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return TableName(defaultTableName)
	}

	db.LogMode(true)

	db.AutoMigrate(autoMigrate...)

	DB = db
	fmt.Printf("数据库链接>>>成功>> %s \n", time.Now().Format(customtype.DefaultTimeLayout))
	return nil
}

//TableName 拼接默认表名
func TableName(str string) (result string) {
	result = config.Global.TablePrefix + "_" + str
	return result
}
