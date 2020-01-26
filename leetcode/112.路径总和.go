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

//这个完全是考虑边界情况
//一个是sum=0，只有根节点，只有根节点和叶子节点
//以及负数的情况，根节点的概念问题

//思路：如果不是叶子节点，则继续看下去，如果是叶子节点,则看sum是否为0
//思路没这么清晰的时候碰到了几个特殊用例，[1,2,3] 1 false和[1] 1 true
//怎么去模拟也没弄出这么简洁明了的思路
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
