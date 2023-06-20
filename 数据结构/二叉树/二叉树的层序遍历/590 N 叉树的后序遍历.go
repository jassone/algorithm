package main

// 同类型590 589 429
//LeetCode 590. N 叉树的后序遍历

// 方法一：递归-官方，推荐
//思路
//递归思路比较简单，N 叉树的前序遍历与二叉树的后序遍历的思路和方法基本一致，可以参考「145. 二叉树的后序遍历」
// 的方法，每次递归时，先递归访问每个孩子节点，然后再访问根节点即可。

//时间复杂度：O(m)，其中 m 为 N 叉树的节点。每个节点恰好被遍历一次。
//空间复杂度：O(m)，递归过程中需要调用栈的开销，平均情况下为 O(logm)，最坏情况下树的深度为 m−1，需要的空间为 O(m−1)，因此空间复杂度为 O(m)。
func postorder1(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, ch := range node.Children {
			dfs(ch)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}

//方法二：迭代-官方
// 思路
//方法一中利用递归来遍历树，实际的递归中隐式利用了栈，在此我们可以直接模拟递归中栈的调用。
// 在后序遍历中从左向右依次先序遍历该每个以子节点为根的子树，然后先遍历节点本身。

//时间复杂度：O(m)，其中 m 为 N 叉树的节点。每个节点恰好被访问一次。
//空间复杂度：O(m)，其中 m 为 N 叉树的节点。如果 N 叉树的深度为 1 则此时栈和哈希表的空间为 O(1)，
// 如果 N 叉树的深度为 m−1 则此时栈和哈希表的空间为 O(m−1)，平均情况下栈和哈希表的空间为 O(logm)，
// 因此空间复杂度为 O(m)。
func postorder2(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{}
	nextIndex := map[*Node]int{}
	node := root
	for len(st) > 0 || node != nil {
		for node != nil {
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
			ans = append(ans, node.Val)
			st = st[:len(st)-1]
			delete(nextIndex, node)
			node = nil
		}
	}
	return
}

//方法三：迭代优化-官方
//在后序遍历中，我们会先从左向右依次后序遍历每个子节点为根的子树，再遍历根节点本身。此时利用栈先进后出的原理，
// 依次从右向左将子节点入栈，这样出栈的时候即可保证从左向右依次遍历每个子树。参考方法二的原理，可以提前将后续
// 需要访问的节点压入栈中。

//时间复杂度：O(m)，其中 m 为 N 叉树的节点。每个节点恰好被访问一次。
//空间复杂度：O(m)，其中 m 为 N 叉树的节点。哈希表的空间为 O(m)，栈的空间与树的深度相同，栈的空间最大为
// O(m−1)，因此空间复杂度为 O(m)。
func postorder3(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{root}
	vis := map[*Node]bool{}
	for len(st) > 0 {
		node := st[len(st)-1]
		// 如果当前节点为叶子节点或者当前节点的子节点已经遍历过
		if len(node.Children) == 0 || vis[node] {
			ans = append(ans, node.Val)
			st = st[:len(st)-1]
			continue
		}
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
		vis[node] = true
	}
	return
}

// 方法四：利用前序遍历反转-官方
// 在前序遍历中，我们会先遍历节点本身，然后从左向右依次先序遍历该每个以子节点为根的子树，而在后序遍历中，
// 需要先从左到右依次遍历每个以子节点为根的子树，然后再访问根节点。

//时间复杂度：O(m)，其中 m 为 N 叉树的节点。每个节点恰好被访问一次。
//空间复杂度：O(m)，其中 m 为 N 叉树的节点。如果 N 叉树的深度为 1 则此时栈的空间为 O(1)，如果 N 叉树的
// 深度为 1 则此时栈的空间为 O(m−1)，平均情况下栈的空间为 O(logm)，因此空间复杂度为
// O(m)。
func postorder4(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		st = append(st, node.Children...)
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return
}
