package model

import (
	"github.com/treblex/go-echo-demo/server/util"
	"strings"

	"github.com/labstack/echo/v4"
)

// Feedback 客户意见反馈
type Feedback struct {
	BaseControll
	Reason  string `json:"reason"`
	Content string `json:"content" gorm:"type:text"`
}

// PointerList 列表实例
func (f *Feedback) PointerList() interface{} {
	return &[]struct {
		*Feedback
		*EmptySystemFiled
	}{}
}

// Pointer 实例
func (f *Feedback) Pointer() interface{} {
	return &Feedback{}
}

// TableName 表名
func (f *Feedback) TableName() string {
	return TableName("feedbacks")
}

// Add 添加文章
func (f *Feedback) Add(c echo.Context) error {
	feedback := &Feedback{}

	if err := c.Bind(feedback); err != nil {
		return util.JSONErr(c, nil, "参数错误")
	}

	if strings.Trim(feedback.Reason, " ") == "" {
		return util.JSONErr(c, nil, "请选择反馈原因")
	}
	if strings.Trim(feedback.Content, " ") == "" {
		return util.JSONErr(c, nil, "请填写反馈描述")
	}

	feedback.Empty()
	return f.BaseControll.DoAdd(c, feedback)
}

// Update 添加文章
func (f *Feedback) Update(c echo.Context) error {
	feedback := &Feedback{}

	if err := c.Bind(feedback); err != nil {
		return util.JSONErr(c, nil, "参数错误")
	}

	feedback.Empty()
	return f.BaseControll.DoUpdate(c, feedback)
}
