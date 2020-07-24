package chat

import (
	"EK-Server/util"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
)

type (
	// Gamer Gamer
	Gamer struct {
		ID          string          `json:"id"`
		Name        string          `json:"name"`
		Ws          *websocket.Conn `json:"-"`
		MessageType int             `json:"-"`
	}
)

// 发送消息
func (g *Gamer) send(msg *Message) (err error) {
	if g.Ws == nil {
		err = errors.New("链接断开")
		return
	}
	fmt.Printf("Msg %+v\n", msg)
	fmt.Printf("user %+v\n", g)

	result := &struct {
		*Message
		IsSelf bool `json:"is_self"`
	}{
		Message: msg,
	}

	// 是否是我自己发的消息
	if msg != nil && msg.From != nil {
		if msg.From.ID == g.ID {
			result.IsSelf = true
		}
	}

	str, err := json.Marshal(result)
	if err != nil {
		return
	}

	err = g.Ws.WriteMessage(websocket.TextMessage, str)

	if err != nil {
		util.Logger.Println(err)
		err = errors.New("writeErr")
	}

	return
}
