package model

import (
	"strings"

	"github.com/lazyfury/go-web-template/model"
	"github.com/lazyfury/go-web-template/response"
)

// MessageTemplate MessageTemplate
type MessageTemplate struct {
	BaseControll
	Template string `json:"template" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
}

// Validator Validator
func (m MessageTemplate) Validator() error {
	if strings.Trim(m.Template, " ") == "" {
		response.Error("模版内容不可空")
	}
	m.Name = strings.Trim(m.Name, " ")
	if m.Name == "" {
		response.Error("模版名称不可空")
	}
	return nil
}

// TableName TableName
func (m MessageTemplate) TableName() string {
	return TableName("message_templates")
}

// Object Object
func (m MessageTemplate) Object() interface{} {
	return &MessageTemplate{}
}

// Objects Objects
func (m MessageTemplate) Objects() interface{} {
	return &[]MessageTemplate{}
}

// Result Result
func (m MessageTemplate) Result(data interface{}) interface{} {
	return data
}

var _ model.Controller = &MessageTemplate{}
