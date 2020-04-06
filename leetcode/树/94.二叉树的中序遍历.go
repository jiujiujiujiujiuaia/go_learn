package main

import "container/list"

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var result []int

//解法一：递归解决
//递归的问题在于，如果树很深，导致栈空间会不停的堆积，因为只有函数结束
//才会释放栈空间
func inorderTraversal1(root *TreeNode) []int {
	result = make([]int, 0)
	solve(root)
	return result
}

func solve(node *TreeNode) {
	if node == nil {
		return
	}
	solve(node.Left)
	result = append(result, node.Val)
	solve(node.Right)
}

//解法二：迭代解决，都是使用嵌套循环解决的，有点烧脑
//这里参考一下评论中颜色标记法，把前中后序三种使用迭代思路串起来了。

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	white, grey := 1, 0
	stack := list.New()
	stack.PushFront(root)
	stack.PushFront(white)
	for stack.Len() != 0 {
		color, node := stack.Front().Value.(int), stack.Front().Next().Value.(*TreeNode)
		stack.Remove(stack.Front())
		stack.Remove(stack.Front())
	if node == nil {
		continue
	}
	
	if color == white {
		stack.PushFront(node.Right)
		stack.PushFront(white)
		stack.PushFront(node)
		stack.PushFront(grey)
		stack.PushFront(node.Left)
		stack.PushFront(white)
	} else {
		result = append(result, node.Val)
	}
	}
	
	return result
}

// @lc code=end
