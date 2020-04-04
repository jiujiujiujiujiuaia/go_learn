package main

import "container/list"

//1.二叉搜索树是左节点<root<右节点
//2.中序是先根节点，然后左节点，再右节点。
//因此最近pop出来的那个值，一定是小于待push的那个值，可以以最小的BST举例
//因此，前序遍历先进栈中，维持一个单调递减的栈，当某个元素大于top，就pop出，然后记录最近的值，如果即将push
func verifyPreorder(preorder []int) bool {
	if len(preorder) == 0 {
		return true
	}
	stack := list.New()
	stack.PushFront(preorder[0])
	last := 0
	for i := 1; i < len(preorder); i++ {
		//（1）校验左子树单调递减的性质
		for stack.Len() != 0 && preorder[i] > stack.Front().Value.(int) {
			last = stack.Front().Value.(int)
			stack.Remove(stack.Front())
		}
		//（2）待push的值一定是右子树的，因此右子树一定大于根或左子树的值
		//否则false
		if preorder[i] < last {
			return false
		}
		stack.PushFront(preorder[i])
	}

	return true
}
