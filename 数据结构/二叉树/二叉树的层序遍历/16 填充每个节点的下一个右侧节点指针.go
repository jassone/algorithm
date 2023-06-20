package main

import "container/list"

// todo
//LeetCode 16.填充每个节点的下一个右侧节点指针

//给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

//struct Node {
//    int val;
//    Node *left;
//    Node *right;
//    Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//初始状态下，所有 next 指针都被设置为 NULL。

//思路:
//本题依然是层序遍历，只不过在单层遍历的时候记录一下本层的头部节点，然后在遍历的时候让前一个节点指向本节点就可以了

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 方法一：层次遍历-官方
//时间复杂度：O(N)。每个节点会被访问一次且只会被访问一次，即从队列中弹出，并建立 next 指针。
//空间复杂度：O(N)。这是一棵完美二叉树，它的最后一个层级包含 N/2 个节点。广度优先遍历的复杂度取决于
// 一个层级上的最大元素数量。这种情况下空间复杂度为 O(N)。
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	// 初始化队列同时将第一层节点加入队列中，即根节点
	queue := []*Node{root}

	// 循环迭代的是层数
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		// 遍历这一层的所有节点
		for i, node := range tmp {
			// 连接
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			// 拓展下一层节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	// 返回根节点
	return root
}

// 方法二：使用已建立的 next 指针 - 官方
//时间复杂度：O(N)，每个节点只访问一次。
//空间复杂度：O(1)，不需要存储额外的节点。
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/solution/tian-chong-mei-ge-jie-dian-de-xia-yi-ge-you-ce-2-4/
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	// 返回根节点
	return root
}

// 卡尔版本
func connect333(root *Node) *Node {
	res := [][]*Node{}
	if root == nil { //防止为空
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []*Node
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []*Node{}        //清空层的数据
	}
	//遍历每层元素,指定next
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i])-1; j++ {
			res[i][j].Next = res[i][j+1]
		}
	}
	return root
}
