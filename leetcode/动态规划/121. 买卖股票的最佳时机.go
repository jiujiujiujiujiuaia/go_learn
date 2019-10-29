package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 2, 6, 3, 8, 1}))
}

//只能完成一笔交易（即一次买入和一次卖出）
func maxProfit(prices []int) int {
	buy, sell := int(^uint(0)>>1), 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		}
		if prices[i] > buy && sell < prices[i]-buy {
			sell = prices[i] - buy
		}
	}
	return sell
}
