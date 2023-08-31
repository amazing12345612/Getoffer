package main

import "sort"

func findUnsortedSubarray(nums []int) int {

	sortNums := append([]int{}, nums...)
	sort.Ints(sortNums)
	l := 0
	r := len(nums) - 1
	for sortNums[l] == nums[l] {
		l++
	}
	for sortNums[r] == nums[r] {
		r--
	}
	return r - l + 1
}
