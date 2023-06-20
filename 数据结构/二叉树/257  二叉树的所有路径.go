package main

import "strconv"

//LeetCode 257. 二叉树的所有路径

//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//说明: 叶子节点是指没有子节点的节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：深度优先搜索-官方
//思路与算法

//最直观的方法是使用深度优先搜索。在深度优先搜索遍历二叉树时，我们需要考虑当前的节点以及它的孩子节点。
// 1 如果当前节点不是叶子节点，则在当前的路径末尾添加该节点，并继续递归遍历该节点的每一个孩子节点。
// 2 如果当前节点是叶子节点，则在当前路径末尾添加该节点后我们就得到了一条从根节点到叶子节点的路径，将该路径加入到答案即可。
//如此，当遍历完整棵二叉树以后我们就得到了所有从根节点到叶子节点的路径。当然，深度优先搜索也可以使用非递归的方式实现，这里不再赘述。
// https://leetcode.cn/problems/binary-tree-paths/solution/er-cha-shu-de-suo-you-lu-jing-by-leetcode-solution/
var paths []string

func binaryTreePaths1(root *TreeNode) []string {
	paths = []string{}
	constructPaths1(root, "")
	return paths
}
func constructPaths1(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			constructPaths1(root.Left, pathSB)
			constructPaths1(root.Right, pathSB)
		}
	}
}

//方法二：广度优先搜索-官方
//思路与算法

//我们也可以用广度优先搜索来实现。我们维护一个队列，存储节点以及根到该节点的路径。一开始这个队列里只有根节点。
// 在每一步迭代中，我们取出队列中的首节点，如果它是叶子节点，则将它对应的路径加入到答案中。如果它不是叶子节点，
// 则将它的所有孩子节点加入到队列的末尾。当队列为空时广度优先搜索结束，我们即能得到答案。
func binaryTreePaths2(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	nodeQueue := []*TreeNode{}
	pathQueue := []string{}
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}

//卡尔
//思路
//这道题目要求从根节点到叶子的路径，所以需要前序遍历，这样才方便让父节点指向孩子节点，找到对应的路径。

//在这道题目中将第一次涉及到回溯，因为我们要把路径记录下来，需要回溯来回退一一个路径在进入另一个路径。

// 1 递归
//我们先使用递归的方式，来做前序遍历。要知道递归和回溯就是一家的，本题也需要回溯。

//****回溯和递归是一一对应的，有一个递归，就要有一个回溯，这么写的话相当于把递归和回溯拆开了， 一个在花括号里，一个在花括号外。***

//所以回溯要和递归永远在一起，世界上最遥远的距离是你在花括号里，而我在花括号外！

func binaryTreePaths222(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			res = append(res, v)
			return
		}
		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}
	travel(root, "")
	return res
}

// 递归 todo ???
func binaryTreePaths333(root *TreeNode) []string {
	var result []string
	traversal333(root, &result, "")
	return result
}
func traversal333(root *TreeNode, result *[]string, pathStr string) {
	//判断是否为第一个元素
	if len(pathStr) != 0 {
		pathStr = pathStr + "->" + strconv.Itoa(root.Val)
	} else {
		pathStr = strconv.Itoa(root.Val)
	}
	//判断是否为叶子节点
	if root.Left == nil && root.Right == nil {
		*result = append(*result, pathStr)
		return
	}
	//左右
	if root.Left != nil {
		traversal333(root.Left, result, pathStr)
	}
	if root.Right != nil {
		traversal333(root.Right, result, pathStr)
	}
}

// 回溯法 todo
func binaryTreePaths555(root *TreeNode) []string {
	var result []string
	var path []int
	traversal55(root, &result, &path)
	return result
}
func traversal55(root *TreeNode, result *[]string, path *[]int) {
	*path = append(*path, root.Val)
	//判断是否为叶子节点
	if root.Left == nil && root.Right == nil {
		pathStr := strconv.Itoa((*path)[0])
		for i := 1; i < len(*path); i++ {
			pathStr = pathStr + "->" + strconv.Itoa((*path)[i])
		}
		*result = append(*result, pathStr)
		return
	}
	//左右
	if root.Left != nil {
		traversal55(root.Left, result, path)
		*path = (*path)[:len(*path)-1] //回溯到上一个节点（因为traversal会加下一个节点值到path中）
	}
	if root.Right != nil {
		traversal55(root.Right, result, path)
		*path = (*path)[:len(*path)-1] //回溯
	}
}

// 2 迭代法
//至于非递归的方式，我们可以依然可以使用前序遍历的迭代方式来模拟遍历路径的过程。

//这里除了模拟递归需要一个栈，同时还需要一个栈来存放对应的遍历路径。
func binaryTreePaths33(root *TreeNode) []string {
	stack := []*TreeNode{}
	paths := make([]string, 0)
	res := make([]string, 0)
	if root != nil {
		stack = append(stack, root)
		paths = append(paths, "")
	}
	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		path := paths[l-1]
		stack = stack[:l-1]
		paths = paths[:l-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return res
}

//文中我明确的说了：回溯就隐藏在traversal(cur->left, path + "->", result);中的 path + "->"。
// 每次函数调用完，path依然是没有加上"->" 的，这就是回溯了。

//如果还不理解的话，可以把
//traversal(cur->left, path + "->", result);

//改成
//string tmp = path + "->";
//traversal(cur->left, tmp, result);
//看看还行不行了，答案是这么写就不行了，因为没有回溯了。
