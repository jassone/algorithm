package main

import "sort"

//LeetCode 222.完全二叉树的节点个数

//给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

//完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
// 并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

//示例 1：
//输入：root = [1,2,3,4,5,6]
//输出：6

//示例 2：
//输入：root = []
//输出：0

//示例 3：
//输入：root = [1]
//输出：1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：二分查找 + 位运算 - 官方
//https://leetcode.cn/problems/count-complete-tree-nodes/solution/wan-quan-er-cha-shu-de-jie-dian-ge-shu-by-leetco-2/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}

// 卡尔
//思路
//本篇给出按照普通二叉树的求法以及利用完全二叉树性质的求法。

//普通二叉树
//首先按照普通二叉树的逻辑来求。

//这道题目的递归法和求二叉树的深度写法类似， 而迭代法，二叉树：层序遍历登场！ (opens new window)遍历模板稍稍修改一下，记录遍历的节点数量就可以了。

//递归遍历的顺序依然是后序（左右中）。

// 方法1：递归
//时间复杂度：O(n)
//空间复杂度：O(log n)，算上了递归系统栈占用的空间

//本题直接就是求有多少个节点，无脑存进数组算长度就行了。
func countNodes11(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}

// 方法2：迭代法
//时间复杂度：O(n)
//空间复杂度：O(n)
// todo

//完全二叉树
//以上方法都是按照普通二叉树来做的。下面看看完全二叉树的情况。

//完全二叉树只有两种情况，情况一：就是满二叉树，情况二：最后一层叶子节点没有满。
// 对于情况一，可以直接用 2^树深度 - 1 来计算，注意这里根节点深度为1。
// 对于情况二，分别递归左孩子，和右孩子，递归到某一深度一定会有左孩子或者右孩子为满二叉树，然后依然可以按照情况1来计算。

// 如果整个树不是满二叉树，就递归其左右孩子，直到遇到满二叉树为止，用公式计算这个子树（满二叉树）的节点数量。

//时间复杂度：O(log n × log n)
//空间复杂度：O(log n)
func countNodes22(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}
	if leftH == rightH {
		return (2 << leftH) - 1
	}
	return countNodes22(root.Left) + countNodes22(root.Right) + 1
}
