package model

import (
	"errors"
	"reflect"
	"strings"

	"github.com/Treblex/go-web-template/model"

	"gorm.io/gorm"
)

// ArticlesRec 文章分类
type ArticlesRec struct {
	BaseControll
	Name string `json:"name" gorm:"unique;not null"`
	Key  string `json:"key"`
	IDs  string `json:"article_ids"`
	Desc string `json:"desc"`
}
type showArticleRec struct {
	ArticlesRec
	// *EmptySystemFiled
	List  []Articles `json:"list" gorm:"-"`
	Count int        `json:"count" gorm:"-"`
}

var _ model.Controller = &ArticlesRec{}

// Object Object
func (a *ArticlesRec) Object() interface{} {
	return &showArticleRec{}
}

// Objects Object
func (a *ArticlesRec) Objects() interface{} {
	return &[]showArticleRec{}
}

// Validator Validator
func (a *ArticlesRec) Validator() error {
	a.Name = strings.Trim(a.Name, " ")
	if a.Name == "" {
		return errors.New("请输入推荐位名称")
	}
	return nil
}

// TableName 表名
func (a *ArticlesRec) TableName() string {
	return TableName("article_rec")
}

// Joins  查询相关文章数据
func (a *ArticlesRec) Joins(db *gorm.DB) *gorm.DB {
	return db
}

func (a *ArticlesRec) getArticle(item *showArticleRec) {
	ids := strings.Split(item.IDs, ",")
	article := &Articles{}
	articles := []Articles{}

	row := DB.Table(article.TableName()).Where("id IN (?)", ids).Find(&articles)
	if row.Error == nil && len(articles) > 0 {
		item.List = articles
	}

	l := len(articles)
	item.Count = l
	if l == 0 {
		item.List = []Articles{}
	}

	// fmt.Println(item)
}

// Result 处理结构
func (a *ArticlesRec) Result(data interface{}) interface{} {
	interf := reflect.ValueOf(data).Elem().Interface()
	arr, ok := interf.([]showArticleRec)
	if ok {
		for i := range arr {
			a.getArticle(&arr[i])
		}
		return arr
	}

	item, ok := interf.(showArticleRec)
	if ok {
		a.getArticle(&item)
		return item
	}
	return data
}
