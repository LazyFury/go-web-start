package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type listModel interface {
	// PointerList return gorm.model数组类型，用户分页查询绑定数据
	PointerList() interface{}
	// Pointer
	Pointer() interface{}
	// TableName 自定义表名
	TableName() string
	// Where 搜索条件
	Search(db *gorm.DB, key string) *gorm.DB
	// 列表的补充条件
	Joins(db *gorm.DB) *gorm.DB
	// 列表，增，查，删，改, 统计
	List(c echo.Context) error
	Detail(c echo.Context) error
	Delete(c echo.Context) error
	Add(c echo.Context) error
	Update(c echo.Context) error
	Count(c echo.Context) error
	Install(g *echo.Group, baseURL string) *echo.Group
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

// Search 搜索
func (b *BaseControll) Search(db *gorm.DB, key string) *gorm.DB {
	return db
}

// Joins 链接
func (b *BaseControll) Joins(db *gorm.DB) *gorm.DB {
	return db
}

// List 数据列表
func (b *BaseControll) List(c echo.Context) error {
	return b.GetList(c, nil)
}

// Detail 详情
func (b *BaseControll) Detail(c echo.Context) error {
	return b.GetDetail(c, "")
}

// ListWithOutPaging 直接取所有数据不分页
func (b *BaseControll) ListWithOutPaging(c echo.Context) error {
	db := DB
	list := b.Model.PointerList()

	row := db.Table(b.Model.TableName()).Find(list)

	if row.Error != nil {
		return util.JSONErr(c, nil, "获取失败")
	}

	return util.JSONSuccess(c, list, "")
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

// Count 统计表
func (b *BaseControll) Count(c echo.Context) error {
	db := DB

	row := db.Table(b.Model.TableName())

	//time: 2006-01-02 15:04:05
	start := c.QueryParam("start")
	end := c.QueryParam("end")
	if start == "" {
		return util.JSONErr(c, nil, "请选择查询开始时间")
	}
	//type: year,month,week,day
	queryType := c.QueryParam("type")
	if queryType == "" {
		queryType = "day"
	}
	queryType = strings.ToLower(queryType)

	// 按时间查询
	startTime, err := time.Parse("2006-01-02 15:04:05", start)
	var endTime time.Time
	if end != "" {
		endTime, err = time.Parse("2006-01-02 15:04:05", end)
	} else {
		endTime = time.Now()
	}

	if err != nil {
		return util.JSONErr(c, err, "时间格式错误")
	}

	row = row.Where("`created_at` BETWEEN ? AND ?", startTime.Format(util.TimeZone()), endTime.Format(util.TimeZone()))

	list := []struct {
		Date       string  `json:"date"`
		Count      int     `json:"count"`
		Offset     int     `json:"offset"`
		GrowthRate float64 `json:"growth_rate"`
		// GrowthRateStr string  `json:"growth_rate_str"`
		List interface{} `json:"list,omitempty"`
	}{}

	dateFormat := "%Y-%m-%d"

	queryTypes := map[string]string{
		"day":   "%Y-%m-%d",
		"week":  "%Y-%u",
		"month": "%Y-%m",
		"year":  "%Y",
	}

	if format := queryTypes[queryType]; format != "" {
		dateFormat = format
	}

	// 统计总数量
	var n int
	row = row.Count(&n)
	// 查询近n (天，周，月) 数据
	row = row.Select("DATE_FORMAT(created_at,'" + dateFormat + "') date,count(*) count")
	row = row.Group("date").Order("date asc").Find(&list)
	// 计算近n天数据的增加或者减少
	var defaultCount int
	for i, item := range list {
		list[i].Offset = item.Count - defaultCount
		if defaultCount > 0 {
			list[i].GrowthRate = math.Floor(float64(list[i].Offset)/float64(defaultCount)*10000) / 100
		} else {
			list[i].GrowthRate = 100
		}
		// list[i].GrowthRateStr = fmt.Sprintf("%.2f%%", list[i].GrowthRate)
		defaultCount = item.Count

		// if item.Count > 0 {
		// 	l := b.Model.PointerList()
		// 	row = DB
		// 	row.Table(b.Model.TableName()).Where("DATE_FORMAT(created_at,'"+dateFormat+"') = ?", item.Date).Find(l)
		// 	list[i].List = &l
		// }
	}

	if row.Error != nil {
		return util.JSONErr(c, nil, "")
	}
	return util.JSONSuccess(c, map[string]interface{}{
		"total": &n,
		"list":  list,
	}, "")
}

// Empty 基础参数id,CreatedAt,UpdatedAt置空，避免更新时修改到
func (b *BaseControll) Empty() {
	b.ID = 0
	b.CreatedAt = customtype.LocalTime{Time: time.Time{}}
	b.UpdatedAt = customtype.LocalTime{Time: time.Time{}}
}

// Install 快速注册路由
func (b *BaseControll) Install(g *echo.Group, baseURL string) *echo.Group {
	route := g.Group(baseURL)

	route.GET("", b.Model.List)
	route.GET("/:id", b.Model.Detail)
	route.POST("", b.Model.Add)
	route.PUT("/:id", b.Model.Update)
	route.DELETE("/:id", b.Model.Delete)
	route.GET("-actions/count", b.Model.Count)

	return route

}
