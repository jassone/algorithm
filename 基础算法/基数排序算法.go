package main

import "fmt"

//算法描述：基数排序类似计数排序，需要额外的空间来记录对应的基数内的数据 额外的空间是有序的，
//最终时间复杂度O(nlogrm),r是基数，r^m=n.当给定 特定的范围，基数排序又可以叫桶排序，
//当以10进制为基数时就是简单的桶排序

// 当前算法能自动

//算法步骤
// 从个位开始排序，从低到高进行递推
// 比较过程中如果遇到高位相同时，顺序不变
//算法分两类
// 低位排序LSD
// 高位排序MSD
func main() {
	var result [3][]int
	myarr := []int{1, 3, 3, 1, 1, 2, 2, 2, 2, 3}
	for i:=0;i<len(myarr);i++ {
		result[myarr[i]-1] = append(result[myarr[i]-1],myarr[i])
	}

	fmt.Println(result)
	// [[1 1 1] [2 2 2 2 2] [3 3 3]] 从小到大
}

// 上面方法有个问题，可能有空数组，比如下面这个demo
func main2() {
	myarr := []int{1, 3, 3, 1, 1, 2, 2, 2, 2, 3, 6}
	maxLen := max(myarr)
	result := make([][]int, maxLen)
	for i:=0; i<len(myarr); i++ {
		result[myarr[i]-1] = append(result[myarr[i]-1],myarr[i])
	}

	var res [][]int
	for _,v := range result {
		if 0 != len(v) {
			res = append(res,v)
		}
	}

	fmt.Println(res)
}
func max(nums []int) int {
	max := 0
	for _,v := range nums {
		if (max < v) {
			max = v
		}
	}

	return max
}