package main

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

//解法二 只用一个栈
//当有更小的值来的时候，我们只需要把之前的最小值入栈，当前更小的值再入栈即可。
//当这个最小值要出栈的时候，下一个值便是之前的最小值了。

//解法一：运用辅助栈（单调栈）注意：两个栈不要用一个node 同时注意重复元素的情况
type Node struct {
	Val  int
	Next *Node
}
type MinStack struct {
	StackDummyHead    *Node
	MinStackDummyHead *Node
	Length            int
}

func Constructor() MinStack {
	node1 := new(Node)
	node2 := new(Node)
	return MinStack{StackDummyHead: node1, MinStackDummyHead: node2, Length: 0}
}

func (this *MinStack) Push(x int) {
	//如果推入的元素比单调栈的栈顶更小 就推入

	if this.Length == 0 {
		node1 := &Node{Val: x}
		node2 := &Node{Val: x}
		this.StackDummyHead.Next, this.MinStackDummyHead.Next = node1, node2
	} else {
		insert(this.StackDummyHead, this.MinStackDummyHead, x)
	}
	this.Length++
}

func (this *MinStack) Pop() {
	//如删除元素比单调栈大 无视
	if this.Length == 0 {
		return
	} else {
		this.Length--
		delete(this.StackDummyHead, this.MinStackDummyHead)
	}
}

func (this *MinStack) Top() int {
	return this.StackDummyHead.Next.Val
}

func (this *MinStack) GetMin() int {
	return this.MinStackDummyHead.Next.Val
}

func delete(StackDummyHead *Node, MinStackDummyHead *Node) {
	if StackDummyHead.Next.Val == MinStackDummyHead.Next.Val {
		MinStackDummyHead.Next = MinStackDummyHead.Next.Next
	}
	StackDummyHead.Next = StackDummyHead.Next.Next

}

func insert(StackDummyHead *Node, MinStackDummyHead *Node, val int) {
	node := &Node{Val: val}
	node.Next = StackDummyHead.Next
	StackDummyHead.Next = node
	if val <= MinStackDummyHead.Next.Val {
		node1 := &Node{Val: val}
		node1.Next = MinStackDummyHead.Next
		MinStackDummyHead.Next = node1
	}
}

// @lc code=end
