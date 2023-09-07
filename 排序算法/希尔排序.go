package main

import "fmt"

// 增量序列折半的希尔排序 todo
func SortShell(list []int) {
	// 数组长度
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入那么交换元素
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SortShell(list)
	fmt.Println(list)
	//
	//list = []int{-1, 5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	//SortShell(list)
	//fmt.Println(list)
	//
	//list = []int{-1, 11, 22, 33}
	//SortShell(list)
	//fmt.Println(list)

}
