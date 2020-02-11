package model

import (
	"github.com/jinzhu/gorm"
)

// API API接口列表
type API struct {
	gorm.Model
	Name string `json:"name"`
	Data string `json:"data"`
}
