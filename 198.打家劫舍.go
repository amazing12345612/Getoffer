package main

func rob(nums []int) int {
	dp := make([]int, len(nums))

	if len(nums) <= 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return max(dp[len(nums)-1], dp[len(nums)-2])
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
