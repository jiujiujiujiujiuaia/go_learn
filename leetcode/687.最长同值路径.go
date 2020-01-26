package main

/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
var res int

func longestUnivaluePath(root *TreeNode) int {
	res = 0
	helpLongestUnivaluePath(root)
	return res
}

func helpLongestUnivaluePath(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := helpLongestUnivaluePath(node.Left)
	right := helpLongestUnivaluePath(node.Right)

	//这个操作非常巧妙，如果没出现断层，那么就是max(left,right)
	//但是出现了断层的话，应该有一边是0，而不是把之前的最大值传上来
	//而且还有一个情况是当父节点等于左右节点时，应该是res = max(res,left + right)
	sameL := 0
	sameR := 0

	if node.Left != nil && node.Left.Val == node.Val {
		sameL = left + 1
	}

	if node.Right != nil && node.Right.Val == node.Val {
		sameR = right + 1
	}

	res = max(res, sameL+sameR)
	return max(sameL, sameR)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
