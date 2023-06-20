package main

import "fmt"

// 208 实现前缀树


type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch] // 把下一个节点赋值给当前节点
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func main() {
	c := Constructor()
	c.Insert("abc")
	c.Insert("abd")
	c.Insert("efg")
	fmt.Println(c.Search("abc"))
	fmt.Println(c.SearchPrefix("ab"))
	fmt.Println(c.StartsWith("ab"))
}