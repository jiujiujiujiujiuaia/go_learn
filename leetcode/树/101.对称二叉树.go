package main

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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

//突然想起来了，可以记一笔
//递归返回标记值，如果题目不能完成，可以快速返回而不是遍历完整个数据（如果数据规模故意很大陷害你）

//递归法
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return helpIsSymmetric(root.Left, root.Right)
	}
}

func helpIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val == right.Val {
		return helpIsSymmetric(left.Left, right.Right) && helpIsSymmetric(left.Right, right.Left)
	} else {
		return false
	}
}

//迭代法(讨巧一点得方法，用栈空间去模拟递归，或者思考其他数据结构来完成)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftRoot, rightRoot := root.Left, root.Right
	for {
		if leftRoot.Val == rightRoot.Val {

		} else {
			return false
		}
	}
}

// @lc code=end
