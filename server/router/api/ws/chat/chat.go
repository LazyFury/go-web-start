package chat

import (
	"math/rand"
	"sync"

	"github.com/Treblex/go-echo-demo/server/util"

	"github.com/gorilla/websocket"
)

// BoradcastChan 广播消息
type BoradcastChan chan *Message

//Chat 聊天室
type Chat struct {
	locker    sync.RWMutex
	broadcast BoradcastChan
	Group     map[*websocket.Conn]*Gamer
	Users     map[string]*Gamer
}

// NewChat 初始化
func NewChat() *Chat {
	chat := &Chat{
		locker:    sync.RWMutex{},
		broadcast: make(BoradcastChan, 100000),
		Group:     make(map[*websocket.Conn]*Gamer, 100),
		Users:     make(map[string]*Gamer),
	}
	go chat.broadcastServer()
	return chat
}

func (c *Chat) broadcastServer() {
	for {
		msg, ok := <-c.broadcast
		if ok {
			c.SendAll(msg)
		}
	}
}

func (c *Chat) pushBroadcastMessage(msg *Message) {
	c.broadcast <- msg
}

// 处理消息
func (c *Chat) handleMessage(msg *UserSubmit, ws *websocket.Conn) {
	user := c.Group[ws]

	switch msg.Action {
	case "join":
	case "ping":
		c.SendTOUser(user, &Message{Action: pingHandle})
	default:
		c.pushBroadcastMessage(&Message{Message: msg.Msg, From: user, Action: allUser})
	}
}
func (c *Chat) getUser(id string, ws *websocket.Conn) *Gamer {
	var u *Gamer
	var ok bool
	//更新ws连接 或者新建用户
	if u, ok = c.Users[id]; ok {
		if u.Ws != ws {
			c.updateUser(u, ws)
		}
	} else {
		u = c.createUser(ws)
	}
	return u
}

// remove 删除链接
func (c *Chat) remove(ws *websocket.Conn) {
	user, ok := c.Group[ws]
	if ok {
		c.pushBroadcastMessage(&Message{Message: "退出房间", From: user, Action: systemNotify})
		c.updateGlobalConfig()
	}
	delete(c.Group, ws)
	user.remove()
}

func (c *Chat) removeByWsConn(ws *websocket.Conn) {
	c.remove(ws)
}

// updateGlobalConfig 更新用户配置
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
	c.pushBroadcastMessage(&Message{From: nil, Action: update, Global: config})
}

// Game 更新用户信息
func (c *Chat) updateUser(user *Gamer, ws *websocket.Conn) {
	c.doUpdateUser(user, ws)
	c.pushBroadcastMessage(&Message{Message: "回来了", From: user, Action: systemNotify})

}

// createUser 创建用户
func (c *Chat) createUser(ws *websocket.Conn) (user *Gamer) {
	var id = util.RandStringBytes(32)
	var name = randName[rand.Intn(len(randName))]
	user = &Gamer{
		ID:   id,
		Name: name,
		Ws:   ws,
	}
	c.doUpdateUser(user, ws)
	c.pushBroadcastMessage(&Message{Message: "加入房间", From: user, Action: systemNotify})
	return
}

// 初始化用户
func (c *Chat) doUpdateUser(user *Gamer, ws *websocket.Conn) {
	c.locker.Lock()
	u := user
	u.Ws = ws
	u.WriteList = make(chan *Message, 1000)
	c.Group[ws] = u
	c.Users[u.ID] = u
	// go u.Write()
	c.locker.Unlock()
	c.SendTOUser(u, &Message{From: u, Action: regUser})
	c.updateGlobalConfig()
}

// SendAll SendAll
func (c *Chat) SendAll(msg *Message) {
	// fmt.Printf("广播用户组：%+v\n\n", c.Group)
	c.locker.Lock()
	for _, v := range c.Group {
		// util.Logger.Printf("%+v\n", v)
		if v.Ws == nil {
			c.remove(v.Ws)
			continue
		}
		c.SendTOUser(v, msg)
	}
	c.locker.Unlock()
}

// SendTOUser SendTOUser
func (c *Chat) SendTOUser(user *Gamer, msg *Message) {
	user.WriteList <- msg
}
