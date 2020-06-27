package model

import (
	"EK-Server/config"
	"EK-Server/util"
	"math"
	"strconv"

	"github.com/jinzhu/gorm"
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

	listModel interface {
		// Pointer return gorm.model数组类型，用户分页查询绑定数据
		Pointer() interface{}
		// TableName 自定义表名
		TableName() string
		// Where 搜索条件
		Search(db *gorm.DB, key string) *gorm.DB
	}
)

// GetList 获取列表
func GetList(c echo.Context, listModel listModel, where interface{}) (err error) {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}
	orderBy := c.QueryParam("order")
	if orderBy == "" {
		orderBy = "created_at desc"
	}
	key := c.QueryParam("key")

	// 转化类型
	p, _ := strconv.Atoi(page)
	size, _ := strconv.Atoi(limit)
	// 请求数据
	list := DataBaselimit(size, p, where, listModel, key, orderBy)
	return util.JSONSuccess(c, list, "")
}

// DataBaselimit  mysql数据分页
func DataBaselimit(limit int, page int, where interface{}, _model listModel, key string, orderBy string) *Result {
	db := DB
	list := _model.Pointer()

	// 初始化数据库对象
	resultModal := db.Table(_model.TableName())
	if where != nil {
		resultModal = resultModal.Where(where)
	}

	resultModal = _model.Search(resultModal, key)
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

//TableName 拼接默认表名
func TableName(str string) (result string) {
	result = config.Global.TablePrefix + "_" + str
	return result
}
