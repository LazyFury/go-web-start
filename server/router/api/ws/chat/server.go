package chat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lazyfury/go-web-start/server/utils"
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
func WsServer(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		utils.Error(err)
	}
	defer ws.Close()

	ws.SetCloseHandler(func(code int, text string) error {
		fmt.Printf("SetCloseHandler Err %d %s ", code, text)
		// chat.removeByWsConn(ws)
		return nil
	})

	id := c.Query("token")
	user := chat.getUser(id, ws)
	// fmt.Println(user)

	go user.Write()
	go received(chat, ws)

	<-user.isDone
}

func received(chat *Chat, ws *websocket.Conn) {
	defer func() {
		chat.removeByWsConn(ws) //SetCloseHandler 在safari无法触发，可能浏览器做了优化，同样的在地址蓝输入链接的时候ws链接就已经建立成功了，不像chrome可以明确触发进入和离开的事件(safari升级解决了这个问题)
	}()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Print("close:", err)
			break
		}
		var info UserSubmit

		if err = json.Unmarshal(message, &info); err != nil {
			continue
		}

		if info.Action != "ping" {
			fmt.Printf("收到消息: %v \n", info.toString())
		}

		chat.handleMessage(&info, ws)
	}
}
