/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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

//时间来不及了，就不细说
//解法一，中序遍历整个树，放在数组中，然后数组指针往后移动
//这样空间复杂度比较高，不符合题意

//解法二，把所有根节点和左节点放在栈中，然后next推出的时候，如果该节点有右节点
//把该节点的所有左节点加入栈中

type BSTIterator struct {

}


func Constructor(root *TreeNode) BSTIterator {

}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {

}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {

}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

