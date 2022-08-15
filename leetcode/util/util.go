package util

import (
	"fmt"
	"strconv"
	"strings"
)

func marshall(str string) {
	str = strings.ReplaceAll(str, `"`, `'`)
	str = strings.ReplaceAll(str, "[", "{")
	str = strings.ReplaceAll(str, "]", "}")
	fmt.Println(str)
}

func PrintArrayItem(nums []int, start, end int, isLeftBoundary bool) {
	msg := ""
	if isLeftBoundary {
		msg += "上半区:nums[%d:%d] = "
	} else {
		msg += "下半区:nums[%d:%d] = "
	}
	for i := start; i <= end; i++ {
		item := strconv.Itoa(nums[i])
		msg += item + " "
	}
	fmt.Printf(msg, start, end)
	fmt.Println()
}
