package main

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
var pathSumRes [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	pathSumRes = make([][]int, 0)
	help(root, sum, []int{})
	return pathSumRes
}

//解法一：真回溯
func help(root *TreeNode, sum int, res []int) {
	if root == nil {
		return
	}

	sum -= root.Val
	res = append(res, root.Val)

	//这个判断条件就是题目变种变化的更改点
	if root.Left == nil && root.Right == nil && sum == 0 {
		//因为pathSumRes加入得是res切片
		//这个切片没发生扩容得时候公用得同一个底层数组
		//避免先加入了正确结果，后来又改掉了res，导致错误
		resTemp := make([]int, len(res))
		copy(resTemp, res)
		pathSumRes = append(pathSumRes, resTemp)
		return
	}

	help(root.Left, sum, res)
	help(root.Right, sum, res)

	//节点回溯
	res = res[:len(res)-1]
}

//解法二：假回溯，通过深拷贝那避免回退
func help2(root *TreeNode, sum int, res []int) {
	if root == nil {
		return
	}

	sum -= root.Val
	res = append(res, root.Val)

	//这个判断条件就是题目变种变化的更改点
	if root.Left == nil && root.Right == nil && sum == 0 {
		pathSumRes = append(pathSumRes, res)
		return
	}

	//这里是采取了深拷贝的方式，无需再处理回溯需要剔除节点的问题了。
	resLeft := make([]int, len(res))
	resRight := make([]int, len(res))
	copy(resLeft, res)
	copy(resRight, res)

	help(root.Left, sum, res)
	help(root.Right, sum, res)
}

// @lc code=end
