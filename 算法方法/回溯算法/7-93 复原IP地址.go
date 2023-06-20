package main

import (
	"fmt"
	"strconv"
)

//LeetCode 93.复原IP地址

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

//有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

//例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
// 和 "192.168@1.1" 是 无效的 IP 地址。

//示例 1：
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]

//示例 2：
//输入：s = "0000"
//输出：["0.0.0.0"]

//0 <= s.length <= 3000
//s 仅由数字组成

// 卡尔
// https://www.programmercarl.com/0093.%E5%A4%8D%E5%8E%9FIP%E5%9C%B0%E5%9D%80.html
// 和 131分割回文很类似

//判断子串是否合法
//主要考虑到如下三点：
// 段位(字符长度大于等于2)以0为开头的数字不合法
// 段位里有非正整数字符不合法
// 段位如果大于255了不合法

func restoreIpAddresses(s string) []string {
	var res, path []string
	backTracking(s, path, 0, &res)
	return res
}
func backTracking(s string, path []string, startIndex int, res *[]string) {
	//终止条件,分割到最后了，且切割后有四个部分
	if startIndex == len(s) && len(path) == 4 {
		tmpIpString := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
		*res = append(*res, tmpIpString)
		return
	}

	for i := startIndex; i < len(s); i++ {
		if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {
			//处理
			path := append(path, s[startIndex:i+1])

			//递归
			backTracking(s, path, i+1, res)

			//回溯
			path = path[:len(path)-1]
		} else { //如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
			return
		}
	}
}
func isNormalIp(s string, startIndex, end int) bool {
	checkInt, _ := strconv.Atoi(s[startIndex : end+1])

	// 比如在0000的判断时，可以是[0.0.0.0]，所以只有当连续字符长度大于等于2时，才去验证首字母不能为零
	if end-startIndex+1 > 1 && s[startIndex] == '0' { //对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}

func main() {
	//str := "25525511135"
	str := "0000"
	fmt.Println(restoreIpAddresses(str))
}

//官方-太复杂
//https://leetcode.cn/problems/restore-ip-addresses/solution/fu-yuan-ipdi-zhi-by-leetcode-solution/
