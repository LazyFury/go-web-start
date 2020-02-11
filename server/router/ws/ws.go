package ws

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

// Init Init
func Init(g *echo.Group) {
	baseURL := "ws"
	app := g.Group(baseURL)

	app.GET("", webscoketHello)
}

func webscoketHello(c echo.Context) (err error) {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			err = websocket.Message.Send(ws, "hello")
			if err != nil {
				log.Fatal(err)
			}

			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s \n", msg)
		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
