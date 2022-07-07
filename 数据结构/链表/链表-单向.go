package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// ****插入时，都是处理最后一个节点的next****
// ****每次循环中判断的都是temp.next的情况，因为第一个node是无data的***

func NewList() *ListNode {
	return &ListNode{}
}

// 第一种插入方法
func insertListNode(node *ListNode, newNode *ListNode) {
	temp := node
	for {
		if temp.Next == nil {
			break
		}

		temp = temp.Next
	}

	temp.Next = newNode
}

// 第二种插入方法
func sortInsertListNode(node *ListNode, newNode *ListNode) {
	temp := node
	flag := true
	for {

		if temp.Next == nil { //最后一个
			break
		} else if temp.Next.Val > newNode.Val { // 说明要插入到其前面
			break
		} else if temp.Next.Val == newNode.Val { // 节点已经存在
			flag = false
			break
		}
		temp = temp.Next
	}

	if !flag {
		fmt.Println("节点已经存在")
	} else {
		newNode.Next = temp.Next
		temp.Next = newNode
	}
}

func listListNode(node *ListNode) {
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

func delListNode(node *ListNode, id int) {
	if node.Next == nil {
		fmt.Println("链表为空")
	}

	temp := node
	flag := false
	//找到要删除的结点的no，和temp的下一个结点的no比较
	for {
		if temp.Next == nil { //说明已经到链表的最后了
			break
		} else if temp.Next.Val == id {
			//说明我们找到这个要删除的结点了
			flag = true
			break
		}
		temp = temp.Next
	}
	if flag {
		temp.Next = temp.Next.Next //要删除这个结点即直接略过这个结点
		// ***temp.next.next可能是nil,也可能不是nil,都无所谓***
		//然后这个被略过的结点会变成垃圾结点，会被链表删除
	} else {
		fmt.Println("要删除的id不存在")
	}
}

func modifyListNode(node *ListNode, id int, newNode *ListNode) {
	temp := node
	flag := false
	//思路：
	//1.找到要修改的结点的no，和temp.next.no作比较
	for {
		if temp.Next == nil { //说明已经到了链表的最后
			break
		} else if temp.Next.Val == id {
			//说明我们已经找到这个要修改的结点了
			flag = true
			break
		}
		temp = temp.Next
	}
	if flag {
		newNode.Next = temp.Next.Next // 将后面的接在新的后面
		temp.Next = newNode
	} else {
		fmt.Println("要修改的结点不存在")
	}
}

func main() {
	//1.先创建一个头结点
	head := NewList()

	//2.创建一个新的ListNode
	stuLisa := &ListNode{
		Val: 1,
	}
	stuBob := &ListNode{
		Val: 2,
	}
	stuNick := &ListNode{
		Val: 3,
	}
	stuMark := &ListNode{
		Val: 4,
	}
	stuMarket := &ListNode{
		Val: 4,
	}
	//3.加入结点
	insertListNode(head, stuLisa)
	insertListNode(head, stuBob)

	//显示链表
	listListNode(head)
	fmt.Println()

	//
	////4.加入结点（第二种方法）
	sortInsertListNode(head, stuMark) //no是4
	sortInsertListNode(head, stuNick) //no是3
	listListNode(head)
	fmt.Println()
	////5.删除结点
	//delListNode(head, 2)
	//listListNode(head)
	fmt.Println()
	////6.修改结点
	modifyListNode(head, 4, stuMarket) // 如果这里 id =2 是不是就乱序了 todo
	listListNode(head)
}
