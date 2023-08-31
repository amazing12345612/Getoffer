package main

func subarraySum(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	m := map[int]int{}
	ans := 0
	for i := 0; i < len(nums); i++ {
		m[nums[i]+k]++
		if nums[i] == k {
			ans++
		}
		ans += m[nums[i]]
	}
	return ans

}
