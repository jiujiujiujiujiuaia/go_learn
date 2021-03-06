package main

import "math"

import "container/list"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start

//注意哦，这个n^2时间复杂度的解法是由n^3优化而来的
//最暴力：枚举所有的两两矩形块，确认好后再在这两块间找到最小矩形块
//解法一：o(n^2) 每次遍历往前增加一个边界，然后以当前范围内最小的矩形为准(木桶效应)
func largestRectangleArea1(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	var max int = heights[0]
	for i := 1; i < len(heights); i++ {
		var minRec int = math.MaxInt32
		//注意哦，这里是从后往前遍历的，
		for j := i; j >= 0; j-- {
			minRec = min(minRec, heights[j])
			curArea := minRec * (i - j + 1)
			if curArea > max {
				max = curArea
			}
		}
	}
	return max
}

//解法二：分治+递归，在一个最小的矩形左边为一组，寻找这一组的最小矩形
//右边为一组，寻找最小的矩形，发现左右两边都是同样的子问题，这样就可以自顶向下的解决
//时间复杂度为o(nlogn)

//解法三：单调栈，比较难想，算法思想是：维护一个递增的栈，如果依然递增，则push，
//如果遇到第一个低峰，则pop并计算最大面积，直到保持递增
func largestRectangleArea(heights []int) int {
	stack := list.New()
	//注意！！这里这个-1非常的关键：例如3，4，5，2，1的高度，当计算完3，4，5，2时
	//栈中依次退出了3，4，5 那么当计算到高度1时，你无法计算以高度11为最小的长度和前面矩形组成的面积
	//这个时候-1就帮助你了。
	stack.PushBack(-1)
	var maxArea int = 0
	for i := 0; i < len(heights); i++ {
		for stack.Back().Value.(int) != -1 && heights[stack.Back().Value.(int)] > heights[i] {
			//前面是一个递增的序列，当出现第一个波谷时，那么可以开始计算
			topIndex := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			//注意！！这个操作是先把顶剔除栈，然后在用top求
			//这里求的是current之前的矩形面积
			maxArea = max(maxArea, heights[topIndex]*(i-stack.Back().Value.(int)-1))

		}
		//current最后入栈
		stack.PushBack(i)
	}

	//2.要么栈中没有了数据，要么只剩下递增的序列
	for stack.Back().Value.(int) != -1 {
		topIndex := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		maxArea = max(maxArea, heights[topIndex]*(len(heights)-stack.Back().Value.(int)-1))

	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
