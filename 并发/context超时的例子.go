package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	go timeOutCtx(ctx)
	<-ctx.Done()
	end := time.Now().Unix()
	fmt.Printf("start-end=[%v]s", end-start)
}

func timeOutCtx(ctx context.Context) {
	for {
		fmt.Println("dododo")
	}
}
