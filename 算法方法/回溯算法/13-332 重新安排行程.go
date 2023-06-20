package main

import "sort"

//LeetCode 332.重新安排行程

//给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，
// 对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。

//提示：

//如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
//所有的机场都用三个大写字母表示（机场代码）。
//假定所有机票至少存在一种合理的行程。
//所有的机票必须都用一次 且 只能用一次。

//示例 1：
//输入：[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
//输出：["JFK", "MUC", "LHR", "SFO", "SJC"]

// todo ，该题目难
// 卡尔
//这道题目有几个难点：
// 一个行程中，如果航班处理不好容易变成一个圈，成为死循环
// 有多种解法，字母序靠前排在前面，让很多同学望而退步，如何该记录映射关系呢 ？
// 使用回溯法（也可以说深搜） 的话，那么终止条件是什么呢？
// 搜索的过程中，如何遍历一个机场所对应的所有机场。

type pair struct {
	target  string
	visited bool
}
type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}

func findItinerary(tickets [][]string) []string {
	result := []string{}
	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string]pairs)
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(pairs, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	}
	for k, _ := range targets {
		sort.Sort(targets[k])
	}
	result = append(result, "JFK")
	var backtracking func() bool
	backtracking = func() bool {
		if len(tickets)+1 == len(result) {
			return true
		}
		// 取出起飞航班对应的目的地
		for _, pair := range targets[result[len(result)-1]] {
			if pair.visited == false {
				result = append(result, pair.target)
				pair.visited = true
				if backtracking() {
					return true
				}
				result = result[:len(result)-1]
				pair.visited = false
			}
		}
		return false
	}

	backtracking()

	return result
}

// 官方
// 本题和「753. 破解保险箱」类似，是力扣平台上为数不多的求解欧拉回路 / 欧拉通路的题目。读者可以配合着进行练习。

func findItinerary(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
