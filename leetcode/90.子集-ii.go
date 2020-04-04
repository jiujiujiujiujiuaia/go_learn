package main

import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start

//值得一提的是，数组可能会出现重复的元素，如1，2，2
//因此子集的数量需要剪枝
//剪枝思路：例如1，2，2，3这个数组
//以2开头去和其他元素组成子集时，其实后面那个2再与其他元素组成子集，其实是重复了
//因此把后面的2剪枝即可。
var res [][]int

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	//(1)先排序
	sort.Ints(nums)
	helpSubsetsWithDup(0, nums, []int{})
	return res
}

func helpSubsetsWithDup(start int, nums []int, path []int) {
	tempPath := make([]int, len(path))
	copy(tempPath, path)
	res = append(res, tempPath)

	for i := start; i < len(nums); i++ {
		//剪枝
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		helpSubsetsWithDup(i+1, nums, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end
