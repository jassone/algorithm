package main

import "container/list"

//LeetCode 226.翻转二叉树

//翻转一棵二叉树。

//题外话
//这道题目是非常经典的题目，也是比较简单的题目（至少一看就会）。
//但正是因为这道题太简单，一看就会，一些同学都没有抓住起本质，稀里糊涂的就把这道题目过了。
//如果做过这道题的同学也建议认真看完，相信一定有所收获！

//可以发现想要翻转它，其实就把每一个节点的左右孩子交换一下就可以了。
//关键在于遍历顺序，前中后序应该选哪一种遍历顺序？ （一些同学这道题都过了，但是不知道自己用的是什么顺序）
//遍历的过程中去翻转每一个节点的左右孩子就可以达到整体翻转的效果。
//注意只要把每一个节点的左右孩子翻转一下，就可以达到整体翻转的效果
//这道题目使用前序遍历和后序遍历都可以，唯独中序遍历不方便(递归的中序遍历是不行的，因为使用递归的中序遍历，某些节点的左右孩子会翻转两次。)，
// 因为中序遍历会把某些节点的左右孩子翻转了两次！但是用栈来遍历(迭代)的中序遍历是可以的，而不是靠指针来遍历，避免了递归法中翻转了两次的情况，
//那么层序遍历可以不可以呢？依然可以的！只要把每一个节点的左右孩子翻转一下的遍历方式都是可以的！

//二叉树解题的大忌就是自己稀里糊涂的过了（因为这道题相对简单），但是也不知道自己是怎么遍历的。
//这也是造成了二叉树的题目“一看就会，一写就废”的原因。
//针对翻转二叉树，我给出了一种递归，三种迭代（两种模拟深度优先遍历，一种层序遍历）的写法，都是之前我们讲过的写法，融汇贯通一下而已。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：递归-官方-后序遍历，推荐
// 时间复杂度：O(N)，其中 N 为二叉树节点的数目。我们会遍历二叉树中的每一个节点，对每个节点而言，我们在常数时间
// 内交换其两棵子树。
// 空间复杂度：O(N)。使用的空间由递归栈的深度决定，它等于当前节点在二叉树中的高度。在平均情况下，二叉树的高度与
// 节点个数为对数关系，即 O(logN)。而在最坏情况下，树形成链状，空间复杂度为 O(N)。
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree1(root.Left)
	right := invertTree1(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 递归版本的前序遍历-卡尔
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree2(root.Left)
	invertTree2(root.Right)

	return root
}

// 迭代版本的前序遍历-卡尔
func invertTree3(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left //交换
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return root
}

// 迭代版本的后序遍历-卡尔
func invertTree4(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			node.Left, node.Right = node.Right, node.Left //交换
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return root
}

// 层序遍历 - 卡尔
func invertTree5(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			e.Left, e.Right = e.Right, e.Left //交换
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root
}
