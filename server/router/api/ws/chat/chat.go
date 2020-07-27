package chat

import (
	"EK-Server/util"
	"fmt"
	"math/rand"
	"sync"

	"github.com/gorilla/websocket"
)

// BoradcastChan 广播消息
type BoradcastChan chan *Message
type users map[string]*Gamer

//Chat 聊天室
type Chat struct {
	locker    sync.RWMutex
	broadcast BoradcastChan
	Group     Group
	Users     users
}

// NewChat 初始化
func NewChat() *Chat {
	chat := &Chat{
		locker:    sync.RWMutex{},
		broadcast: make(BoradcastChan, 1000),
		Group:     make(Group, 100),
		Users:     make(users),
	}
	go chat.broadcastServer()
	return chat
}

func (c *Chat) broadcastServer() {
	for {
		msg := <-c.broadcast
		fmt.Printf("发送广播消息：%+v\n", msg)
		c.SendAll(msg)
	}
}

func (c *Chat) pushBroadcastMessage(msg *Message) {
	c.locker.Lock()
	c.broadcast <- msg
	c.locker.Unlock()
}

// 处理消息
func (c *Chat) handleMessage(msg *UserSubmit, ws *websocket.Conn) {

	user := c.getUser(msg.ID, ws)

	switch msg.Action {
	case "join":
	case "ping":
	default:
		c.pushBroadcastMessage(&Message{Message: msg.Msg, From: user, Action: allUser})
	}
}
func (c *Chat) getUser(id string, ws *websocket.Conn) *Gamer {
	var u Gamer
	//更新ws连接 或者新建用户
	if u, ok := c.Users[id]; ok {
		if u.Ws != ws {
			c.updateUser(u, ws)
		} else {
			u = c.Group[ws]
		}
	} else {
		u = c.createUser(ws)
	}
	return &u
}

func (c *Chat) remove(ws *websocket.Conn) {
	user, ok := c.Group[ws]
	delete(c.Group, ws)
	if ok {
		c.pushBroadcastMessage(&Message{Message: "退出房间", From: user, Action: systemNotify})
	}

	defer user.Ws.Close()

}

func (c *Chat) removeByWsConn(ws *websocket.Conn) {
	c.remove(ws)
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
	c.pushBroadcastMessage(&Message{From: nil, Action: update, Global: config})
}

// Game
func (c *Chat) updateUser(user *Gamer, ws *websocket.Conn) {

	c.locker.Lock()
	u := &Gamer{
		ID:   user.ID,
		Name: user.Name,
		Ws:   ws,
	}
	c.Group[ws] = u
	c.Users[u.ID] = u
	c.locker.Unlock()

	c.pushBroadcastMessage(&Message{Message: "回来了", From: u, Action: systemNotify})
	c.SendTOUser(u, &Message{From: u, Action: regUser})
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
	c.Group[ws] = user
	c.Users[id] = user
	c.pushBroadcastMessage(&Message{Message: "加入房间", From: user, Action: systemNotify})
	c.updateGlobalConfig()
	c.SendTOUser(user, &Message{From: user, Action: regUser})
	return
}

// SendAll SendAll
func (c *Chat) SendAll(msg *Message) {
	fmt.Printf("广播用户组：%+v\n\n", c.Group)
	for _, v := range c.Group {
		// util.Logger.Printf("%+v\n", v)
		if v.Ws == nil {
			c.remove(v.Ws)
			continue
		}
		c.SendTOUser(v, msg)
	}
}

// SendTOUser SendTOUser
func (c *Chat) SendTOUser(user *Gamer, msg *Message) {
	c.locker.Lock()
	err := user.send(msg)
	if err != nil && err.Error() == "writeErr" {
		fmt.Println(err)
		c.remove(user.Ws)
	}
	c.locker.Unlock()
}
