package chat

import (
	"EK-Server/util"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
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
func WsServer(c echo.Context) (err error) {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		util.Logger.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	ws.SetCloseHandler(func(code int, text string) error {
		util.Logger.Printf("Err %v %v", code, text)
		return nil
	})

	for {
		//处理接收消息  保存conn对象，在需要的时候推送消息给用户
		_, message, err := ws.ReadMessage()
		if err != nil {
			util.Logger.Println("close:", err)
			chat.removeByWsConn(ws) //SetCloseHandler 在safari无法触发，可能浏览器做了优化，同样的在地址蓝输入链接的时候ws链接就已经建立成功了，不像chrome可以明确触发进入和离开的事件
			break
		}
		info := UserSubmit{}

		if err = json.Unmarshal(message, &info); err != nil {
			util.Logger.Println(message)
			util.Logger.Println(info)
			continue
		}

		util.Logger.Printf("收到消息: %v", info.toString())

		// readMessage(info, ws)
		chat.handleMessage(&info, ws)
	}

	return

}