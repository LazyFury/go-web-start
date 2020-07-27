package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 1000; i++ {
		wsURI := "ws://127.0.0.1:8080/api/v1/ws"
		origin := "http://127.0.0.1:8080/"

		conn, err := websocket.Dial(wsURI, "", origin)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = conn.Write([]byte(fmt.Sprintf("{action:\"join\",id:\"%s\"}", i)))
		if err != nil {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer conn.Close()
			for {
				buf := make([]byte, 512)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}
				str := fmt.Sprintf("%v", n)
				fmt.Println(str)
			}
		}()

		time.Sleep(time.Millisecond * 100)
	}

	wg.Wait()
}
