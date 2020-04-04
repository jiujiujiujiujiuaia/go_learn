package main

/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start

//镜像法
//n=2  00  10  11  01  00  10  11  01
//n=3 000 010 011 001 100 110 111 101
//g（n）的前半部分等于g(n-1),后半部分等于g(n-1)+2^n
func grayCode(n int) []int {
	ans := make([]int, 1)
	ans[0] = 0
	for i := 0; i < n; i++ {
		add := 1 << i
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]+add)
		}
	}
	return ans
}

// @lc code=end
