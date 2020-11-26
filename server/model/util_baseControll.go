package model

import (
	"math"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/Treblex/go-echo-demo/server/utils/customtype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model Model
type Model interface {
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
	List(c *gin.Context)
	Detail(c *gin.Context)
	Delete(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Count(c *gin.Context)
	// 快速注册路由
	Install(g *gin.RouterGroup, baseURL string) *gin.RouterGroup

	// 处理列表返回结果
	Result(data interface{}, userID uint) interface{}

	// 是否公开数据，私有数据查询是需要用户id 或者超级管理员权限
	// 用户id字段统一为UserID json:user_id
	IsPublic() bool
}

// BaseControll 空方法用户数据模型继承方法
type BaseControll struct {
	ID        uint                 `json:"id" gorm:"primary_key"`
	Code      string               `json:"code"`
	CreatedAt customtype.LocalTime `json:"created_at"`
	UpdatedAt customtype.LocalTime `json:"updated_at"`
	DeletedAt *time.Time           `json:"deleted_at,omitempty" sql:"index"`
	Model     Model                `json:"-" gorm:"-"`
}

// EmptySystemFiled 置空
type EmptySystemFiled struct {
	A string `json:"created_at,omitempty"`
	B string `json:"updated_at,omitempty"`
}

func (b *BaseControll) model() Model {
	if b.Model == nil {
		b.Model = &Null{}
	}
	return b.Model
}

// SetController SetController
func (b *BaseControll) SetController(m Model) {
	b.Model = m
}

// IsPublic IsPublic
func (b *BaseControll) IsPublic() bool { return true }

// Search 搜索
func (b *BaseControll) Search(db *gorm.DB, key string) *gorm.DB {
	return db
}

// Joins 链接
func (b *BaseControll) Joins(db *gorm.DB) *gorm.DB {
	return db
}

// Result 处理列表返回结果
func (b *BaseControll) Result(data interface{}, userID uint) interface{} {
	return data
}

// List 数据列表
func (b *BaseControll) List(c *gin.Context) {
	b.GetList(c, nil)
}

// Detail 详情
func (b *BaseControll) Detail(c *gin.Context) {
	b.GetDetail(c, "")
}

// Add Add
func (b *BaseControll) Add(c *gin.Context) {
	panic("不可添加")
}

// Update Update
func (b *BaseControll) Update(c *gin.Context) {
	panic("不可修改")
}

// Delete 删除数据
func (b *BaseControll) Delete(c *gin.Context) {
	b.DoDelete(c)
}

// DoDelete DoDelete
func (b *BaseControll) DoDelete(c *gin.Context) {
	db := DB
	id := c.Param("id")
	if id == "" {
		panic("参数错误")
	}

	p := b.model().Pointer()
	row := db.Table(b.model().TableName()).Where(map[string]interface{}{
		"id": id,
	}).Delete(p)

	if row.Error != nil {
		panic("删除失败")
	}

	if row.RowsAffected <= 0 {
		panic("删除失败,数据不存在")
	}

	c.JSON(http.StatusOK, utils.JSONSuccess("删除成功", nil))
}

// ListWithOutPaging 直接取所有数据不分页
func (b *BaseControll) ListWithOutPaging(where interface{}) interface{} {
	db := DB
	list := b.model().PointerList()

	row := db.Table(b.model().TableName())
	if where != nil {
		row = row.Where(where)
	}
	orderBy := "created_at desc,id desc"
	row = row.Order(orderBy)

	row = b.model().Joins(row)

	row = row.Find(list)

	if row.Error != nil {
		return []interface{}{}
	}

	return list
}

// GetList 获取列表
func (b *BaseControll) GetList(c *gin.Context, where interface{}) {

	orderBy := c.Query("order")
	if orderBy == "" {
		orderBy = "created_at desc,id desc"
	}
	key := c.Query("key")

	// 转化类型
	page, size := GetPagingParams(c)

	// 请求数据
	list := b.model().PointerList()
	listModel := DB.GetObjectsOrEmpty(list, where, func(db *gorm.DB) *gorm.DB {
		return db.Order(orderBy)
	})

	if !b.model().IsPublic() {
		listModel.Model = listModel.Model.Where(map[string]interface{}{})
	}

	listModel.Model = b.model().Search(listModel.Model, key)
	if err := listModel.Paging(page, size, func(db *gorm.DB) *gorm.DB {
		return b.model().Joins(db).Select([]string{"*"})
	}); err != nil {
		// pass; result empty array
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("", listModel.Result))
}

// GetDetail 获取某一条数据
func (b *BaseControll) GetDetail(c *gin.Context, recordNotFoundTips string) {

	if recordNotFoundTips == "" {
		recordNotFoundTips = "内容不存在"
	}

	id := c.Param("id")
	if id == "" {
		panic("参数错误")
	}

	p := b.model().Pointer()
	where := map[string]interface{}{
		"id": id,
	}
	if err := DB.GetObjectOrNotFound(p, where, func(db *gorm.DB) *gorm.DB {
		return b.model().Joins(db).Select([]string{"*"})
	}); err != nil {
		panic(recordNotFoundTips)
	}

	p = b.model().Result(p, 0)

	c.JSON(http.StatusOK, utils.JSONSuccess("", p))
}

// DoAdd 添加 需要实现绑定json的部分以及自定义的验证
// 必须重写 需要调用empty避免关键字段修改
func (b *BaseControll) DoAdd(c *gin.Context, data interface{}) {
	db := DB

	elem := reflect.ValueOf(data).Elem()
	code := elem.FieldByName("Code")
	code.SetString(uuid.New().String())

	row := db.Create(data)

	if row.Error != nil {
		panic(row.Error)
	}

	if row.RowsAffected <= 0 {
		panic("添加失败，没有更改")
	}

	c.JSON(http.StatusOK, utils.JSONSuccess("添加成功", nil))
}

// DoUpdate 更新数据  需要实现绑定json的部分以及自定义的验证
// 必须重写 需要调用empty避免关键字段修改
func (b *BaseControll) DoUpdate(c *gin.Context, data interface{}) {
	id := c.Param("id")
	if id == "" {
		panic("参数错误")
	}

	row := DB.Table(b.model().TableName()).Where(map[string]interface{}{
		"id": id,
	}).Updates(data)

	if row.Error != nil {
		panic(row.Error)
	}

	if row.RowsAffected <= 0 {
		panic("没有更改")
	}

	c.JSON(http.StatusOK, utils.JSONSuccess("更新成功", nil))
}

// Count 统计表
func (b *BaseControll) Count(c *gin.Context) {
	db := DB

	row := db.Table(b.model().TableName())

	//time: 2006-01-02 15:04:05 开始时间必选，结束时间判空位当前时间
	start := c.Query("start")
	end := c.Query("end")
	if start == "" {
		panic("请选择查询开始时间")
	}
	//type: year,month,week,day
	queryType := c.Query("type")
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
		panic(err)
	}

	row = row.Where("`created_at` BETWEEN ? AND ?", startTime, endTime)

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
	var n int64
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
		// 	l := b.model().PointerList()
		// 	row = DB
		// 	row.Table(b.model().TableName()).Where("DATE_FORMAT(created_at,'"+dateFormat+"') = ?", item.Date).Find(l)
		// 	list[i].List = &l
		// }
	}

	if row.Error != nil {
		panic(row.Error)
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("", map[string]interface{}{
		"total": &n,
		"list":  list,
	}))
}

// Empty 基础参数id,CreatedAt,UpdatedAt置空，避免更新时修改到
func (b *BaseControll) Empty() {
	b.ID = 0
	b.CreatedAt = customtype.LocalTime{Time: time.Time{}}
	b.UpdatedAt = customtype.LocalTime{Time: time.Now()}
}

// Install 快速注册路由
func (b *BaseControll) Install(g *gin.RouterGroup, baseURL string) *gin.RouterGroup {
	route := g.Group(baseURL)

	route.GET("", b.model().List)
	route.GET("/:id", b.model().Detail)
	route.POST("", b.model().Add)
	route.PUT("/:id", b.model().Update)
	route.DELETE("/:id", b.model().Delete)
	// route.GET("-actions/count", b.model().Count)

	return route

}

// HasOne 避免重复
func (b *BaseControll) HasOne(where interface{}) bool {
	return (DB.Table(b.model().TableName()).Where(where).First(b.model().Pointer()).Error == nil)
}
