package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	cur := 0
	jumpCnt := 0
	if len(nums) == 1 {
		return 0
	}
	for cur < len(nums) {
		maxBound := cur + nums[cur]
		max := cur
		i := cur
		for ; i <= Min(maxBound, len(nums)-1); i++ {
			if max < nums[i] {
				cur = i
				max = nums[i]
			}
		}
		jumpCnt++
	}
	return jumpCnt
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
