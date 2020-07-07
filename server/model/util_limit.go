package model

import (
	"math"
)

type (

	// Result 分页方法返回结果
	Result struct {
		Count     int         `json:"count"`
		PageSize  int         `json:"page"`
		PageCount int         `json:"pageCount"`
		PageNow   int         `json:"pageNow"`
		List      interface{} `json:"list"`
	}
)

// DataBaselimit  mysql数据分页
func DataBaselimit(limit int, page int, where interface{}, _model listModel, key string, orderBy string) *Result {
	db := DB
	list := _model.PointerList()

	// 初始化数据库对象
	resultModal := db.Table(_model.TableName())
	if where != nil {
		resultModal = resultModal.Where(where)
	}

	resultModal = _model.Search(resultModal, key)
	// 总数
	var count int
	countRow := resultModal
	countRow.Select("count(id)").Where("deleted_at is NULL").Count(&count)
	// 绑定总数
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
