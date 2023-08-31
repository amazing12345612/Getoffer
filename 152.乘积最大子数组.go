package main

import "math"

func maxProduct(nums []int) int {
	n := len(nums)
	dpmin := make([]int, n)
	dpmax := make([]int, n)
	if len(nums) <= 1 {
		return nums[0]
	}
	dpmax[0] = nums[0]
	dpmin[0] = nums[0]
	ans := math.MinInt32
	for i := 1; i < len(nums); i++ {

		dpmin[i] = min(min(dpmax[i-1]*nums[i], dpmin[i-1]*nums[i]), nums[i])
		dpmax[i] = max(max(dpmin[i-1]*nums[i], dpmax[i-1]*nums[i]), nums[i])
		if dpmax[i] > ans {
			ans = dpmax[i]
		}
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
