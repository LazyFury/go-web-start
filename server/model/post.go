package model

import (
	"EK-Server/util/structtype"

	"github.com/jinzhu/gorm"
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
