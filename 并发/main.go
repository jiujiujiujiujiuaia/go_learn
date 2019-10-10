package main

import (
	"fmt"
	"sync"
	"time"
)

//go程序并发执行的有点诡异啊,竟然main函数会早于协程结束(暂时没理解为什么，只知道如何让main等待)
func main() {
	golangNotify()
}

var commonInt = 1

//1.golang 中倡导用通信来共享数据，而不是通过共享数据来进行通信
//2.比如有一个监控线程在后台跑着，主线程可以发消息通知监控线程停止，在java中就只有共享内存然后轮询查看某个flag有没有改变
//golang中是用通信的方式解决
//3.select的多路复用，它的每个case都是一个通信操作，原则就是哪个case的通信操作可以执行就执行哪个，
//如果同时有多个可以执行的case，那么就随机选择一个执行
//如果没有default的话，select是会阻塞的
func golangNotify() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

//缓冲信道和非缓冲信道的区别
//非缓冲信道只负责数据的流入流出，不负责存储，因此只流入或者只流出就会导致死锁
//缓冲信道有了储存数据的能力，不超过其存储的上限就不会阻塞goroutine，到了上限，就要从信道中取出来了，一旦没有取出来，就死锁
func bufferChannelAndNoBufferChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//打开下面一行的注释就会死锁
	//ch <- 3
}

//常见的主线程阻塞，等待其他线程完成任务
func otherDoWork(ch chan int) {
	fmt.Println("其他线程开始任务")
	time.Sleep(3 * time.Second)
	fmt.Println("其他线程完成任务")
	ch <- 0
}

func fakeMain() {
	ch := make(chan int, 1)
	fmt.Println("main等待其他线程完成任务")
	go otherDoWork(ch)
	<-ch
	fmt.Println("切换回main")

}

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

//golang中的信道可以设置是读取还是写入还是默认的双向
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
