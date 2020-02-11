package model

import (
	"github.com/jinzhu/gorm"
)

// API API接口列表
type API struct {
	gorm.Model
	Name string `json:"name"`
	Data string `json:"data"`
	Cid  string `json:"cid"`
}

// APICate api分类
type APICate struct {
	gorm.Model
	Name string `json:"name"`
	Desc string `json:"desc"`
}
