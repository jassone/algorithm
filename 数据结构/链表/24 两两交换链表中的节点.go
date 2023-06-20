package main

import "fmt"

//LeetCode 24. 两两交换链表中的节点

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题
//（即，只能进行节点交换）。

// 方法1：迭代-官方,推荐
//时间复杂度：O(n)，其中 n 是链表的节点数量。需要对每个节点进行更新指针的操作。
//空间复杂度：O(1)。
func swapPairs1(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

// 递归版本-官方
//时间复杂度：O(n)，其中 n 是链表的节点数量。需要对每个节点进行更新指针的操作。
//空间复杂度：O(n)，其中 n 是链表的节点数量。空间复杂度主要取决于递归调用的栈空间。
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs2(newHead.Next)
	newHead.Next = head
	return newHead
}

func main() {
	List := NewList()

	var node *ListNode
	cur := List
	for i := 1; i <= 4; i++ {
		node = &ListNode{Val: i}
		cur.Next = node
		cur = cur.Next
	}
	display(List)

	// *****特别注意******
	// func里面只处理了逻辑上的链表，入参和出参的时候要特殊处理下
	List.Next = swapPairs1(List.Next)
	display(List)

	List.Next = swapPairs2(List.Next)
	display(List)
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
