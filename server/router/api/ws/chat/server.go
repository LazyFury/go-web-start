package chat

import (
	"github.com/treblex/go-echo-demo/server/util"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: time.Duration(10000),
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	chat = NewChat()
)

// WsServer server
func WsServer(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		util.Logger.Print("upgrade:", err)
		return err
	}
	defer util.Recover()
	defer ws.Close()

	ws.SetCloseHandler(func(code int, text string) error {
		fmt.Printf("\nSetCloseHandler Err %v %v \n", code, text)
		// chat.removeByWsConn(ws)
		return nil
	})

	id := c.QueryParam("token")
	user := chat.getUser(id, ws)
	// fmt.Println(user)

	go user.Write()
	go received(chat, ws)

	<-user.isDone
	return nil
}

func received(chat *Chat, ws *websocket.Conn) {
	defer util.Recover()

	defer func() {
		chat.removeByWsConn(ws) //SetCloseHandler 在safari无法触发，可能浏览器做了优化，同样的在地址蓝输入链接的时候ws链接就已经建立成功了，不像chrome可以明确触发进入和离开的事件(safari升级解决了这个问题)
	}()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("close:", err)
			break
		}
		var info UserSubmit

		if err = json.Unmarshal(message, &info); err != nil {
			continue
		}

		if info.Action != "ping" {
			util.Logger.Printf("收到消息: %v \n", info.toString())
		}

		chat.handleMessage(&info, ws)
	}
}
