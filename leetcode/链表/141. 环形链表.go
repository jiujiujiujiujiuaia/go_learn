package main

//这个题目很简单，如果是一个环形赛道，跑的快的最终一定会比超过跑的慢的一圈
//因为快的是慢的速度两倍，所以当慢的走完时，如果是环，则快的应该等于慢的，如果不是环，则不等
//这里写一个do（）while（）会更优雅，但是golang不支持
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
