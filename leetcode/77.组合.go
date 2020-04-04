package main

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
var res [][]int

func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	helpCombine(n, k, 0, 1, []int{})
	return res
}

func helpCombine(n, k, cnt, start int, path []int) {
	if cnt == k {
		tempPath := make([]int, len(path))
		copy(tempPath, path)
		res = append(res, tempPath)
		return
	}

	for i := start; i <= n; i++ {
		path = append(path, i)
		helpCombine(n, k, cnt+1, i+1, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end
