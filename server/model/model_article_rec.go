package model

import (
	"strings"

	"github.com/Treblex/go-echo-demo/server/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// ArticlesRec 文章分类
type ArticlesRec struct {
	BaseControll
	Name string `json:"name"`
	Key  string `json:"key"`
	Desc string `json:"desc"`
}

// PointerList 列表
func (a *ArticlesRec) PointerList() interface{} {
	return &[]struct {
		*ArticlesRec
		Count int `json:"count"`
	}{}
}

// Pointer 实例
func (a *ArticlesRec) Pointer() interface{} {
	return &struct {
		*ArticlesRec
		*EmptySystemFiled
	}{}
}

// TableName 表名
func (a *ArticlesRec) TableName() string {
	return TableName("article_rec")
}

// Joins  查询相关文章数据
func (a *ArticlesRec) Joins(db *gorm.DB) *gorm.DB {
	return db
}

// List 分页
func (a *ArticlesRec) List(c echo.Context) error {
	return util.JSONSuccess(c, a.BaseControll.ListWithOutPaging(nil), "")
}

// Add 添加分类
func (a *ArticlesRec) Add(c echo.Context) error {
	cate := &ArticlesRec{}

	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	cate.Name = strings.Trim(cate.Name, " ")
	if cate.Name == "" {
		return util.JSONErr(c, nil, "名称不可空")
	}

	cate.Empty()
	return a.BaseControll.DoAdd(c, cate)
}

// Update 添加分类
func (a *ArticlesRec) Update(c echo.Context) error {
	cate := &ArticlesRec{}

	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	cate.Empty()
	return a.BaseControll.DoUpdate(c, cate)
}

// Delete 删除
func (a *ArticlesRec) Delete(c echo.Context) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}
	article := &Articles{}
	if hasArticle := db.Model(article).Where(map[string]interface{}{"cate_id": id}).Find(article).RowsAffected; hasArticle > 0 {
		return util.JSONErr(c, nil, "该推荐位下还有文章，不能删除")
	}
	return a.BaseControll.Delete(c)
}
