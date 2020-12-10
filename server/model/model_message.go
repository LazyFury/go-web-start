package model

import (
	"fmt"

	"gorm.io/gorm"
)

// MsgType MsgType
type MsgType int

// MsgSendStatus MsgSendStatus
type MsgSendStatus int

const (
	// MsgTemplateArticleLike 文章点赞
	MsgTemplateArticleLike MsgType = iota + 1
)
const (
	// MsgSendStatusOK 发送成功
	MsgSendStatusOK MsgSendStatus = iota + 1
	// MsgSendStatusFail 发送失败
	MsgSendStatusFail
	// MsgSendStatusWait 等待发送
	MsgSendStatusWait
)

// MsgTemplate MsgTemplate
var MsgTemplate = map[MsgType]string{
	MsgTemplateArticleLike: `#{username}点赞了你的文章,快去看看吧!`,
}

// Message 客户意见反馈
// 暂定 如果存sql数据量太多，后期尝试redis之类的
// TODO:尝试改改结构
type Message struct {
	BaseControll
	FromUser   uint                   `json:"from_id" gorm:"not null;comment:'操作人id，点赞 评论我的人'"`
	ToUser     uint                   `json:"to_user_id" gorm:"not null"`
	TemplateID MsgType                `json:"template_id" gorm:"not null"`
	Params     map[string]interface{} `json:"params" gorm:"type:'text';not null;comment:'定义参数 cover template"`
	URL        string                 `json:"url"`
	Status     MsgSendStatus          `json:"status" gorm:"default:1;not null"`
}

// Validator Validator
func (m Message) Validator() error {
	if m.ToUser == 0 {
		panic("userid不可空")
	}
	if m.TemplateID == 0 {
		panic("消息类型不可空")
	}
	return nil
}

// TableName TableName
func (m Message) TableName() string {
	return TableName("messages")
}

// Object Object
func (m Message) Object() interface{} {
	return &Message{}
}

// Objects Objects
func (m Message) Objects() interface{} {
	return &[]Message{}
}

// Result Result
func (m Message) Result(data interface{}) interface{} {
	return data
}

// Joins Joins
func (m Message) Joins(db *gorm.DB) *gorm.DB {
	article := &Articles{}
	db = db.Joins(fmt.Sprintf("left join (select id aid,title article_name from `%s`) a1 on a1.aid = `%s`.`param_id`", article.TableName(), m.TableName()))
	return db
}

var _ Controller = &Message{}
