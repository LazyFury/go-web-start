package model

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/Treblex/go-echo-demo/server/utils/customtype"
	"github.com/Treblex/go-web-template/xmodel"
	"gorm.io/gorm"
)

// MsgType MsgType
type MsgType int

// MsgSendStatus MsgSendStatus
type MsgSendStatus int

const (
	// MsgSendStatusOK 发送成功
	MsgSendStatusOK MsgSendStatus = iota + 1
	// MsgSendStatusFail 发送失败
	MsgSendStatusFail
	// MsgSendStatusWait 等待发送
	MsgSendStatusWait
)

// Message 客户意见反馈
type Message struct {
	BaseControll
	FromUser   uint                  `json:"from_user_id" gorm:"column:from_user_id;not null;comment:'操作人id，点赞 评论我的人'"`
	ToUser     uint                  `json:"to_user_id" gorm:"not null"`
	TemplateID MsgType               `json:"template_id" gorm:"not null"`
	Params     customtype.JSONObject `json:"params" gorm:"type:text;not null;comment:'定义参数 cover template'"`
	URL        string                `json:"url"`
	Status     MsgSendStatus         `json:"status" gorm:"default:1;not null"`
}

// SelectMessage SelectMessage
type SelectMessage struct {
	Message
	// 不是sql字段
	Content      string `json:"content" gorm:"-"`
	FromUserName string `json:"from_username" gorm:"->;column:from_username"`
	// 不再json显示
	Template string `json:"-" gorm:"->;"`
}

// Validator Validator
func (m Message) Validator() error {
	if m.ToUser == 0 {
		utils.Error("接收者不可空")
	}
	if m.FromUser == 0 {
		utils.Error("发送者id不可空")
	}
	if m.TemplateID == 0 {
		utils.Error("消息模版id不可空")
	}
	return nil
}

// TableName TableName
func (m Message) TableName() string {
	return TableName("messages")
}

// Object Object
func (m Message) Object() interface{} {
	return &SelectMessage{}
}

// Objects Objects
func (m Message) Objects() interface{} {
	return &[]SelectMessage{}
}

// Result Result
func (m Message) Result(data interface{}) interface{} {
	r, _ := regexp.Compile("#{(.*?)}")
	arr, ok := reflect.ValueOf(data).Elem().Interface().([]SelectMessage)
	if ok {
		for i, msg := range arr {
			var msgStr = msg.Template
			str := r.ReplaceAllStringFunc(msgStr, func(s string) string {
				s = strings.TrimPrefix(s, "#{")
				s = strings.TrimSuffix(s, "}")

				// 提供默认关键字
				switch s {
				case "username":
					return msg.FromUserName
				}

				// 自定义关键字
				if param, ok := msg.Params.JSON[s]; ok {
					if str, ok := param.(string); ok {
						return str
					}
				}
				return "#{" + s + "}" + "is null"
			})
			arr[i].Content = str
		}
		return arr
	}
	return data
}

// Joins Joins
func (m Message) Joins(db *gorm.DB) *gorm.DB {
	user := &User{}
	db = db.Joins(fmt.Sprintf("left join (select id uid,name from_username from `%s` where deleted_at is null) u1 on u1.`uid`=`%s`.`from_user_id`", user.TableName(), m.TableName()))
	template := &MessageTemplate{}
	db = db.Joins(fmt.Sprintf("left join (select id tid,template from `%s` where deleted_at is null) t1 on t1.`tid`=`%s`.`template_id`", template.TableName(), m.TableName()))
	return db
}

var _ xmodel.Controller = &Message{}
