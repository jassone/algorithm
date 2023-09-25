package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 计数排序--升序
func ContingSortAsc(slice []int) {
	// 创建map统计0-999每个数出现的次数
	m := make(map[int]int)
	// 遍历待排序的数据，统计结果
	for _, v := range slice {
		m[v]++
	}
	// 借助map，统计排序的数据重新赋值为原序列
	slice = slice[0:0] // 将原序列清空
	for i := 0; i < 1000; i++ {
		// 数据出现的次数：m[i]的值
		for j := 0; j < m[i]; j++ {
			slice = append(slice, i) // 重新赋值
		}
	}
}

// 计数排序--降序
func ContingSortDesc(slice []int) {
	// 创建map统计0-999每个数出现的次数
	m := make(map[int]int)
	// 遍历待排序的数据，统计结果
	for _, v := range slice {
		m[v]++
	}
	// 借助map，统计排序的数据重新赋值为原序列
	slice = slice[0:0] // 将原序列清空
	for i := 999; i >= 0; i-- {
		// for i := 0; i < 1000; i++ {
		// 数据出现的次数：m[i]的值
		for j := 0; j < m[i]; j++ {
			slice = append(slice, i) // 重新赋值
		}
	}
}

func main() {
	slice := make([]int, 0)
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成100个1000以内的随机数
	for i := 1; i <= 100; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	fmt.Println("原数据：", slice)
	ContingSortAsc(slice)
	fmt.Println("计数排序升序：", slice)
	//ContingSortDesc(slice)
	//fmt.Println("计数排序降序：", slice)
}
