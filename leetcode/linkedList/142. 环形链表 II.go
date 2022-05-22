package linkedlist

//这个利用哈希表查找出是环形链表的第一个节点是比较直观的
//官方题解上给出了 另一种思路 延续了环形链表1中判断是否是环形链表快慢指针的思路
//https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode/
func detectCycle(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	dummyHead := head
	for dummyHead != nil {
		if ok, _ := nodeMap[dummyHead]; ok {
			return dummyHead
		}
		nodeMap[dummyHead] = true
		dummyHead = dummyHead.Next
	}
	return nil
}

//解法二 也就是快慢指针去寻找，可以参考notability上和287题解法
