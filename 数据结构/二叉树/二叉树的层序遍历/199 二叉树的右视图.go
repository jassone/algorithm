package main

import "container/list"

//LeetCode 199.二叉树的右视图

// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

//思路
// 即看到每一层最右边的节点
// 层序遍历的时候，判断是否遍历到单层的最后面的元素，如果是，就放进result数组中，随后返回result就可以了。

//都一样
//时间复杂度 : (n)。 每个节点最多进队列一次，出队列一次，因此广度优先搜索的复杂度为线性。
//空间复杂度 : O(n)。每个节点最多进队列一次，所以队列长度最大不不超过 n，所以这里的空间代价为 O(n)。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：广度优先
func rightSideView1(root *TreeNode) []int {
	arr := []*TreeNode{root}
	res := []int{}

	if root == nil {
		return res
	}
	for len(arr) > 0 {
		size := len(arr)
		res = append(res, arr[size-1].Val)
		for i := 0; i < size; i++ {
			node := arr[i]

			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}
		arr = arr[size:]
	}
	return res
}

// 卡尔版本
// 先取每一层节点，再取每一层最后一个节点
func rightSideView2(root *TreeNode) []int {
	queue := list.New()
	res := [][]int{}
	var finaRes []int
	if root == nil {
		return finaRes
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
	//取每一层的最后一个元素
	for i := 0; i < len(res); i++ {
		finaRes = append(finaRes, res[i][len(res[i])-1])
	}
	return finaRes
}

// 其他方法：广度优先
func rightSideView2222(root *TreeNode) (res []int) {
	queue := make([]*TreeNode, 0)
	if root == nil {
		return
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		first := true
		for size := len(queue); size != 0; size-- {
			pop := queue[0]
			queue = queue[1:]
			if first {
				res = append(res, pop.Val)
				first = false
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
		}
	}
	return
}

// 其他方法：深度优先
func rightSideView33(root *TreeNode) (res []int) {
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(res) < depth {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 1)
	return
}

func main() {

}
