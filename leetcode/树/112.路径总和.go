package main

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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

//边界情况：(1)root=nil sum = 0 false
//(2)[1] 1 true
//(3)[1,2] 1 false

//由于这里严格按照定义，必须要叶子节点到根节点，那么有两种情况
//既是根节点又是叶子节点
//根节点和叶子节点不相同


func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

// @lc code=end
