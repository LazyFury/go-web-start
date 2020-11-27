package controller

import (
	"io"
	"net/http"

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
