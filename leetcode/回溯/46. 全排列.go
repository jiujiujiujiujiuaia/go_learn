package main

import "fmt"

var result [][]int

func permute(nums []int) [][]int {
	result = [][]int{}
	return recursive1(nums, []int{})
}

func recursive1(nums []int, oneRes []int) [][]int {
	if len(nums) == 0 {
		result = append(result, oneRes)
		return result
	}

	for i := 0; i < len(nums); i++ {
		newOneRes := append(oneRes, nums[i])
		tempNum := make([]int, len(nums))
		copy(tempNum, nums)
		recursive1(append(tempNum[:i], tempNum[i+1:]...), newOneRes)
	}

	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
}
