package array

/**
	title：长度最小的子数组
	solution:滑动窗口，注意考虑左指针移动的限制规则
**/

func minSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	left, right := 0, 0
	n := len(nums)
	res := n + 1
	sum := 0

	for right < n {
		sum += nums[right]
		// 如果和>=target
		if sum >= target {
			// 移动左指针直到sum < target
			for left <= right && sum >= target {
				// 如果更新答案
				if right-left+1 < res {
					res = right - left + 1
				}
				sum -= nums[left]
				left++
			}
		}
		right++
	}

	if res == n+1 {
		return 0
	}
	return res

}
