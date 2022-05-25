package array

/**
	title：有序数组的平方
	solution:
	双指针
	创建一个长度为n的切片
	从后往前插入
**/

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	k := n - 1
	for k >= 0 {

		a, b := nums[left]*nums[left], nums[right]*nums[right]

		if a < b {
			right--
			res[k] = b
		} else {
			left++
			res[k] = a
		}
		k--
	}

	return res

}
