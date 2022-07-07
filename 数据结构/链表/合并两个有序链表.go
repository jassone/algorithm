package main

import "fmt"

type LinkNode struct {
	no int
	next *LinkNode
}

func NewList() *LinkNode {
	return &LinkNode{}
}

func display(node *LinkNode) {
	if node.next == nil {
		fmt.Println("链表为空")
	} else {
		temp := node
		for {
			if temp.next == nil{
				break
			}
			fmt.Println(temp.next.no)

			temp = temp.next
		}
	}
}

func merge(list1, list2 *LinkNode) *LinkNode {
	newList := NewList()
	pre := newList

	temp1 := list1
	temp2 := list2
	for  {
		if temp1.next == nil && temp2.next == nil {
			break
		}

		t1Num := temp1.next.no
		t2Num := temp2.next.no

		if t1Num < t2Num {
			pre.next = temp1.next
			temp1 = temp1.next.next
		} else {
			pre.next = temp2.next
			temp2 = temp2.next.next
		}

		pre = pre.next
	}

    return newList
}

func main() {

	cur1 := NewList()
	var node *LinkNode
	for i := 1; i <= 3; i++ {
		node = &LinkNode{no:i}
		cur1.next = node
		cur1 = cur1.next
	}

	cur2 := NewList()
	var node2 *LinkNode
	for i := 5; i <= 9; i++ {
		node2 = &LinkNode{no:i}
		cur2.next = node2
		cur2 = cur2.next
	}

	s := merge(cur1,cur2)
	fmt.Println(s)
	display(merge(cur1,cur2))
}
