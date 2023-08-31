package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	ans := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
				if dp[i] > ans {
					ans = dp[i]
				}
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
