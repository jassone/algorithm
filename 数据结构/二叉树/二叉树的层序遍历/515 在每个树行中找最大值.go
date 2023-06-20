package main

import (
	"container/list"
	"math"
)

// todo
//LeetCode 515.在每个树行中找最大值

// 要在二叉树的每一行中找到最大的值。

//思路：
//层序遍历，取每一层的最大值

// 方法1：深度优先遍历
//解题思路
// 定义递归函数 traverse
// 确定终止条件，也就是当前遍历的这个节点是空，就直接 return
// 确定每一次递归的执行逻辑，定义一个 depth 记录层数，如果当前层数对应的 res 有值就比较两者大小，
//  没有就把当前 val 当作最大值添加进去
// 向下递归遍历左右节点，直到结束

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues1(root *TreeNode) []int {
	// 结束条件
	if root == nil {
		return nil
	}
	//迭代
	result := make([]int, 0)
	//通过queue对层次数据的维护
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		queueLength := len(queue)
		max := math.MinInt64
		for i := 0; i < queueLength; i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}
			// 在遍历本层的基础上，同时把下一层的节点记录到queue中
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// 把每一层的最大值，添加到result
		result = append(result, max)
		// 更新queue，保证数据为下一层所需的节点数据
		queue = queue[queueLength:]
	}
	return result
}

// 方法2：广度优先遍历（层序遍历）
//解题思路：
// 定义一个队列 queue 来存放节点
// 遍历队列，挨个节点从头部取出
// 记录取出的节点数值 val 中的最大值 max，如果有左右节点再加入队列，供下一次遍历使用
// 最后队列长度为 0 即遍历结束
func largestValues2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	dfsLargestValues(root, &res, 0)
	return res
}
func dfsLargestValues(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, math.MinInt32)
	}
	(*res)[level] = int(math.Max(float64((*res)[level]), float64(root.Val)))
	if root.Left != nil {
		dfsLargestValues(root.Left, res, level+1)
	}
	if root.Right != nil {
		dfsLargestValues(root.Right, res, level+1)
	}
}

//卡尔版
func largestValues333(root *TreeNode) []int {
	res := [][]int{}
	var finRes []int
	if root == nil { //防止为空
		return finRes
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	//层次遍历
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据
	}
	//找到每层的最大值
	for i := 0; i < len(res); i++ {
		finRes = append(finRes, max(res[i]...))
	}
	return finRes
}
func max(vals ...int) int {
	max := int(math.Inf(-1)) //负无穷
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
