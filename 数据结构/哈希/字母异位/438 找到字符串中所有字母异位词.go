package main

import "fmt"

//LeetCode 438. 找到字符串中所有字母异位词

//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

//异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

//示例 1:
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

// 方法一：滑动窗口-官方
//算法
//在算法的实现中，我们可以使用数组来存储字符串 pp 和滑动窗口中每种字母的数量。

//细节
//当字符串 s 的长度小于字符串 p 的长度时，字符串 s 中一定不存在字符串 p 的异位词。但是因为字符串 s 中无法构造长度
// 与字符串 p 的长度相同的窗口，所以这种情况需要单独处理。

//时间复杂度：O(m+(n−m)×Σ)，其中 n 为字符串 s 的长度，m 为字符串 p 的长度，Σ 为所有可能的字符数。
// 我们需要 O(m) 来统计字符串 p 中每种字母的数量；需要 O(m) 来初始化滑动窗口；需要判断
// n−m+1 个滑动窗口中每种字母的数量是否与字符串 p 中每种字母的数量相同，每次判断需要 )O(Σ) 。
// 因为 s 和 p 仅包含小写字母，所以 Σ=26。
//空间复杂度：O(Σ)。用于存储字符串 p 和滑动窗口中每种字母的数量。
func findAnagrams1(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	//fmt.Println(sCount, pCount)
	// 两个都是[1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	//因为之前已经将pLen长度的字符串处理了，所以这里可以直接比较起来
	if sCount == pCount {
		ans = append(ans, 0)
	}

	// 开始处理从0位置滑动窗口
	for i, ch := range s[:sLen-pLen] { //"cbaebab"
		// 将s的窗口往后滑动一下再比较
		sCount[ch-'a']-- // 从头开始滑动，减一个，下面再加一个，如果还是相等，则就符合要求
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

//方法二：优化的滑动窗口-官方-有点繁琐
//思路和算法
//在方法一的基础上，我们不再分别统计滑动窗口和字符串 p 中每种字母的数量，而是统计滑动窗口和字符串 p 中每种字母
//数量的差；并引入变量 differ 来记录当前窗口与字符串 p 中数量不同的字母的个数，并在滑动窗口
//的过程中维护它。
//在判断滑动窗口中每种字母的数量与字符串 p 中每种字母的数量是否相同时，只需要判断 differ 是否为零即可。

func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	count := [26]int{}
	for i, ch := range p {
		count[s[i]-'a']++
		count[ch-'a']--
	}

	differ := 0
	for _, c := range count {
		if c != 0 {
			differ++
		}
	}
	if differ == 0 { // 正好都抵消掉了
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		if count[ch-'a'] == 1 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[ch-'a'] == 0 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[ch-'a']--

		if count[s[i+pLen]-'a'] == -1 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[s[i+pLen]-'a'] == 0 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[s[i+pLen]-'a']++

		if differ == 0 {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams1(s, p))
}
