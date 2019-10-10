package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	//先是建立了一个新的上下文
	ctx, cancel := context.WithCancel(context.Background())
	//建立了一个子的上下文
	valueCtx := context.WithValue(ctx, key, "【监控1】")

	go watch(valueCtx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	fmt.Println("进入携程前")
	select {
	case <-ctx.Done():
		//取出值
		fmt.Println(ctx.Value(key), "监控退出，停止了...")
		return
		//default:
		//	//取出值
		//	fmt.Println(ctx.Value(key), "goroutine监控中...")
		//	time.Sleep(2 * time.Second)
	}
	fmt.Println("进入携程后")
}
