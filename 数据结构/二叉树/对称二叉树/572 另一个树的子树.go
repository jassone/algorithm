package main

import "math"

//LeetCode 572.另一个树的子树

//给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。
// 如果存在，返回 true ；否则，返回 false 。

//二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：转换为两个节点是否相等的问题

// 方法一：深度优先搜索暴力匹配-官方
//思路和算法
//这是一种最朴素的方法——深度优先搜索枚举 ss 中的每一个节点，判断这个点的子树是否和 tt 相等。
// 如何判断一个节点的子树是否和 tt 相等呢，我们又需要做一次深度优先搜索来检查，即让两个指针一开始先
// 指向该节点和 tt 的根，然后「同步移动」两根指针来「同步遍历」这两棵树，判断对应位置是否相等。
// https://leetcode.cn/problems/subtree-of-another-tree/solution/ling-yi-ge-shu-de-zi-shu-by-leetcode-solution/
func isSubtree1(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree1(s.Left, t) || isSubtree1(s.Right, t)
}
func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}

// 方法二：深度优先搜索序列上做串匹配-官方-太难了
func isSubtree2(s *TreeNode, t *TreeNode) bool {
	maxEle := math.MinInt32
	getMaxElement(s, &maxEle)
	getMaxElement(t, &maxEle)
	lNull := maxEle + 1
	rNull := maxEle + 2

	sl, tl := getDfsOrder(s, []int{}, lNull, rNull), getDfsOrder(t, []int{}, lNull, rNull)
	return kmp(sl, tl)
}
func kmp(s, t []int) bool {
	sLen, tLen := len(s), len(t)
	fail := make([]int, sLen)
	for i := 0; i < sLen; i++ {
		fail[i] = -1
	}
	for i, j := 1, -1; i < tLen; i++ {
		for j != -1 && t[i] != t[j+1] {
			j = fail[j]
		}
		if t[i] == t[j+1] {
			j++
		}
		fail[i] = j
	}

	for i, j := 0, -1; i < sLen; i++ {
		for j != -1 && s[i] != t[j+1] {
			j = fail[j]
		}
		if s[i] == t[j+1] {
			j++
		}
		if j == tLen-1 {
			return true
		}
	}
	return false
}
func getDfsOrder(t *TreeNode, list []int, lNull, rNull int) []int {
	if t == nil {
		return list
	}
	list = append(list, t.Val)
	if t.Left != nil {
		list = getDfsOrder(t.Left, list, lNull, rNull)
	} else {
		list = append(list, lNull)
	}

	if t.Right != nil {
		list = getDfsOrder(t.Right, list, lNull, rNull)
	} else {
		list = append(list, rNull)
	}
	return list
}
func getMaxElement(t *TreeNode, maxEle *int) {
	if t == nil {
		return
	}
	if t.Val > *maxEle {
		*maxEle = t.Val
	}
	getMaxElement(t.Left, maxEle)
	getMaxElement(t.Right, maxEle)
}
