package game

import (
	"EK-Server/util"
	"fmt"
	"math/rand"

	"github.com/gorilla/websocket"
)

type (
	// ID 用户id
	ID string
	// Gamer Gamer
	Gamer struct {
		ID          string
		Name        string `json:"name"`
		Ws          *websocket.Conn
		MessageType int
	}
	// Group Group
	Group map[*websocket.Conn]*Gamer
	// Message 消息
	Message struct {
		ID      string `json:"id"`
		Message string `json:"msg"`
		Action  string `json:"action"`
	}
	// Cast 广播
	Cast struct {
		Msg string
		UID string
	}
	// Room 游戏房间
	Room struct {
		ID       string
		redTeam  []string //红队
		blueTeam []string //蓝队
		MaxCount int      //每队最大人数
	}
)

var (
	group     = Group{}
	broadcast = make(chan Cast)
	randName  = []string{"西门吹雪", "陆小凤", "章北海", "搬山", "斜岭", "摸金", "吃瓜群众", "花满楼", "崇华", "小柿子", "xixi"}
	userList  = map[string]*Gamer{}
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
		group.sendAll(c.Msg)
	}
}

// Game
func updateUser(user *Gamer, ws *websocket.Conn) {
	user.Ws = ws
	group[ws] = user
	broadcast <- Cast{Msg: user.Name + " 重新回到游戏", UID: user.ID}
	user.send(user)
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
	broadcast <- Cast{Msg: user.Name + " 加入游戏", UID: user.ID}
	user.send(user)
	return
}

// 发送消息
func (g *Gamer) send(data interface{}) (msg string, err error) {

	if str, ok := data.(string); ok {
		err = g.Ws.WriteMessage(1, []byte(str))
	} else {
		err = g.Ws.WriteJSON(data)
	}

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
func (g *Group) sendAll(msg interface{}) {
	var player []string
	for _, v := range *g {
		player = append(player, v.Name)
	}
	for _, v := range *g {
		_, _ = v.send(map[string]interface{}{
			"msg":        msg,
			"count":      len(group),
			"OnLineUser": player,
		})
	}

}
func (g Group) remove(ws *websocket.Conn) {
	user := g[ws]
	broadcast <- Cast{Msg: "用户 " + user.Name + " 退出房间"}
	delete(g, ws)
	ws.Close()
}

// Message
func (m *Message) toString() (str string) {
	var tmp = `ws Message:{
					  "id":%s,
					  "message":%s,
					  "action":%s
					}
				`
	return fmt.Sprintf(tmp, m.ID, m.Message, m.Action)
}
