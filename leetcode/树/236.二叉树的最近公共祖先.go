package main

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

//这个函数就是寻找这两个节点
//如果这两个节点任意一个等于根节点，则返回
//如果都不等于，那么就分别在左右子树搜索这两个节点，如果分别在两边
//则返回root，如果一个

//解法一：递归的思路
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//1.如果p q任意一个节点等于根节点，这个节点就是祖先
	if p.Val == root.Val {
		return root
	}

	if q.Val == root.Val {
		return root
	}

	//2.如果根节点不是，则分别去左右两边搜索
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//3.如果左右节点都找到了，则根节点就是，如果一边空一边不是，则直接返回非空的
	if left != nil && right != nil {
		return root
	} else if left == nil && right != nil {
		return right
	} else {
		return left
	}
}

//解法二：迭代的思路
//时间来不及了，说说思路

//首先利用一个hashmap存放着node->parent 直到同时找到p q
//然后用一个set存放所有p的祖先
//去q中寻找，看是否有q的祖先和p的祖先相同，那么那个就是最近的祖先

// @lc code=end
