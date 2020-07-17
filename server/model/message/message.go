package message

import (
	"EK-Server/model"
	"fmt"

	"github.com/labstack/gommon/color"
)

// 点赞类型声明
const (
	LIKE    int = 1
	UNLIKE  int = 2
	COMMENT int = 3
	ORDER   int = 4
	BUY     int = 5
	ADDCART int = 6
	SHARE   int = 7
)

// MessageType 操作类型
var MessageType = map[int]string{
	LIKE:    "点赞",
	UNLIKE:  "取消点赞",
	COMMENT: "评论",
	ORDER:   "用户下单",
	BUY:     "用户付款",
	ADDCART: "添加购物车",
	SHARE:   "分享",
}

var (
	missingParam = ">> 缺少，或参数错误  【%s】"
)

// AddUserActionLog 添加用户操作日志
func AddUserActionLog(data map[string]interface{}) {
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
		return fmt.Errorf(missingParam, "actionType")
	}
	remark, ok := data["remark"].(string)
	if !ok {
		return fmt.Errorf(missingParam, "remark")
	}

	db := model.DB
	msg := &model.Message{
		FromID:    fromID,
		OrderID:   orderID,
		ArticleID: articleID,
		Action:    MessageType[actionType],
		Remark:    remark,
	}

	db.NewRecord(msg)
	err := db.Create(msg).Error
	return err
}
