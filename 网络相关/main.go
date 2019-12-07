package main

import "github.com/go_learn/网络相关/server"

import "github.com/go_learn/网络相关/client"

type student struct {c
	age  int
	name string
}

func main() {
	go server.ListenAndServer()
	go client.SendPackage()
	select {}
}
