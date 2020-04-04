package main

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

//解法二：迭代解决
func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	
}

// @lc code=end
