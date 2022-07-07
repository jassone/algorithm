package main

import "fmt"

//LeetCode 19.删除链表的倒数第N个节点

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//进阶：你能尝试使用一趟扫描实现吗？

//输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5] 示例 2：

//方法一：计算链表长度-官方
//时间复杂度：O(L)，其中 LL 是链表的长度。
//空间复杂度：O(1)。
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

//方法二：栈-官方
//时间复杂度：O(L)，其中 L 是链表的长度。
//空间复杂度：O(L)，其中 L 是链表的长度。主要为栈的开销。
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// 方法三：双指针-官方-秒
//时间复杂度：O(L)，其中 L 是链表的长度。
//空间复杂度：O(1)。
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ****插入时，都是处理最后一个节点的next****
// ****每次循环中判断的都是temp.next的情况，因为第一个node是无data的***
func NewList() *ListNode {
	return &ListNode{}
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
	fmt.Println()
}
