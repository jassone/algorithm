package main

import (
	"fmt"
	"math"
)

// LeetCode 162. 寻找峰值

//峰值元素是指其值严格大于左右相邻值的元素。给你一个整数数组 nums，找到峰值元素并返回其索引。
//数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//你可以假设 nums[-1] = nums[n] = -∞ 。你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

//示例：
//输入：nums = [1,2,3,1]
//输出：3 是峰值元素，你的函数应该返回其索引 2。

//输入：nums = [1,2,1,3,5,6,4]
//输出：1 或 5
//解释：你的函数可以返回索引 1，其峰值元素为 2；
//     或者返回索引 5， 其峰值元素为 6。

// 方法1 ：遍历
// 本题首先最容易想到的就是遍历数组，找出第一个符合条件的元素就可以了，这很简单，
// 但是不满足O(logN)的时间复杂度要求
//时间复杂度：O(n)，其中 n 是数组 nums 的长度。
//空间复杂度：O(1)。
func findPeakElement(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	// 修复下，如果数组是{1,2,3}，则不符合
	if idx == len(nums)-1 {
		idx = -1
	}

	return
}

//遍历2
func f2(nums []int) int {
	var max int
	n := len(nums)
	// 数组长度为1，那第一个就是最大，位置为0
	if n == 1 {
		max = 0
	} else if n == 2 {
		// 数组长度为2，比较一下，如果第2个大于第1个，位置1为最大，其他都是位置0最大
		if nums[1] > nums[0] {
			max = 1
		} else {
			max = 0
		}
	} else if n > 2 {
		// 数组长度大于2，遍历数组，如果有大于左边并且大于右边的，记下位置并且跳出遍历
		for i := 1; i < n-1; i++ {
			if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
				max = i
				break
			} else if nums[1] < nums[0] {
				// 如果没有符合上面的要求，就剩下一直增的数组或者一直减的数组两种情况，只需要比较第一和第二个数
				max = 0
				// 第一个数大，那位置0就是最大的
			} else if nums[1] > nums[0] {
				max = n - 1
				// 第二个数大，那位置n-1就是最大的
			}
		}
	}

	return max
}

// 方法2：二分法，推荐
//定义高低两个指针low，high，分别指向数组首尾
//1 定义中间指针mid=(low+high)/2我们知道二分法的结束标志是low>=high，当low==high的时候，算法肯定结束
//2 然后我们每次判断中间值，如果中间值处于下降状态，那么可以肯定中间值的左边有比中间值大的数，
// 而数组两头又是非常小的数，所以峰值肯定在左边，所以我们二分法的下一次开始就是让high=mid
//3 如果中间值处于上升状态，那么峰值肯定就在中间值的右边，所以二分法的下一次开始就是让low=mid+1
//你需要仔细感受high=mid和low=mid+1这两个等式，揣摩细微之处，可以发现low随着mid移动，而high最多等于mid，
//仔细揣摩他们的区别
func f1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

// 二分法-官方
//时间复杂度：O(logn)，其中 nn 是数组 nums 的长度。
//空间复杂度：O(1)。
func f3(nums []int) int {
	n := len(nums)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1

	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func main() {
	//arr := []int{1,2,3,1} // 3
	arr := []int{1, 2, 3, 1} // 3
	fmt.Println(f1(arr), arr[f1(arr)])

	fmt.Println(f2(arr), arr[f2(arr)])

	fmt.Println(f3(arr), arr[f3(arr)])

	fmt.Println(f4(arr), arr[f4(arr)])
}

// 其他的
func f4(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {

		return len(nums) - 1
	}

	l, r := 0, len(nums)-1

	return find(nums, l, r)
}

func find(nums []int, l, r int) int {
	mid := (l + r) / 2
	// 满足局部最大
	if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
		return mid
	}

	if nums[mid-1] > nums[mid] {
		return find(nums, l, mid)
	} else {
		return find(nums, mid, r)
	}
}
