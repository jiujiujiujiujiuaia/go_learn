package main

import (
	"fmt"
	"sync"
	"time"
)

//go程序并发执行的有点诡异啊
//竟然main函数会早于协程结束
func main() {
	//启动注释就强制只有一个处理器,不启动这个程序是两个处理器执行, 那就是并行
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=1;i<100;i++{
			time.Sleep(1)
			fmt.Println("A: ",i)
		}
	}()

	go func(){
		defer wg.Done()
		for i:=1;i<100;i++{
			time.Sleep(1)
			fmt.Println("B",i)
		}
	}()
	wg.Wait()
}
