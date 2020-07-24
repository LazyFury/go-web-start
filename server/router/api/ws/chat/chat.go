package chat

import (
	"EK-Server/util"
	"fmt"
	"math/rand"
	"sync"

	"github.com/gorilla/websocket"
)

//Chat 聊天室
type Chat struct {
	locker    sync.RWMutex
	broadcast BoradcastChan
	Group     Group
	Users     map[string]*Gamer
}

// NewChat 初始化
func NewChat() *Chat {
	chat := &Chat{
		locker:    sync.RWMutex{},
		broadcast: make(BoradcastChan),
		Group:     make(Group),
	}
	go chat.broadcastServer()
	return chat
}
func (c *Chat) broadcastServer() {
	for {
		msg := <-c.broadcast
		fmt.Printf("%+v\n", msg)
		c.SendAll(msg)
	}
}

// 处理消息
func (c *Chat) handleMessage(msg *UserSubmit, ws *websocket.Conn) {
	checkUser := func() {
		//更新ws连接 或者新建用户
		if u, ok := c.Group[msg.ID]; ok {
			c.updateUser(u, ws)
		} else {
			c.createUser(ws)
		}
		return
	}
	user := c.Group[msg.ID]

	switch msg.Action {
	case "join":
		checkUser()
	case "ping":
	default:
		c.broadcast.Write(&Message{Message: msg.Msg, From: user, Action: allUser})
	}
}

func (c *Chat) remove(id string) {
	user, ok := c.Group.remove(id)
	if ok {
		c.broadcast.Write(&Message{Message: "退出房间", From: user, Action: systemNotify})
	}
	user.Ws.Close()
}

func (c *Chat) updateGlobalConfig() {
	l := len(c.Group)
	onlineUser := []map[string]string{}

	for _, gamer := range c.Group {
		onlineUser = append(onlineUser, map[string]string{
			"id":   gamer.ID,
			"name": gamer.Name,
		})
	}

	var config = map[string]interface{}{
		"count":      l,
		"onlineUser": onlineUser,
	}
	c.broadcast.Write(&Message{From: nil, Action: update, Global: config})
}

// Game
func (c *Chat) updateUser(user *Gamer, ws *websocket.Conn) {
	u := c.Group[user.ID]
	u.Ws = ws
	c.broadcast.Write(&Message{Message: "回来了", From: u, Action: systemNotify})

	u.send(&Message{From: u, Action: regUser})
	c.updateGlobalConfig()

}

func (c *Chat) createUser(ws *websocket.Conn) (user *Gamer) {
	var id = util.RandStringBytes(32)
	var name = randName[rand.Intn(len(randName))]
	user = &Gamer{
		ID:   id,
		Name: name,
		Ws:   ws,
	}
	c.Group[id] = user
	c.broadcast.Write(&Message{Message: "加入房间", From: user, Action: systemNotify})
	c.updateGlobalConfig()

	user.send(&Message{From: user, Action: regUser})
	return
}

// SendAll SendAll
func (c *Chat) SendAll(msg *Message) {
	for _, v := range c.Group {
		// util.Logger.Printf("%+v\n", v)
		if v.Ws == nil {
			c.remove(v.ID)
			continue
		}
		c.SendTOUser(v, msg)
	}
}

// SendTOUser SendTOUser
func (c *Chat) SendTOUser(user *Gamer, msg *Message) {
	err := user.send(msg)
	if err != nil && err.Error() == "writeErr" {
		fmt.Println(err)
		c.remove(user.ID)
	}
}
