package main

import (
	"context"
	"fmt"
	"time"
)

//虽然select+chan的方法也可以，但是如果携程链条关系复杂，比如说一个req进来，起了一个
//携程，然后携程又起了携程，这种关系使用ctx就比较方便了
//因为子ctx会被父ctx
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}
