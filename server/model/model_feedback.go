package model

import (
	"strings"

	"github.com/Treblex/go-web-start/server/utils"
	"github.com/Treblex/go-web-template/xmodel"
)

// Feedback 客户意见反馈
type Feedback struct {
	BaseControll
	Reason  string `json:"reason"`
	Content string `json:"content" gorm:"type:text"`
}

// Validator Validator
func (f *Feedback) Validator() error {
	if strings.Trim(f.Reason, " ") == "" {
		utils.Error("请选择反馈原因")
	}
	if strings.Trim(f.Content, " ") == "" {
		utils.Error("请填写反馈描述")
	}
	return nil
}

// Object Object
func (f *Feedback) Object() interface{} {
	return &Feedback{}
}

// Objects Objects
func (f *Feedback) Objects() interface{} {
	return &[]Feedback{}
}

// Result Result
func (f *Feedback) Result(data interface{}) interface{} {
	return data
}

var _ xmodel.Controller = &Feedback{}

// TableName 表名
func (f *Feedback) TableName() string {
	return TableName("feedbacks")
}
