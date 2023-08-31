package main

import "sort"

func nextPermutation(nums []int) {

	r := len(nums) - 1
	for ; r >= 1; r-- {
		if nums[r-1] < nums[r] {
			break
		}
	}
	if r == 0 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}
	i := len(nums)
	for ; i > r; i-- {
		if nums[i] > nums[r-1] {
			break
		}
	}
	nums[r-1], nums[i] = nums[i], nums[r-1]
	sort.Ints(nums[r:])
	return

}
