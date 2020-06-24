package model

import (
	"EK-Server/config"
	"fmt"
	"math"
	"strconv"

	"github.com/labstack/echo"
)

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

// GetList 获取列表
func GetList(c echo.Context) (list *Result, err error) {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	// 转化类型
	p, _ := strconv.Atoi(page)
	// 请求数据
	list = QuickLimit(p, map[string]interface{}{}, &[]Post{}, "posts")
	return
}

// QuickLimit 分页方法省略参数
func QuickLimit(page int, where interface{}, list interface{}, tableName string) *Result {
	return DataBaselimit(10, page, where, list, tableName, "created_at desc")
}

// DataBaselimit  mysql数据分页
func DataBaselimit(limit int, page int, where interface{}, list interface{}, tableName string, orderBy string) *Result {
	db := DB
	// 初始化数据库对象
	resultModal := db.Table(TableName(tableName)).Where(where)
	// 总数
	var count int
	// 绑定总数
	resultModal.Count(&count)
	// 查询绑定用户列表
	resultModal.Offset(limit * (page - 1)).Limit(limit).Order(orderBy).Find(list)
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
	// 显示的最大页码  超出8页是会显示省略号隐藏一部分
	var pagingSize int = 8
	var pagingHalf int = int(math.Ceil(float64(pagingSize / 2)))
	// 建立切片
	arr = make([]Paging, l)
	// 填充内容
	for i := range arr {
		active := (page == i+1)
		arr[i].Href = fmt.Sprintf("%s%d", href, i+1)
		if active {
			arr[i].Href = "javascript:;"
		}
		arr[i].Name = fmt.Sprintf("%d", i+1)
		arr[i].Active = active
	}
	// 处理超出隐藏
	if l > 10 {
		start := page - pagingHalf
		end := page + pagingHalf
		//默认 首....2,3,4,5.....尾
		if start < 0 { //首,1,2,3,4...尾
			start = 0
			end = pagingSize
		} else if end > l { //首...3,4,5,6,尾
			start = l - pagingSize
			end = l
		}

		arr = arr[start:end]
	}

	// 添加首尾
	arr = append([]Paging{{Href: fmt.Sprintf("%s%d", href, 1), Name: "首页", Active: page == 1}}, arr...)
	arr = append(arr, Paging{Href: fmt.Sprintf("%s%d", href, l), Name: "末页", Active: page == l})
	return
}

//TableName 拼接默认表名
func TableName(str string) (result string) {
	result = config.Global.TablePrefix + "_" + str
	return result
}
