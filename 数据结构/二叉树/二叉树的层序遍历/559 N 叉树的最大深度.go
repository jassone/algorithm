package main

// 同类型 104 559 110
//LeetCode 559.n叉树的最大深度

//给定一个 n 叉树，找到其最大深度。
//最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

//思路：
//依然可以提供递归法和迭代法，来解决这个问题，思路是和二叉树思路一样的。

//迭代法
//依然是层序遍历，

type Node struct {
	Val      int
	Children []*Node
}

// 方法一：深度优先搜索-官方，推荐
// 时间复杂度：O(n)，其中 n 为 N 叉树节点的个数。每个节点在递归中只被遍历一次。
// 空间复杂度：O(height)，其中 height 表示 N 叉树的高度。递归函数需要栈空间，而栈空间取决于递归的深度，
// 因此空间复杂度等价于 NN 叉树的高度。
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		if childDepth := maxDepth1(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}

// 方法二：广度优先搜索-官方
// 时间复杂度：O(n)，其中 n 为 N 叉树的节点个数。与方法一同样的分析，每个节点只会被访问一次。
// 空间复杂度：此方法空间的消耗取决于队列存储的元素数量，其在最坏情况下会达到 O(n)。
func maxDepth2(root *Node) (ans int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, node := range q {
			queue = append(queue, node.Children...)
		}
		ans++
	}
	return
}
