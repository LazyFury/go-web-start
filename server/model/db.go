package model

import (
	"EK-Server/config"
	"fmt"
	"math"
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
	db.AutoMigrate(&User{}, &WechatOauth{}, &API{}, &APICate{}, &Goods{}, &GoodsCate{}, &Post{})

	return db
}

type (
	// PageParams PageParams
	PageParams struct {
		Page  int    `json:"page"`
		Limit int    `json:"limit"`
		Order string `json:"order"`
	}
	// Result 分页方法返回结果
	Result struct {
		Count     int         `json:"count"`
		PageSize  int         `json:"page"`
		PageCount int         `json:"pageCount"`
		PageNow   int         `json:"pageNow"`
		List      interface{} `json:"list"`
	}
)

// QuickLimit 分页方法省略参数
func QuickLimit(page int, where interface{}, list interface{}, tableName string) *Result {
	return DataBaselimit(10, page, where, list, tableName, "created_at desc")
}

// DataBaselimit  mysql数据分页
func DataBaselimit(limit int, page int, where interface{}, list interface{}, tableName string, orderBy string) *Result {
	db := DB
	// 初始化数据库对象
	userModal := db.Table(TableName(tableName)).Where(where)
	// 总数
	var count int
	// 绑定总数
	userModal.Count(&count)
	// 查询绑定用户列表
	userModal.Offset(limit * (page - 1)).Limit(limit).Order(orderBy).Find(list)
	var pageCount int = int(math.Ceil(float64(count) / float64(limit)))
	if list == nil {
		list = []string{}
	}
	return &Result{
		Count:     count,
		PageCount: pageCount,
		PageNow:   page,
		PageSize:  limit,
		List:      list,
	}
}

//Paging 分页类型
type Paging struct {
	Href   string
	Name   string
	Active bool
}

//GeneratePaging 生成分页html数组
func GeneratePaging(l int, page int, href string) (arr []Paging) {
	var pagingSize int = 8
	var pagingHalf int = int(math.Ceil(float64(pagingSize / 2)))
	arr = make([]Paging, l)

	for i := range arr {
		active := (page == i+1)
		arr[i].Href = fmt.Sprintf("%s%d", href, i+1)
		if active {
			arr[i].Href = "javascript:;"
		}
		arr[i].Name = fmt.Sprintf("%d", i+1)
		arr[i].Active = active
	}

	if l > 10 {
		start := page - pagingHalf
		end := page + pagingHalf
		if start < 0 {
			start = 0
			end = pagingSize
		} else if end > l {
			start = l - pagingSize
			end = l
		}
		arr = arr[start:end]
	}
	arr = append([]Paging{{Href: fmt.Sprintf("%s%d", href, 1), Name: "首页", Active: page == 1}}, arr...)
	arr = append(arr, Paging{Href: fmt.Sprintf("%s%d", href, l), Name: "末页", Active: page == l})
	return
}

//TableName 拼接默认表名
func TableName(str string) (result string) {
	result = config.Global.TablePrefix + "_" + str
	return result
}
