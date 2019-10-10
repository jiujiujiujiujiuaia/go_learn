package main

import (
	"fmt"
	_ "go_learn/init顺序/dic/dic2"
)

func init() {
	fmt.Println("main int")
}

func main() {

}
