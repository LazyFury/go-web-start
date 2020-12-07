package controller

import (
	"io"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BaseInterface ControllerInterface
type BaseInterface interface {
	ListAll(c *gin.Context)
	ListPaging(c *gin.Context)
	Detail(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Count(c *gin.Context)
	// Install Install
	Install(g *gin.RouterGroup, path string)
}

type (
	// Controller Controller
	Controller struct {
		DB     *model.GormDB
		Model  model.Controller
		isUser bool //用户可操作的数据，仅查看自己的数据，admin查看全部
	}
)

var _ BaseInterface = &Controller{}

// Install Install
func (t *Controller) Install(g *gin.RouterGroup, path string) {
	Install(g, t, path)
}

// Install Install
func Install(g *gin.RouterGroup, c BaseInterface, path string) {
	route := g.Group(path)
	g.GET(path+"-all", c.ListAll)
	g.GET(path+"-count", c.Count)
	route.GET("", c.ListPaging)
	route.GET("/:id", c.Detail)
	route.POST("", c.Add)
	route.PATCH("/:id", c.Update)
	route.DELETE("/:id", c.Delete)
}

// Update Update
func (t *Controller) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic("请输入id")
	}
	obj := t.Model.Object().(model.Controller)
	where := map[string]interface{}{
		"id": id,
	}
	if err := t.DB.Where(where).First(obj).Error; err != nil {
		panic(err)
	}
	if err := c.ShouldBind(obj); err != nil {
		panic(err)
	}
	if err := obj.Validator(); err != nil {
		panic(err)
	}
	if err := t.DB.Where(where).Updates(obj).Error; err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, utils.JSONSuccess("更新成功", nil))
}

// Delete Delete
func (t *Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic("请输入id")
	}
	obj := t.Model.Object().(model.Controller)
	row := t.DB.Where(map[string]interface{}{
		"id": id,
	}).Delete(obj)
	if err := row.Error; err != nil {
		panic(err)
	}
	if row.RowsAffected == 0 {
		panic("数据不存在!")
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("删除成功", nil))
}

// Add Add
func (t *Controller) Add(c *gin.Context) {
	obj := t.Model.Object().(model.Controller)
	if err := c.ShouldBind(obj); err != nil {
		if err == io.EOF {
			panic("没有传入参数，请使用post json传入参数")
		}
		panic(err)
	}

	if err := obj.Validator(); err != nil {
		panic(err)
	}
	if err := t.DB.Create(obj).Error; err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, utils.JSONSuccess("添加成功", obj))
}

// ListPaging ListPaging
func (t *Controller) ListPaging(c *gin.Context) {
	t.DefaultListPaging(c, nil)
}

// DefaultListPaging DefaultListPaging
func (t *Controller) DefaultListPaging(c *gin.Context, midd model.Middleware) {
	ListPaging(c, t.Model.Objects(), t.DB,
		t.Model.Result,
		func(db *gorm.DB) *gorm.DB { return db.Select([]string{"*"}) },
		func(db *gorm.DB) *gorm.DB {
			return t.Model.Joins(db)
		}, midd, t.checkUser(c))
}

func (t *Controller) checkUser(c *gin.Context) model.Middleware {
	return func(db *gorm.DB) *gorm.DB {
		user := model.GetUserOrEmpty(c)
		if t.isUser {
			db = db.Where(map[string]interface{}{
				"user_id": user.ID,
			})
		}
		return db
	}
}

// ListAll ListAll
func (t *Controller) ListAll(c *gin.Context) {
	obj := t.Model.Objects()
	objModel := t.DB.GetObjectsOrEmpty(obj, nil, func(db *gorm.DB) *gorm.DB {
		return t.Model.Joins(db).Select([]string{"*"})
	}, t.checkUser(c))
	if err := objModel.All(); err != nil {
		panic(err)
	}
	obj = t.Model.Result(obj)
	c.JSON(http.StatusOK, utils.JSONSuccess("", obj))
}

// Detail Detail
func (t *Controller) Detail(c *gin.Context) {
	obj := t.Model.Object()
	id := c.Param("id")
	if id == "" {
		panic("请传入ID")
	}
	if err := t.DB.GetObjectOrNotFound(obj, map[string]interface{}{
		"id": id,
	}, func(db *gorm.DB) *gorm.DB {
		return t.Model.Joins(db).Select([]string{"*"})
	}, t.checkUser(c)); err != nil {
		panic(utils.NotFound)
	}
	obj = t.Model.Result(obj)
	c.JSON(http.StatusOK, utils.JSONSuccess("", obj))
}

// Count Count
func (t *Controller) Count(c *gin.Context) {

	row := t.DB.Table(t.Model.TableName())

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

// ListPaging 处理通用 page size orderby search
func ListPaging(c *gin.Context, obj interface{}, db *model.GormDB, result func(data interface{}) interface{}, pagingMidd model.Middleware, midd ...model.Middleware) {
	objModel := db.GetObjectsOrEmpty(obj, nil, func(db *gorm.DB) *gorm.DB {
		for _, mid := range midd {
			if mid != nil {
				db = mid(db)
			}
		}
		return db
	})

	page, size := model.GetPagingParams(c)

	if err := objModel.Paging(page, size, pagingMidd); err != nil {
		panic(err)
	}

	objModel.Result.List = result(objModel.Result.List)
	c.JSON(http.StatusOK, utils.JSONSuccess("", objModel.Result))
}
