package model

import (
	"EK-Server/config"
	"fmt"
	"time"

	// _ 数据库驱动
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB DB
var DB *gorm.DB

// InitDB InitDB
func InitDB(DataBaseConfig string) *gorm.DB {
	t := time.Now().Format("2006年01-02 15:04:05")
	fmt.Printf("数据库链接>>>>>>>> %s \n", t)
	db, err := gorm.Open("mysql", DataBaseConfig)
	// db, err := gorm.Open("sqlite3", "config/database.db")
	if err != nil {
		panic(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.Global.TablePrefix + "_" + defaultTableName
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &WechatOauth{}, &Goods{}, &GoodsCate{}, &Post{})

	return db
}
