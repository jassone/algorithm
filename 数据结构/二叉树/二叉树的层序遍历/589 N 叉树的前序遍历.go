package main

// 同类型590 589 429
//LeetCode 589. N 叉树的前序遍历

//方法一：递归-官方，推荐
//时间复杂度：O(m)，其中 m 为 N 叉树的节点。每个节点恰好被遍历一次。
//空间复杂度：O(m)，递归过程中需要调用栈的开销，平均情况下为 O(logm)，最坏情况下树的深度为 m−1，
// 此时需要的空间复杂度为 O(m)。
func preorder1(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return
}

// 方法二：迭代- 官方
//方法一中利用递归来遍历树，实际的递归中隐式调用了栈，在此我们可以直接模拟递归中栈的调用。在前序遍历中，
// 我们会先遍历节点本身，然后从左向右依次先序遍历该每个以子节点为根的子树。

//时间复杂度：O(m)，其中 m 为 N 叉树的节点。每个节点恰好被访问一次。
//空间复杂度：O(m)，其中 m 为 N 叉树的节点。题目中用到哈希表来记录节点的子节点访问记录，哈希表的存储空间等于树
// 的深度，如果 N 叉树的深度为 1 则此时栈与哈希表的空间均为 O(1)，如果 N 叉树的深度为 m−1 则此时栈与哈希表的
// 空间为 O(m−1)，平均情况下栈与哈希表的空间为 O(logm)，因此空间复杂度为 O(m)。
func preorder2(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{}
	nextIndex := map[*Node]int{}
	node := root
	for len(st) > 0 || node != nil {
		for node != nil {
			ans = append(ans, node.Val)
			st = append(st, node)
			if len(node.Children) == 0 {
				break
			}
			nextIndex[node] = 1
			node = node.Children[0]
		}
		node = st[len(st)-1]
		i := nextIndex[node]
		if i < len(node.Children) {
			nextIndex[node] = i + 1
			node = node.Children[i]
		} else {
			st = st[:len(st)-1]
			delete(nextIndex, node)
			node = nil
		}
	}
	return
}

// 方法三：迭代优化-官方
//思路
//在前序遍历中，我们会先遍历节点本身，然后从左向右依次先序遍历该每个以子节点为根的子树，此时利用栈先进后出的原理，
// 依次从右向左将子节点入栈，这样出栈的时候即可保证从左向右依次遍历每个子树。参考方法二的原理，可以提前将后续需要
// 访问的节点压入栈中，这样就可以避免记录每个节点的子节点访问数量。

//时间复杂度：O(m)，其中 m 为 N 叉树的节点。每个节点恰好被访问一次。
//空间复杂度：O(m)，其中 m 为 N 叉树的节点。如果 N 叉树的深度为 1 则此时栈的空间为 O(m−1)，如果 N 叉树的
// 深度为 m−1 则此时栈的空间为 O(1)，平均情况下栈的空间为 O(logm)，因此空间复杂度为 O(m)。
func preorder4(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
	}
	return
}
