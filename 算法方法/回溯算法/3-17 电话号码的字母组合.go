package main

import "fmt"

//LeetCode 17.电话号码的字母组合

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

// 卡尔
// https://www.programmercarl.com/0017.%E7%94%B5%E8%AF%9D%E5%8F%B7%E7%A0%81%E7%9A%84%E5%AD%97%E6%AF%8D%E7%BB%84%E5%90%88.html#%E6%95%B0%E5%AD%97%E5%92%8C%E5%AD%97%E6%AF%8D%E5%A6%82%E4%BD%95%E6%98%A0%E5%B0%84
// ****如果是多个集合取组合，各个集合之间相互不影响，那么就不用startIndex。******

var res = make([]string, 0)
var digitsMap = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func letterCombinations555(digits string) []string {
	lenth := len(digits)
	if lenth == 0 || lenth > 4 {
		return nil
	}
	recursion555("", digits, 0)
	return res
}
func recursion555(tempString, digits string, Index int) { //index表示第几个数字
	if len(tempString) == len(digits) { //终止条件，字符串长度等于digits的长度
		res = append(res, tempString)
		return
	}
	tmpK := digits[Index] - '0' // 将index指向的数字转为int（确定下一个数字）
	letter := digitsMap[tmpK]   // 取数字对应的字符集
	for i := 0; i < len(letter); i++ {
		tempString = tempString + string(letter[i]) //拼接结果
		recursion555(tempString, digits, Index+1)
		//tempString = tempString[:len(tempString)-1] //回溯
	}
}

func main() {
	str := "23"
	//fmt.Println(letterCombinations555(str))
	fmt.Println(letterCombinations1(str))
}

// 方法1-回溯-官方
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}
func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
			// 注意这里并没有给combination赋值，所以不需要再显示回溯
		}
	}
}
