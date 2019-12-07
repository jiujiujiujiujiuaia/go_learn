package client

import (
	"fmt"
	"net"
	"time"
)

func SendPackage() {
	time.Sleep(time.Second * 1)
	fmt.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("dial ok")

	time.Sleep(time.Second * 2)
	data := "hello golang"
	conn.Write([]byte(data))

	time.Sleep(time.Second * 10000)
}
