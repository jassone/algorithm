package main

import (
	"container/list"
	"math"
)

//LeetCode 111.二叉树的最小深度

//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

//说明: 叶子节点是指没有子节点的节点。

//相对于 104.二叉树的最大深度 ，本题还也可以使用层序遍历的方式来解决，思路是一样的。
//需要注意的是，只有当左右孩子都为空的时候，才说明遍历的最低点了。如果其中一个孩子为空则不是最低点

//看完了这篇104.二叉树的最大深度 (opens new window)，再来看看如何求最小深度。
//直觉上好像和求最大深度差不多，其实还是差不少的。
//遍历顺序上依然是后序遍历（因为要比较递归返回之后的结果），但在处理中间节点的逻辑上，最大深度很容易理解，
// 最小深度可有一个误区，即第一个元素不是最小深度哦。

//这就重新审题了，题目中说的是：最小深度是从根节点到最近叶子节点的最短路径上的节点数量。，注意是叶子节点。
//什么是叶子节点，左右孩子都为空的节点才是叶子节点！

// 求二叉树的最小深度和求二叉树的最大深度的差别主要在于处理左右孩子不为空的逻辑。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：深度优先搜索-官方 && 卡尔，推荐
//时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
//空间复杂度：O(H)，其中 H 是树的高度。空间复杂度主要取决于递归时栈空间的开销，最坏情况下，树呈现链状，
// 空间复杂度为 O(N)。平均情况下树的高度与节点数的对数正相关，空间复杂度为 O(logN)。
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth1(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth1(root.Right), minD)
	}
	return minD + 1
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 方法二：广度优先搜索-官方
// 需要注意的是，只有当左右孩子都为空的时候，才说明遍历的最低点了。如果其中一个孩子为空则不是最低点

//时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
//空间复杂度：O(N)，其中 N 是树的节点数。空间复杂度主要取决于队列的开销，队列中的元素个数不会超过树的节点数。
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}

// 卡尔版
func minDepth333(root *TreeNode) int {
	ans := 0
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil { //当前节点没有左右节点，则代表此层是最小层
				return ans + 1 //返回当前层 ans代表是上一层
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans++ //记录层数

	}
	return ans + 1
}

// 卡尔总结
//注意这里最小深度是从根节点到最近叶子节点的最短路径上的节点数量。注意是叶子节点。

//什么是叶子节点，左右孩子都为空的节点才是叶子节点！

//求二叉树的最小深度和求二叉树的最大深度的差别主要在于处理左右孩子不为空的逻辑。

//注意到这一点之后 递归法和迭代法 都可以参照104题目写出来。
