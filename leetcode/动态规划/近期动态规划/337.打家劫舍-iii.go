package main

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
var TreeNodeMap map[*TreeNode]int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	TreeNodeMap = make(map[*TreeNode]int)
	helpRob(root)
	return TreeNodeMap[root]
}

func helpRob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if val, ok := TreeNodeMap[root]; ok {
		return val
	}

	//1.选择偷root，则左右子树皆不可以偷了
	chooseStealRoot := root.Val
	if root.Left != nil {
		chooseStealRoot += helpRob(root.Left.Left) + helpRob(root.Left.Right)
	}
	if root.Right != nil {
		chooseStealRoot += helpRob(root.Right.Left) + helpRob(root.Right.Right)
	}

	//2.选择不偷root，可以偷左右子树
	NotChooseStealRoot := helpRob(root.Left) + helpRob(root.Right)

	//3.比较大小
	TreeNodeMap[root] = max(chooseStealRoot, NotChooseStealRoot)
	return TreeNodeMap[root]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
