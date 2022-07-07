package main

import "container/list"

//LeetCode 117.填充每个节点的下一个右侧节点指针II

//给定一个二叉树
//struct Node {
//	int val;
//	Node *left;
//	Node *right;
//	Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//初始状态下，所有 next 指针都被设置为 NULL。

//思路：
//这道题目说是二叉树，但116题目说是完整二叉树，其实没有任何差别，一样的代码一样的逻辑一样的味道

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//方法一：层次遍历-官方
//记树上的点的个数为 N。
//时间复杂度：O(N)。我们需要遍历这棵树上所有的点，时间复杂度为 O(N)。
//空间复杂度：O(N)。即队列的空间代价。
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

// 方法二：使用已建立的 next 指针-官方
//时间复杂度：O(N)。分析同「方法一」。
//空间复杂度：O(1)。
func connect2(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}

// 卡尔版本

func connect3(root *Node) *Node {
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
