package controller

import (
	"net/http"

	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// BaseInterface ControllerInterface
type BaseInterface interface {
	ListAll(c echo.Context) error
	ListPaging(c echo.Context) error
	Detail(c echo.Context) error
	Add(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type (
	// Controller Controller
	Controller struct {
		DB    *model.GormDB
		Model model.Model
		obj   interface{}
	}
)

// var _ BaseInterface = &Controller{}
func (t *Controller) object() {
	if t.obj == nil {
		t.obj = t.Model.Pointer()
	}
}

func (t *Controller) objects() {
	if t.obj == nil {
		t.obj = t.Model.Pointer()
	}
}

// ListPaging ListPaging
func (t *Controller) ListPaging(c *gin.Context) {
	ListPaging(c, t.Model.PointerList(), t.DB,
		func(db *gorm.DB) *gorm.DB { return t.Model.Joins(db) },
		func(db *gorm.DB) *gorm.DB { return db.Select([]string{"*"}) })
}

// ListPaging ListPaging
func ListPaging(c *gin.Context, obj interface{}, db *model.GormDB, midd model.Middleware, pagingMidd model.Middleware) {
	objModel := db.GetObjectsOrEmpty(obj, nil, func(db *gorm.DB) *gorm.DB {
		return midd(db)
	})

	page, size := model.GetPagingParams(c)

	if err := objModel.Paging(page, size, pagingMidd); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("", objModel.Result))
}

// ListAll ListAll
func (t *Controller) ListAll(c *gin.Context) {
	t.objects()

	objModel := t.DB.GetObjectsOrEmpty(t.obj, nil, func(db *gorm.DB) *gorm.DB {
		return t.Model.Joins(db)
	})
	if err := objModel.All(); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("", t.obj))
}

// Detail Detail
func (t *Controller) Detail(c *gin.Context) {
	t.object()
	id := c.Param("id")
	if id == "" {
		panic("请传入ID")
	}
	if err := t.DB.GetObjectOrNotFound(t.obj, map[string]interface{}{
		"id": id,
	}); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, utils.JSONSuccess("", t.obj))
}
