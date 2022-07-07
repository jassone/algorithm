package main

//LeetCode 383.赎金信
//给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//如果可以，返回 true ；否则返回 false 。
//magazine 中的每个字符只能在 ransomNote 中使用一次。

//输入：ransomNote = "aa", magazine = "aab"
//输出：true

// 思路分析
//只需要满足字符串 magazine 中的每个英文字母 (’a’-’z’) 的统计次数都大于等于 ransomNote 中相同字母的统计次数即可。

// 方法1：统计字符-官方
//时间复杂度：O(m+n)，其中 m 是字符串 ransomNote 的长度，n 是字符串 magazine 的长度，我们只需要遍历两个字符一次即可。
//空间复杂度：O(∣S∣)，S 是字符集，这道题中 S 为全部小写英语字母，因此 ∣S∣=26。
func canConstruct(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {

}
