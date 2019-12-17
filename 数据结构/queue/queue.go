package main

import (
	"container/list"
	"fmt"
)

func main() {
	queueStruct()
}

func queueStruct() {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	getAllQueen(queue)
	queue.Remove(queue.Front())
	getAllQueen(queue)
}

func getAllQueen(queue *list.List) {
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
