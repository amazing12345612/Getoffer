package main

func minArray(numbers []int) int {

	var dfs func(nums []int, l, r int) int

	dfs = func(nums []int, l, r int) int {
		if l >= r {
			return nums[l]
		}
		mid := (r - l) / 2
		if nums[mid] < nums[r] {
			return dfs(nums, l, mid)
		} else if nums[mid] > nums[r] {
			return dfs(nums, mid+1, r)
		} else {
			return min(dfs(nums, l, mid), dfs(nums, mid, r))
		}
	}
	return dfs(numbers, 0, len(numbers)-1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
