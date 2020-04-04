package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	helpSubSet(0, nums, []int{})
	return res
}

func helpSubSet(start int, nums []int, path []int) {
	tempPath := make([]int,len(path))
	copy(tempPath,path)
	res = append(res, tempPath)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		helpSubSet(i+1, nums, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end
