package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。

//返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。

//示例 1：
//输入: words = ["i", "love", "leetcode", "i", "love", "coding"], k = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。

//方法一：哈希表 + 排序
//思路及算法
//我们可以预处理出每一个单词出现的频率，然后依据每个单词出现的频率降序排序，最后返回前 k 个字符串即可。

//具体地，我们利用哈希表记录每一个字符串出现的频率，然后将哈希表中所有字符串进行排序，排序时，
//如果两个字符串出现频率相同，那么我们让两字符串中字典序较小的排在前面，否则我们让出现频率较高的排在前面。
//最后我们只需要保留序列中的前 k 个字符串即可。

func topKFrequent1(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

//方法二：优先队列  todo
//思路及算法

//对于前k 大或前k 小这类问题，有一个通用的解法：优先队列。

//在本题中，我们可以创建一个小根优先队列（顾名思义，就是优先队列顶端元素是最小元素的优先队列）。我们将每一个字符串插入到优先队列中，如果优先队列的大小超过了
//k，那么我们就将优先队列顶端元素弹出。这样最终优先队列中剩下的k 个元素就是前k 个出现次数最多的单词。
type pair struct {
	w string
	c int
}
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.c < b.c || a.c == b.c && a.w > b.w }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func topKFrequent2(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	h := &hp{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}


func main() {
	arr :=[]string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent1(arr,2))
	
}
