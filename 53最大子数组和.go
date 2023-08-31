package main

func maxSubArray(nums []int) int {
	ans := nums[0]
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

// func maxSubArray(nums []int) int {
//
//		for i := 1; i < len(nums); i++ {
//			nums[i] += nums[i-1]
//		}
//		ans := nums[0]
//		for i := 0; i < len(nums); i++ {
//			ans = max(ans, nums[i])
//			for j := 0; j < i; j++ {
//				ans = max(nums[i]-nums[j], ans)
//			}
//		}
//		return ans
//	}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
