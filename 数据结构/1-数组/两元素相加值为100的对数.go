package main

import "fmt"

// 给定一个数组(元素不重复)，求两个数之和=给定值sum的所有组合个数。
// 【方法一】穷举法
//  从数组中任意找两个数，看其和是否=sum。时间复杂度O(N^2)

// 【方法二】】hash表法
//  只需要遍历一遍数组，非常高效
func f(arr []int) [][]int {
	var result [][]int
	var resultDetail []int
	chaMap := make(map[int]int) // 差集hash
	for k,v := range arr  {
		if index,ok := chaMap[v] ; ok {
			resultDetail = []int{} // 重置
			resultDetail = append(resultDetail,arr[index],arr[k])
			result = append(result,resultDetail)
		} else {
			chaMap[100-v] = k
		}
	}

	return result
}

func main() {
	arr := []int{10,20,30,90,80}
	fmt.Println(f(arr))
}
