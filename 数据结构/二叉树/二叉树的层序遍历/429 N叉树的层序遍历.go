package main

import "container/list"

// 同类型590 589 429
//LeetCode 429 N 叉树的层序遍历
// 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。

type Node struct {
	Val      int
	Children []*Node
}

//方法一：广度优先搜索，推荐

//思路与算法
//1 对于「层序遍历」的题目，我们一般使用广度优先搜索。在广度优先搜索的每一轮中，我们会遍历同一层的所有节点。
//2 具体地，我们首先把根节点 root 放入队列中，随后在广度优先搜索的每一轮中，我们首先记录下当前队列中包含的节点个数
// （记为 cnt），即表示上一层的节点个数。在这之后，我们从队列中依次取出节点，直到取出了上一层的全部 cnt 个节点为止。
// 当取出节点 cur 时，我们将 cur 的值放入一个临时列表，再将 cur 的所有子节点全部放入队列中。
//3 当这一轮遍历完成后，临时列表中就存放了当前层所有节点的值。这样一来，当整个广度优先搜索完成后，我们就可以得到层序遍历的结果。

//细节
//需要特殊判断树为空的情况。

//时间复杂度：O(n)，其中 n 是树中包含的节点个数。在广度优先搜索的过程中，我们需要遍历每一个节点恰好一次。
//空间复杂度：O(n)，即为队列需要使用的空间。在最坏的情况下，树只有两层，且最后一层有 n−1 个节点，此时就需要 O(n) 的空间。
func levelOrder1(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*Node{root} // 把顶层节点放到队列中
	for q != nil {
		level := []int{} // 定义临时层数组
		tmp := q
		q = nil // 每一层都重置下
		for _, node := range tmp {
			level = append(level, node.Val) // 当前层的所有原始值放进临时层数组中
			q = append(q, node.Children...) // 当前层下的所有子节点放入到队列中
		}
		ans = append(ans, level) //  当前层处理完放入结果数组中
	}
	return
}

// 卡尔版本
func levelOrder(root *Node) [][]int {
	queue := list.New()
	res := [][]int{} //结果集
	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len() //记录当前层的数量
		var tmp []int
		for T := 0; T < length; T++ { //该层的每个元素：一添加到该层的结果集中；
			// 二找到该元素的下层元素加入到队列中，方便下次使用
			myNode := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, myNode.Val)
			for i := 0; i < len(myNode.Children); i++ {
				queue.PushBack(myNode.Children[i])
			}
		}
		res = append(res, tmp)
	}
	return res
}
