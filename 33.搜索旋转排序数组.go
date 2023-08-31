package main

func search(nums []int, target int) int {

	n := len(nums) - 1
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		mid := (l + r) / 2
		if l > r {
			return -1
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[n-1] {
			if nums[mid] > target && nums[n-1] < target {
				return dfs(l, mid-1)
			} else {
				return dfs(mid+1, r)
			}

		} else {
			if nums[mid] < target && nums[n-1] > target {
				return dfs(mid+1, r)
			} else {
				return dfs(l, mid-1)
			}
		}

	}
	return -1

}
