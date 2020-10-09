package model

import (
	"fmt"
	"strings"

	"github.com/Treblex/go-echo-demo/server/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// ArticlesCate 文章分类
type ArticlesCate struct {
	BaseControll
	Name string `json:"name"`
	Key  string `json:"key"`
	Desc string `json:"desc"`
}

// NewArticleCate 文章分类
func NewArticleCate() *ArticlesCate {
	c := &ArticlesCate{}
	c.BaseControll.Model = c
	return c
}

// PointerList 列表
func (a *ArticlesCate) PointerList() interface{} {
	return &[]struct {
		*ArticlesCate
		Count    int `json:"count"`
		TagCount int `json:"tag_count"`
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
	db = db.Select("*")
	article := &Articles{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) count,`cate_id` from `%s` where `%s`.`deleted_at` IS NULL group by `cate_id`) a1 on `%s`.`id`=`a1`.`cate_id`", article.TableName(), article.TableName(), a.TableName()))

	tag := &ArticlesTag{}
	db = db.Joins(fmt.Sprintf("left join (select count(id) tag_count,`cate_id` `tag_cate_id` from `%s` group by `tag_cate_id`) a2 on `%s`.`id`=`a2`.`tag_cate_id`", tag.TableName(), a.TableName()))
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
	return a.BaseControll.DoAdd(c, cate)
}

// Update 添加分类
func (a *ArticlesCate) Update(c echo.Context) error {
	cate := &ArticlesCate{}

	if err := c.Bind(cate); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	cate.Empty()
	return a.BaseControll.DoUpdate(c, cate)
}

// Delete 删除
func (a *ArticlesCate) Delete(c echo.Context) error {
	db := DB
	id := c.Param("id")
	if id == "" {
		return util.JSONErr(c, nil, "参数错误")
	}
	article := &Articles{}
	if hasArticle := db.Model(article).Where(map[string]interface{}{"cate_id": id}).Find(article).RowsAffected; hasArticle > 0 {
		return util.JSONErr(c, nil, "该分类下还有文章，不能删除")
	}

	tag := &ArticlesTag{}
	if hasTag := db.Model(tag).Where(map[string]interface{}{"cate_id": id}).Find(tag).RowsAffected; hasTag > 0 {
		return util.JSONErr(c, nil, "该分类下有标签,无法删除")
	}

	return a.BaseControll.Delete(c)
}
