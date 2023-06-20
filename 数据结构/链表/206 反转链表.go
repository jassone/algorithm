package main

import "fmt"

//https://www.programmercarl.com/0206.翻转链表.html

//LeetCode 206.反转链表
//题意：反转一个单链表。
//示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL

// 方法1：双指针法/迭代-官方,推荐
//总结下来
// 切断前先存储下一个节点
// 把指向反一下
// 把处理好的链表头往后移一下
// 把待反转的链表头往后移动一下
//时间复杂度：O(n)，其中 n 是链表的长度。需要遍历链表一次。
//空间复杂度：O(1)。
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode // 已经反转处理好的链表头，为nil
	curr := head       // 当前待反转的链表头
	for curr != nil {
		next := curr.Next // 临时存储待转换结点的下一个节点
		curr.Next = prev  // 将待转换结点的下一个节点指向已经反转好的链表头，第一次循环prev为nill
		prev = curr       // 把已经反转好的链表头往后移动一位
		curr = next       // 当前待反转的链表头往后移动一位
	}
	return prev
}

// 方法2：递归-官方
//时间复杂度：O(n)，其中 n 是链表的长度。需要对链表的每个节点进行反转操作。
//空间复杂度：O(n)，其中 n 是链表的长度。空间复杂度主要取决于递归调用的栈空间，最多为 n 层。
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	// 下面两步将前后两个节点反一下，连接起来
	head.Next.Next = head
	head.Next = nil

	return newHead // 所以每次归的时候都是返回最后一个节点
}

func main() {
	List := NewList()

	var node *ListNode
	cur := List
	for i := 1; i <= 3; i++ {
		node = &ListNode{Val: i}
		cur.Next = node
		cur = cur.Next
	}
	display(List)

	// *****特别注意******
	// func里面只处理了逻辑上的链表，入参和出参的时候要特殊处理下
	List.Next = reverseList1(List.Next)
	display(List)

	//List.Next = reverseList2(List.Next)
	//display(List)
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
