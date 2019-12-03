package main

func main() {

}

//这个题目虽然简单，但是和常见的删除节点有些不一样。
//一般的删除节点是：a-b-c,a.next=b.next，这个题目是把c复制给b
//当然也是因为这个题目有限制条件，不会给最后一个节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
