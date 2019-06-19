/*
9.整数快速划分
*/
package code

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(8))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	list := make([]int, 0)
	for x > 0 {
		list = append(list, x%10)
		x /= 10
	}
	for i, j := 0, len(list)-1; i <= j; {
		if list[i] != list[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}
