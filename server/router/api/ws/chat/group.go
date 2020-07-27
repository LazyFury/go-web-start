package chat

import "github.com/gorilla/websocket"

// Group Group
type Group map[*websocket.Conn]*Gamer

// Group
func (g Group) hasKey(ws *websocket.Conn) (hasKey bool) {
	_, hasKey = g[ws]
	return
}
