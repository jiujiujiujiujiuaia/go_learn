package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

type ListNode struct {
      Val int
      Next *ListNode
}

//Label:[链表遍历][新增链表节点][dummy节点使用]
//Corner Cases:
//这是一道经典的链表题，[链表遍历][新增链表节点][dummy节点使用]
//链表题的难点在于遍历链表的时候的条件
//当循环条件是currentNode != nil, 意味着所有节点都会被遍历到
//当循环条件是currentNode.Next !=nil,意味着最后一个节点不会遍历到，最后一个lastNode.next == nil
//只能通过currentNode.Next.Val访问

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := new(ListNode)
	//1.复制指针
	dummyNode := headNode
	dummyNodeL1 := l1
	dummyNodeL2 := l2
	carryOne := 0
	for (dummyNodeL1 != nil && dummyNodeL2 != nil){
		dummyNode.Val = (dummyNodeL1.Val + dummyNodeL2.Val + carryOne) % 10
		carryOne = (dummyNodeL1.Val + dummyNodeL2.Val + carryOne) / 10

		if(dummyNodeL1.Next == nil || dummyNodeL2.Next == nil){
			break
		}else{
			nextNode := new(ListNode)
			dummyNode.Next = nextNode
			dummyNode = nextNode
		}
		dummyNodeL1 = dummyNodeL1.Next
		dummyNodeL2 = dummyNodeL2.Next
	}

	for (dummyNodeL1.Next != nil){
		nextNode := new(ListNode)
		nextNode.Val = (dummyNodeL1.Next.Val + carryOne) % 10 
		carryOne = (dummyNodeL1.Next.Val + carryOne) / 10
		dummyNode.Next = nextNode
		dummyNode = nextNode
		dummyNodeL1 = dummyNodeL1.Next
	}

	for (dummyNodeL2.Next != nil){
		nextNode := new(ListNode)
		nextNode.Val = (dummyNodeL2.Next.Val + carryOne) % 10 
		carryOne = (dummyNodeL2.Next.Val + carryOne) / 10
		dummyNode.Next = nextNode
		dummyNode = nextNode
		dummyNodeL2 = dummyNodeL2.Next
	}

	if (carryOne != 0){
		nextNode := new(ListNode)
		nextNode.Val = carryOne
		dummyNode.Next = nextNode
	}

	return headNode
}
// @lc code=end

