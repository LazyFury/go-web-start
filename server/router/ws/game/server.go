package game

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: time.Duration(10000),
	}
)

// WsServer server
func WsServer(c echo.Context) (err error) {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		//处理接收消息  保存conn对象，在需要的时候推送消息给用户
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("close:", err)
			group.remove(ws) //SetCloseHandler 有点在safari延时
			break
		}
		info := Message{}

		if err = json.Unmarshal(message, &info); err != nil {
			log.Println(err)
		}

		log.Printf("收到消息: %v", info.toString())

		readMessage(info, ws)

		ws.SetCloseHandler(func(code int, text string) error {
			log.Printf("Err %v %v", code, text)
			return nil
		})
	}
	return

}

// 读取用户信息
func readMessage(info Message, ws *websocket.Conn) {
	var user = &Gamer{}
	if !group.hasKey(ws) {
		//更新链接 或者新建用户
		fmt.Println(userList)
		if u, ok := userList[info.ID]; ok {
			updateUser(u, ws)
			user = u
		} else {
			user = createUser(ws)

		}
		return
	}
	user = group[ws]
	log.Println(user)
	user.send(fmt.Sprintf("serve收到消息：%+v", info)) //以获取到用户 其他操作
}
