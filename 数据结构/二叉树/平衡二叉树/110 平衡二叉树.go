package main

// 同类型 104 559 110
//LeetCode 110.平衡二叉树

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

//示例 1:
//给定二叉树 [3,9,20,null,null,15,7]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 因为求深度可以从上到下去查 所以需要前序遍历（中左右），而高度只能从下到上去查，所以只能后序遍历（左右中）

//为什么104.二叉树的最大深度中求的是二叉树的最大深度，也用的是后序遍历。
//那是因为代码的逻辑其实是求的根节点的高度，而根节点的高度就是这棵树的最大深度，所以才可以使用后序遍历。
//在104.二叉树的最大深度中，如果真正求取二叉树的最大深度，代码应该写成前序遍历。

//方法一：自顶向下的递归-官方
//https://leetcode.cn/problems/balanced-binary-tree/solution/ping-heng-er-cha-shu-by-leetcode-solution/
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs1(height1(root.Left)-height1(root.Right)) <= 1 && isBalanced1(root.Left) && isBalanced1(root.Right)
}
func height1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max1(height1(root.Left), height1(root.Right)) + 1
}
func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs1(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 方法二：自底向上的递归-官方
//时间复杂度：O(n)，其中 n 是二叉树中的节点个数。使用自底向上的递归，每个节点的计算高度和判断是否平衡都只
// 需要处理一次，最坏情况下需要遍历二叉树中的所有节点，因此时间复杂度是 O(n)。
//空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度主要取决于递归调用的层数，递归调用的层数不会超过 nn。
func isBalanced2(root *TreeNode) bool {
	return height2(root) >= 0
}
func height2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height2(root.Left)
	rightHeight := height2(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs2(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max2(leftHeight, rightHeight) + 1
}
func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs2(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 卡尔
//本题思路
// 1 递归
// 此时大家应该明白了既然要求比较高度，必然是要后序遍历。

// 2 迭代
//在104.二叉树的最大深度 (opens new window)中我们可以使用层序遍历来求深度，但是就不能直接用层序遍历来求高度了，这就体现出求高度和求深度的不同。
//本题的迭代方式可以先定义一个函数，专门用来求高度。
//这个函数通过栈模拟的后序遍历找每一个节点的高度（其实是通过求传入节点为根节点的最大深度来求的高度）

func isBalanced22(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced22(root.Left) || !isBalanced22(root.Right) {
		return false
	}
	LeftH := maxdepth(root.Left) + 1
	RightH := maxdepth(root.Right) + 1
	if abs(LeftH-RightH) > 1 {
		return false
	}
	return true
}
func maxdepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxdepth(root.Left), maxdepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//当然此题用迭代法，其实效率很低，因为没有很好的模拟回溯的过程，所以迭代法有很多重复的计算。
//虽然理论上所有的递归都可以用迭代来实现，但是有的场景难度可能比较大。
//例如：都知道回溯法其实就是递归，但是很少人用迭代的方式去实现回溯算法！
//因为对于回溯算法已经是非常复杂的递归了，如果在用迭代的话，就是自己给自己找麻烦，效率也并不一定高。

//总结
//通过本题可以了解求二叉树深度 和 二叉树高度的差异，求深度适合用前序遍历，而求高度适合用后序遍历。
//本题迭代法其实有点复杂，大家可以有一个思路，也不一定说非要写出来。
//但是递归方式是一定要掌握的！

//讲了这么多二叉树题目的迭代法，有的同学会疑惑，迭代法中究竟什么时候用队列，什么时候用栈？
//如果是模拟前中后序遍历就用栈，如果是适合层序遍历就用队列，当然还是其他情况，那么就是 先用队列试试行不行，不行就用栈。
