package main

import "fmt"

// LeetCode 977.有序数组的平方

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按
//非递减顺序 排序。

//示例 1：
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]

//方法一：直接排序
//思路与算法
//最简单的方法就是将数组 nums 中的数平方后直接排序。
// 这个时间复杂度是 O(n + nlogn)， 可以说是O(nlogn)的时间复杂度。

//方法三：双指针-官方-妙，推荐
//思路与算法
//同样地，我们可以使用两个指针分别指向位置 0 和 n−1，每次比较两个指针对应的数，
//选择较大的那个逆序放入答案并移动指针。这种方法无需处理某一指针移动至边界的情况，
//可以仔细思考其精髓所在。
//时间复杂度：O(n)，其中 n 是数组 nums 的长度。
//空间复杂度：O(1)。除了存储答案的数组以外，我们只需要维护常量空间。
func f3(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
}

// 方法2：双指针+归并排序-官方
//思路与算法
//方法一没有利用「数组 nums 已经按照升序排序」这个条件。显然，如果数组
//nums 中的所有数都是非负数，那么将每个数平方后，数组仍然保持升序；如果数组
//nums 中的所有数都是负数，那么将每个数平方后，数组会保持降序。

//这样一来，如果我们能够找到数组 nums 中负数与非负数的分界线，那么就可以用类似
//「归并排序」的方法了。具体地，我们设 neg 为数组 nums 中负数
//与非负数的分界线，也就是说 nums[0] 到 nums[neg] 均为负数，而 nums[neg+1]
//到 nums[n−1]  均为非负数。当我们将数组 nums 中的数平方后，
//那么 nums[0] 到 nums[neg] 单调递减，
//nums[neg+1] 到 nums[n−1] 单调递增。

//由于我们得到了两个已经有序的子数组，因此就可以使用归并的方法进行排序了。
//具体地，使用两个指针分别指向位置 neg 和 neg+1，
//每次比较两个指针对应的数，选择较小的那个放入答案并移动指针。当某一指针移至边界时，
//将另一指针还未遍历到的数依次放入答案。
//时间复杂度：O(n)，其中 n 是数组 nums 的长度。
//空间复杂度：O(1)。除了存储答案的数组以外，我们只需要维护常量空间
func f2(nums []int) []int {
	n := len(nums)
	lastNegIndex := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}

	ans := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		if i < 0 { // 负数读完了
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n { // 非负数读完了
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] { // 两边都没读完
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}

	return ans
}

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println(f2(arr))
	fmt.Println(f3(arr))
}
