package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {

		dp[i] = math.MaxInt32
		for _, v := range coins {
			if i-v >= 0 && dp[i-v] != -1 {
				dp[i] = min(1+dp[i-v], dp[i])
			}
		}
		if dp[i] == math.MaxInt32 {
			dp[i] = -1
		}
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return b
	}
	return a
}
