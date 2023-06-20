package main

//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// “对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x
// 的深度尽可能大（一个节点也可以是它自己的祖先）。”

//输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出：3
//解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

// 代码随想录
// https://www.programmercarl.com/0235.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88.html

// 方法1：递归-官方-推荐
// 采用递归回溯的方法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil // 传入的节点本身就是空的，比如根节点为空，或者子节点最后一个节点的left或者right都是为空的
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root // 这时候表示找到了其中一个节点
	}
	// 后续遍历
	// 左
	left := lowestCommonAncestor(root.Left, p, q)
	// 右
	right := lowestCommonAncestor(root.Right, p, q)
	// 下面都是中
	if left != nil && right != nil {
		return root // 说明找到了，则往上回溯
	}
	// 下面两行表示如果left或right不为空则返回其中一个，如果都为空则返回空
	if left == nil {
		return right // 这里right也可能为空
	}
	return left //表示left下面有符合条件的节点
}

//方法二：存储父节点-官方
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}
