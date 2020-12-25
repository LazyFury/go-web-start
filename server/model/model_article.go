package model

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/Treblex/go-web-start/server/utils/customtype"
	"github.com/Treblex/go-web-template/xmodel"

	"gorm.io/gorm"
)

type (
	// Articles 文章模型
	Articles struct {
		BaseControll
		Title   string           `json:"title" gorm:"not null"`
		Desc    string           `json:"desc"`
		Author  string           `json:"author" gorm:"DEFAULT:'佚名'"`
		Content string           `json:"content" gorm:"type:text"`
		Email   string           `json:"email"`
		Cover   string           `json:"cover" gorm:"DEFAULT:'/static/images/default.jpg'"`
		Tag     customtype.Array `json:"tag" gorm:"type:varchar(255);not null"`
		CateID  uint             `json:"cate_id" gorm:"column:cate_id;not null"`
		UserID
	}
)

// Validator 验证
func (a *Articles) Validator() error {
	a.Title = strings.Trim(a.Title, " ")
	if a.Title == "" {
		return errors.New("文章标题不可空")
	}
	if a.CateID == 0 {
		return errors.New("请选择文章分类")
	}

	if len(a.Tag) == 0 {
		return errors.New("请至少选择一个标签")
	}
	return nil
}

// Object 自身
func (a *Articles) Object() interface{} {
	return &struct {
		Articles
		CateName string `json:"cate_name" gorm:"->"`
	}{}
}

// Objects 自身列表
func (a *Articles) Objects() interface{} {
	return &[]struct {
		Articles
		CateName string `json:"cate_name"  gorm:"->"`
		A        string `json:"content,omitempty"`
	}{}
}

var _ xmodel.Controller = &Articles{}

// TableName 表名
func (a *Articles) TableName() string {
	return TableName("articles")
}

// Search 搜索
func (a *Articles) Search(db *gorm.DB, key string) *gorm.DB {
	if key != "" {
		db = db.Where("`title` like ?", "%"+key+"%").Or("`desc` like ?", "%"+key+"%").Or("`tag` like ?", "%"+key+"%")
	}
	return db
}

// Joins 查询分类名
func (a *Articles) Joins(db *gorm.DB) *gorm.DB {
	// left join 需要手动拼接字段了,gorm默认都是slect tableName.*
	// db = db.Select("id,title,cate_id,cate_name,created_at,updated_at,author,email,cover,tag,`like`,`desc`")
	cate := &ArticlesCate{}
	db = db.Joins(fmt.Sprintf("left join (select name cate_name,id cid from `%s`) cate on cate.cid=`%s`.`cate_id`", cate.TableName(), a.TableName()))
	return db
}

// Result sql查询完成之后调用这里绑定额外的数据
func (a *Articles) Result(data interface{}) interface{} {

	item, ok := reflect.ValueOf(data).Elem().Interface().(Articles)
	if ok {
		return item
	}

	return data
}
