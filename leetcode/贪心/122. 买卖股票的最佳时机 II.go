package main

import "fmt"

func main() {
	fmt.Println(maxProfit2([]int{6, 1, 3, 2, 4, 7}))
}

//这个方法的思路就是计算出所有的波峰波谷
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	j := 0
	//神来之笔，可能存在最后一个波谷和波峰没计算完循环就结束了
	//因此加一个波谷，然这个波折计算出来，免得又额外加if else循环判断
	prices = append(prices, 0)
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			res += prices[i-1] - prices[j]
			j = i
		}
	}
	return res
}

//但是实际上这道题考察的是贪心，因为一支股票可以多次购买，当我把一支股票以2元的价格卖出的时候，再以2元的价格买入，等待后面更大再卖出
//实质上与上面那个解法是一样的：波峰（a3）-波谷（a1）= (a2-a1) + (a3-a2)
func greedMaxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
