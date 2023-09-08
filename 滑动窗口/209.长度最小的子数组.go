package 滑动窗口

import "math"

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt8
	l := 0
	r := 0
	tmp := 0
	for ; r < len(nums); r++ {
		if tmp < target {
			tmp += nums[r]
		}
		for tmp-nums[l] >= target && l <= r {
			tmp -= nums[l]
			l++
		}
		if tmp >= target && r-l+1 < ans {
			ans = r - l + 1
		}
	}
	return ans

}
