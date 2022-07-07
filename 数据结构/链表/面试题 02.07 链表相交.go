package main

import "fmt"

//LeetCode 面试题 02.07. 链表相交
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，
//返回 null 。
//题目数据 保证 整个链式结构中不存在环。
//注意，函数返回结果后，链表必须 保持其原始结构 。

//思路
// 简单来说，就是求两个链表交点节点的指针。 这里要注意，交点不是数值相等，而是指针相等。

//方法一：哈希集合-官方
//时间复杂度：O(m+n)，其中 m 和 n 是分别是链表 headA 和 headB 的长度。需要遍历两个链表各一次。
//空间复杂度：O(m)，其中 m 是链表 headA 的长度。需要使用哈希集合存储链表 headA 中的全部节点。
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

//方法二：双指针-官方-不容易想到
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/solution/intersection-of-two-linked-lists-shuang-zhi-zhen-l/
//考虑构建两个节点指针 A​ , B 分别指向两链表头节点 headA , headB ，做如下操作：

//指针 A 先遍历完链表 headA ，再开始遍历链表 headB ，当走到 node 时，共走步数为：
//a + (b - c)

//指针 B 先遍历完链表 headB ，再开始遍历链表 headA ，当走到 node 时，共走步数为：
//b + (a - c)

//如下式所示，此时指针 A , B 重合，并有两种情况：
//a + (b - c) = b + (a - c)

//若两链表 有 公共尾部 (即 c > 0c>0 ) ：指针 A , B 同时指向「第一个公共节点」node 。
//若两链表 无 公共尾部 (即 c = 0c=0 ) ：指针 A , B 同时指向 nullnull 。
//因此返回 A 即可。

//时间复杂度：O(m+n)，其中 m 和 n 是分别是链表 headA 和 headB 的长度。两个指针同时遍历两个链表，
// 每个指针遍历两个链表各一次。
//空间复杂度：O(1)。
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa // 要么是公共节点头，要么是nil
}

//方法3：双指针-理解起来容易
//求出两个链表的长度，并求出两个链表长度的差值，然后让curA移动到，和curB 末尾对齐的位置，
// 此时我们就可以比较curA和curB是否相同，如果不相同，同时向后移动curA和curB，如果遇到curA == curB，
// 则找到交点。 否则循环退出返回空指针。
//时间复杂度：$O(n + m)
//空间复杂度：$O(1)
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}

	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func main() {
	var node *ListNode
	ListA := NewList()

	cur := ListA
	for i := 1; i <= 3; i++ {
		node = &ListNode{Val: i}
		cur.Next = node
		cur = cur.Next
	}
	display(ListA)

	ListB := NewList()
	cur = ListB
	for i := 5; i <= 7; i++ {
		node = &ListNode{Val: i}
		cur.Next = node
		cur = cur.Next
	}
	cur.Next = ListA.Next
	display(ListB)

	//// *****特别注意******
	//// func里面只处理了逻辑上的链表，入参和出参的时候要特殊处理下
	ListAA := NewList()
	ListAA.Next = getIntersectionNode1(ListA.Next, ListB.Next)
	display(ListAA)

	ListBB := NewList()
	ListBB.Next = getIntersectionNode2(ListA.Next, ListB.Next)
	display(ListBB)

	ListCC := NewList()
	ListCC.Next = getIntersectionNode3(ListA.Next, ListB.Next)
	display(ListCC)
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
