package main /*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 */

import "container/list"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {

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
			stack.PushFront(node)
			stack.PushFront(grey)
			stack.PushFront(node.Right)
			stack.PushFront(white)
			stack.PushFront(node.Left)
			stack.PushFront(white)
		} else {
			result = append(result, node.Val)
		}
	}

	return result

}

// @lc code=end
