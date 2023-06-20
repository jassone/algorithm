package main

import (
	"container/list"
	"fmt"
)

//LeetCode 144.二叉树的前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：递归-官方
//时间复杂度：O(n)，其中 n 是二叉树的节点数。每一个节点恰好被遍历一次。
//空间复杂度：O(n)，为递归过程中栈的开销，平均情况下为 O(logn)，最坏情况下树呈现链状，为O(n)。
func preorderTraversal1(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)

	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

//方法二：迭代
//思路
//递归的实现就是：每一次递归调用都会把函数的局部变量、参数值和返回地址等压入调用栈中，然后递归返回的时候，
// 从栈顶弹出上一次递归的各项参数，所以这就是递归为什么可以返回上一层位置的原因。
//此时大家应该知道我们用栈也可以是实现二叉树的前后中序遍历了。

// ****两种方式是等价的，区别在于递归的时候隐式地维护了一个栈，而我们在迭代的时候需要显式地将
// 这个栈模拟出来，其他都相同。****

//前序遍历是中左右，每次先处理的是中间节点，那么先将根节点放入栈中，然后将右孩子加入栈，再加入左孩子。
//为什么要先加入右孩子，再加入左孩子呢？ 因为这样出栈的时候才是中左右的顺序。

// *****要点：要访问的元素和要处理的元素顺序是一致的，都是中间节点。*****

//时间复杂度：O(n)，其中 n 是二叉树的节点数。每一个节点恰好被遍历一次。
//空间复杂度：O(n)，为递归过程中栈的开销，平均情况下为 O(logn)，最坏情况下树呈现链状，为O(n)。

// 官方
func preorderTraversal2(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		// 一次大循环中，弹出一个栈顶数据
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

// 卡尔版本-不建议
func preorderTraversal21(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

//方法三：Morris 遍历-官方 todo
//Morris 遍历的核心思想是利用树的大量空闲指针，实现空间开销的极限缩减。其前序遍历规则总结如下：

//  1新建临时节点，令该节点为 root；
//  2如果当前节点的左子节点为空，将当前节点加入答案，并遍历当前节点的右子节点；
//  3如果当前节点的左子节点不为空，在当前节点的左子树中找到当前节点在中序遍历下的前驱节点：
//    3-1如果前驱节点的右子节点为空，将前驱节点的右子节点设置为当前节点。然后将当前节点加入答案，
//       并将前驱节点的右子节点更新为当前节点。当前节点更新为当前节点的左子节点。
//    3-2如果前驱节点的右子节点为当前节点，将它的右子节点重新设为空。当前节点更新为当前节点的右子节点。
// 重复步骤 2 和步骤 3，直到遍历结束。

//这样我们利用 Morris 遍历的方法，前序遍历该二叉树，即可实现线性时间与常数空间的遍历。

//时间复杂度：O(n)，其中 n 是二叉树的节点数。没有左子树的节点只被访问一次，有左子树的节点被访问两次。
//空间复杂度：O(1)。只操作已经存在的指针（树的空闲指针），因此只需要常数的额外空间。
func preorderTraversal3(root *TreeNode) (vals []int) {
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				vals = append(vals, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			vals = append(vals, p1.Val)
		}
		p1 = p1.Right
	}
	return
}

func main() {
	node := &TreeNode{Val: 5}
	node.Left = &TreeNode{Val: 4}
	node.Left.Left = &TreeNode{Val: 1}
	node.Left.Right = &TreeNode{Val: 2}

	node.Right = &TreeNode{Val: 6}
	node.Right.Left = &TreeNode{Val: 7}
	node.Right.Right = &TreeNode{Val: 8}

	fmt.Println(preorderTraversal1(node))
	fmt.Println(preorderTraversal2(node))
	fmt.Println(preorderTraversal21(node))
}
