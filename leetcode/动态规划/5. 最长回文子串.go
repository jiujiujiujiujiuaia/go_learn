package main

import "fmt"

var cnt1, cnt2 = 0, 0

func main() {
	a := "dsadsabaaabaa"
	fmt.Println(longestPalindrome(a))
}

//动态规划的解法，使用了二维数组空间
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	longestStr := string(s[0])
	longestLen := 1
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for r := 1; r < len(s); r++ {
		for l := 0; l < r; l++ {
			if s[r] == s[l] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > longestLen {
					longestLen = r - l + 1
					longestStr = s[l : r+1]
				}
			}
		}
	}
	return longestStr
}

//计算字符串s有多少种子字符串，打印出来
//递归的办法
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

//迭代的办法
func Caculate2(s string) int {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(string(s[i : j+1]))
			cnt2++
		}
	}
	return cnt2
}
