package main

import "container/list"

// todo
//LeetCode 107.二叉树的层次遍历 II

//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

// 方法一：广度优先搜索-官方
// 其实就是和102一样，最后多反转一下
//时间复杂度：O(n)，其中 n 是二叉树中的节点个数。每个节点访问一次，结果列表使用链表的结构时，在结果列表头部添加一层节点值
// 的列表的时间复杂度是 O(1)，因此总时间复杂度是 O(n)。
//空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度取决于队列开销，队列中的节点个数不会超过 n。
func levelOrderBottom1(root *TreeNode) [][]int {
	levelOrder := [][]int{}
	if root == nil {
		return levelOrder
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, level)
	}
	for i := 0; i < len(levelOrder)/2; i++ {
		levelOrder[i], levelOrder[len(levelOrder)-1-i] = levelOrder[len(levelOrder)-1-i], levelOrder[i]
	}
	return levelOrder
}

// 卡尔版
func levelOrderBottom2(root *TreeNode) [][]int {
	queue := list.New()
	res := [][]int{}
	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	//反转结果集
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

func main() {

}
