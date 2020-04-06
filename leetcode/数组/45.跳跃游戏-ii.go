package main

import "fmt"

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	var cnt int = 0
	for i := 0; i < len(nums)-1; {
		furtherStep := nums[i]
		canJumpFurther := 0
		nextPoint := 0

		//1.如果当前起始点可以直接到达，则返回
		if i+furtherStep >= len(nums)-1 {
			cnt++
			return cnt
		}

		//2.当前起始点无法到达，则寻找范围内能抵达最远的点（nums[i+j]+j）
		for j := 1; j <= furtherStep; j++ {
			if i+j < len(nums) && nums[i+j]+j >= canJumpFurther {
				fmt.Println(i, j)
				canJumpFurther = nums[i+j] + j
				nextPoint = i + j
			}
		}
		i = nextPoint
		cnt++
	}
	return cnt
}

// @lc code=end
