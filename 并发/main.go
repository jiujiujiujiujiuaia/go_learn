package main

import (
	"fmt"
	"sync"
	"time"
)

//go程序并发执行的有点诡异啊
//竟然main函数会早于协程结束
//golang 中倡导用通信来共享数据，而不是通过共享数据来进行通信
func main() {
	syncLock()
}

var commonInt = 1

//互斥锁的使用
func syncLock() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10000; i++ {
			lock.Lock()
			commonInt += 1
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			lock.Lock()
			commonInt += 1
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(commonInt)
}

func chanTypeUsage() {
	var ch1 chan int   //双向
	var ch2 chan<- int //单向写入
	var ch3 <-chan int //单项读取
	ch1 = ch1
	ch2 = ch2
	ch3 = ch3
}

func randomSend1Or0() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}

		i := <-ch
		fmt.Println("received : ", i)
		time.Sleep(1e8)
	}
}

func productAndConsumer() {
	ch := make(chan int, 10)

	//produce
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	//consumer
	go func() {
		for val := range ch {
			fmt.Printf("consume : %v\n", val)
		}
	}()
	time.Sleep(1e9)
}

func channelUsage() {
	theMine := []string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	//寻找金矿
	go func(mine []string) {
		for _, v := range mine {
			fmt.Println("send " + v)
			oreChan <- v

		}
	}(theMine)

	//接收金矿，可以看到这里是阻塞读取
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * 5)
			found := <-oreChan
			fmt.Println("receive " + found)
		}
	}()

	<-time.After(time.Second * 10)
}

func waitGroupUsage() {
	//启动注释就强制只有一个处理器,不启动这个程序是两个处理器执行, 那就是并行
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			time.Sleep(1)
			fmt.Println("A: ", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			time.Sleep(1)
			fmt.Println("B", i)
		}
	}()
	wg.Wait()
}
