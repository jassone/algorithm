package main

import "container/list"

//LeetCode 404.左叶子之和
//计算给定二叉树的所有左叶子之和。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一：深度优先搜索-官方
//时间复杂度：O(n)，其中 n 是树中的节点个数。
//空间复杂度：O(n)。空间复杂度与深度优先搜索使用的栈的最大深度相关。在最坏的情况下，树呈现链式结构，深度为 O(n)，对应的空间复杂度也为 O(n)。
func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func dfs(node *TreeNode) (ans int) {
	if node.Left != nil {
		if isLeafNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeafNode(node.Right) {
		ans += dfs(node.Right)
	}
	return
}
func sumOfLeftLeaves1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

//方法二：广度优先搜索-官方
//时间复杂度：O(n)，其中 n 是树中的节点个数。
//空间复杂度：O(n)。空间复杂度与广度优先搜索使用的队列需要的容量相关，为 O(n)。
func isLeafNode2(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
func sumOfLeftLeaves2(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			if isLeafNode2(node.Left) {
				ans += node.Left.Val
			} else {
				q = append(q, node.Left)
			}
		}
		if node.Right != nil && !isLeafNode2(node.Right) {
			q = append(q, node.Right)
		}
	}
	return
}

//卡尔
//思路
//首先要注意是判断左叶子，不是二叉树左侧节点，所以不要上来想着层序遍历。

//因为题目中其实没有说清楚左叶子究竟是什么节点，那么我来给出左叶子的明确定义：如果左节点不为空，
// 且左节点没有左右孩子，那么这个节点的左节点就是左叶子

//那么判断当前节点是不是左叶子是无法判断的，必须要通过节点的父节点来判断其左孩子是不是左叶子。
//如果该节点的左节点不为空，该节点的左节点的左节点为空，该节点的左节点的右节点为空，则找到了一个左叶子，

//递归法
//递归的遍历顺序为后序遍历（左右中），是因为要通过递归函数的返回值来累加求取左叶子数值之和。。
func sumOfLeftLeaves11(root *TreeNode) int {
	var res int
	findLeft1(root, &res)
	return res
}
func findLeft1(root *TreeNode, res *int) {
	//左节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*res = *res + root.Left.Val
	}
	if root.Left != nil {
		findLeft1(root.Left, res)
	}
	if root.Right != nil {
		findLeft1(root.Right, res)
	}
}

//迭代法
//本题迭代法使用前中后序都是可以的，只要把左叶子节点统计出来，就可以了。
func sumOfLeftLeaves22(root *TreeNode) int {
	var res int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
				res = res + node.Left.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}

//总结
//这道题目要求左叶子之和，其实是比较绕的，因为不能判断本节点是不是左叶子节点。

//此时就要通过节点的父节点来判断其左孩子是不是左叶子了。

//平时我们解二叉树的题目时，已经习惯了通过节点的左右孩子判断本节点的属性，而本题我们要通过节点的父节点
// 判断本节点的属性。
