package main

import (
	"fmt"
	"net"
)

var ConnMap map[string]*net.TCPConn

func checkErr(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			//用户退出
			fmt.Println("用户推出了")
			return 0
		}
		fmt.Println("错误")
		return -1
	}
	return 1
}

func say(tcpConn *net.TCPConn) {
	for {
		//读取一个客户端发送过来的数据
		data := make([]byte, 128)
		total, err := tcpConn.Read(data)

		fmt.Println(string(data[:total]), err)

		flag := checkErr(err)
		if flag == 0 {
			//退出整个循环
			break
		}

		//广播形式，向各个客户端发送数据
		for _, conn := range ConnMap {
			if conn.RemoteAddr().String() == tcpConn.RemoteAddr().String() {
				//不向数据输入的客户端发送消息
				continue
			}
			conn.Write(data[:total])
		}
	}
}

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	/*
	   map 定义完后，还要make? (哪些数据类型定义完后，还要make?)
	   http://stackoverflow.com/questions/27267900/runtime-error-assignment-to-entry-in-nil-map
	*/
	ConnMap = make(map[string]*net.TCPConn)

	for {

		tcpConn, _ := tcpListener.AcceptTCP()
		defer tcpConn.Close()

		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		fmt.Println("连接的客服端信息:", tcpConn.RemoteAddr().String())

		go say(tcpConn)
	}
}
