package main

func subsets(nums []int) [][]int {

	ans := make([][]int, 0)

	path := make([]int, 0)
	n := len(nums)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int{}, path...))
		}
		dfs(i + 1)
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]

	}

	dfs(0)
	return ans
}
