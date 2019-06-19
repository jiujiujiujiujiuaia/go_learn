package code

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(a))
}

func canJump(nums []int) bool {
	last := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}

	if last == 0 {
		return true
	} else {
		return false
	}
}
