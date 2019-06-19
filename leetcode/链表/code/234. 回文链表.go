/*
234：
	1.slow-fast法找到链表中间节点
	2.if判断中是node.next != nil 还是 node ！= nil 取决于特定的情况
*/
package code

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一：把这道题变成之前做过的，遍历一遍用数组存储，然后双指针（可能不是面试官想考察的点）
func IsPalindrome1(head *ListNode) bool {
	arrays := make([]int, 0)
	for head != nil {
		arrays = append(arrays, head.Val)
		head = head.Next
	}

	for i, j := 0, len(arrays)-1; i <= j; {
		if arrays[i] != arrays[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

//解法二：先用slow-fast法找到中间节点，然后反转后半部分，然后比较
func IsPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow := head, head
	cnt := 0
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		cnt++
		if fast.Next != nil {
			fast = fast.Next
		} else {
			break
		}
	}
	tail := reverse(slow)
	for ; cnt > 0; cnt-- {
		if tail.Val != head.Val {
			return false
		} else {
			tail = tail.Next
			head = head.Next
		}
	}
	return true
}
func reverse(head *ListNode) *ListNode {
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
