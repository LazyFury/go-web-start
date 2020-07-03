package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"
	"strconv"
	"strings"

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

	// 尝试列表或者详情隐藏部分隐私字段
	showArticle struct {
		*Articles
		AnyThing string `json:"updated_at1,omitempty"`
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
	return &[]showArticle{}
}

// Pointer 实例
func (article *Articles) Pointer() interface{} {
	return &showArticle{}
}

// TableName 表名
func (article *Articles) TableName() string {
	return TableName("articles")
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

// Detail 文章详情
func (article *Articles) Detail(c echo.Context) error {
	return article.BaseControll.GetDetail(c, "文章不存在")
}

// Add 添加
func (article *Articles) Add(c echo.Context) error {
	a := &Articles{}

	if err := c.Bind(a); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	if strings.Trim(a.Title, " ") == "" {
		return util.JSONErr(c, nil, "文章标题不可空")
	}

	a.Empty()

	return article.BaseControll.Add(c, a)
}

// Update Update
func (article *Articles) Update(c echo.Context) error {
	a := &Articles{}

	if err := c.Bind(a); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	a.Empty()
	return article.BaseControll.Update(c, a)
}
