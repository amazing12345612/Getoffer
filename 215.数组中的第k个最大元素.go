package main

func findKthLargest(nums []int, k int) int {
	h := len(nums)
	build(nums, h)
	for i := 1; i < k; i++ {
		nums[0], nums[h-i-1] = nums[h-i-1], nums[0]
		build(nums, h-i)
	}
	return nums[0]
}

func build(nums []int, h int) {

	for i := h / 2; i >= 0; i-- {
		down(nums, i, h)
	}
}

func down(nums []int, i, h int) {
	t := i
	for t < h {
		z := t
		j1 := 2*t + 1
		if j1 < h && nums[j1] > nums[t] {
			t = j1
		}
		j2 := 2*t + 2
		if j2 < h && nums[j2] > nums[t] {
			t = j2
		}
		if t == z {
			break
		}
		nums[z], nums[t] = nums[t], nums[z]
	}
}
