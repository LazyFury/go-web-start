package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
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
		response.Error("请选择反馈原因")
	}
	if strings.Trim(f.Content, " ") == "" {
		response.Error("请填写反馈描述")
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

var _ model.Controller = &Feedback{}

// TableName 表名
func (f *Feedback) TableName() string {
	return TableName("feedbacks")
}
