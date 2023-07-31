package main

func exchange(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		for l < r && nums[l]%2 != 0 {
			l++
		}
		for r > l && nums[r]%2 != 1 {
			r--
		}
		if l >= r {
			return nums
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
