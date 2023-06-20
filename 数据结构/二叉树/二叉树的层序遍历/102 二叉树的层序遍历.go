package main

import (
	"container/list"
	"fmt"
)

//LeetCode 102.二叉树的层序遍历

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//      3
// 9        20
//        15  7

// https://www.programmercarl.com/0102.二叉树的层序遍历.html

//接下来我们再来介绍二叉树的另一种遍历方式：层序遍历。
// 层序遍历一个二叉树。就是从左到右一层一层的去遍历二叉树。这种遍历的方式和我们之前讲过的都不太一样。
// 需要借用一个辅助数据结构即队列来实现，*****队列先进先出，符合一层一层遍历的逻辑，而是用栈先进后出适合
//  模拟深度优先遍历也就是递归的逻辑。*****
// 而这种层序遍历方式就是图论中的广度优先遍历，只不过我们应用在二叉树上。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：递归写法，推荐
// 时间复杂度：每个点进队出队各一次，故渐进时间复杂度为 O(n)。
// 空间复杂度：队列中元素的个数不超过 nn 个，故渐进空间复杂度为 O(n)
var res [][]int

func levelOrder1(root *TreeNode) [][]int {
	// 1.递归函数及参数
	// 从0开始遍历
	traversal(root, 0)
	return res
}
func traversal(node *TreeNode, depth int) {
	// 2.递归终止条件
	if node == nil {
		return
	}
	// 3.递归逻辑
	// 初始化，depth个元素
	if len(res) == depth {
		res = append(res, []int{})
	}
	res[depth] = append(res[depth], node.Val)
	// 遍历节点的左右子节点，depth+1
	if node.Left != nil {
		traversal(node.Left, depth+1)
	}
	if node.Right != nil {
		traversal(node.Right, depth+1)
	}
}

// 方法二：广度优先搜索-官方
// *****出队列的一种方式是重新赋值给队列*****

//记树上所有节点的个数为 n。
// 时间复杂度：每个点进队出队各一次，故渐进时间复杂度为 O(n)。
// 空间复杂度：队列中元素的个数不超过 nn 个，故渐进空间复杂度为 O(n)
func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}        // 当前需要循环的队列(current)
	for i := 0; len(q) > 0; i++ { // 随着循环进行，q会越来越大，直到q为空数组
		ret = append(ret, []int{}) // 每一层执行一次，加一个一维数组字段
		p := []*TreeNode{}         // 存储最新temp临时队列
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val) // 取当前层的值

			// 将当前层的所有节点的左右节点都存进去
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

// 卡尔版本
func levelOrder22(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
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
	return res
}

func main() {
	node := &TreeNode{Val: 5}
	node.Left = &TreeNode{Val: 4}
	node.Left.Left = &TreeNode{Val: 1}
	node.Left.Right = &TreeNode{Val: 2}

	node.Right = &TreeNode{Val: 6}
	node.Right.Left = &TreeNode{Val: 7}
	node.Right.Right = &TreeNode{Val: 8}

	fmt.Println(levelOrder1(node))
	fmt.Println(levelOrder2(node))
}
