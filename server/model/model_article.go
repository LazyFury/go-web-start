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
)

// PointerList 列表
func (a *Articles) PointerList() interface{} {
	return &[]showArticle{}
}

// Pointer 实例
func (a *Articles) Pointer() interface{} {
	return &showArticle{}
}

// Search 搜索
func (a *Articles) Search(db *gorm.DB, key string) *gorm.DB {
	if key != "" {
		return db.Where("`title` like ?", "%"+key+"%").Or("`desc` like ?", "%"+key+"%")
	}
	return db
}

//List 文章列表
func (a *Articles) List(c echo.Context) error {
	cid := c.QueryParam("cid")
	if cid != "" {
		cateID, err := strconv.Atoi(cid)
		if err == nil && cateID > 0 {
			return a.BaseControll.GetList(c, &Articles{CID: cateID})
		}
	}
	return a.BaseControll.GetList(c, nil)
}

// Detail 文章详情
func (a *Articles) Detail(c echo.Context) error {
	return a.BaseControll.GetDetail(c, "文章不存在")
}

// Add 添加
func (a *Articles) Add(c echo.Context) error {
	article := &Articles{}

	if err := c.Bind(article); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	article.Title = strings.Trim(article.Title, " ")
	if article.Title == "" {
		return util.JSONErr(c, nil, "文章标题不可空")
	}

	if article.CID == 0 {
		return util.JSONErr(c, nil, "请选择文章分类")
	}

	article.Empty()

	return a.BaseControll.Add(c, article)
}

// Update Update
func (a *Articles) Update(c echo.Context) error {
	article := &Articles{}

	if err := c.Bind(article); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	a.Empty()
	return a.BaseControll.Update(c, a)
}
