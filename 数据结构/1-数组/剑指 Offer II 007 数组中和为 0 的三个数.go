package main

import (
	"fmt"
	"sort"
)

//LeetCode 剑指 Offer II 007. 数组中和为 0 的三个数

//给定一个包含n个整数的数组nums，判断nums中是否存在三个元素 a ，b ，c ，
// 使得 a + b + c = 0 ？请找出所有和为 0 且不重复的三元组。
//示例：
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]

// 双指针法
//这道题目算是二数之和的升级版，所以我们也可以采用双指针法来求解。那三个数如何采用双指针法呢，
//其实很简单，我们先把一个数固定下来，然后另外两个数再使用双指针去寻找不就好了。按照惯例，
//我们首先对数组进行排序，然后固定第一个数first，假设为nums[i]，然后再使用双指针法在数组中
//寻找两数之和等于-nums[i]即可。由于题目要求所求的三元组是不重复的，所以需要判断去掉重复解。
//重复解主要有以下两种情况。

func f1(arr []int) [][]int {
	length := len(arr)
	if length < 3 {
		return [][]int{}
	}

	// 排序
	sort.Sort(sort.IntSlice(arr))
	//fmt.Println(arr)

	res := [][]int{}
	detail := []int{}

	right := length - 1
	left := 0

	var first, target int

	for k, _ := range arr {

		first = arr[k]
		if first > 0 { //第一个数大于0，由于第二个、第三个数都大于第一个数，所以不可能相加等于0
			continue
		}

		//已经查找过了，所以不需要再继续寻找，直接跳过
		if k > 0 && first == arr[k-1] {
			continue
		}

		//第三个数，开始时指向数组的最右端
		target = 0
		target -= first
		left = k + 1
		//right := length-1
		//fmt.Println(k,v,left,right)

		// 开始循环查找
		for left < right {
			//fmt.Println(left,right,target)
			if arr[left]+arr[right] < target {
				left++
			} else if arr[left]+arr[right] > target {
				right--
			} else {
				detail = detail[0:0]
				res = append(res, append(detail, first, arr[left], arr[right]))

				//如果left和left+1对于的元素相同，由于left已经添加到result中了
				//为了避免重复，我们跳过相同的元素
				if arr[left] == arr[left+1] {
					left++
				}
				// 同理
				if arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return res
}

// 方法2：排序+双指针-官方 todo
//时间复杂度：O(N^2)其中 N 是数组 nums 的长度。
//空间复杂度：O(logN)。我们忽略存储答案的空间，额外的排序的空间复杂度为 O(logN)。
//然而我们修改了输入的数组 nums，在实际情况下不一定允许，因此也可以看成使用了一个额外的数组存储了
//nums 的副本并进行排序，空间复杂度为 O(N)。
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func main() {
	arr := []int{-1, -1, 0, 1, 1, 1, 2, -1, -4}
	fmt.Println(f1(arr))
	fmt.Println(threeSum(arr))
}
