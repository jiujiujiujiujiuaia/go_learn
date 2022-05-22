package linkedlist

import "fmt"

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	//1.先算出整体链表的长度和尾节点
	dummyHead := new(ListNode)
	dummyHead.Next = head
	tail := dummyHead
	length := 0
	for tail.Next != nil {
		tail = tail.Next
		length++
	}
	fmt.Println(length)

	//2. 移动可能会出现大于length的情况
	k = k % length

	fast, slow := dummyHead, dummyHead
	//距离+1 slow节点找到的就是倒数第k+1个节点
	//再多说一句，由于已经知道了长度，可以不用这么麻烦的去找到那个节点
	for i := 0; i <= k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	//3.找到了新的头节点，连成一个环再断开
	tail.Next = head
	newHead := slow.Next
	slow.Next = nil
	return newHead
}

// @lc code=end
