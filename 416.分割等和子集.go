package main

//func canPartition(nums []int) bool {
//	sum := 0
//	for _, v := range nums {
//		sum += v
//	}
//	if sum%2 != 0 {
//		return false
//	}
//	sum = sum / 2
//
//	var dfs func(i, c int) bool
//	dfs = func(i, c int) bool {
//		if c == 0 {
//			return true
//		}
//		if i >= len(nums) || nums[i] > c {
//			return false
//		}
//		return dfs(i+1, c) || dfs(i+1, c-nums[i])
//	}
//	return dfs(0, sum)
//
//}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	//拿到第i件物品时，背包容量为c时最大价值
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])

		}
	}
	return dp[sum] == sum
}

func lastStoneWeightII(stones []int) int {
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if stones[x] == 0 || stones[y] == 0 {
			return
		}

		stones[x], stones[y] = abs(x-y), abs(x-y)

		for i := 0; i < len(stones); i++ {
			for j := 0; j < len(stones) && j != i; j++ {
				dfs(i, j)
			}
		}
	}

	return 0
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
