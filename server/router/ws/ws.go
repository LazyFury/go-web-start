package ws

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "ws"
	app := g.Group(baseURL)

	app.GET("", hello)
}

func hello(c echo.Context) (err error) {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			log.Fatal(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", msg)
	}

}
