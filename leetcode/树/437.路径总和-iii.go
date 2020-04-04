package main

import "fmt"

/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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

//2020/3/15思路：遍历到任何节点时，（1）加入该节点本身（2）加入该节点和前面每一节点的和

//思路：这个路径总和和前面不同的点在于：
//1.没有限制必须要到叶子节点，也就是单个节点，或者中间节点，只要等于sum了，都可以那么就要把中间结果保存起来
//2.就是当中间两个节点为0的时候，这个中间态也要保存起来，下面也有可能和为0的节点能够利用到这条路径
var cnt int

func pathSum(root *TreeNode, sum int) int {
	cnt = 0
	help(root, sum, []int{})
	return cnt
}

//这个题目没办法通过res = res[:len(res)-1]进行回溯，因为
//有可能切片扩容了，你再去修改不会影响其他的了。

//解法一：伪回溯
func help1(root *TreeNode, sum int, res []int) {
	//递归终止条件
	if root == nil {
		return
	}

	//0.res数组中存放的都是和sum的差值，假设到了节点n，那么res数组
	//存放的是sum-node(n-1),sum-node(n-1)-node(n-2)...,sum-node(n-1)..-node(1)
	//根据这个就可以得出，第n层的某个节点，res数组的长度是n-1

	//1.先判断当前节点是否等于sum
	temp := sum - root.Val
	if temp == 0 {
		cnt++
	}

	//2.循环遍历之前的节点，判断是否可以组成一条等于sum的路径。
	//无论是否等于0，都要进行更新
	for i := 0; i < len(res); i++ {
		var result int
		if result = res[i] - root.Val; result == 0 {
			fmt.Println(root.Val)
			cnt++
		}
		res[i] = result
	}

	//3.把当前节点和sum的差值放入数组中
	res = append(res, temp)

	//回溯的时候不用作处理(时间和空间的代价都大，
	//看有没有更好的办法？)
	resLeft := make([]int, len(res))
	resRight := make([]int, len(res))
	copy(resLeft, res)
	copy(resRight, res)

	help(root.Left, sum, resLeft)
	help(root.Right, sum, resRight)
}

//解法二：进阶一点得伪回溯(比上面少了一半的空间和时间)
func help(root *TreeNode, sum int, res []int) {
	//递归终止条件
	if root == nil {
		return
	}

	//0.res数组中存放的都是和sum的差值，假设到了节点n，那么res数组
	//存放的是sum-node(n-1),sum-node(n-1)-node(n-2)...,sum-node(n-1)..-node(1)
	//根据这个就可以得出，第n层的某个节点，res数组的长度是n-1

	//1.先判断当前节点是否等于sum
	temp := sum - root.Val
	if temp == 0 {
		cnt++
	}

	//2.循环遍历之前的节点，判断是否可以组成一条等于sum的路径。
	//无论是否等于0，都要进行更新
	var current []int
	for i := 0; i < len(res); i++ {
		var result int
		if result = res[i] - root.Val; result == 0 {
			cnt++
		}
		current = append(current, result)
	}

	//3.把当前节点和sum的差值放入数组中
	current = append(current, temp)

	//回溯的时候不用作处理(时间和空间的代价都大，
	//看有没有更好的办法？)
	// resLeft := make([]int, len(res))
	// resRight := make([]int, len(res))
	// copy(resLeft, res)
	// copy(resRight, res)

	help(root.Left, sum, current)
	help(root.Right, sum, current)

}

// @lc code=end
