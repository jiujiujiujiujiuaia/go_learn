package main

import "container/list"

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start

//解法一：暴力，暴力是nk

//解法二：维护一个双端队列，这个队列单调递减。
//每个元素从尾巴进，如果这个元素大于队列得元素就删除。然后从尾巴插入
//这样这个队列从头出来的节点就是当前最大得

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	deque := list.New()
	output := make([]int, len(nums)-k+1)
	//1.预处理，找到第一个窗口的最大值，并初始化双端队列
	maxIndex := 0
	for i := 0; i < k; i++ {
		cleanDeque(k, i, nums, deque)
		deque.PushBack(i)
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	output[0] = nums[maxIndex]

	//插入都是从尾巴插入，由于清理函数从尾巴开始清理，
	//所以从尾插入一定是有序的，并且保证了头是最大
	//并且每一次遍历至少队列中还有一个节点
	for i := k; i < len(nums); i++ {
		cleanDeque(k, i, nums, deque)
		deque.PushBack(i)
		output[i-k+1] = nums[deque.Front().Value.(int)]
	}
	return output
}

//去除不是当前窗口的元素
//去除比当前元素小的元素
func cleanDeque(windowSize, current int, nums []int, deque *list.List) {

	//每次只有一种可能链表头是窗口的边界
	if deque.Len() != 0 &&
		current-windowSize == deque.Front().Value.(int) {
		deque.Remove(deque.Front())
	}

	//从尾巴开始遍历，如果比当前元素小就删除，保证从尾插入是有序的
	for deque.Len() != 0 && nums[deque.Back().Value.(int)] < nums[current] {
		deque.Remove(deque.Back())
	}
}

// @lc code=end
