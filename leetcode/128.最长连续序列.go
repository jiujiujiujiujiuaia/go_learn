package main

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start

//没有时间拉

//解法一：暴力
//枚举以数组中每个数为开头，然后去寻找比他大一的数
//o（n^3）

//解法二：排序
//先将数组排序，然后再遍历数组一遍，如果后一个数
//大于前一个数加1，则继续判断，否则计算一下是否超过最大长度

//解法三：使用额外空间去高效寻找
//先将数字都存在hash表中，然后一个数子比他小1的数字不存在
//那么这个数就是起始的头，然后去寻找它的下一个值

func longestConsecutive(nums []int) int {

}

// @lc code=end
