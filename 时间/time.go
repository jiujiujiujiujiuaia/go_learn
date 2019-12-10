package main

import (
	"fmt"
	"time"
)

func main() {
	startSecond := time.Now().Unix()
	startNans := time.Now().UnixNano()
	startMs := time.Now().UnixNano() / 1e6
	time.Sleep(1 * time.Second)
	endMs := time.Now().UnixNano() / 1e6
	endNans := time.Now().UnixNano()
	endSecond := time.Now().Unix()

	fmt.Printf("时间戳（秒）：%v;\n", endSecond-startSecond)
	fmt.Printf("时间戳（纳秒）：%v;\n", endNans-startNans)
	fmt.Printf("时间戳（毫秒）：%v;\n", endMs-startMs)
}
