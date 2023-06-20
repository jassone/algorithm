package main

import "fmt"

//LeetCode 459.重复的子字符串 todo

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

//示例 1:
//输入: "abab"
//输出: True
//解释: 可由子字符串 "ab" 重复两次构成。

//示例 2:
//输入: "aba"
//输出: False

//示例 3:
//输入: "abcabcabcabc"
//输出: True
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

// 注意这里题目要看清：一定要是子串拼接构成父串

// 方法1：枚举-官方
//思路与算法
//如果一个长度为 nn 的字符串 ss 可以由它的一个长度为 n′的子串 s′重复多次构成，那么：
// 1 n 一定是 n′ 的倍数；
// 2 s′ 一定是 s 的前缀；
// 3 对于任意的 i \in [n', n)i∈[n′,n)，有 s[i] = s[i-n']s[i]=s[i−n′]。

// 也就是说，ss 中长度为 n′ 的前缀就是 s′ ，并且在这之后的每一个位置上的字符 s[i]，都需要与它之前的
// 第 n′ 个字符 s[i-n']s[i−n′] 相同。
// 因此，我们可以从小到大枚举 n′ ，并对字符串 s 进行遍历，进行上述的判断。

//时间复杂度：O(n^2)，其中 n 是字符串 s 的长度。枚举 i 的时间复杂度为 O(n)，遍历 s 的时间复杂度为 O(n)，
// 相乘即为总时间复杂度。
//空间复杂度：O(1)。
func repeatedSubstringPattern1(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

// 方法2：kmp-官方
//时间复杂度：O(n)，其中 n 是字符串 s 的长度。
//空间复杂度：O(n)。
func repeatedSubstringPattern2(s string) bool {
	return kmp(s+s, s)
}
func kmp(query, pattern string) bool {
	n, m := len(query), len(pattern)
	fail := make([]int, m)
	for i := 0; i < m; i++ {
		fail[i] = -1
	}
	for i := 1; i < m; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	match := -1
	for i := 1; i < n-1; i++ {
		for match != -1 && pattern[match+1] != query[i] {
			match = fail[match]
		}
		if pattern[match+1] == query[i] {
			match++
			if match == m-1 {
				return true
			}
		}
	}
	return false
}

// 方法3：kmp算法-优化后的版-官方
//时间复杂度：O(n)，其中 n 是字符串 s 的长度。
//空间复杂度：O(n)。
func repeatedSubstringPattern3(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	// next[n-1]  最长相同前后缀的长度
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}

	return false
}

func main() {
	str := "abab"
	fmt.Println(repeatedSubstringPattern1(str))

}
