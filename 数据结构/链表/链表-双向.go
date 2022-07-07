package main

import "fmt"

//链表分类
// 单链表
// 双链表
// 环链表

//时间复杂度
// 修改（增删改）时间复杂度 O(1)
// 查询时间复杂度O(n)

type Node struct {
	Data      int
	NextPoint *Node
	PrePoint  *Node
}
type LinkedList struct {
	head    *Node
	tail    *Node
}

func LinkNew() *LinkedList {
	return new(LinkedList)
}

func showLinkedList(link *LinkedList) {
	if link.head.NextPoint == nil{
		fmt.Println("链表为空")
	}

	currentNode := link.head
	for {
		if currentNode.NextPoint == nil {
			break
		}
		fmt.Println("Node:", currentNode.NextPoint.Data)

		currentNode = currentNode.NextPoint
	}
}

func insertNode(link *LinkedList, node *Node) {
	if link.head == nil {
		link.head = node
		link.tail = node
	} else {
		// 这两行处理真实链表数据
		link.tail.NextPoint = node
		node.PrePoint = link.tail

		// 这一行就尾部数据替换为当前的数据
		link.tail = node
	}
}

func main() {
	//data := []int{1, 21, 31, 51, 62, 2, 3, 42, 33, 12, 12}
	data := []int{1, }
	link := LinkNew()
	var currentNode *Node
	for i := 0; i < len(data); i++ {
		currentNode = new(Node)
		currentNode.Data = data[i]
		insertNode(link, currentNode)
	}
	showLinkedList(link)
}