package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"main/config"
	"time"
)

// DB DB
var DB = getDB()

func getDB() *gorm.DB {
	t := time.Now().Format("2006年01-02 15:04:05")
	fmt.Printf("数据库链接>>>>>>>> %s", t)
	db, err := gorm.Open("mysql", config.DataBase)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	// defer db.Close()
	return db
}
