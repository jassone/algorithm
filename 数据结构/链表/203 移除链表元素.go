package main

import "fmt"

//LeetCode 203. 移除链表元素
type ListNode struct {
	Val  int
	Next *ListNode
}

func display(node *ListNode) {
	if node.Next == nil {
		fmt.Println("链表为空")
	}

	temp := node
	for {
		if temp.Next == nil {
			break
		}
		fmt.Printf("[%d]", temp.Next.Val)
		temp = temp.Next
	}
}

//方法一：递归-官方
//时间复杂度：O(n)，其中 n 是链表的长度。递归过程中需要遍历链表一次。
//空间复杂度：O(n)，其中 n 是链表的长度。空间复杂度主要取决于递归调用栈，最多不会超过 n 层。
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

//方法二：迭代-官方
//重点是引入一个哑节点
//也可以用迭代的方法删除链表中所有节点值等于特定值的节点。
//时间复杂度：O(n)，其中 n 是链表的长度。需要遍历链表一次。
//空间复杂度：O(1)。
func removeElements2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func main() {
	head := &ListNode{}
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 2,
	}
	node4 := &ListNode{
		Val: 3,
	}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	target := 2
	//removeElements(head, target)
	removeElements2(head, target)
	display(head)

}
