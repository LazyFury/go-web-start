package model

import "github.com/jinzhu/gorm"

type (
	// Post 文章模型
	Post struct {
		gorm.Model
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		Author  string `json:"author"`
		Context string `json:"context"`
		Email   string `json:"email"`
	}
)
