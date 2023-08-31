package main

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix))
	}
	ans := 0
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
			}
			if i-1 >= 0 && j-1 >= 0 && dp[i][j] == 1 {
				for k := 1; dp[i-1][j] == k && dp[i][j-1] == k && dp[i-1][j-1] == k; {
					dp[i][j]++
				}
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}

		}
	}
	return ans
}
