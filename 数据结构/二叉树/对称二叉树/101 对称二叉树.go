package 对称二叉树

//LeetCode 101. 对称二叉树

// 给定一个二叉树，检查它是否是镜像对称的。

//首先想清楚，判断对称二叉树要比较的是哪两个节点，要比较的可不是左右节点！
//对于二叉树是否对称，要比较的是根节点的左子树与右子树是不是相互翻转的，理解这一点就知道了其实我们要比较的是两个树
// （这两个树是根节点的左右子树），所以在递归遍历的过程中，也是要同时遍历两棵树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//本题遍历只能是“后序遍历”，因为我们要通过递归函数的返回值来判断两个子树的内侧节点和外侧节点是否相等。

//正是因为要遍历两棵树而且要比较内侧和外侧节点，所以准确的来说是一个树的遍历顺序是左右中，一个树的遍历顺序是右左中。

//但都可以理解算是后序遍历，尽管已经不是严格上在一个树上进行遍历的后序遍历了。

//其实后序也可以理解为是一种回溯，当然这是题外话，讲回溯的时候会重点讲的。

// 思路1：递归法

// 思路2：迭代法
//这道题目我们也可以使用迭代法，但要注意，这里的迭代法可不是前中后序的迭代写法，因为本题的本质是判断两个树是否是相互翻转的，其实已经不是所谓二叉树遍历的前中后序的关系了。

//这里我们可以使用队列来比较两个树（根节点的左右子树）是否相互翻转，（注意这不是层序遍历）

// 2-1使用队列

// 2-2使用栈
//细心的话，其实可以发现，这个迭代法，其实是把左右两个子树要比较的元素顺序放进一个容器，然后成对成对的取出来进行比较，那么其实使用栈也是可以的。

//只要把队列原封不动的改成栈就可以了，

// 官方 && 卡尔-递归，推荐
//假设树上一共 n 个节点。
//时间复杂度：这里遍历了这棵树，渐进时间复杂度为 O(n)。
//空间复杂度：这里的空间复杂度和递归使用的栈空间有关，这里递归层数不超过 n，故渐进空间复杂度为 O(n)。
func isSymmetric33(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}
func defs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs(left.Left, right.Right) && defs(right.Left, left.Right)
}

// 迭代法-官方
//时间复杂度：O(n)，同「方法一」。
//空间复杂度：这里需要用一个队列来维护节点，每个节点最多进队一次，出队一次，队列中最多不会超过 n 个点，故渐进空间复杂度为 O(n)。
func isSymmetric2(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}

// 卡尔-迭代
func isSymmetric44(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, right.Left, left.Right)
	}
	return true
}

// 总结
//这道题目的本质是要比较两个树（这两个树是根节点的左右子树），遍历两棵树而且要比较内侧和外侧节点，
// 所以准确的来说是一个树的遍历顺序是左右中，一个树的遍历顺序是右左中。
//而本题的迭代法中我们使用了队列，需要注意的是这不是层序遍历，而且仅仅通过一个容器来成对的存放我们
// 要比较的元素，认识到这一点之后就发现：用队列，用栈，甚至用数组，都是可以的。
