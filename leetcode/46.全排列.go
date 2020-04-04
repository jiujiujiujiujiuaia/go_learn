package main

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
var permuteResult [][]int

func permute(nums []int) [][]int {
	permuteResult = make([][]int, 0)
	helpPermute(nums, []int{})
	return permuteResult
}

func helpPermute(nums []int, path []int) {
	if len(nums) == len(path) {
		res := make([]int, len(path))
		copy(res, path)
		permuteResult = append(permuteResult, res)
		return
	}

	for i := 0; i < len(nums); i++ {
		if check(nums[i], path) {
			continue
		} else {
			path = append(path, nums[i])
			helpPermute(nums, path)
			path = path[:len(path)-1]
		}
	}
}

func check(num int, path []int) bool {
	for _, val := range path {
		if val == num {
			return true
		}
	}
	return false
}

// @lc code=end
