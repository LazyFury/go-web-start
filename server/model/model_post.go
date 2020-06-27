package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	// Articles 文章模型
	Articles struct {
		gorm.Model
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
		gorm.Model
		ParentID int    `json:"parent_id"`
		Name     string `json:"name"`
	}
)

// Pointer 列表
func (Article *Articles) Pointer() interface{} {
	return &[]Articles{}
}

// TableName 表名
func (Article *Articles) TableName() string {
	return TableName("posts")
}

// Search 搜索
func (Article *Articles) Search(db *gorm.DB, key string) *gorm.DB {
	if key != "" {
		return db.Where("`title` like ?", "%"+key+"%").Or("`desc` like ?", "%"+key+"%")
	}
	return db
}

//List 文章列表
func (Article *Articles) List(c echo.Context) error {
	cid := c.QueryParam("cid")
	if cid != "" {
		cateID, err := strconv.Atoi(cid)
		if err == nil && cateID > 0 {
			return GetList(c, Article, &Articles{CID: cateID})
		}
	}
	return GetList(c, Article, nil)
}

// Delete 删除文章
func (Article *Articles) Delete(c echo.Context) error {
	id := c.Param("id")
	return util.JSONSuccess(c, id, "删除成功")
}

// Detail 文章详情
func (Article *Articles) Detail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return util.JSONErr(c, err, "参数错误")
	}
	fmt.Println(id)
	p := &Articles{Model: gorm.Model{ID: uint(id)}}
	if DB.First(p).RecordNotFound() {
		return util.JSONErr(c, nil, "文章不存在")
	}
	return util.JSONSuccess(c, p, "")
}
