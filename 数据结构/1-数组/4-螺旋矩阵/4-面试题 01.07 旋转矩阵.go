package main

import "fmt"

//LeetCode 面试题 01.07. 旋转矩阵

//有一个 NxN 整数矩阵，请编写一个算法，将矩阵顺时针旋转90度。
//要求：时间复杂度O(N^2)，空间复杂度是O(N^2)。

//进阶：时间复杂度是O(N^2)，空间复杂度是O(1)

//示例：
//1 2
//3 4
//变成
//3 1
//4 2

// 方法一：使用辅助数组-官方
//旋转后，第一行的第 i 个元素在旋转后恰好是倒数第一列的第 i 个元素。对于第二行的元素也是如此，
//在旋转后变成倒数第二列的元素，并且第二行的第i个元素在旋转后恰好是倒数第二列的第i个元素。
//所以，我们可以得出规律，对于矩阵中第 i 行的第 j 个元素，在旋转后，它出现在倒数第 i 列的第 j 个位置，
//即对于矩阵中的 matrix[i] [j] 元素，在旋转后，它的新位置为 matrix [j] [n-i-1]。

//我们申请一个大小为 n * n 的新矩阵，来临时存储旋转后的结果。我们通过遍历matrix中的所有元素，
//根据上述规则将元素存放到新矩阵中的对应位置。在遍历完成后，再将新矩阵中复制到原矩阵即可。下面我们来
//时间复杂度是O(N^2)，
//空间复杂度O(N^2)。
func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-i-1] = v
		}
	}
	//copy(matrix, tmp) // 拷贝 tmp 矩阵每行的引用
	return tmp
}

// 方法2：原地替换法 todo
//对于arr中的元素，我们使用公式temp[j] [n - i - 1] = arr[i] [j]进行旋转，
//如果不申请辅助矩阵，我们直接把元素 arr[i] [j]，放到矩阵 arr[j] [n - i - 1]位置，
//原矩阵中的arr[j] [n - i - 1]元素就被覆盖了，这显然不是我们要的结果。

//一次可以原地交换四个位置，所以：
// 1 当n为偶数时，我们需要选取 n^2 / 4 = (n/2) * (n/2)个元素进行原地交换操作，
//   可以将该图形分为四块，可以保证不重复、不遗漏旋转所有元素；
// 2 当n为奇数时，由于中心的位置经过旋转后位置不变，我们需要选取 (n^2-1)/4=(n-1)/2 * (n+1) /2
//   个元素进行原地交换操作，我们以5*5的矩阵为例，可以按照以下方式划分，进而保证不重复、不遗漏的旋转所有元素。
//时间复杂度是O(n^2)，
//空间复杂度是O(1)。
func f2(arr [][]int) [][]int {
	//矩阵的大小
	n := len(arr)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			// 进行一轮原地旋转，旋转4个元素
			temp := arr[i][j]
			arr[i][j] = arr[n-j-1][i]
			arr[n-j-1][i] = arr[n-i-1][n-j-1]
			arr[n-i-1][n-j-1] = arr[j][n-i-1]
			arr[j][n-i-1] = temp
		}
	}

	return arr
}

// 方法3：用翻转代替旋转-官方
//时间复杂度：O(N^2)，其中 N 是 matrix 的边长。对于每一次翻转操作，我们都需要枚举矩阵中一半的元素。
//空间复杂度：O(1)。为原地翻转得到的原地旋转。
func rotate3(matrix [][]int) [][]int {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(rotate(arr))
	fmt.Println(f2(arr))
	fmt.Println(rotate3(arr))

}
