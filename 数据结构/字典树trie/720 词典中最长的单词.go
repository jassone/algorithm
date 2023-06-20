package main

import (
	"fmt"
	"sort"
)

// 720. 词典中最长的单词

// 给出一个字符串数组 words 组成的一本英语词典。返回 words 中最长的一个单词，该单词是由 words 
// 词典中其他单词逐步添加一个字母组成。
// 若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。

//示例 1：
//输入：words = ["w","wo","wor","worl", "world"]
//输出："world"
//解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。

//示例 2：
//输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
//输出："apple"
//解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"

//方法一：哈希集合
//思路和算法

//定义「符合要求的单词」如下：
//空字符串是符合要求的单词；

//在符合要求的单词的末尾添加一个字母，得到的新单词是符合要求的单词。

//这道题要求返回数组
//words
//words 中的最长的符合要求的单词，如果有多个最长的符合要求的单词则返回其中字典序最小的单词。以下将返回值称为「答案」。

//为了方便处理，需要将数组
//words
//words 排序，排序的规则是首先按照单词的长度升序排序，如果单词的长度相同则按照字典序降序排序。排序之后，可以确保当遍历到每个单词时，比该单词短的全部单词都已经遍历过，且每次遇到符合要求的单词一定是最长且字典序最小的单词，可以直接更新答案。

//将答案初始化为空字符串。使用哈希集合存储所有符合要求的单词，初始时将空字符串加入哈希集合。遍历数组
//words
//words，对于每个单词，判断当前单词去掉最后一个字母之后的前缀是否在哈希集合中，如果该前缀在哈希集合中则当前单词是符合要求的单词，将当前单词加入哈希集合，并将答案更新为当前单词。

//遍历结束之后，返回答案。

// todo
func longestWord1(words []string) (longest string) {
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})

	candidates := map[string]struct{}{"": {}}
	for _, word := range words {
		if _, ok := candidates[word[:len(word)-1]]; ok {
			longest = word
			candidates[word] = struct{}{}
		}
	}
	return longest
}

// 方法二：字典树
//预备知识
//该方法需要使用字典树。如果读者对字典树不了解，建议首先阅读「208. 实现 Trie (前缀树) 的官方题解」，在理解字典树的实现之后继续阅读。

//思路和算法

//由于符合要求的单词的每个前缀都是符合要求的单词，因此可以使用字典树存储所有符合要求的单词。

//创建字典树，遍历数组
//words
//words 并将每个单词插入字典树。当所有的单词都插入字典树之后，将答案初始化为空字符串，再次遍历数组
//words
//words，判断每个单词是否是符合要求的单词，并更新答案。如果一个单词是符合要求的单词，则比较当前单词与答案，如果当前单词的长度大于答案的长度，或者当前单词的长度等于答案的长度且当前单词的字典序小于答案的字典序，则将答案更新为当前单词。

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil || !node.children[ch].isEnd {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// todo
func longestWord2(words []string) (longest string) {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	for _, word := range words {
		if t.Search(word) && (len(word) > len(longest) || len(word) == len(longest) && word < longest) {
			longest = word
		}
	}
	return longest
}

func main() {
	arr := []string{"w","wo","wor","worl", "world",}
	fmt.Println(longestWord2(arr))
	
}
