package main

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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


// func isValidBST(root *TreeNode) bool {
// 	res = make([]int,0)
// 	transformBSTTOArray(root)
// 	fmt.Println(res)
// 	return checkArrayIsSorted()
	// if root == nil {
	// 	return true
	// }

	// if root.Left == nil && root.Right == nil {
	// 	return true
	// }

	// var l, r bool
	// if root.Left == nil {
	// 	l = true
	// } else if root.Val > root.Left.Val {
	// 	l = true
	// }

	// if root.Right == nil {
	// 	r = true
	// } else if root.Val < root.Right.Val {
	// 	r = true
	// }

	// if l && r {
	// 	return isValidBST(root.Left) && isValidBST(root.Right)
	// } else {
	// 	return false
	// }

// }



//解法一，中序遍历整棵树，然后把节点保存在数组中判断是否有序数组
//这种做法有一个缺点是，错误情况还是要遍历完BST
var res []int
func isValidBST1(root *TreeNode) bool {
	res = make([]int,0)
	transformBSTTOArray(root)
	fmt.Println(res)
	return checkArrayIsSorted()
}

func transformBSTTOArray(root *TreeNode){
	if root == nil{
		return 
	}

	transformBSTTOArray(root.Left)
	res = append(res,root.Val)
	transformBSTTOArray(root.Right)
}

func checkArrayIsSorted()bool{
	for i := 0;i<len(res) - 1;i++{
		if res[i] >= res[i+1]{
			return false
		}
	}
	return true
}

//解法二：前一种解法使用了list，完全可以使用空间为1的复杂度
var pre int
func isValidBST(root *TreeNode) bool {
	pre = -2147483649
	return helpIsValidBST(root)
}

//这道题是很经典的递归思想，只处理当前
//放在这道题目来说，我这样考虑，我左边节点都ok，那么前一个节点不能大于我，然后再看右边节点
//是否ok，如果都ok，那么返回true。最后对异常情况/递归终止情况讨论。

//这个题目最关键的是pre这个值，先把pre=左节点，然后到根节点。
//根节点必须大于左节点，然后pre=根节点，右节点必须要大于根节点然后就是pre=右节点
func helpIsValidBST(root *TreeNode)bool{

	if root == nil{
		return true
	}

	if !helpIsValidBST(root.Left){
		return false
	}

	if pre >= root.Val {
		return false
	}

	pre = root.Val
	if !helpIsValidBST(root.Right){
		return false
	}

	return true

}
// @lc code=end
