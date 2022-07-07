package main

import (
	"fmt"
)

type LinkNode struct {
	no int
	name string
	prev *LinkNode
	next *LinkNode
}

//显示节点
func ShowNode(head *LinkNode) {
	//创建辅助节点
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}
	for{
		fmt.Printf("[%d,%s,%s] ==>",temp.next.no,temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

//给双向链表插入一个节点
func InsertNode(head *LinkNode, newNode *LinkNode) {
	temp := head
	for{
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.prev = temp
}
//双向链表有序添加
func InsertOrderNode(head *LinkNode,newNode *LinkNode) {
	//创建辅助节点
	temp := head
	flag := true
	for {
		if temp.next == nil {
			//说 明到链表的最后
			break
		} else if temp.next.no > newNode.no {
			//插入到temp后面
			break
		} else if temp.next.no == newNode.no {
			flag = false //说明链表中以存在该节点
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("当前节点以存在",newNode.no)
		return
	} else {
		newNode.next = temp.next
		newNode.prev = temp
		if temp.next != nil {
			temp.next.prev = newNode
		}
		temp.next = newNode
	}
}
//双向链表的删除
func DeleteNode(head *LinkNode,no int) {
	temp := head
	flag := false
	for{
		if temp.next == nil { //到达链表的最后
			break
		} else if temp.next.no == no {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag { //说明找到
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.prev = temp
		}
	} else {
		fmt.Println("没有要删除的节点")
	}
}

func LinkNew() *LinkNode{
	return &LinkNode{}
}

func main() {
	head := LinkNew()
	person1 := &LinkNode{
		no:1,
		name:"张三",
	}
	person2 := &LinkNode{
		no:2,
		name:"李四",
	}
	person3 := &LinkNode{
		no:3,
		name:"王五",
	}
	InsertOrderNode(head,person3)
	InsertOrderNode(head,person2)
	InsertOrderNode(head,person1)
	DeleteNode(head,2)
	ShowNode(head)

	InsertNode(head,person1)
	// InsertNode(head,person2)
	// ShowNode(head)
}