package linkedlist

//这个题意不是很好理解
//这里的node是要删除的节点
//一般的删除是从头节点开始，一个一个遍历删除

//题目的边界是1.删除的是否是尾节点2.链表长度多长
func deleteNode(node *ListNode) {
	//复制下一个节点的值
	//然后删除下一个节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
