package model

import (
	"EK-Server/util/structtype"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	// Post 文章模型
	Post struct {
		gorm.Model
		Title   string           `json:"title" gorm:"not null"`
		Desc    string           `json:"desc"`
		Author  string           `json:"author" gorm:"DEFAULT:'佚名'"`
		Content string           `json:"content" gorm:"type:text"`
		Email   string           `json:"email"`
		Cover   string           `json:"cover" gorm:"DEFAULT:'/static/images/default.jpg'"`
		Tag     structtype.Array `json:"tag" gorm:"type:varchar(255)"`
	}
)

//List 文章列表
func (post *Post) List(c echo.Context) (posts *Result, err error) {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	// 转化类型
	p, _ := strconv.Atoi(page)
	// 请求数据
	posts = QuickLimit(p, &Post{}, &[]Post{}, "posts")
	return
}
