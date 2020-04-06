package main

import "sort"

import "math"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
var permuteResult [][]int

func permuteUnique(nums []int) [][]int {
	permuteResult = make([][]int, 0)
	sort.Ints(nums)
	helpPermuteUnique(nums, []int{})
	return permuteResult
}

func helpPermuteUnique(nums []int, path []int) {
	if len(nums) == len(path) {
		res := make([]int, 0)
		for _, val := range path {
			res = append(res, nums[val])
		}
		permuteResult = append(permuteResult, res)
		return
	}

	last := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if check(i, path) {
			continue
		} else if i > 0 && last == nums[i] {
			continue
		} else {
			last = nums[i]
			path = append(path, i)
			helpPermuteUnique(nums, path)
			path = path[:len(path)-1]
		}
	}
}

func check(index int, path []int) bool {
	for _, val := range path {
		if val == index {
			return true
		}
	}
	return false
}

// @lc code=end
