package server

import (
	"fmt"
	"net"
)

func ListenAndServer() {
	fmt.Println("start")
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 10)
		fmt.Println("start read from conn")
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read error", err)
			return
		}
		fmt.Println("data=", string(buf[:n]))
	}
}
