package message

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
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
	LIKE:    "点赞了你的文章",
	UNLIKE:  "取消点赞了你的文章",
	COMMENT: "评论了你的文章",
	ORDER:   "提交了订单，请尽快支付",
	BUY:     "订单付款成功！",
	ADDCART: "商品已成功加入购物车",
	SHARE:   "分享",
}

// Action Action
type Action int

//MarshalJSON MarshalJSON
func (a Action) MarshalJSON() ([]byte, error) {
	if a <= 0 {
		return []byte("\"\""), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", MessageType[int(a)])), nil
}

//UnmarshalJSON UnmarshalJSON
func (a Action) UnmarshalJSON(b []byte) error {
	a = Action(int(binary.BigEndian.Uint64(b)))
	return nil
}

//Value Value
func (a *Action) Value() (driver.Value, error) {
	return a, nil
}

// Scan Scan
func (a *Action) Scan(v interface{}) error {
	switch fmt.Sprintf("%s", reflect.TypeOf(v)) {
	case "int64":
		_val, ok := v.(int64)
		if ok {
			*a = Action(int(_val))
			return nil
		}
	case "[]uint8":
		value, ok := v.([]uint8)
		if ok {
			s := string(value)
			if s == "" {
				return nil
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf(err.Error())
			}
			*a = Action(i)
			return nil
		}
	}

	return nil
}
