package main

import "fmt"

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
var cnt1, cnt2 = 0, 0

//待解决：1.滚动数组法（降低空间复杂度）2.中心扩展法 3.马拉车算法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	//1.预处理，初始化dp的二维数组
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	//2.解决当len(s) = 1 时的情形，初始化自底向上的底！
	longestStr := string(s[0])
	longestLength := 1

	for r := 1; r < len(s); r++ {
		for l := 0; l < r; l++ {
			//递推公式，注意这里也解决了aa 和aba的情形（r-l<2）
			if s[r] == s[l] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > longestLength {
					longestLength = r - l + 1

					//注意:string的切片是[)是这样一个区间
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

// @lc code=end
