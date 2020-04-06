package main

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start

//正常数字:123 123.123 1e4

//[]表示里面可以有也可以没有  (1)A[.B] (2)A[e|EC] (3)A[.B][e|EC] (4).B[e|EC]
//AC表示是有符号整数 B表示是无符号整数
func isNumber(s string) bool {

}

func scanInterger(s string) bool {
	if s[0] == '+' || s[0] == '-' {
		return scanUnsignedInterger(s[1:])
	}
	return scanUnsignedInterger(s)
}

func scanUnsignedInterger(s string) bool {
	p := 0
	for s != "" && s[p] >= '0' && s[p] <= '9' {
		p++
	}
}

// @lc code=end
