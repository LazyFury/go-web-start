package model

import (
	"EK-Server/util/customtype"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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
		Tag     customtype.Array `json:"tag" gorm:"type:varchar(255)"`
		CID     int              `json:"cate_id" gorm:"column:cate_id"`
	}

	// ArticlesCate 文章分类
	ArticlesCate struct {
		BaseControll
		ParentID int    `json:"parent_id"`
		Name     string `json:"name"`
	}
)

// PointerList 列表
func (article *Articles) PointerList() interface{} {
	return &[]Articles{}
}

// Pointer 实例
func (article *Articles) Pointer() interface{} {
	return &Articles{}
}

// TableName 表名
func (article *Articles) TableName() string {
	return TableName("posts")
}

// Search 搜索
func (article *Articles) Search(db *gorm.DB, key string) *gorm.DB {
	if key != "" {
		return db.Where("`title` like ?", "%"+key+"%").Or("`desc` like ?", "%"+key+"%")
	}
	return db
}

//List 文章列表
func (article *Articles) List(c echo.Context) error {
	cid := c.QueryParam("cid")
	if cid != "" {
		cateID, err := strconv.Atoi(cid)
		if err == nil && cateID > 0 {
			return article.BaseControll.GetList(c, &Articles{CID: cateID})
		}
	}
	return article.BaseControll.GetList(c, nil)
}
