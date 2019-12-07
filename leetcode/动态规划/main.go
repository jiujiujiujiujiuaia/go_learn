package main

import (
	"fmt"
)

//以完成 5 53 120 121 122 221
//未完成 0  70 152 300
func main() {
	param := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(LengthOfLIS(param))

	param = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaxSubArray(param))

	param = []int{-2, 0, -1}
	fmt.Println(MaxProduct(param))

	a := "abcd"
	fmt.Println(Caculate2(a))
	fmt.Println("----")
	fmt.Println(Caculate1(a))
}
