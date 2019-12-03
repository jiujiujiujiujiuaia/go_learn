package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2}))
}

//有点意思，典型的空间换时间，前缀和之后又学到了新的东西，前缀积，后缀积。
func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	post := make([]int, len(nums))
	post[len(post)-1] = 1
	for i := len(post) - 2; i > -1; i-- {
		post[i] = post[i+1] * nums[i+1]
	}

	output := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		output[i] = pre[i] * post[i]
	}
	return output
}
