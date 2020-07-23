package game

import (
	"EK-Server/util"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sync"

	"github.com/gorilla/websocket"
)

type (
	// ID 用户id
	ID string
	// Gamer Gamer
	Gamer struct {
		ID          string          `json:"id"`
		Name        string          `json:"name"`
		Ws          *websocket.Conn `json:"-"`
		MessageType int             `json:"-"`
	}
	// Group Group
	Group map[*websocket.Conn]*Gamer
	// Message 消息
	Message struct {
		From    *Gamer                 `json:"from"`
		To      string                 `json:"to"`
		Message string                 `json:"msg,omitempty"`
		Action  string                 `json:"action"`
		Global  map[string]interface{} `json:"global,omitempty"`
		IsSelf  bool                   `json:"is_self,omitempty"`
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
)

var (
	group     = Group{}
	broadcast = make(chan *Message)
	randName  = []string{"西门吹雪", "陆小凤", "章北海", "搬山", "斜岭", "摸金", "吃瓜群众", "花满楼", "崇华", "小柿子", "xixi"}
	userList  = map[string]*Gamer{}

	rlock sync.Mutex
)

// 在数组中
func inArray(arr []interface{}, item interface{}) (inArr bool) {
	index := -1
	for i, x := range arr {
		if item == x {
			index = i
		}
	}
	return index > -1
}

// Push Push推送公共消息
func Push() {
	for {
		c := <-broadcast
		util.Logger.Print(c)
		group.sendAll(c)
	}
}
func updateGlobalConfig() {
	l := len(group)
	onlineUser := []map[string]string{}

	for _, gamer := range group {
		onlineUser = append(onlineUser, map[string]string{
			"id":   gamer.ID,
			"name": gamer.Name,
		})
	}

	var config = map[string]interface{}{
		"count":      l,
		"onlineUser": onlineUser,
	}
	rlock.Lock()
	broadcast <- &Message{From: nil, To: "all", Action: update, Global: config}
	rlock.Unlock()
}

// Game
func updateUser(user *Gamer, ws *websocket.Conn) {
	group[ws] = &Gamer{
		ID:   user.ID,
		Name: user.Name,
		Ws:   ws,
	}

	rlock.Lock()
	broadcast <- &Message{Message: "回来了", From: user, To: "all", Action: systemNotify}
	rlock.Unlock()

	updateGlobalConfig()

	group[ws].send(&Message{From: group[ws], To: "", Action: regUser})
}
func createUser(ws *websocket.Conn) (user *Gamer) {
	var id = util.RandStringBytes(32)
	var name = randName[rand.Intn(len(randName))]
	user = &Gamer{
		ID:   id,
		Name: name,
		Ws:   ws,
	}
	group[ws] = user
	userList[user.ID] = user
	rlock.Lock()
	broadcast <- &Message{Message: "加入房间", From: user, To: "all", Action: systemNotify}
	rlock.Unlock()
	updateGlobalConfig()

	user.send(&Message{From: user, To: "", Action: regUser})
	return
}

// 发送消息
func (g *Gamer) send(msg *Message) (err error) {
	if g.Ws == nil {
		err = errors.New("链接断开")
		return
	}

	str, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = g.Ws.WriteMessage(websocket.TextMessage, str)

	if err != nil {
		util.Logger.Println(err)
		// safai 浏览器只能在这里检测到用户刷新页面断开
		group.remove(g.Ws)
	}

	return
}

// Group
func (g Group) hasKey(ws *websocket.Conn) (hasKey bool) {
	_, hasKey = g[ws]
	return
}
func (g *Group) sendAll(msg *Message) {
	for _, v := range *g {
		util.Logger.Printf("%+v\n", v)
		if v.Ws == nil {
			g.remove(v.Ws)
			continue
		}
		v.send(msg)
	}

}
func (g Group) remove(ws *websocket.Conn) {
	if ws == nil {
		return
	}
	user, ok := g[ws]
	if ok {

		rlock.Lock()
		broadcast <- &Message{Message: "退出房间", From: user, To: "all", Action: systemNotify}
		rlock.Unlock()
		delete(g, ws)
	}
	ws.Close()
}

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
