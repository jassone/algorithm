package main

import "sort"

//LeetCode 15题. 三数之和

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
// 请你找出所有满足条件且不重复的三元组。

//注意： **** 答案中不可以包含重复的三元组。******

//示例：
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]

//思路
//注意[0， 0， 0， 0] 这组数据
// 本题与 1. 两数之和 类似，是非常经典的面试题，但是做法不尽相同。

// 方法1：暴力
// 三层循环
//时间复杂度 O(n^3)
//空间复杂度O(1)

// 方法2：哈希解法

//1 两层for循环就可以确定 a 和b 的数值了，可以使用哈希法来确定 0-(a+b) 是否在 数组里出现过，其实这个思路是正确的，
// 但是我们有一个非常棘手的问题，就是题目中说的不可以包含重复的三元组。
//2 把符合条件的三元组放进vector中，然后再去重，这样是非常费时的，很容易超时，也是这道题目通过率如此之低的根源所在。
//3 去重的过程不好处理，有很多小细节，如果在面试中很难想到位。
//4 时间复杂度可以做到 O(n^2)，但还是比较费时的，因为不好做剪枝操作。
//5 大家可以尝试使用哈希法写一写，就知道其困难的程度了。

//方法2:排序 + 双指针-官方

//时间复杂度：O(N^2)，其中 N 是数组 nums 的长度。
//  排序O(nlogn),搜索O(N^2)，所以综合下来是O(N^2)。
//空间复杂度：O(log N)。我们忽略存储答案的空间，额外的排序的空间复杂度为 O(logN)。然而我们修改了输入的数组
// nums，在实际情况下不一定允许，因此也可以看成使用了一个额外的数组存储了 nums 的副本并进行排序，空间复杂度为 O(N)。
func threeSum3(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] { // 这样就可以避免重复了
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

// 卡尔的解法
func threeSum444(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main() {

}
