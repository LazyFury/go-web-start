package model

import (
	"EK-Server/util"
	"strconv"

	"github.com/labstack/echo"
)

// BaseControll 空方法用户数据模型继承方法
type BaseControll struct {
	Model listModel
}

// GetList 获取列表
func (b *BaseControll) GetList(c echo.Context, where interface{}) (err error) {
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
	list := DataBaselimit(size, p, where, b.Model, key, orderBy)
	return util.JSONSuccess(c, list, "")
}

// Detail 获取某一条数据
func (b *BaseControll) Detail(c echo.Context, recordNotFoundTips string) error {
	db := DB
	if recordNotFoundTips == "" {
		recordNotFoundTips = "内容不存在"
	}

	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}

	p := b.Model.Pointer()
	where := map[string]interface{}{
		"id": id,
	}
	if db.Where(where).First(p).RecordNotFound() {
		return util.JSONErr(c, nil, recordNotFoundTips)
	}
	return util.JSONSuccess(c, p, "")
}

// Delete 删除数据
func (b *BaseControll) Delete(c echo.Context) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}

	p := b.Model.Pointer()
	row := db.Where(map[string]interface{}{
		"id": id,
	}).Delete(p)

	if row.Error != nil {
		return util.JSONErr(c, nil, "删除失败")
	}

	if row.RowsAffected <= 0 {
		return util.JSONErr(c, nil, "删除失败,数据不存在")
	}

	return util.JSONSuccess(c, nil, "删除成功")
}
