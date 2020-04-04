package main

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

//边界情况考虑 树中是否有重复的元素（题目说了没有）

//注意一下切片的操作，切片是一个前闭后开的区间[n,m)因此m要大于数组最后一个索引+1，才表示的是整个数组
//也就是你从数组切一个数，那就是a[n:n+1]，不切数字就是a[n:n]

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	rootIndex := find(preorder[0], inorder)

	//中序数组中根节点的索引，表示的是有多少个左子树。
	//因此左子树的前序是1:rootIndex+1(切片的特性)，中序是从0到rootIndex
	//右子树同理
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

func find(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}

// func main() {
// 	a := []int{1, 2, 3}
// 	b := a[0:len(a)]
// 	// fmt.Println(a[2:])
// 	fmt.Println(len(b))
// }

// @lc code=end
