package main

import "container/list"

// 同类型 104 559 110
//LeetCode 104.二叉树的最大深度

//给定一个二叉树，找出其最大深度。
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//说明: 叶子节点是指没有子节点的节点。

//思路：
//使用迭代法的话，使用层序遍历是最为合适的，因为最大的深度就是二叉树的层数，和层序遍历的方式极其吻合。

// 1 递归法
// 本题可以使用前序（中左右），也可以使用后序遍历（左右中），使用前序求的就是深度，使用后序求的是高度。

// 而根节点的高度就是二叉树的最大深度，所以本题中我们通过后序求的根节点高度来求的二叉树最大深度。

// 2 迭代法
// 使用迭代法的话，使用层序遍历是最为合适的，因为最大的深度就是二叉树的层数，和层序遍历的方式极其吻合。

// 在二叉树中，一层一层的来遍历二叉树，记录一下遍历的层数就是二叉树的深度，

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：深度优先搜索-官方
//时间复杂度：O(n)，其中 n 为二叉树节点的个数。每个节点在递归中只被遍历一次。
//空间复杂度：O(height)，其中 height 表示二叉树的高度。递归函数需要栈空间，而栈空间取决于递归的深度
// 因此空间复杂度等价于二叉树的高度。
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法二：广度优先搜索-官方
//时间复杂度：O(n)，其中 n 为二叉树的节点个数。与方法一同样的分析，每个节点只会被访问一次。
//空间复杂度：此方法空间的消耗取决于队列存储的元素数量，其在最坏情况下会达到 O(n)。
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

// 卡尔版本
func maxDepth333(root *TreeNode) int {
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
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans++ //记录深度，其他的是层序遍历的板子
	}
	return ans
}

// 卡尔总结
//本题可以使用前序，也可以使用后序遍历（左右中），使用前序求的就是深度，使用后序呢求的是高度。
// 使用了前序（中左右）的遍历顺序，这才是真正求深度的逻辑！

//而根节点的高度就是二叉树的最大深度，所以本题中我们通过后序求的根节点高度来求的二叉树最大深度，
// 所以当前题目中使用的是后序遍历。
