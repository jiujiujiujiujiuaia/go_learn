package main

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (37.24%)
 * Likes:    276
 * Dislikes: 0
 * Total Accepted:    81K
 * Total Submissions: 217.2K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 *
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 *
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 *
 * 示例 1:
 *
 * 输入: 4
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842...,
 * 由于返回类型是整数，小数部分将被舍去。
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 0, x/2
	for l <= r {
		mid := l + (r-l)/2
		number := mid * mid
		if number == x {
			return mid
		} else if number > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	//简单分析下这里为什么返回r，最终的结果肯定是l=r=mid
	//然后如果target>mid，那么就是l=mid+1,但是因为是开平方，一定是小于l+1的平方应该返回r
	//如果target<mid那么就是r=mid-1，那么也是返回r
	//以上分析的原因是任何一个数的平方一定大于整数i,小于等于i+1
	// if l > r {
	// 	return r
	// } else {
	// 	return r
	// }
	return r
}

// @lc code=end
