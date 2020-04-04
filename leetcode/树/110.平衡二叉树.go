package main

import "math"

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
const NOT_BALANCE_BINARY_TREE = -1

//平衡二叉树:一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。

//2020/3/15 应该弄一个全局变量这样才可以避免已经找到了非平衡二叉树，但是要把所有节点都遍历一遍
func isBalanced(root *TreeNode) bool {
	if judgeIsBalanceTree(root) == NOT_BALANCE_BINARY_TREE {
		return false
	} else {
		return true
	}
}

func judgeIsBalanceTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := judgeIsBalanceTree(root.Left)
	r := judgeIsBalanceTree(root.Right)
	if l == NOT_BALANCE_BINARY_TREE || r == NOT_BALANCE_BINARY_TREE ||
		math.Abs(float64(l-r)) > 1 {
		return NOT_BALANCE_BINARY_TREE
	}
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
