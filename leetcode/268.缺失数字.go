package main

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
//这个题目特有意思，综合了很多内容
//思路一：排序
//由于是0-n个数，那么索引0指向的肯定是0，如果不是0，那说明缺失0

//思路二：额外空间
//可以使用hashset两次遍历

//思路三：位运算
//异或：相同为0，相异为1，那么索引^val，缺失的那个值就是最终的结果

//思路四：等差数列
//我们可以求出来0-n的所有值，然后减去数组中的元素，即可以找出那个值
//关键点来了！如果n很大，求等差数列导致类型溢出了怎么办？
//讨巧的方法是用索引-val可以绕开这个问题

//面试流程：位运算>等差数列>排序>额外空间
//为什么位运算比乘法除法快？因为乘法除法最终都是使用位运算解决的，还不如直接位运算

func missingNumber(nums []int) int {
	nums = append(nums, 0)
	res := 0
	for i := 0; i < len(nums); i++ {
		res += (i - nums[i])
	}
	return res
}

// @lc code=end
