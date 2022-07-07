package main

import "fmt"

// 算法描述：冒泡算法，数组中前一个元素和后一个元素进行比较如果大于或者小于 前者就进行交换，
// 最终返回最大或者最小都冒到数组的最后序列时间复杂度为 O(n^2).

func maopao(s []int) []int {
	len := len(s)
	for i:=0;i<(len-1);i++ {
		for j:=i+1;j<len;j++ {
			if (s[i] > s[j]) {
				s[i],s[j] = s[j],s[i]
			}
		}
	}

	return s
}

func main() {
	s := []int{1, 4, 2, 9, 6, 5, 7, 3, 8}
	result := maopao(s)
	fmt.Println(result)
}