package model

import (
	"strings"

	"github.com/gin-gonic/gin"
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
func (f *Feedback) Add(c *gin.Context) {
	feedback := &Feedback{}

	if err := c.Bind(feedback); err != nil {
		panic("参数错误")
	}

	if strings.Trim(feedback.Reason, " ") == "" {
		panic("请选择反馈原因")
	}
	if strings.Trim(feedback.Content, " ") == "" {
		panic("请填写反馈描述")
	}

	feedback.Empty()
	f.BaseControll.DoAdd(c, feedback)
}

// Update 添加文章
func (f *Feedback) Update(c *gin.Context) {
	feedback := &Feedback{}

	if err := c.Bind(feedback); err != nil {
		panic("参数错误")
	}

	feedback.Empty()
	f.BaseControll.DoUpdate(c, feedback)
}
