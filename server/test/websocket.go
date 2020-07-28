package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 200; i++ {
		wsURI := "wss://go.abadboy.cn/api/v1/ws"
		origin := "https://go.abadboy.cn/"

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
					return
				}
				str := fmt.Sprintf("%v", n)
				fmt.Println(str)

			}
		}()

		time.Sleep(time.Millisecond * 200)
	}

	wg.Wait()
	os.Exit(0)
}
