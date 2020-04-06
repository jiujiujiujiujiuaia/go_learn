package main

import "sort"

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start

//多说一句，类似这种的数组题目，时间要求n并且空间要求1，那么最优的肯定是位运算

//解法一：排序，然后遍历整个数组，从最小值1开始，如果数组存在1，那么就去找2
//直到遍历完数组，就知道哪个值没有找到
//nlogn
func firstMissingPositive1(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	sort.Ints(nums)
	MinNotZeroNumber := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == MinNotZeroNumber {
			MinNotZeroNumber++
		}
	}
	return MinNotZeroNumber
}

//解法二：hash表
//如果说数组长度为n那么最小值肯定是1-n之间
func firstMissingPositive2(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	numberMap := make(map[int]int)
	for _, num := range nums {
		numberMap[num] = num
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := numberMap[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}

//解法三：最优的解法
//如果说数组长度为n那么最小值肯定是1-n之间，或者说n+1
func firstMissingPositive(nums []int) int {

	for i := 0; i < len(nums); i++ {
		//首先nums[i]满足在1-len范围内，其次要保证他想要放到的位置上的那个值不等于本身
		//要不然会有死循环的
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			index := nums[i] - 1
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}

// @lc code=end
