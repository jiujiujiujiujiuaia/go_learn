package main /*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (72.94%)
 * Likes:    476
 * Dislikes: 0
 * Total Accepted:    63.9K
 * Total Submissions: 87.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

// @lc code=start
var permuteResult [][]int

func permute(nums []int) [][]int {
	permuteResult = make([][]int, 0)
	recurisive(nums, []int{})
	return permuteResult
}

func recurisive(nums []int, value []int) {
	if len(nums) == 0 {
		permuteResult = append(permuteResult, value)
	}

	var newValue []int
	for i := 0; i < len(nums); i++ {
		tempNum := make([]int, len(nums))
		copy(tempNum, nums)
		newValue = append(value, nums[i])
		recurisive(append(tempNum[:i], tempNum[i+1:]...), newValue)
	}
}

// @lc code=end
