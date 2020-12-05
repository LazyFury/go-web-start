package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ArticlesCate 文章分类
type ArticlesCate struct {
	BaseControll
	Name string `json:"name" gorm:"not null;unique"`
	Key  string `json:"key"`
	Desc string `json:"desc"`
}

type selectArticleCate struct {
	ArticlesCate
	Count    int `json:"count" gorm:"->"`
	TagCount int `json:"tag_count" gorm:"->"`
}

var _ Controller = &ArticlesCate{}

// Validator Validator
func (a *ArticlesCate) Validator() error {
	a.Name = strings.Trim(a.Name, " ")
	if a.Name == "" {
		panic("分类名称不可空")
	}
	return nil
}

// Object Object
func (a *ArticlesCate) Object() interface{} {
	return &selectArticleCate{}
}

// Objects Objects
func (a *ArticlesCate) Objects() interface{} {
	return &[]selectArticleCate{}
}

//Result Result
func (a *ArticlesCate) Result(obj interface{}) interface{} {
	return obj
}

// TableName 表名
func (a *ArticlesCate) TableName() string {
	return TableName("article_cates")
}

// Joins  查询相关文章数据
func (a *ArticlesCate) Joins(db *gorm.DB) *gorm.DB {
	article := &Articles{}
	db = db.Joins(fmt.Sprintf("left join (select count(1) count,`cate_id` from `%s` where `%s`.`deleted_at` IS NULL group by `cate_id`) a1 on `%s`.`id`=`a1`.`cate_id`", article.TableName(), article.TableName(), a.TableName()))

	tag := &ArticlesTag{}
	db = db.Joins(fmt.Sprintf("left join (select count(1) tag_count,`cate_id` `tag_cate_id` from `%s` group by `tag_cate_id`) a2 on `%s`.`id`=`a2`.`tag_cate_id`", tag.TableName(), a.TableName()))
	return db
}
