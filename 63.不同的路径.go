package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	dp := make([][]int, len(obstacleGrid)+1)
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	for i := 0; i < len(obstacleGrid)+1; i++ {
		dp[i] = make([]int, len(obstacleGrid[0])+1)
	}
	for i := 0; i < len(obstacleGrid); i++ {
		if dp[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 0; i < len(obstacleGrid[0]); i++ {
		if dp[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	return dp[m-1][n-1]

}