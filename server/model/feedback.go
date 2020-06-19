package model

import "github.com/jinzhu/gorm"

// FeedBack 客户意见反馈
type FeedBack struct {
	gorm.Model
	Reason  string `json:"reason"`
	Content string `json:"content" gorm:"type:text"`
}
