package main

// leetcode  1171. 从链表中删去总和值为零的连续节点

// 给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。
//
//删除完毕后，请你返回最终结果链表的头节点。

//你可以返回任何满足题目要求的答案。
//
//（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）
//
//示例 1：
//输入：head = [1,2,-1,-1,1]
//输出：[1,1]

type ListNode struct {
	Val  int
	Next *ListNode
}

// 一次遍历,推荐
func removeZeroSumSublists2(head *ListNode) *ListNode {
	m := make(map[int]*ListNode) // m记录：key为遍历过的节点Val之和, value为当前节点的地址
	dummy := &ListNode{          // 假头
		Val:  0,
		Next: head,
	}
	m[0] = dummy
	total := 0
	for cur := head; cur != nil; cur = cur.Next {
		total += cur.Val
		if appear, ok := m[total]; ok { // 如果遍历过的节点之和之前出现过，那么则出现了和为0的段，那就删掉喽
			prevTotal := total
			for p := appear.Next; p != cur; p = p.Next { // 删掉中间段
				prevTotal += p.Val
				delete(m, prevTotal)
			}
			appear.Next = cur.Next
		} else {
			m[total] = cur
		}
	}

	return dummy.Next
}

// 哈希表存前缀和  官方
func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	seen := map[int]*ListNode{}
	prefix := 0

	for node := dummy; node != nil; node = node.Next {
		prefix += node.Val
		seen[prefix] = node
	}

	prefix = 0
	for node := dummy; node != nil; node = node.Next {
		prefix += node.Val
		node.Next = seen[prefix].Next
	}
	return dummy.Next
}
