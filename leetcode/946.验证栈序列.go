package main

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start

//边界情况得讨论：1.栈中如果出现重复元素就不能用这个算法了
//pushed=[1，2，1，3],poped=[1, 3, 2, 1],它按push(1), push(2), push(1), pop(1), push(3), pop(3), pop(2), pop(1)
//就是对的呀，但是这个算法就是false

func validateStackSequences(pushed []int, popped []int) bool {
	
	stack := list.New()
	var j int
	for i := 0; i < len(pushed); i++ {
		stack.PushFront(pushed[i])
		for stack.Front().Value.(int) == popped[j] {
			j++
			stack.Remove(stack.Front())
		}
	}
	return j == len(popped)
}

// @lc code=end
