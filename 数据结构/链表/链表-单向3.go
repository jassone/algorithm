package main

import (
	"fmt"
	"reflect"
	"log"
	"errors"
)

type ElementType interface{}

// 和go的双向链表接近了
type LinkNode struct {
	Data ElementType
	Next *LinkNode
}

type LinkedList struct {
	Head *LinkNode
}

func NewNode(data ElementType, next *LinkNode) *LinkNode {
	return &LinkNode{data, next}
}

func NewList() *LinkedList {
	head := &LinkNode{0, nil}

	return &LinkedList{head}
}

//处理链表头节点的数据大小
func (list *LinkedList) sizeListInc() {
	v := reflect.ValueOf((*list.Head).Data).Int()//v := reflect.ValueOf((list.Head).Data).Int()
	list.Head.Data = v + 1
}

func (list *LinkedList) sizeListDec() {
	v := reflect.ValueOf((*list.Head).Data).Int()
	if v == 0 {
		return
	}
	list.Head.Data = v - 1
}

//在链表的后边添加节点
func (list *LinkedList) Append(node *LinkNode) {
	if node == nil {
		log.Panic("待添加的节点不能是nil节点")
	}

	current := list.Head//即头节点
	if current.Next == nil {
		current.Next = node
		list.sizeListInc()
		fmt.Println(reflect.ValueOf((*list.Head).Data).Int())
		return
	}

	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	list.sizeListInc()

	fmt.Println(reflect.ValueOf((*list.Head).Data).Int())
}

//在链表前面添加节点
func (list *LinkedList) Prepend(node *LinkNode) {
	if node == nil {
		log.Panic("待添加的节点不能是nil节点")
	}

	current := list.Head//即头节点
	node.Next = current.Next
	current.Next = node
	list.sizeListInc()

	return
}

//遍历节点
func (list *LinkedList) Print() {
	fmt.Println("开始遍历链表")
	if list.Head.Next == nil {
		fmt.Println("list 是空链表")
		return
	}

	//测试节点数量
	fmt.Println("head node's Data: ", reflect.ValueOf((*list.Head).Data).Int())

	i := 1
	current := list.Head.Next//第一个节点
	for current.Next != nil {
		fmt.Printf("Node %d:    Data = %v, current node's address = %p, Next = %p\n", i, current.Data, current, current.Next)
		current = current.Next
		i++
	}
	fmt.Printf("Node %d:    Data = %v, current node's address = %p, Next = %p\n", i, current.Data, current, current.Next)
}

//查找指定值的节点，仅查找符合查找条件的第一个值，并返回该节点
func (list *LinkedList) Find(element ElementType) (*LinkNode, bool){
	if list.Head.Next == nil {
		log.Println("链表为空")
		return nil, false
	}

	current := list.Head//即头节点
	for current.Next != nil {
		current = current.Next
		if current.Data == element {
			return current, true
		}
	}
	log.Printf("未查到节点值为%s的节点\n", element)

	return nil, false
}

//删除指定节点
func (list *LinkedList) Remove(element ElementType) error {
	current := list.Head
	for current.Next != nil {
		if current.Next.Data == element {
			current.Next = current.Next.Next
			list.sizeListDec()
			return nil
		}
		current = current.Next
	}

	return errors.New(fmt.Sprintf("Not find the node that value is %v, fail in Removing node", element))
}

func main() {
	list := NewList()

	fmt.Println("\n开始测试在链表后面追加节点")
	var node *LinkNode = &LinkNode{}
	for i := 1; i < 5; i++ {
		node = &LinkNode{Data: fmt.Sprintf("node%d", i)}
		list.Append(node)
	}
	list.Print()

	fmt.Println("\n开始测试在链表前面插入节点")
	node = &LinkNode{Data:"node5"}
	list.Prepend(node)
	list.Print()

	fmt.Println("\n开始测试查找链表节点")
	if linkNodePointer, ok := list.Find("node3");  ok != false {
		fmt.Printf("find node:    Data = %v, node's address = %p, Next = %p\n", linkNodePointer.Data, linkNodePointer, linkNodePointer.Next)
	}

	fmt.Println("\n开始测试删除节点")
	if err := list.Remove("node"); err != nil {
		fmt.Println("删除节点失败")
		list.Print()
	} else {
		fmt.Println("成功删除指定节点")
		list.Print()
	}
}