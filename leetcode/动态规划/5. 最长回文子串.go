package main

import "fmt"

var cnt1, cnt2 = 0, 0

func longestPalindrome(s string) string {
	return ""
}

//计算字符串s有多少种子字符串，打印出来
func Caculate1(s string) int {
	i, j := 0, len(s)-1
	if i >= j {
		fmt.Println(s)
		cnt1++
		return cnt1
	}
	for ; i < len(s); i++ {
		cnt1++
		fmt.Println(s[i:])
	}
	for ; j >= 1; j-- {
		cnt1++
		fmt.Println(s[:j])
	}
	return Caculate1(s[1 : i-1])
}

func Caculate2(s string) int {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(string(s[i : j+1]))
			cnt2++
		}
	}
	return cnt2
}
