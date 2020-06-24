package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"

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
		Tag     customtype.Array `json:"tag" gorm:"type:varchar(255)"`
	}
)

//List 文章列表
func (post *Post) List(c echo.Context) error {
	posts, err := GetList(c)
	if err != nil {
		return util.JSONErr(c, nil, err.Error())
	}
	return util.JSONSuccess(c, posts, "")
}

// Delete 删除文章
func (post *Post) Delete(c echo.Context) error {
	id := c.Param("id")
	return util.JSONSuccess(c, id, "删除成功")
}

// Detail 文章详情
func (post *Post) Detail(c echo.Context) error {
	return util.JSONSuccess(c, nil, "")
}
