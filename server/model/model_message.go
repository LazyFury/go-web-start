package model

import (
	"EK-Server/util"

	"github.com/labstack/echo"
)

// Message 客户意见反馈
// 暂定 如果存sql数据量太多，后期尝试redis之类的
type Message struct {
	BaseControll
	FromID    uint   `json:"from_id"`
	OrderID   uint   `json:"order_id"`
	ArticleID uint   `json:"article_id"`
	Action    string `json:"action"`
	Remark    string `json:"remark"`
}

// PointerList PointerList
func (m *Message) PointerList() interface{} {
	return &[]Message{}
}

// Pointer Pointer
func (m *Message) Pointer() interface{} {
	return &Message{}
}

// TableName TableName
func (m *Message) TableName() string {
	return TableName("messages")
}

// Add Add
func (m *Message) Add(c echo.Context) error {
	return util.JSONErr(c, nil, "不可手动添加")
}

// Update Update
func (m *Message) Update(c echo.Context) error {
	return util.JSONErr(c, nil, "记录日志，不可修改")
}

// Delete Delete
func (m *Message) Delete(c echo.Context) error {
	return util.JSONErr(c, nil, "不可删除")
}
