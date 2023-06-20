package main

import (
	"fmt"
	"sort"
)

// todo
//LeetCode 56. 合并区间
// https://leetcode-cn.com/problems/merge-intervals/solution/he-bing-qu-jian-by-leetcode-solution/

//数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

//输入：intervals = [ [1，3]，[2，6]，[8，10]，[15，18] ]
//输出：[ [1，6]，[8，10]，[15，18] ]
//解释：区间 [1，3] 和 [2，6]
//重叠，将它们合并为 [1，6]

//归纳为三种情况
// ----1---            -----1-----          ------1-----
//           ----2---         ----2----      -----2---

//首先，我们用数组 merged 存储最终的答案。然后我们将第一个区间加入 merged 数组中，
//并按顺序依次考虑之后的每个区间：

//  1 如果当前区间的左端点在数组 merged 中最后一个区间的右端点之后，即上图中的第1种情况，
//  那么它们不会重合。我们可以直接将这个区间加入数组 merged 的末尾；

//  2 否则，它们是有重合部分的，即上图中的第二、三种情况，我们需要用当前区间的右端点更新数组
//  merged 中最后一个区间的右端点，将其置为二者的较大值。

//该算法的时间复杂度是O(nlogn)，其中n是区间的数量，除去排序的开销，我们只需要一次线性扫描，
//所以主要的时间开销是排序的 O(nlogn)。空间复杂度是O(logn)。
func f1(arr [][]int) [][]int {
	// 将区间数组按照左端点进行升序排序
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] //按照每行的第一个元素排序
	})

	//fmt.Println(arr)

	res := [][]int{}
	cur := arr[0]
	for k, v := range arr {
		if 0 == k {
			continue
		}

		if cur[1] < v[0] {
			if 1 == k {
				res = append(res, cur)
			}
			res = append(res, v)
			cur = v
		} else {
			if cur[1] < v[1] {
				cur = []int{cur[0], v[1]}
				res = append(res, []int{cur[0], v[1]})
			}
		}
		//fmt.Println(res)
	}

	return res
}

func main() {
	arr := [][]int{{1, 3}, {2, 6}, {15, 18}, {3, 5}, {8, 10}}
	fmt.Println(f1(arr))
}
