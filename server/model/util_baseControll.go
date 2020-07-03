package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// 说明一下这个文件里的方法，这个接口暂时没用到
type controll interface {
	// 列表
	List(c echo.Context) error
	// 详情
	Detail(c echo.Context) error
	// 删除
	Delete(c echo.Context) error
	// 添加
	Add(c echo.Context, data interface{}) error
	// 更新
	Update(c echo.Context, data interface{}) error

	GetList(c echo.Context, where interface{}) error
	GetDetail(c echo.Context, recordNotFoundTips string) error

	// 置空对象
	Empty()
}

type listModel interface {
	// PointerList return gorm.model数组类型，用户分页查询绑定数据
	PointerList() interface{}
	// Pointer
	Pointer() interface{}
	// TableName 自定义表名
	TableName() string
	// Where 搜索条件
	Search(db *gorm.DB, key string) *gorm.DB
}

// BaseControll 空方法用户数据模型继承方法
type BaseControll struct {
	ID        uint                 `json:"id" gorm:"primary_key"`
	CreatedAt customtype.LocalTime `json:"created_at"`
	UpdatedAt customtype.LocalTime `json:"updated_at"`
	DeletedAt *time.Time           `json:"deleted_at,omitempty" sql:"index"`
	Model     listModel            `json:"-" gorm:"-"`
}

// EmptySystemFiled 置空
type EmptySystemFiled struct {
	A string `json:"created_at,omitempty"`
	B string `json:"updated_at,omitempty"`
}

// List 数据列表
func (b *BaseControll) List(c echo.Context) error {
	return b.GetList(c, nil)
}

// Detail 详情
func (b *BaseControll) Detail(c echo.Context) error {
	return b.GetDetail(c, "")
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
		orderBy = "created_at desc,id desc"
	}
	key := c.QueryParam("key")

	// 转化类型
	p, _ := strconv.Atoi(page)
	size, _ := strconv.Atoi(limit)
	// 请求数据
	list := DataBaselimit(size, p, where, b.Model, key, orderBy)
	return util.JSONSuccess(c, list, "")
}

// GetDetail 获取某一条数据
func (b *BaseControll) GetDetail(c echo.Context, recordNotFoundTips string) error {
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
	if db.Table(b.Model.TableName()).Where(where).First(p).RecordNotFound() {
		return util.JSONErr(c, nil, recordNotFoundTips)
	}
	return util.JSONSuccess(c, p, "")
}

// Delete 删除数据 无需重复实现
func (b *BaseControll) Delete(c echo.Context) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}

	p := b.Model.Pointer()
	row := db.Table(b.Model.TableName()).Where(map[string]interface{}{
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

// Add 添加 需要实现绑定json的部分以及自定义的验证
func (b *BaseControll) Add(c echo.Context, data interface{}) error {
	db := DB

	db.NewRecord(data)
	row := db.Create(data)

	if row.Error != nil {
		return util.JSONErr(c, row.Error.Error(), "添加失败")
	}

	if row.RowsAffected <= 0 {
		return util.JSONErr(c, nil, "添加失败，没有更改")
	}

	return util.JSONSuccess(c, nil, "提交成功")
}

// Update 更新数据  需要实现绑定json的部分以及自定义的验证
func (b *BaseControll) Update(c echo.Context, data interface{}) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}

	row := db.Table(b.Model.TableName()).Where(map[string]interface{}{
		"id": id,
	}).Update(data)

	if row.Error != nil {
		return util.JSONErr(c, row.Error.Error(), "更新失败")
	}

	if row.RowsAffected <= 0 {
		return util.JSONErr(c, nil, "没有更改")
	}

	return util.JSONSuccess(c, nil, "更新成功")
}

// Empty 基础参数id,CreatedAt,UpdatedAt置空，避免更新时修改到
func (b *BaseControll) Empty() {
	b.ID = 0
	b.CreatedAt = customtype.LocalTime{Time: time.Time{}}
	b.UpdatedAt = customtype.LocalTime{Time: time.Time{}}
}
