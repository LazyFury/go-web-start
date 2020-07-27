package chat

import (
	"fmt"
)

type (
	// ID 用户id
	ID string

	// Message 消息
	Message struct {
		From    *Gamer                 `json:"from,omitempty"`
		Message string                 `json:"msg,omitempty"`
		Action  string                 `json:"action"`
		Global  map[string]interface{} `json:"global,omitempty"`
	}
	// Room 游戏房间
	Room struct {
		ID       string
		redTeam  []string //红队
		blueTeam []string //蓝队
		MaxCount int      //每队最大人数
	}

	// UserSubmit 用户提交的数据
	UserSubmit struct {
		ID     string
		Msg    string
		Action string
	}
)

const (
	systemNotify = "SystemNotify"
	toUser       = "MsgToUser"
	regUser      = "regUser"
	allUser      = "allUser"
	update       = "update"
	pingHandle   = "pingHandle"
)

var (
	// broadcast = make(chan *Message)
	randName = []string{"西门吹雪", "陆小凤", "章北海", "搬山", "斜岭", "摸金", "吃瓜群众", "花满楼", "崇华", "小柿子", "xixi"}
)

// Message
func (u *UserSubmit) toString() (str string) {
	var tmp = `ws Message:{
					  "id":%s,
					  "message":%s,
					  "action":%s
					}
				`
	return fmt.Sprintf(tmp, u.ID, u.Msg, u.Action)
}
