package main

import "container/list"

//LeetCode 637.二叉树的层平均值

//给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

//思路
//就是层序遍历的时候把一层求个总和在取一个均值。

// 方法一：深度优先搜索-官方
//时间复杂度：O(n)，其中 n 是二叉树中的节点个数。
// 深度优先搜索需要对每个节点访问一次，对于每个节点，维护两个数组的时间复杂度都是 O(1)，因此深度优先搜索的时间复杂度
// 是 O(n)。
// 遍历结束之后计算每层的平均值的时间复杂度是 O(h)，其中 h 是二叉树的高度，任何情况下都满足 h≤n。
// 因此总时间复杂度是 O(n)。
//空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度取决于两个数组的大小和递归调用的层数，两个数组的大小都等于
// 二叉树的高度，递归调用的层数不会超过二叉树的高度，最坏情况下，二叉树的高度等于节点个数。

type data struct{ sum, count int }

func averageOfLevels1(root *TreeNode) []float64 {
	levelData := []data{}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(levelData) {
			levelData[level].sum += node.Val
			levelData[level].count++
		} else {
			levelData = append(levelData, data{node.Val, 1})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)

	averages := make([]float64, len(levelData))
	for i, d := range levelData {
		averages[i] = float64(d.sum) / float64(d.count)
	}
	return averages
}

// 方法二：广度优先搜索-官方
//时间复杂度：O(n)，其中 n 是二叉树中的节点个数。
// 广度优先搜索需要对每个节点访问一次，时间复杂度是 O(n)。
// 需要对二叉树的每一层计算平均值，时间复杂度是 O(h)，其中 h 是二叉树的高度，任何情况下都满足 h≤n。
// 因此总时间复杂度是 O(n)。
//空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度取决于队列开销，队列中的节点个数不会超过 n。

func averageOfLevels(root *TreeNode) (averages []float64) {
	nextLevel := []*TreeNode{root}
	for len(nextLevel) > 0 {
		sum := 0
		curLevel := nextLevel
		nextLevel = nil
		for _, node := range curLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(len(curLevel)))
	}
	return
}

// 广度优先-卡尔版本
func averageOfLevels3333(root *TreeNode) []float64 {
	res := [][]int{}
	var finRes []float64
	if root == nil { //防止为空
		return finRes
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
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
	//计算每层的平均值
	length := len(res)
	for i := 0; i < length; i++ {
		var sum int
		for j := 0; j < len(res[i]); j++ {
			sum += res[i][j]
		}
		tmp := float64(sum) / float64(len(res[i]))
		finRes = append(finRes, tmp) //将平均值放入结果集合
	}
	return finRes
}
