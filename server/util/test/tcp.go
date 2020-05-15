package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:52000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go server(conn)

	}
}

func server(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("welcome 请输入:"))
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%v", string(buf[0:n]))
		handle()
		buf = buf[0 : n-1]
		buf = append(buf, []byte("?\n请输入:")...)
		conn.Write(buf)

	}
}

func handle() {}
