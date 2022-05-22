/*
206
	1.经典的反转链表，思路就是模拟pre,cur,next，想象这三个节点要反转的话应该怎么办
	2.注意go语言中，new()一个结构体会有默认值，var struct 就是空nil
*/
package linkedlist

func ReverseList(head *ListNode) *ListNode {
	var next, pre *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
