package model

import (
	"EK-Server/util"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// ArticlesCate 文章分类
type ArticlesCate struct {
	BaseControll
	Name string `json:"name"`
	Key  string `json:"key"`
	Desc string `json:"desc"`
}

// PointerList 列表
func (a *ArticlesCate) PointerList() interface{} {
	return &[]struct {
		*ArticlesCate
		*EmptySystemFiled
		Count int `json:"count"`
	}{}
}

// Pointer 实例
func (a *ArticlesCate) Pointer() interface{} {
	return &struct {
		*ArticlesCate
		*EmptySystemFiled
	}{}
}

// TableName 表名
func (a *ArticlesCate) TableName() string {
	return TableName("article_cates")
}

// Joins  查询相关文章数据
func (a *ArticlesCate) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("`name`,`desc`,`key`,`id`,a1.count")
	articleTable := TableName("articles")
	db = db.Joins(fmt.Sprintf("left join (select count(id) count,`cate_id` from `%s` group by `cate_id`) a1 on `%s`.`id`=`a1`.`cate_id`", articleTable, a.TableName()))
	return db
}

// List 分页
func (a *ArticlesCate) List(c echo.Context) error {
	return util.JSONSuccess(c, a.BaseControll.ListWithOutPaging(nil), "")
}

// Add 添加分类
func (a *ArticlesCate) Add(c echo.Context) error {
	cate := &ArticlesCate{}

	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	cate.Name = strings.Trim(cate.Name, " ")
	if cate.Name == "" {
		return util.JSONErr(c, nil, "分类名称不可空")
	}

	cate.Empty()
	return a.BaseControll.Add(c, cate)
}

// Update 添加分类
func (a *ArticlesCate) Update(c echo.Context) error {
	cate := &ArticlesCate{}

	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	cate.Empty()
	return a.BaseControll.Update(c, cate)
}
