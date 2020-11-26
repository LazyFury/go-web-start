package model

import (
	"fmt"
	"reflect"

	"github.com/Treblex/go-echo-demo/server/utils/customtype/message"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"github.com/labstack/gommon/color"
	"gorm.io/gorm"
)

// Message 客户意见反馈
// 暂定 如果存sql数据量太多，后期尝试redis之类的
type Message struct {
	BaseControll
	UserID    uint           `json:"uid" gorm:"comment:'当前用户id，使用我的uid查询跟我有关当消息'"`
	FromID    uint           `json:"from_id" gorm:"comment:'操作人id，点赞 评论我的人'"`
	Action    message.Action `json:"action"`
	OrderID   uint           `json:"order_id"`
	ArticleID uint           `json:"article_id"`
}

type selectMessage struct {
	BaseControll

	UserID   uint   `json:"uid"`
	FromID   uint   `json:"from_id"`
	UserName string `json:"from_name"`

	Action message.Action `json:"action"`
	// 订单
	OrderID uint        `json:"order_id,omitempty"`
	Order   interface{} `json:"order,omitempty"`
	// 文章
	ArticleID uint          `json:"article_id,omitempty"`
	Articles  selectArticle `json:"article,omitempty"`
}

// NewMessage NewMessage
func NewMessage() *Message {
	m := &Message{}
	m.BaseControll.Model = m
	return m
}

// PointerList PointerList
func (m *Message) PointerList() interface{} {
	return &[]selectMessage{}
}

// Pointer Pointer
func (m *Message) Pointer() interface{} {
	return &selectMessage{}
}

// TableName TableName
func (m *Message) TableName() string {
	return TableName("messages")
}

// IsPublic IsPublic
func (m *Message) IsPublic() bool { return true }

// Joins Joins
func (m *Message) Joins(db *gorm.DB) *gorm.DB {
	db = db.Select("*")

	user := &User{}
	db = db.Joins(fmt.Sprintf("left join (select ifnull(`name`,'未知用户') `user_name`,`id` `user_id` from `%s`) u1 on `u1`.`user_id`=`%s`.`from_id`", user.TableName(), m.TableName()))

	// article := &Articles{}
	// db = db.Joins(fmt.Sprintf("left join (select `title` `article_title`,`id` `article_id`,`desc` `article_desc` from `%s`) t2 on `t2`.`article_id`=`%s`.`article_id`", article.TableName(), m.TableName()))
	return db
}
func (m *Message) getMoreField(v *selectMessage) {
	db := DB
	// 绑定文章信息
	a := selectArticle{}
	row := db.Table(a.TableName()).Where(map[string]interface{}{
		"id": v.ArticleID,
	})
	row = a.Joins(row)
	row.First(&a)
	// 修改对象
	v.Articles = a
}

// Result 处理返回值
func (m *Message) Result(data interface{}, userID uint) interface{} {
	var val, ok = reflect.ValueOf(data).Elem().Interface().([]selectMessage)
	// 如果是列表
	if ok {
		for i := range val {
			m.getMoreField(&val[i])
		}
		return val
	}
	// 如果是详情
	item, ok := reflect.ValueOf(data).Elem().Interface().(selectMessage)
	fmt.Println(item)
	if ok {
		m.getMoreField(&item)
		return item
	}

	return data
}

// Add Add
func (m *Message) Add(c *gin.Context) {
	panic("不可手动添加")
}

// Update Update
func (m *Message) Update(c *gin.Context) {
	panic("记录日志，不可修改")
}

// Delete Delete
func (m *Message) Delete(c *gin.Context) {
	panic("不可删除")
}

var (
	missingParam = ">> 缺少，或参数错误  【%s】"
)

// AddArticleLog 添加文章操作记录
func (m *Message) AddArticleLog(fromID uint, articleID uint, action int) {
	m.AddUserActionLog(map[string]interface{}{
		"fromID":    fromID,
		"articleID": articleID,
		"action":    action,
	})
}

// AddUserActionLog 添加用户操作日志
func (m *Message) AddUserActionLog(data map[string]interface{}) {
	err := writeLogs(data)
	if err != nil {
		color.Printf(color.Red("\n用户通知记录 Err:%s \n"), err.Error())
	}
	return
}
func writeLogs(data map[string]interface{}) error {

	fromID, ok := data["fromID"].(uint)
	if !ok {
		return fmt.Errorf(missingParam, "fromID")
	}
	var orderID uint
	// 订单id
	if data["orderID"] != nil {
		orderID, ok = data["orderID"].(uint)
		if !ok {
			return fmt.Errorf(missingParam, "orderID")
		}
	}
	var articleID uint
	// 文章id
	if data["articleID"] != nil {
		articleID, ok = data["articleID"].(uint)
		if !ok {
			return fmt.Errorf(missingParam, "articleID")
		}
	}

	actionType, ok := data["action"].(int)
	if !ok {
		return fmt.Errorf(missingParam, "action")
	}

	db := DB
	msg := &Message{
		FromID:    fromID,
		OrderID:   orderID,
		ArticleID: articleID,
		Action:    message.Action(actionType),
		BaseControll: BaseControll{
			Code: uuid.New().String(),
		},
	}

	err := db.Create(msg).Error
	return err
}
