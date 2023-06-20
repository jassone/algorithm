package main

import (
	"fmt"
	"strings"
)

//LeetCode 71. 简化路径

//给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为更加简洁的规范路径。

//在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；
// 两者都可以是复杂相对路径的组成部分。任意多个连续的斜杠（即，'//'）都被视为单个斜杠 '/' 。 对于此问题，任何其他格式
// 的点（例如，'...'）均被视为文件/目录名称。

//请注意，返回的 规范路径 必须遵循下述格式：
// 始终以斜杠 '/' 开头。
// 两个目录名之间必须只有一个斜杠 '/' 。
// 最后一个目录名（如果存在）不能 以 '/' 结尾。
// 此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。

//返回简化后得到的 规范路径 。

//示例 4：
//输入：path = "/a/./b/../../c/"
//输出："/c"

//方法一：栈-官方
//时间复杂度：O(n)，其中 n 是字符串 path 的长度。
//空间复杂度：O(n)。我们需要 O(n) 的空间存储 names 中的所有字符串。
func simplifyPath(path string) string {
	// 整理下栈的实现
	stack := []string{}
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]// 去掉末尾最后一个目录
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	str := "/a/./b/../../c/"
	fmt.Println(simplifyPath(str))
}
