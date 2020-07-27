package chat

import (
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
		WriteList   chan *Message   `json:"-"`
		isDone      chan bool       `json:"-"`
	}
)

func (g *Gamer) Write() {
	fmt.Printf("初始化用户监听消息\n")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	defer g.remove()

	for {
		if msg, ok := <-g.WriteList; ok {
			// fmt.Printf("监听发消息:%v \n", msg)
			err := g.send(msg)
			if err != nil {
				fmt.Println(err)
				break
			}
		} else {
			break
		}

	}
}

func (g *Gamer) remove() {
	g.Ws.Close()
	close(g.WriteList)
	g.isDone <- true
}

// 发送消息
func (g *Gamer) send(msg *Message) (err error) {
	if g.Ws == nil {
		err = errors.New("链接断开")
		return
	}

	// fmt.Printf("发送消息内容%v\n\n", struct {
	// 	Ws *websocket.Conn
	// 	*Message
	// }{
	// 	Ws:      g.Ws,
	// 	Message: msg,
	// })

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
		err = errors.New("writeErr")
	}
	return
}
