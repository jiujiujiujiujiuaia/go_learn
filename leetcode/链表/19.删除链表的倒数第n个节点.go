package main

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

//拿到这个题后，可以沟通（1）n会大于链表的长度吗
//（2）数据规模有多大
//（3）时间复杂度上有要求吗
//（4）可以先出一个简单的版本的再慢慢优化吗

//总结：1.哑节点解决头节点被删除的情况
//	    2.快慢指针表示间隔
//		3.当判断条件是while node != nil 表示的是遍历整个链表
//	      当判断条件是while node.Next != nil 表示的是遍历前n-1个结点

//2020/3/10

//解法一：暴力模拟版本
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//1.三种情况，被删除的节点在头，尾，中间.
	//引入哑节点解决
	//如果被删除节点是头节点需要判断的问题

	total := 0
	dummyHead := new(ListNode)
	dummyHead.Next = head
	first := head

	//2.看链表有多长
	for first != nil {
		total++
		first = first.Next
	}

	//3.由于引入了一个节点，因此total-=n是被删除节点前一个节点
	first = dummyHead
	total -= n
	for total > 0 {
		total--
		first = first.Next
	}

	//4.找到被删除节点前一个即可
	first.Next = first.Next.Next
	return dummyHead.Next
}

//解法二：只遍历一遍算法
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head

	//1.引入双指针，由于倒数第几个节点是以末尾节点为锚
	//双指针的间隔为n，当快指针到达末尾了，则慢指针距离快n
	first, second := dummyHead, dummyHead

	//2.距离+1，为了找到被删除节点前一个节点
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummyHead.Next
}

// @lc code=end
