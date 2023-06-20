package main

import "fmt"

//LeetCode 707.设计链表
//在链表类中实现这些功能：
//get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index
// 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index
// 小于0，则在头部插入节点。
//deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

//循环双链表
type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

//仅保存哑节点，pre-> rear, next-> head
/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	rear := &Node{ // 想象一下rear(Node的指针值)值为0x01
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear // 想象一下rear.Next的值为0x01
	rear.Pre = rear  // 想象一下rear.Next的值为0x01
	return MyLinkedList{rear}

	// 所以下面的打印值都一样，都在打转， &{-1 0xc00011c018 0xc00011c018}
	//fmt.Println(MyLinkedList.dummy)
	//fmt.Println(MyLinkedList.dummy.Next)
	//fmt.Println(MyLinkedList.dummy.Next.Next)
	//fmt.Println(MyLinkedList.dummy.Next.Next.Next)
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	//head == this, 遍历完全
	for head != this.dummy && index > 0 {
		index--
		head = head.Next
	}
	//否则, head == this, 索引无效
	if 0 != index {
		return -1
	}
	return head.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val: val,
		//head.Next指向原头节点
		Next: dummy.Next,
		//head.Pre 指向哑节点
		Pre: dummy,
	}

	//更新原头节点
	dummy.Next.Pre = node
	//更新哑节点
	dummy.Next = node
	//以上两步不能反
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	rear := &Node{
		Val: val,
		//rear.Next = dummy(哑节点)
		Next: dummy,
		//rear.Pre = ori_rear
		Pre: dummy.Pre,
	}

	//ori_rear.Next = rear
	dummy.Pre.Next = rear
	//update dummy
	dummy.Pre = rear
	//以上两步不能反
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.Next
	//head = MyLinkedList[index]
	//fmt.Println(head)
	//fmt.Println(this.dummy)
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index > 0 {
		return
	}
	node := &Node{
		Val: val,
		//node.Next = MyLinkedList[index]
		Next: head,
		//node.Pre = MyLinkedList[index-1]
		Pre: head.Pre,
	}
	//MyLinkedList[index-1].Next = node
	head.Pre.Next = node
	//MyLinkedList[index].Pre = node
	head.Pre = node
	//以上两步不能反
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	//链表为空
	if this.dummy.Next == this.dummy {
		return
	}
	head := this.dummy.Next
	//head = MyLinkedList[index]
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	//验证index有效
	if index == 0 {
		//MyLinkedList[index].Pre = index[index-2]
		head.Next.Pre = head.Pre
		//MyLinedList[index-2].Next = index[index]
		head.Pre.Next = head.Next
		//以上两步顺序无所谓
	}
}

func main() {
	pre := Constructor()
	fmt.Println(pre.dummy)
	fmt.Println(pre.dummy.Pre)
	fmt.Println(pre.dummy.Next)
	fmt.Println(pre.dummy.Pre.Pre)
	fmt.Println(pre.dummy.Next.Next)
	fmt.Println(pre.dummy.Pre.Pre.Pre)
	fmt.Println(pre.dummy.Next.Next.Next)
	fmt.Println(pre.dummy.Pre.Pre.Pre.Pre)
	fmt.Println(pre.dummy.Next.Next.Next.Next)
	fmt.Println()

	pre.AddAtHead(2)
	pre.AddAtHead(1)
	fmt.Println(pre.dummy)
	fmt.Println(pre.dummy.Pre)
	fmt.Println(pre.dummy.Next)
	fmt.Println(pre.dummy.Pre.Pre)
	fmt.Println(pre.dummy.Next.Next.Pre)
	fmt.Println(pre.dummy.Next.Next)
	fmt.Println(pre.dummy.Pre.Pre.Pre)
	fmt.Println()

	fmt.Println(pre.dummy.Next.Next.Next)
	fmt.Println(pre.dummy.Pre.Pre.Pre.Pre)
	fmt.Println(pre.dummy.Next.Next.Next.Next)

	//pre.AddAtIndex(1, 1)
	//fmt.Println(pre.dummy.Next.Next.Val)
}
